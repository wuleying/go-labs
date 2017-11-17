package main

import (
    "log"
    "strings"
    "strconv"
    "fmt"

    "github.com/robfig/cron"
    "github.com/PuerkitoBio/goquery"
    _ "github.com/go-sql-driver/mysql"

    "go-labs/silver-monitor/src/model"
    "go-labs/silver-monitor/src/util"
)

// 全局配置
var config util.Config
var err error

func main() {
    // 保存pid
    util.SavePid("silver-monitor-server.pid")

    config, _ = util.InitConfig();
    if err != nil {
        log.Fatal("Init config failed: ", err.Error())
    }

    // 初始化模型
    model.InitModel(config)

    crontab := cron.New()

    crontab.AddFunc(config.Setting["schedule"], func() {
        getPrice();
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
        if (i >= 2 && i <= 6) {
            prices[i] = strings.TrimSpace(tag.Text())
        }
    })

    id := model.LogSaveData(prices)

    if (prices[2] == "") {
        log.Print("get price failed")
        return
    }

    price, _ := strconv.ParseFloat(prices[2], 64)

    alert_price, _ := strconv.ParseFloat(config.Setting["alert_price"], 64)

    if (price <= alert_price) {
        go func() {
            email_err := util.SendMail(
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