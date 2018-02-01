package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"github.com/wuleying/go-labs/silver-president/src/model"
	"github.com/wuleying/go-labs/silver-president/src/util"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 全局配置
var config util.Config
var err error

var timeNow = time.Now().Local()

func main() {
	// 保存pid
	util.FileSavePid("silver-president.pid")

	config, _ = util.ConfigInit()
	if err != nil {
		log.Panic("Init config failed: ", err.Error())
	}

	crontab := cron.New()

	crontab.AddFunc(config.Setting["schedule"], func() {
		getData()
	})

	crontab.Start()

	select {}
}

// 获取数据
func getData() {
	// 抓取目标页
	var target_url string = fmt.Sprintf(config.Setting["target_url"], timeNow.Format("20060102"))

	resp, err := http.Get(target_url)

	if err != nil {
		log.Panic("Get target url context failed: ", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Panic("Read context failed: ", err.Error())
	}

	var jsonDataList []util.JsonData

	err = json.Unmarshal(body, &jsonDataList)
	if err != nil {
		log.Panic("Josn decode failed: ", err.Error())
	}

	// 获取返回的数据数量
	jsonDataListLen := len(jsonDataList)

	if jsonDataListLen <= 0 {
		return
	}

	// 初始化模型
	db := model.Init(config)
	defer db.Close()

	// 获取当天已写入的数据
	dataList := model.CalendarGetAll(db, timeNow.Format("2006-01-02"))

	// 获取当天已写入的数据数量
	dataListLen := len(dataList)

	if jsonDataListLen <= dataListLen {
		return
	}

	for _, value := range jsonDataList {
		for _, v := range dataList {
			// 过滤已入库数据
			if value.Title == v.Title {
				continue
			}
		}

		insertId := model.CalendarSaveData(db, value)
		log.Printf("Insert id: [%d], title: [%s]", insertId, value.Title)

		util.Notification(fmt.Sprintf("%s  [%s]", value.Title, value.InputDate))
	}
}
