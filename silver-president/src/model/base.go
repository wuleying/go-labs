package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/wuleying/go-labs/silver-president/src/util"
	"log"
)

// 初始化数据库配置项
func Init(config util.Config) (db *sqlx.DB) {
	db, err := sqlx.Connect(config.Database["driver"],
		fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s",
			config.Database["user"],
			config.Database["passwd"],
			config.Database["protocol"],
			config.Database["host"],
			config.Database["port"],
			config.Database["dbname"],
			config.Database["charset"]))

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
