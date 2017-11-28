package util

import (
	"flag"
	"time"
)

// 全局常量
const (
	VERSION         = "0.0.2"
	DEV             = "dev"
	PROD            = "prod"
	ApplicationJSON = "application/json"
	ApplicationXML  = "application/xml"
	TextHTML        = "text/html"
	TextXML         = "text/xml"
)

// 全局变量
var (
	ROOT_DIR      = FileGetParentDirectory(FileGetCurrentDirectory())
	SRC_DIR       = ROOT_DIR + "/src"
	PIDS_DIR      = ROOT_DIR + "/pids"
	TEMPLATES_DIR = SRC_DIR + "/views"
	CURRENT_TIME  = time.Now().String()
	CONFIG_PATH   = flag.String("config", "config/config-dev.ini", "config file path")
)
