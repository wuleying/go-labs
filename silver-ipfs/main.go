package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/go-labs/silver-ipfs/util"
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

	command := "ipfs"
	params := []string{"add", "~/Desktop/1.png"}
	err, result := util.ExecCommand(command, params)

	if err != nil {
		clog.Fatal(2, err.Error())
	}

	clog.Info(result)
}
