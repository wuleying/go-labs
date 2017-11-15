package util

import (
    "flag"
    "log"

    "github.com/Unknwon/goconfig"
)

type Config struct {
    Setting  map[string]string
    Database map[string]string
    Email    map[string]string
    Manager  map[string]string
}

// 初始化配置
func InitConfig() (Config, error) {
    var config Config

    flag.Parse()

    goconfig, err := goconfig.LoadConfigFile(*CONFIG_PATH)

    if err != nil {
        log.Printf("Read config file failed: %s", err)
        return config, err
    }

    log.Printf("Load config file success: %s", *CONFIG_PATH)

    config.Setting, _ = goconfig.GetSection("setting")
    config.Database, _ = goconfig.GetSection("database")
    config.Email, _ = goconfig.GetSection("email")
    config.Manager, _ = goconfig.GetSection("manager")

    return config, nil
}