package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jmoiron/sqlx"
	"github.com/robfig/cron"

	"github.com/wuleying/go-labs/silver-monitor/src/model"
	"github.com/wuleying/go-labs/silver-monitor/src/util"
)

// 全局配置
var config util.Config
var err error
var db *sqlx.DB

func main() {
	// 保存pid
	util.FileSavePid("silver-monitor-server.pid")

	config, _ = util.ConfigInit()
	if err != nil {
		log.Fatal("Init config failed: ", err.Error())
	}

	// 初始化模型
	db = model.Init(config)

	crontab := cron.New()

	crontab.AddFunc(config.Setting["schedule"], func() {
		getPrice()
	})

	crontab.Start()

	select {}
}

// 获取实时银价
func getPrice() {
	// 抓取目标页
	var url string = config.Setting["tendency_url"]

	doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Print(err)
		return
	}

	selection := doc.Find("#TABLE1 tbody tr").Eq(2).Find("td")

	var prices = make(map[int]string)

	selection.Each(func(i int, tag *goquery.Selection) {
		if i >= 2 && i <= 6 {
			prices[i] = strings.TrimSpace(tag.Text())
		}
	})

	id := model.LogSaveData(db, prices)

	if prices[2] == "" {
		log.Print("get price failed")
		return
	}

	price, _ := strconv.ParseFloat(prices[2], 64)

	alert_price, _ := strconv.ParseFloat(config.Setting["alert_price"], 64)

	if price <= alert_price {
		go func() {
			email_err := util.EmailSend(
				config.Email["user"],
				config.Email["passwd"],
				fmt.Sprintf("%s:%s", config.Email["host"], config.Email["port"]),
				config.Email["to"],
				config.Email["subject"],
				fmt.Sprintf("当前价格:%f", price))

			if email_err != nil {
				log.Printf("Send email failed: %s. Current price:%f", email_err, price)
				return
			}
		}()
	}

	log.Printf("monitor id:%d", id)
}
