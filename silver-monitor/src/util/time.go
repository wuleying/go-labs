package util

import "time"

var timeLayout string = "2006-01-02 15:04:05"

// 日期转UNIX时间戳
func TimeGetUnixTimestamp(date string) (int64) {
    // 获取时区
    loc, _ := time.LoadLocation("Local")

    theTime, _ := time.ParseInLocation(timeLayout, date, loc)
    return theTime.Unix()
}