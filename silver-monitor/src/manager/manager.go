package main

import (
    "log"
    "flag"
    "io"
    "fmt"
    "net/http"

    "go-labs/silver-monitor/src/common"
)

// 全局配置
var config common.Config

func HomeServer(response http.ResponseWriter, request *http.Request) {
    io.WriteString(response, "hello, world!\n")
}

func DataServer(response http.ResponseWriter, request *http.Request) {
    io.WriteString(response, "hello, data!\n")
}

func main() {
    common.SavePid("./pid/silver-monitor-manager.pid");

    // 命令行参数，配置文件路径
    config_path := flag.String("config", "config/config.ini", "config file path")

    flag.Parse()
    config, _ = common.InitConfig(*config_path);

    http.HandleFunc("/", HomeServer)
    http.HandleFunc("/data", DataServer)

    err := http.ListenAndServe(fmt.Sprintf(":%s", config.Manager["port"]), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}