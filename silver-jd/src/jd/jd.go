package jd

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	sjson "github.com/bitly/go-simplejson"
	"github.com/go-clog/clog"
	"go-labs/silver-jd/src/util"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
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

func (jd *JingDong) Login(args ...interface{}) error {
	clog.Info(strSeperater)

	if jd.validateLogin(URLForQR[4]) {
		clog.Info("无需重新登录")
		return nil
	}

	var (
		err   error
		qrImg string
	)

	clog.Info("请打开京东手机客户端，准备扫码登陆:")
	jd.jar.Clean()

	if err = jd.loginPage(URLForQR[0]); err != nil {
		return err
	}

	if qrImg, err = jd.loadQRCode(URLForQR[1]); err != nil {
		return err
	}

	// just start, do not wait it complete
	if err = jd.runCommand(qrImg); err != nil {
		clog.Info("打开二维码图片失败: %+v.", err)
		return err
	}

	if err = jd.waitForScan(URLForQR[2]); err != nil {
		return err
	}

	if err = jd.validateQRToken(URLForQR[3]); err != nil {
		return err
	}

	return nil
}

func (jd *JingDong) validateLogin(URL string) bool {
	var (
		err  error
		req  *http.Request
		resp *http.Response
	)

	if req, err = http.NewRequest("GET", URL, nil); err != nil {
		clog.Info("请求（%+v）失败: %+v", URL, err)
		return false
	}

	jd.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		// disable redirect
		return http.ErrUseLastResponse
	}

	defer func() {
		// restore to default
		jd.client.CheckRedirect = nil
	}()

	if resp, err = jd.client.Do(req); err != nil {
		clog.Info("需要重新登录: %+v", err)
		return false
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		clog.Info("需要重新登录")
		return false
	}

	clog.Trace("Response Data: %s", string(data))
	return true
}

func (jd *JingDong) loginPage(URL string) error {
	var (
		err  error
		req  *http.Request
		resp *http.Response
	)

	if req, err = http.NewRequest("GET", URL, nil); err != nil {
		clog.Info("请求（%+v）失败: %+v", URL, err)
		return err
	}

	applyCustomHeader(req, DefaultHeaders)

	if resp, err = jd.client.Do(req); err != nil {
		clog.Info("请求登录页失败: %+v", err)
		return err
	}

	defer resp.Body.Close()
	return nil
}

func (jd *JingDong) loadQRCode(URL string) (string, error) {
	var (
		err  error
		req  *http.Request
		resp *http.Response
	)

	u, _ := url.Parse(URL)
	q := u.Query()
	q.Set("appid", strconv.Itoa(133))
	q.Set("size", strconv.Itoa(147))
	q.Set("t", strconv.FormatInt(time.Now().Unix()*1000, 10))
	u.RawQuery = q.Encode()

	if req, err = http.NewRequest("GET", u.String(), nil); err != nil {
		clog.Error(0, "请求（%+v）失败: %+v", URL, err)
		return "", err
	}

	applyCustomHeader(req, DefaultHeaders)
	if resp, err = jd.client.Do(req); err != nil {
		clog.Error(0, "下载二维码失败: %+v", err)
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		clog.Error(0, "http status : %d/%s", resp.StatusCode, resp.Status)
	}

	// from mime get QRCode image type
	//  content-type:image/png
	//
	filename := qrCodeFile + ".png"
	mt, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if typ, e := mime.ExtensionsByType(mt); e == nil {
		filename = qrCodeFile + typ[0]
	}

	dir, _ := os.Getwd()
	filename = filepath.Join(dir, filename)
	clog.Trace("QR Image: %s", filename)

	file, _ := os.Create(filename)
	defer file.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		clog.Error(0, "下载二维码失败: %+v", err)
		return "", err
	}

	return filename, nil
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

func (jd *JingDong) validateQRToken(URL string) error {
	var (
		err  error
		req  *http.Request
		resp *http.Response
	)

	u, _ := url.Parse(URL)
	q := u.Query()
	q.Set("t", jd.token)
	u.RawQuery = q.Encode()

	if req, err = http.NewRequest("GET", u.String(), nil); err != nil {
		clog.Info("请求（%+v）失败: %+v", URL, err)
		return err
	}
	if resp, err = jd.client.Do(req); err != nil {
		clog.Error(0, "二维码登陆校验失败: %+v", err)
		return nil
	}

	//
	// 京东有时候会认为当前登录有危险，需要手动验证
	// url: https://safe.jd.com/dangerousVerify/index.action?username=...
	//
	if resp.Header.Get("P3P") == "" {
		var res struct {
			ReturnCode int    `json:"returnCode"`
			Token      string `json:"token"`
			URL        string `json:"url"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&res); err == nil {
			if res.URL != "" {
				verifyURL := res.URL
				if !strings.HasPrefix(verifyURL, "https:") {
					verifyURL = "https:" + verifyURL
				}
				clog.Error(2, "安全验证: %s", verifyURL)
				jd.runCommand(verifyURL)
			}
		}
		return fmt.Errorf("login failed")
	}

	if resp.StatusCode == http.StatusOK {
		clog.Info("登陆成功, P3P: %s", resp.Header.Get("P3P"))
	} else {
		clog.Info("登陆失败")
		err = fmt.Errorf("%+v", resp.Status)
	}

	resp.Body.Close()
	return nil
}

// 签到
func (jd *JingDong) VipSignIn() error {
	u, _ := url.Parse(URLSignIn)
	q := u.Query()
	q.Set("token", jd.token)
	u.RawQuery = q.Encode()

	if _, err := http.NewRequest("GET", u.String(), nil); err != nil {
		clog.Info("签到（%+v）失败: %+v", URLSignIn, err)
		return err
	}

	clog.Info("签到成功.")

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
		//clog.Trace("Encoding: %+v", resp.Header.Get("Content-Encoding"))
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
