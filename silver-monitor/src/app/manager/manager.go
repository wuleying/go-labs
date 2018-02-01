package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/wuleying/go-labs/silver-monitor/src/model"
	"github.com/wuleying/go-labs/silver-monitor/src/util"
)

// 首页数据结构体
type HomeData struct {
	LogData []*model.LogReport
}

// 全局配置
var config util.Config
var err error
var db *sqlx.DB

func main() {
	util.FileSavePid("silver-monitor-manager.pid")

	config, err = util.ConfigInit()
	if err != nil {
		log.Fatal("Init config failed: ", err.Error())
	}

	// 初始化模型
	db = model.Init(config)

	http.HandleFunc("/", HomeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(fmt.Sprintf("%s/%s", util.ROOT_DIR, "src/static")))))

	err := http.ListenAndServe(fmt.Sprintf(":%s", config.Manager["port"]), nil)
	if err != nil {
		log.Fatal("Listen and serve failed: ", err.Error())
	}
}

// 首页
func HomeHandler(response http.ResponseWriter, request *http.Request) {
	var data HomeData
	template, err := template.ParseFiles(util.TEMPLATES_DIR + "/manager/home.html")

	if err != nil {
		log.Fatal("Load template failed: ", err.Error())
		return
	}

	currentTime := time.Now().Unix()
	startTime := currentTime - (100 * 86400)

	start := time.Unix(startTime, 0).Format("2006-01-02")
	end := time.Unix(currentTime, 0).Format("2006-01-02")

	// 解析请求参数
	request.ParseForm()
	start_param := request.Form.Get("start")
	end_param := request.Form.Get("end")

	if len(start_param) > 0 {
		start = start_param
	}

	if len(end_param) > 0 {
		end = end_param
	}

	data.LogData = model.LogReportList(db, start, end)

	template.Execute(response, data)
}
