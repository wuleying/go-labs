package util

import (
	"log"
	"os"
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
