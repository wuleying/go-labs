package orm

import (
    "time"
    "log"

    "github.com/server-nado/orm"
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

// 数据落地
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
        log.Print(err.Error())

        return 0
    }

    return id
}

// 获取日志列表
func LogList(data_type string) ([]*Log) {
    logs := []*Log{}
    logModel := new(Log)

    err := logModel.Objects(logModel).All(&logs)

    if err != nil {
        log.Print(err.Error())

        return logs
    }

    return logs
}