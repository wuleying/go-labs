package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/go-labs/silver-ipfs/ipfs"
	"os"
)

func init() {
	if err := clog.New(clog.CONSOLE, clog.ConsoleConfig{
		Level:      clog.INFO,
		BufferSize: 100,
	}); err != nil {
		fmt.Printf("[INFO] Init console log failed. error %+v.", err)
		os.Exit(1)
	}
}

func main() {
	defer clog.Shutdown()

	/*
		fileHash, err := ipfs.AddObject("/Users/luoliang/Desktop/test.txt")
		if err != nil {
			clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
		}

		clog.Info("fileHash = %s", fileHash)
	*/

	// QmXsjqFzpz5e7qC2fkPb12HiMPtj81BXrJBfC5zWkJRPcP
	object, _ := ipfs.GetObject("QmXsjqFzpz5e7qC2fkPb12HiMPtj81BXrJBfC5zWkJRPcP")

	clog.Info("fileHash = %s", object.FileHash)
}
