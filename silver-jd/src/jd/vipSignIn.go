package jd

import (
	"github.com/go-clog/clog"
	"net/http"
	"net/url"
)

// 签到
func (jd *JingDong) VipSignIn() error {
	u, _ := url.Parse(URLSignIn)
	q := u.Query()
	q.Set("token", jd.token)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		clog.Info("Sign in（%+v）failed: %+v", URLSignIn, err.Error())
		return err
	}

	resp, err := jd.client.Do(req)
	if err != nil {
		clog.Info("Sign in（%+v）failed: %+v", URLSignIn, err.Error())
		return err
	}

	if resp.StatusCode == http.StatusOK {
		clog.Info("Sign in（%+v) success.", URLSignIn)
	}

	return nil
}
