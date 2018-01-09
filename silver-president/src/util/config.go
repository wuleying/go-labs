package util

import (
	"flag"
	"log"

	"github.com/Unknwon/goconfig"
)

type Config struct {
	Setting  map[string]string
	Database map[string]string
}

// 初始化配置
func ConfigInit() (Config, error) {
	var config Config

	flag.Parse()

	goconfig, err := goconfig.LoadConfigFile(*CONFIG_PATH)

	if err != nil {
		log.Printf("Read config file failed: %s", err)
		return config, err
	}

	log.Printf("Load config file success: %s", *CONFIG_PATH)

	config.Setting, _ = goconfig.GetSection("setting")
	config.Database, _ = goconfig.GetSection("database")

	return config, nil
}
