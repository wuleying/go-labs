package util

import (
	"time"
)

// 全局变量
var (
	ROOT_DIR     = FileGetParentDirectory(FileGetCurrentDirectory()) // 根目录
	CURRENT_TIME = time.Now().String()                               // 当前时间
)
