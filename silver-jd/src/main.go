package main

import (
	"flag"
	"fmt"
	"github.com/go-clog/clog"
	"go-labs/silver-jd/src/jd"
	"os"
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

func init() {
	if err := clog.New(clog.CONSOLE, clog.ConsoleConfig{
		Level:      clog.INFO,
		BufferSize: 100},
	); err != nil {
		fmt.Printf("init console log failed. error %+v.", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	defer clog.Shutdown()

	jd := jd.NewJingDong(jd.JDConfig{
		Period:     time.Millisecond * time.Duration(*period),
		ShipArea:   *area,
		AutoRush:   *rush,
		AutoSubmit: *order,
	})

	defer jd.Release()
	if err := jd.Login(); err == nil {
		// 登录成功
	}
}
