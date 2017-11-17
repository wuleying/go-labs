package model

import (
    "time"
    "log"

    "github.com/server-nado/orm"
    _ "github.com/go-sql-driver/mysql"
)

// 日志结构体
type Log struct {
    orm.DBHook
    Id              int64 `field:"id" auto:"true" index:"pk"`
    PriceBid        string `field:"price_bid"`
    PriceSell       string `field:"price_sell"`
    PriceMiddle     string `field:"price_middle"`
    PriceMiddleHigh string `field:"price_middle_high"`
    PriceMiddleLow  string `field:"price_middle_low"`
    InsertTime      string `field:"insert_time"`
}

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