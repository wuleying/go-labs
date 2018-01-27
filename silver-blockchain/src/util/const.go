package util

import (
	"time"
)

// 全局常量
const (
	VERSION              = byte(0x00)
	ADDRESS_CHECKSUM_LEN = 4

	// 数据库路径
	DB_FILE_PATH = "db/silver-blockchain-%s.db"
	// Last hash key
	LAST_HASH_KEY = "lastHash"
	// Bucket名称
	BLOCK_BUCKET_NAME = "blocks"
	// 创世币数据
	GENESIS_COIN_BASEDATA = "hello luoliang"

	// 权限
	FILE_READ_MODE  = 0644
	FILE_WRITE_MODE = 0666
	DIR_READ_MODE   = 0755
	DIR_WRITE_MODE  = 0777
)

// 全局变量
var (
	ROOT_DIR     = FileGetParentDirectory(FileGetCurrentDirectory()) // 根目录
	CURRENT_TIME = time.Now().String()                               // 当前时间
)
