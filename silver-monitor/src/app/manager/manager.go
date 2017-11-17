package main

import (
    "log"
    "go-labs/silver-monitor/src/model"
    "go-labs/silver-monitor/src/util"
)

// 全局配置
var config util.Config
var err error

func main() {
    util.SavePid("silver-monitor-manager.pid");

    config, err = util.InitConfig();
    if err != nil {
        log.Fatal("Init config failed: ", err.Error())
    }

    db := model.Init(config)

    print(db)
}

