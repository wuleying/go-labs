package main

import (
    "log"
    "fmt"
    "net/http"
    "html/template"

    "go-labs/silver-monitor/src/util"
)

// 全局配置
var config util.Config

// 首页数据结构体
type HomeData struct {
    Type string
}

func main() {
    util.SavePid("silver-monitor-manager.pid");

    config, _ = util.InitConfig();

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

    template, err := template.ParseFiles(util.TEMPLATES_DIR + "/manager/home.html")

    if err != nil {
        log.Fatal("Load template failed: ", err.Error())
        return
    }

    // 解析请求参数
    request.ParseForm();

    data.Type = request.Form.Get("type")

    template.Execute(response, data)
}