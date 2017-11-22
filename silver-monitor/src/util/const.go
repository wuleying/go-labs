package util

import "flag"

// 全局常量
const (
	// 系统版本
	VERSION = "0.0.2"

	DEV  = "dev"
	PROD = "prod"

	// http头
	ApplicationJSON = "application/json"
	ApplicationXML  = "application/xml"
	TextHTML        = "text/html"
	TextXML         = "text/xml"
)

// 全局变量
var (
	// 项目根目录
	ROOT_DIR = FileGetParentDirectory(FileGetCurrentDirectory())
	// 源代码目录
	SRC_DIR = ROOT_DIR + "/src"
	// PID文件目录
	PIDS_DIR = ROOT_DIR + "/pids"

	// 模块目录
	MODULES_DIR = SRC_DIR + "/modules"
	// 模板目录
	TEMPLATES_DIR = SRC_DIR + "/views"

	// 命令行参数，配置文件路径
	CONFIG_PATH = flag.String("config", "config/config-dev.ini", "config file path")
)
