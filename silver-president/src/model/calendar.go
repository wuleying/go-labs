package model

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/wuleying/go-labs/silver-president/src/util"
)

// 日历数据结构体
type Calendar struct {
	Id         int64  `db:"id"`
	Title      string `db:"title"`
	Summary    string `db:"summary"`
	Url        string `db:"url"`
	OriginUrl  string `db:"origin_url"`
	OriginName string `db:"origin_name"`
	ImageUrl   string `db:"image_url"`
	ImageTitle string `db:"image_title"`
	InputDate  string `db:"input_date"`
	InsertDate string `db:"insert_date"`
	InsertTime string `db:"insert_time"`
}

// 日志数据落地
func CalendarSaveData(db *sqlx.DB, jsonData util.JsonData) int64 {
	currentDate := time.Now().Local().Format("2006-01-02")
	currentTime := time.Now().Local().Format("2006-01-02 15:04:05")

	calendarState := `INSERT INTO calendar (title, summary, url, origin_url, origin_name,
	                    image_url, image_title, input_date, insert_date, insert_time)
	                    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	id, err := db.MustExec(calendarState, jsonData.Title, jsonData.Summary,
		jsonData.Url, jsonData.OriginUrl, jsonData.OriginName, jsonData.ImageUrl, jsonData.ImageTitle,
		jsonData.InputDate, currentDate, currentTime).LastInsertId()

	if err != nil {
		log.Printf("Insert data failed. %s", err.Error())
	}

	return id
}

// 获取某天全部数据
func CalendarGetAll(db *sqlx.DB, date string) []*Calendar {
	lists := []*Calendar{}

	sql := "SELECT * FROM `calendar` WHERE insert_date = ? ORDER BY `insert_time` ASC"
	db.Select(&lists, sql, date)
	return lists
}
