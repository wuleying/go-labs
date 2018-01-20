package jd

import (
	"compress/gzip"
	"fmt"
	sjson "github.com/bitly/go-simplejson"
	"github.com/go-clog/clog"
	"go-labs/silver-jd/src/util"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
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
	cookieFile   = "cookies/jd.cookies"
	qrCodeFile   = "cookies/jd.qr"
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

func (jd *JingDong) runCommand(strCmd string) error {
	var err error
	var cmd *exec.Cmd

	// for different platform
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", strCmd)
	case "linux":
		cmd = exec.Command("eog", strCmd)
	default:
		cmd = exec.Command("open", strCmd)
	}

	// just start, do not wait it complete
	if err = cmd.Start(); err != nil {
		if runtime.GOOS == "linux" {
			cmd = exec.Command("gnome-open", strCmd)
			return cmd.Start()
		}
		return err
	}
	return nil
}

func (jd *JingDong) waitForScan(URL string) error {
	var (
		err    error
		req    *http.Request
		resp   *http.Response
		wlfstk string
	)

	for _, c := range jd.jar.Cookies(nil) {
		if c.Name == "wlfstk_smdl" {
			wlfstk = c.Value
			break
		}
	}

	u, _ := url.Parse(URL)
	q := u.Query()
	q.Set("callback", "jQuery123456")
	q.Set("appid", strconv.Itoa(133))
	q.Set("token", wlfstk)
	q.Set("_", strconv.FormatInt(time.Now().Unix()*1000, 10))
	u.RawQuery = q.Encode()

	if req, err = http.NewRequest("GET", u.String(), nil); err != nil {
		clog.Info("请求（%+v）失败: %+v", URL, err)
		return err
	}

	// mush have
	req.Host = "qr.m.jd.com"
	req.Header.Set("Referer", "https://passport.jd.com/new/login.aspx")
	applyCustomHeader(req, DefaultHeaders)

	for retry := 50; retry != 0; retry-- {
		if resp, err = jd.client.Do(req); err != nil {
			clog.Info("二维码失效：%+v", err)
			break
		}

		if resp.StatusCode == http.StatusOK {
			respMsg := string(responseData(resp))
			resp.Body.Close()

			n1 := strings.Index(respMsg, "(")
			n2 := strings.Index(respMsg, ")")

			var js *sjson.Json
			if js, err = sjson.NewJson([]byte(respMsg[n1+1 : n2])); err != nil {
				clog.Error(0, "解析响应数据失败: %+v", err)
				clog.Trace("Response data  : %+v", respMsg)
				clog.Trace("Response Header: %+v", resp.Header)
				break
			}

			code := js.Get("code").MustInt()
			if code == 200 {
				jd.token = js.Get("ticket").MustString()
				clog.Info("token : %+v", jd.token)
				break
			} else {
				clog.Info("%+v : %s", code, js.Get("msg").MustString())
				time.Sleep(time.Second * 3)
			}
		} else {
			resp.Body.Close()
		}
	}

	if jd.token == "" {
		err = fmt.Errorf("未检测到QR扫码结果")
		return err
	}

	return nil
}

func applyCustomHeader(req *http.Request, header map[string]string) {
	if req == nil || len(header) == 0 {
		return
	}

	for key, val := range header {
		req.Header.Set(key, val)
	}
}

func responseData(resp *http.Response) []byte {
	if resp == nil {
		return nil
	}

	var reader io.Reader
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
	default:
		reader = resp.Body
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		clog.Error(0, "读取响应数据失败: %+v", err)
		return nil
	}

	return data
}
