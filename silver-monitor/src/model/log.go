package model

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

// 日志结构体
type Log struct {
	Id              int64  `db:"id"`
	PriceBid        string `db:"price_bid"`
	PriceSell       string `db:"price_sell"`
	PriceMiddle     string `db:"price_middle"`
	PriceMiddleHigh string `db:"price_middle_high"`
	PriceMiddleLow  string `db:"price_middle_low"`
	InsertTime      string `db:"insert_time"`
}

type LogReport struct {
	Date  string `db:"date"`
	Price string `db:"price"`
}

// 日志列表
func LogReportList(db *sqlx.DB, start string, end string) []*LogReport {
	logReport := []*LogReport{}
	sql := "SELECT DATE_FORMAT(`insert_time`, \"%Y-%m-%d\") AS `date`, AVG(`price_bid`) AS `price` FROM `log` WHERE `insert_time` >= ? AND `insert_time` <= ? GROUP BY `date` ORDER BY `date`"
	db.Select(&logReport, sql, start, end)
	return logReport
}

// 日志数据落地
func LogSaveData(db *sqlx.DB, prices map[int]string) int64 {
	currentTime := time.Now().Local().Format("2006-01-02 15:04:05")

	logState := `INSERT INTO log (price_bid, price_sell, price_middle, price_middle_high, price_middle_low, insert_time) VALUES (?, ?, ?, ?, ?, ?)`

	id, err := db.MustExec(logState, prices[2], prices[3], prices[4], prices[5], prices[6], currentTime).LastInsertId()

	if err != nil {
		log.Printf("insert log failed. %s", err.Error())
	}

	return id
}
