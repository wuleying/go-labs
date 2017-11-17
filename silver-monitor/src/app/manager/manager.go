package main

import (
    "log"
    "go-labs/silver-monitor/src/util"
    "go-labs/silver-monitor/src/model"
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

    // 初始化模型
    model.InitModel(config)

    model.LogList()
}

