package main

import (
    "io"
    "net/http"
    "log"
    "flag"
    "github.com/Unknwon/goconfig"
    "strconv"
    "os"
    "fmt"
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
    getPid();

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

// 获取pid
func getPid() {
    file, err := os.OpenFile("./pid/silver-monitor-manager.pid", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)

    if err != nil {
        log.Fatal("open file failed.", err.Error())
    }

    defer file.Close()

    pid := os.Getpid()

    log.Printf("pid:%d", pid)

    file.WriteString(strconv.Itoa(pid))
}