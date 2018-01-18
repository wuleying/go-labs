package main

import (
	"flag"
	"github.com/go-clog/clog"
	"go-labs/silver-jd/src/jd"
	"log"
	"time"
)

const (
	AreaBeijing = "1_72_2799_0"
)

var (
	area   = flag.String("area", AreaBeijing, "ship location string, default to Beijing")
	period = flag.Int("period", 500, "the refresh period when out of stock, unit: ms.")
	rush   = flag.Bool("rush", false, "continue to refresh when out of stock.")
	order  = flag.Bool("order", false, "submit the order to JingDong when get the Goods.")
)

func main() {
	jd := jd.NewJingDong(jd.JDConfig{
		Period:     time.Millisecond * time.Duration(*period),
		ShipArea:   *area,
		AutoRush:   *rush,
		AutoSubmit: *order,
	})

	clog.Info("hehe")

	defer jd.Release()
	if err := jd.Login(); err == nil {
		// 登录成功
		log.Println("Success.")
	}
}
