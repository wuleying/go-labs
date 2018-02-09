package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"os"
)

func init() {
	consoleConfig := clog.ConsoleConfig{
		Level:      clog.INFO,
		BufferSize: 100,
	}

	if err := clog.New(clog.CONSOLE, consoleConfig); err != nil {
		fmt.Printf("Init console log failed. error %+v.", err)
		os.Exit(1)
	}
}

func main() {
	defer clog.Shutdown()

	clog.Info("hello world.")
}
