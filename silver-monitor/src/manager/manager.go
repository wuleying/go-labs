package main

import (
    "io"
    "net/http"
    "log"
    "flag"
    "fmt"

    "github.com/Unknwon/goconfig"
    "go-labs/silver-monitor/src/common"
)

type Config struct {
    database map[string]string
    manager  map[string]string
}

// 全局配置项
var config Config
// 命令行参数，配置文件路径
var config_path = flag.String("config", "config/config.ini", "config file path")

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
}

func TestServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, test!\n")
}

func main() {
    common.SavePid("./pid/silver-monitor-server.pid");

    initConfig();

    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/test", TestServer)

    err := http.ListenAndServe(fmt.Sprintf(":%s", config.manager["port"]), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

// 初始化配置
func initConfig() {
    flag.Parse()
    goconfig, err := goconfig.LoadConfigFile(*config_path)

    if err != nil {
        log.Printf("Read config file failed: %s", err)
        return
    }

    log.Printf("Load config file success: %s", *config_path)

    config.database, _ = goconfig.GetSection("database")
    config.manager, _ = goconfig.GetSection("manager")
}

