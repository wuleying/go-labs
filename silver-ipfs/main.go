package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/go-labs/silver-ipfs/utils"
	"html/template"
	"io"
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

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(fmt.Sprintf("%s/%s", utils.ROOT_DIR, "static")))))

	err := http.ListenAndServe(fmt.Sprintf(":%s", "10099"), nil)
	if err != nil {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}
}

// 首页
func homeHandler(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles(utils.TEMPLATES_DIR + "/home.html")

	if err != nil {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, err.Error())
		return
	}

	template.Execute(response, nil)
}

// 上传
func uploadHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	if request.Method == "GET" {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, "Must post method.")
	}

	file, handle, err := request.FormFile("file")

	if err != nil {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, "Get file info failure.")
	}

	f, err := os.OpenFile(utils.ROOT_DIR+"/files/"+handle.Filename, os.O_WRONLY|os.O_CREATE, utils.FILE_WRITE_MODE)
	io.Copy(f, file)

	if err != nil {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, "Copy file failure.")
	}

	defer f.Close()
	defer file.Close()

	fmt.Println("upload success")
}
