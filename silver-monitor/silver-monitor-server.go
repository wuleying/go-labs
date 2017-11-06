package main

import (
    "log"
    "strings"
    "time"
    "net/smtp"
    "os"
    "strconv"
    "bytes"
    "fmt"
    "flag"

    "github.com/robfig/cron"
    "github.com/PuerkitoBio/goquery"
    "github.com/server-nado/orm"
    "github.com/Unknwon/goconfig"
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

type Config struct {
    setting  map[string]string
    database map[string]string
    email    map[string]string
}

// 全局配置项
var config Config
// 命令行参数，配置文件路径
var config_path = flag.String("config", "config.ini", "config file path")

func main() {
    getPid()

    initConfig();

    crontab := cron.New()

    crontab.AddFunc(config.setting["schedule"], func() {
        getPrice();
    })

    crontab.Start()

    select {}
}

// 初始化配置
func initConfig() {
    flag.Parse()
    goconfig, err := goconfig.LoadConfigFile(*config_path)

    if err != nil {
        log.Printf("Read config file failed: %s", err)
        return
    }

    log.Printf("Load config file success: %s", *config_path)

    config.setting, _ = goconfig.GetSection("setting")
    config.database, _ = goconfig.GetSection("database")
    config.email, _ = goconfig.GetSection("email")
}

// 获取实时银价
func getPrice() {
    // 抓取目标页
    var url string = config.setting["tendency_picture_url"]

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

    alert_price, _ := strconv.ParseFloat(config.setting["alert_price"], 64)

    if (price <= alert_price) {
        go func() {
            email_err := sendMail(config.email["user"], config.email["passwd"],
                fmt.Sprintf("%s:%s", config.email["host"], config.email["port"]),
                config.email["to"], config.email["subject"], fmt.Sprintf("当前价格:%f", price))

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
    orm.NewDatabase("default", config.database["driver"],
        fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s",
        config.database["user"],
        config.database["passwd"],
        config.database["protocol"],
        config.database["host"],
        config.database["port"],
        config.database["name"],
        config.database["charset"]))
    orm.SetDebug(false)

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
        log.Print(err)
    }

    return id
}

// 获取pid
func getPid() {
    file, err := os.OpenFile("./pid", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)

    if err != nil {
        log.Fatal("open file failed.", err.Error())
    }

    defer file.Close()

    pid := os.Getpid()

    log.Printf("pid:%d", pid)

    file.WriteString(strconv.Itoa(pid))
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