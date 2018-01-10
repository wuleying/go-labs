package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-labs/silver-president/src/model"
	"go-labs/silver-president/src/util"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 全局配置
var config util.Config
var err error
var db *sqlx.DB

var timeNow = time.Now().Local()

func main() {
	// 保存pid
	util.FileSavePid("silver-monitor-server.pid")

	config, _ = util.ConfigInit()
	if err != nil {
		log.Fatal("Init config failed: ", err.Error())
	}

	// 初始化模型
	db = model.Init(config)

	getData()
}

// 获取数据
func getData() {
	// 抓取目标页
	var target_url string = fmt.Sprintf(config.Setting["target_url"], timeNow.Format("20060102"))

	resp, err := http.Get(target_url)

	if err != nil {
		log.Fatal("Get target url context failed: ", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Read context failed: ", err.Error())
	}

	var jsonDataList []util.JsonData

	err = json.Unmarshal(body, &jsonDataList)
	if err != nil {
		log.Fatal("Josn decode failed: ", err.Error())
	}

	// 获取返回的数据数量
	jsonDataListLen := len(jsonDataList)

	if jsonDataListLen <= 0 {
		return
	}

	// 获取当天已写入的数据
	dataList := model.CalendarGetAll(db, timeNow.Format("2006-01-02"))

	// 获取当天已写入的数据数量
	dataListLen := len(dataList)

	if jsonDataListLen == dataListLen {
		return
	}

	for _, value := range jsonDataList {
		for _, v := range dataList {
			// 过滤已入库数据
			if value.Title == v.Title && value.InputDate == v.InputDate {
				continue
			}
		}

		insertId := model.CalendarSaveData(db, value)
		log.Printf("Insert id [%d]", insertId)

		util.Notification(fmt.Sprintf("%s  [%s]", value.Title, value.InputDate))

	}
}
