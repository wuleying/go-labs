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

type JsonData struct {
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Url        string `json:"url"`
	OriginUrl  string `json:"origin_url"`
	OriginName string `json:"origin_name"`
	ImageUrl   string `json:"image_url"`
	ImageTitle string `json:"image_title"`
	InputDate  string `json:"input_date"`
}

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

	getData()
}

// 获取数据
func getData() {
	// 抓取目标页
	var target_url string = fmt.Sprintf(config.Setting["target_url"], time.Now().Local().Format("20060102"))

	resp, err := http.Get(target_url)

	if err != nil {
		log.Fatal("Init config failed: ", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var jsonDatas []JsonData

	err = json.Unmarshal(body, &jsonDatas)
	if err != nil {
		fmt.Println("error:", err)
	}

	util.Notification(jsonDatas[0].Title)
}
