package utils

import (
	"time"
)

// 全局变量
var (
	ROOT_DIR     = FileGetParentDirectory(FileGetCurrentDirectory()) // 根目录
	SRC_DIR      = ROOT_DIR + "/src"                                 // 源代码目录
	CURRENT_TIME = time.Now().String()                               // 当前时间
)
