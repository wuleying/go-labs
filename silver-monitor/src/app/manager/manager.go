package main

import (
    "log"
    "fmt"
    "net/http"
    "html/template"
    "time"
    "github.com/jmoiron/sqlx"

    "go-labs/silver-monitor/src/util"
    "go-labs/silver-monitor/src/model"
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
    util.SavePid("silver-monitor-manager.pid");

    config, err = util.InitConfig();
    if err != nil {
        log.Fatal("Init config failed: ", err.Error())
    }

    // 初始化模型
    db = model.Init(config)

    http.HandleFunc("/", HomeHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/static"))))

    err := http.ListenAndServe(fmt.Sprintf(":%s", config.Manager["port"]), nil)
    if err != nil {
        log.Fatal("Listen and serve failed: ", err.Error())
    }
}

// 首页
func HomeHandler(response http.ResponseWriter, request *http.Request) {
    var data HomeData;
    var start string
    var end string
    template, err := template.ParseFiles(util.TEMPLATES_DIR + "/manager/home.html")

    if err != nil {
        log.Fatal("Load template failed: ", err.Error())
        return
    }

    currentTime := time.Now().Unix()
    startTime := currentTime - (100 * 86400)

    // 解析请求参数
    request.ParseForm();
    start_param := request.Form.Get("start")
    end_param := request.Form.Get("end")

    if len(start_param) > 0 {
        start = start_param
    } else {
        start = time.Unix(startTime, 0).Format("2006-01-02")
    }

    if len(end_param) > 0 {
        end = end_param
    } else {
        end = time.Unix(currentTime, 0).Format("2006-01-02")
    }

    data.LogData = model.LogReportList(db, start, end)

    template.Execute(response, data)
}