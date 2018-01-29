package util

import (
	"math"
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
	// UTXO Bucket名称
	UTXO_BUCKET_NAME = "chainState"

	// 创世币数据
	GENESIS_COIN_BASEDATA = "hello luoliang"

	// 挖矿难度
	MINE_TARGET_BITS = 18
	// 挖矿奖励
	MINE_SUBSIDY = 10
	// Number once最大值
	MAX_NONCE = math.MaxInt64

	// 权限
	FILE_READ_MODE  = 0644
	FILE_WRITE_MODE = 0666
	DIR_READ_MODE   = 0755
	DIR_WRITE_MODE  = 0777
)

// 全局变量
var (
	// 根目录
	ROOT_DIR = FileGetParentDirectory(FileGetCurrentDirectory())
	// 当前时间
	CURRENT_TIME = time.Now().String()
)
