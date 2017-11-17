package model

import (
    "time"
    "fmt"
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

func LogList() {
    logModel := new(Log)

    logs := []*Log{}

    logModel.Objects(logModel).SetTable("log")

    if err := logModel.Objects(logModel).Filter("Id__lt", 10).All(&logs); err == nil {
        for _, u := range logs {
            fmt.Println(u.Id, u.PriceSell)
        }
    }
}