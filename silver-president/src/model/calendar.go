package model

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

// 日志结构体
type Log struct {
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
func calendarSaveData(db *sqlx.DB, title, summary, url, origin_url, origin_name, image_url, image_title, input_date string) int64 {
	currentDate := time.Now().Local().Format("2006-01-02")
	currentTime := time.Now().Local().Format("2006-01-02 15:04:05")

	calendarState := `INSERT INTO log (title, summary, url, origin_url, origin_name,
	                    image_url, image_title, input_date, insert_date, insert_time)
	                    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	id, err := db.MustExec(calendarState, title, summary,
		url, origin_url, origin_name, image_url, image_title,
		input_date, currentDate, currentTime).LastInsertId()

	if err != nil {
		log.Printf("Insert data failed. %s", err.Error())
	}

	return id
}
