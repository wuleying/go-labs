package common

import (
    "flag"
)

// 全局常量
const (
    VERSION = "1.0.0"
    DEV = "dev"
    PROD = "prod"
)

// http头
const (
    ApplicationJSON = "application/json"
    ApplicationXML = "application/xml"
    TextXML = "text/xml"
)

var (
    // 项目根目录
    ROOT_DIR = GetParentDirectory(GetCurrentDirectory())
    // 源代码目录
    SRC_DIR = ROOT_DIR + "/src"
    // PID文件目录
    PIDS_DIR = ROOT_DIR + "/pids"

    // 模块目录
    MODULES_DIR = SRC_DIR + "/modules"
    // 模板目录
    TEMPLATES_DIR = SRC_DIR + "/templates"


    // 命令行参数，配置文件路径
    CONFIG_PATH = flag.String("config", "config/config-prod.ini", "config file path")
)