package handlers

import (
	"github.com/go-clog/clog"
	"github.com/wuleying/go-labs/silver-ipfs/utils"
	"html/template"
	"net/http"
)

// 首页
func Home(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles(utils.TEMPLATES_DIR + "/home.html")

	if err != nil {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, err.Error())
		return
	}

	template.Execute(response, nil)
}
