package main

import (
    "log"
    "io"
    "fmt"
    "net/http"
    "html/template"

    "go-labs/silver-monitor/src/common"
)

// 全局配置
var config common.Config

// 首页
func HomeHandler(response http.ResponseWriter, request *http.Request) {
    template, err := template.ParseFiles(common.TEMPLATES_DIR + "/manager/home.html")

    if err != nil {
        log.Fatal("Load template failed: ", err.Error())
        return
    }

    template.Execute(response, "Hello world")

    io.WriteString(response, common.TEMPLATES_DIR + "/manager/home.html")
}

// 数据
func DataHandler(response http.ResponseWriter, request *http.Request) {
    io.WriteString(response, "hello, data!\n")
}

func main() {
    common.SavePid("./pid/silver-monitor-manager.pid");

    config, _ = common.InitConfig();

    http.HandleFunc("/", HomeHandler)
    http.HandleFunc("/data", DataHandler)

    err := http.ListenAndServe(fmt.Sprintf(":%s", config.Manager["port"]), nil)
    if err != nil {
        log.Fatal("Listen and serve failed: ", err.Error())
    }
}