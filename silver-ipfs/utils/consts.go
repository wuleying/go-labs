package utils

import (
	"time"
)

// 全局常量
const (
	VERSION = "0.0.1"

	// 权限
	DB_FILE_MODE    = 0600
	FILE_READ_MODE  = 0644
	FILE_WRITE_MODE = 0666
	DIR_READ_MODE   = 0755
	DIR_WRITE_MODE  = 0777

	// Clog skip 级别
	CLOG_SKIP_DEFAULT      = 0
	CLOG_SKIP_DISPLAY_INFO = 2

	// 数据库路径
	DB_FILE_PATH = "db/silver-ipfs.db"
	// Bucket name
	BLOCK_BUCKET_NAME = "silver-ipfs"
)

// 全局变量
var (
	// 根目录
	ROOT_DIR = FileGetParentDirectory(FileGetCurrentDirectory())
	// 视图目录
	TEMPLATES_DIR = ROOT_DIR + "/views"
	// 当前时间
	CURRENT_TIME = time.Now().String()
)
