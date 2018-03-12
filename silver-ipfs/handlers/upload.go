package handlers

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/go-labs/silver-ipfs/ipfs"
	"github.com/wuleying/go-labs/silver-ipfs/utils"
	"io"
	"net/http"
	"os"
	"path"
)

// 上传
func Upload(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	if request.Method == "GET" {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, "Must post method.")
	}

	file, handle, err := request.FormFile("file")

	if err != nil {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, "Get file info failure.")
	}

	tmpFile := utils.ROOT_DIR + "/files/" + handle.Filename

	fileHandle, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE, utils.FILE_WRITE_MODE)
	io.Copy(fileHandle, file)

	if err != nil {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, "Copy file failure.")
	}

	defer file.Close()
	defer fileHandle.Close()

	fileHash, err := ipfs.AddObject(tmpFile)

	if err != nil {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, "Add IPFS failure.")
	}

	// 获取文件扩展名
	fileSuffix := path.Ext(handle.Filename)

	// todo 检查文件扩展名

	newFile := utils.ROOT_DIR + "/files/" + fileHash + fileSuffix
	err = os.Rename(tmpFile, newFile)

	if err != nil {
		clog.Fatal(utils.CLOG_SKIP_DISPLAY_INFO, "Move file failure.")
	}

	fmt.Println(newFile)
}
