package util

import (
	"time"
)

// 全局常量
const (
	VERSION = "0.0.1"

	// 权限
	FILE_READ_MODE  = 0644
	FILE_WRITE_MODE = 0666
	DIR_READ_MODE   = 0755
	DIR_WRITE_MODE  = 0777

	// clog skip级别
	CLOG_SKIP_DEFAULT      = 0
	CLOG_SKIP_DISPLAY_INFO = 2
)

// 全局变量
var (
	// 根目录
	ROOT_DIR = FileGetParentDirectory(FileGetCurrentDirectory())
	// 当前时间
	CURRENT_TIME = time.Now().String()
)