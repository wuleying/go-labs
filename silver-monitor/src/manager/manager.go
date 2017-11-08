package main

import (
    "io"
    "net/http"
    "log"
    "flag"
    "fmt"

    "go-labs/silver-monitor/src/common"
)

// 全局配置
var config common.Config

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
}

func TestServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, test!\n")
}

func main() {
    common.SavePid("./pid/silver-monitor-manager.pid");

    // 命令行参数，配置文件路径
    var config_path = flag.String("config", "config/config.ini", "config file path")
    config, _ = common.InitConfig(*config_path);

    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/test", TestServer)

    err := http.ListenAndServe(fmt.Sprintf(":%s", config.Manager["port"]), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}