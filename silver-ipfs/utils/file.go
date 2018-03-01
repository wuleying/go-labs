package utils

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 获取当前目录
func FileGetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("Get current directory failed.", err.Error())
	}

	return strings.Replace(dir, "\\", "/", -1)
}

// 获取上级目录
func FileGetParentDirectory(dirctory string) string {
	return StringSubstr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

// 获取文件名称
func FileGetName(filePath string) string {
	return path.Base(filePath)
}

// 藜取文件大小
func FileGetSize(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(CLOG_SKIP_DISPLAY_INFO, err)
	}

	fileSize := fileInfo.Size()
	return fileSize
}
