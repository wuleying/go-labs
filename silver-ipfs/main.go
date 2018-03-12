package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/go-labs/silver-ipfs/handlers"
	"github.com/wuleying/go-labs/silver-ipfs/utils"
	"net/http"
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

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/upload", handlers.Upload)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(fmt.Sprintf("%s/%s", utils.ROOT_DIR, "static")))))

	err := http.ListenAndServe(fmt.Sprintf(":%s", "10099"), nil)
	if err != nil {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}
}
