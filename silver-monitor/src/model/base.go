package model

import (
    "fmt"
    "strconv"

    "github.com/server-nado/orm"
    _ "github.com/go-sql-driver/mysql"

    "go-labs/silver-monitor/src/util"
)

// 初始化数据库配置项
func Init(config util.Config) {
    orm.NewDatabase("default",
        config.Database["driver"],
        fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s",
            config.Database["user"],
            config.Database["passwd"],
            config.Database["protocol"],
            config.Database["host"],
            config.Database["port"],
            config.Database["dbname"],
            config.Database["charset"]))

    debug, _ := strconv.ParseBool(config.Database["debug"])

    orm.SetDebug(debug)
}