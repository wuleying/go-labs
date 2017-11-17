package model

import (
    "github.com/jmoiron/sqlx"
)

// 日志结构体
type Log struct {
    Id              int64 `db:"id"`
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
func LogReportList(db *sqlx.DB, start string, end string) ([]*LogReport) {
    logReport := []*LogReport{}
    sql := "SELECT DATE_FORMAT(`insert_time`, \"%Y-%m-%d\") AS `date`, AVG(`price_bid`) AS `price` FROM `log` WHERE `insert_time` > ? AND `insert_time` < ? GROUP BY `date` ORDER BY `date`"
    db.Select(&logReport, sql, start, end)
    return logReport
}

/*
// 日志数据落地
func LogSaveData(prices map[int]string) (int64) {
    currentTime := time.Now().Local()

    logModel := new(Log)

    logModel.PriceBid = prices[2]
    logModel.PriceSell = prices[3]
    logModel.PriceMiddle = prices[4]
    logModel.PriceMiddleHigh = prices[5]
    logModel.PriceMiddleLow = prices[6]
    logModel.InsertTime = currentTime.Format("2006-01-02 15:04:05")
    logModel.Objects(logModel).SetTable("log")
    _, id, err := logModel.Objects(logModel).Save()

    if err != nil {
        log.Fatalf("Save log failed. %s", err.Error())
    }

    return id
}

// 日志列表
func LogList() ([]*Log) {
    logModel := new(Log)
    logs := []*Log{}

    model := logModel.Objects(logModel)
    model.SetTable("log")
    model.All(&logs)
    return logs
    //logModel.Objects(logModel).Filter("InsertTime__gt", "2017-10-17 00:00:00").All(&logs)
}
*/