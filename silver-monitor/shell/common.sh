#!/usr/bin/env bash

# 项目名称
PROJECT_NAME="silver-monitor"
# 服务端名称
SERVER_NAME=server
# 管理端名称
MANAGER_NAME=manager

# 根目录
ENV_ROOT_DIR=$(cd "$(dirname "$1")" || exit; pwd)
# src目录
ENV_SRC_DIR="$ENV_ROOT_DIR"/src
# bin目录
ENV_BIN_DIR="$ENV_ROOT_DIR"/bin
# pid目录
ENV_PID_DIR="$ENV_ROOT_DIR"/pid

# 编译参数缺省值
BUILD_MODE="server"

# 控制台输出
_info() {
    printf "[%-5s] %s\n" "${FUNCNAME[1]}" "$1"
}

# 错误信息
_error() {
    printf "[%-5s] Error: %s \n" "${FUNCNAME[1]}" "$1"
    exit
}