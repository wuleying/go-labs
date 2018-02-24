package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/go-labs/silver-ipfs/commands"
	"github.com/wuleying/go-labs/silver-ipfs/util"
	"os"
)

func init() {
	if err := clog.New(clog.CONSOLE, clog.ConsoleConfig{
		Level:      clog.TRACE,
		BufferSize: 100,
	}); err != nil {
		fmt.Printf("[INFO] Init console log failed. error %+v.", err)
		os.Exit(1)
	}
}

func main() {
	defer clog.Shutdown()

	err, out := commands.AddFile("/Users/luoliang/Desktop/test.txt")
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	err, out = commands.CatFile("QmXsjqFzpz5e7qC2fkPb12HiMPtj81BXrJBfC5zWkJRPcP")
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	clog.Trace(out)
}
