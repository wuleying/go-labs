package main

import (
    "log"
    "strings"
    "time"
    "net/smtp"
    "strconv"
    "bytes"
    "fmt"

    "github.com/robfig/cron"
    "github.com/PuerkitoBio/goquery"
    "github.com/server-nado/orm"
    _ "github.com/go-sql-driver/mysql"

    "go-labs/silver-monitor/src/common"
    "go-labs/silver-monitor/src/models"
)

// 全局配置
var config common.Config

func main() {
    // 保存pid
    common.SavePid("silver-monitor-server.pid")

    config, _ = common.InitConfig();

    crontab := cron.New()

    crontab.AddFunc(config.Setting["schedule"], func() {
        getPrice();
    })

    crontab.Start()

    select {}
}

// 获取实时银价
func getPrice() {
    // 抓取目标页
    var url string = config.Setting["tendency_url"]

    doc, err := goquery.NewDocument(url)

    if err != nil {
        log.Print(err)
        return
    }

    selection := doc.Find("#TABLE1 tbody tr").Eq(2).Find("td")

    var prices = make(map[int]string)

    selection.Each(func(i int, tag *goquery.Selection) {
        if (i >= 2 && i <= 6) {
            prices[i] = strings.TrimSpace(tag.Text())
        }
    })

    id := saveData(prices)

    if (prices[2] == "") {
        log.Print("get price failed")
        return
    }

    price, _ := strconv.ParseFloat(prices[2], 64)

    alert_price, _ := strconv.ParseFloat(config.Setting["alert_price"], 64)

    if (price <= alert_price) {
        go func() {
            email_err := sendMail(
                config.Email["user"],
                config.Email["passwd"],
                fmt.Sprintf("%s:%s", config.Email["host"], config.Email["port"]),
                config.Email["to"],
                config.Email["subject"],
                fmt.Sprintf("当前价格:%f", price))

            if email_err != nil {
                log.Printf("Send email failed: %s. Current price:%f", email_err, price)
                return
            }
        }()
    }

    log.Printf("monitor id:%d", id)
}

// 数据落地
func saveData(prices map[int]string) (int64) {
    orm.NewDatabase("default",
            config.Database["driver"],
            fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s",
            config.Database["user"],
            config.Database["passwd"],
            config.Database["protocol"],
            config.Database["host"],
            config.Database["port"],
            config.Database["name"],
            config.Database["charset"]))
    orm.SetDebug(false)

    currentTime := time.Now().Local()

    logModel := new(models.Log)

    logModel.PriceBid = prices[2]
    logModel.PriceSell = prices[3]
    logModel.PriceMiddle = prices[4]
    logModel.PriceMiddleHigh = prices[5]
    logModel.PriceMiddleLow = prices[6]
    logModel.InsertTime = currentTime.Format("2006-01-02 15:04:05")
    logModel.Objects(logModel).SetTable("log")
    _, id, err := logModel.Objects(logModel).Save()

    if err != nil {
        log.Print(err)
    }

    return id
}

// 发送邮件
func sendMail(user string, password string, host string, to string, subject string, body string) error {
    hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", user, password, hp[0])
    send_to := strings.Split(to, ";")

    buffer := bytes.NewBuffer(nil)

    boudary := "SILVER_MONITOR"

    msg := fmt.Sprintf("To:%s\r\n" +
    "From:%s\r\n" +
    "Subject:%s\r\n" +
    "Content-Type:multipart/mixed;Boundary=\"%s\"\r\n" +
    "Mime-Version:1.0\r\n" +
    "Date:%s\r\n" +
    "\r\n\r\n--%s\r\n" +
    "Content-Type:text/plain;charset=utf-8\r\n\r\n%s\r\n",
        to, user, subject, boudary, time.Now().String(), boudary, body)

    buffer.WriteString(msg)

    err := smtp.SendMail(host, auth, user, send_to, buffer.Bytes())

    return err
}