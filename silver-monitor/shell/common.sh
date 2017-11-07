#!/usr/bin/env bash

# shellcheck disable=SC2034

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
# config目录
ENV_CONFIG_DIR="$ENV_ROOT_DIR"/config
# shell目录
ENV_SHELL_DIR="$ENV_ROOT_DIR"/shell

# 服务端pid文件名称
FILE_NAME_SERVER_PID="silver-monitor-server.pid"
# 管理端pid文件名称
FILE_NAME_MANAGER_PID="silver-monitor-manager.pid"

# 编译参数缺省值
BUILD_MODE="server"

# 编译参数
if [ -n "$1" ]; then
    BUILD_MODE=$(echo "$1" | tr '[:upper:]' '[:lower:]')
fi

# 编译参数校验
if [[ "$BUILD_MODE" != "server" && "$BUILD_MODE" != "manager" ]]; then
    _error "Build mode must be 'server' or 'manager'."
fi

# 控制台输出
_info() {
    printf "[%-5s] %s\n" "${FUNCNAME[1]}" "$1"
}

# 错误信息
_error() {
    printf "[%-5s] Error: %s \n" "${FUNCNAME[1]}" "$1"
    exit
}