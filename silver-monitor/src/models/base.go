package models

import (
    "fmt"

    "github.com/server-nado/orm"
    "go-labs/silver-monitor/src/utils"
    "strconv"
)

// 初始化数据库配置项
func InitModel(config utils.Config) {
    orm.NewDatabase("default",
        config.Database["driver"],
        fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s",
            config.Database["user"],
            config.Database["passwd"],
            config.Database["protocol"],
            config.Database["host"],
            config.Database["port"],
            config.Database["name"],
            config.Database["charset"]))

    debug, _ := strconv.ParseBool(config.Database["debug"])

    orm.SetDebug(debug)
}
