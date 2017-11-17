package model

import (
    "fmt"
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "go-labs/silver-monitor/src/util"
)

// 初始化数据库配置项
func Init(config util.Config) (*gorm.DB) {
    connectParams := fmt.Sprintf("%s:%s@/%s?charset=%s",
        config.Database["user"],
        config.Database["passwd"],
        config.Database["dbname"],
        config.Database["charset"])

    db, err := gorm.Open(config.Database["driver"], connectParams)

    if err != nil {
        log.Printf("Connect mysql failed. %s", connectParams)
    }

    return db
}
