package util

import (
	"flag"
	"time"
)

// 全局常量
const (
	VERSION         = "0.0.2"            // 版本
	DEV             = "dev"              // 开发环境
	PROD            = "prod"             // 生产环境
	ApplicationJSON = "application/json" // json header
	ApplicationXML  = "application/xml"  // xml header
	TextHTML        = "text/html"        // text html header
	TextXML         = "text/xml"         // text xml header
)

// 全局变量
var (
	ROOT_DIR      = FileGetParentDirectory(FileGetCurrentDirectory())                  // 根目录
	SRC_DIR       = ROOT_DIR + "/src"                                                  // 源代码目录
	PIDS_DIR      = ROOT_DIR + "/pids"                                                 // pid文件目录
	TEMPLATES_DIR = SRC_DIR + "/views"                                                 // 视图目录
	CURRENT_TIME  = time.Now().String()                                                // 当前时间
	CONFIG_PATH   = flag.String("config", "config/config-dev.ini", "config file path") // 配置文件路径
)
