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

	object := ipfs.NewObject()
	object.Save("/Users/luoliang/Desktop/test.txt")

	clog.Info("%x", object)
}
