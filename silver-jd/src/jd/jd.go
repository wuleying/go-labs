package jd

import (
	"fmt"
	"github.com/go-clog/clog"
	"go-labs/silver-jd/src/util"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	URLSKUState    = "https://c0.3.cn/stocks"
	URLGoodsDets   = "http://item.jd.com/%s.html"
	URLGoodsPrice  = "http://p.3.cn/prices/mgets"
	URLAdd2Cart    = "https://cart.jd.com/gate.action"
	URLChangeCount = "http://cart.jd.com/changeNum.action"
	URLCartInfo    = "https://cart.jd.com/cart.action"
	URLOrderInfo   = "http://trade.jd.com/shopping/order/getOrderInfo.action"
	URLSubmitOrder = "http://trade.jd.com/shopping/order/submitOrder.action"
	URLSignIn      = "https://vip.jd.com/common/signin.html"
)

var (
	URLForQR = [...]string{
		"https://passport.jd.com/new/login.aspx",
		"https://qr.m.jd.com/show",
		"https://qr.m.jd.com/check",
		"https://passport.jd.com/uc/qrCodeTicketValidation",
		"http://home.jd.com/getUserVerifyRight.action",
	}

	DefaultHeaders = map[string]string{
		"User-Agent":      "Chrome/51.0.2704.103",
		"ContentType":     "application/json",
		"Connection":      "keep-alive",
		"Accept-Encoding": "gzip, deflate",
		"Accept-Language": "zh-CN,zh;q=0.8",
	}

	maxNameLen   = 40
	cookieFile   = fmt.Sprintf("%s/%s", util.ROOT_DIR, "cookies/jd.cookies")
	qrCodeFile   = fmt.Sprintf("%s/%s", util.ROOT_DIR, "cookies/jd.qr")
	strSeperater = strings.Repeat("+", 60)
)

// JDConfig ...
type JDConfig struct {
	Period     time.Duration // refresh period
	ShipArea   string        // shipping area
	AutoRush   bool          // continue rush when out of stock
	AutoSubmit bool          // whether submit the order
}

// SKUInfo ...
type SKUInfo struct {
	ID        string
	Price     string
	Count     int    // buying count
	State     string // stock state 33 : on sale, 34 : out of stock
	StateName string // "现货" / "无货"
	Name      string
	Link      string
}

// JingDong wrap jing dong operation
type JingDong struct {
	JDConfig
	client *http.Client
	jar    *util.SimpleJar
	token  string
}

func NewJingDong(option JDConfig) *JingDong {
	jd := &JingDong{
		JDConfig: option,
	}

	jd.jar = util.NewSimpleJar(util.JarOption{
		JarType:  util.JarGob,
		Filename: cookieFile,
	})

	if err := jd.jar.Load(); err != nil {
		clog.Error(0, "加载Cookies失败: %s", err)
		jd.jar.Clean()
	}

	jd.client = &http.Client{
		Timeout: time.Minute,
		Jar:     jd.jar,
	}

	return jd
}

func (jd *JingDong) Release() {
	if jd.jar != nil {
		if err := jd.jar.Persist(); err != nil {
			log.Panic("Failed to persist cookiejar. error %+v.", err)
		}
	}
}
