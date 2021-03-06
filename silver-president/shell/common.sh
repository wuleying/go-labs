#!/usr/bin/env bash

# shellcheck disable=SC2034

# 项目名称
PROJECT_NAME="silver-president"

# 根目录
ENV_ROOT_DIR="$(cd "$(dirname "$0")" && cd .. && pwd)"
# src目录
ENV_SRC_DIR="$ENV_ROOT_DIR"/src
# bin目录
ENV_BIN_DIR="$ENV_ROOT_DIR"/bin
# pid目录
ENV_PID_DIR="$ENV_ROOT_DIR"/pids
# config目录
ENV_CONFIG_DIR="$ENV_ROOT_DIR"/config

# 编译参数缺省值 dev/test/prod
BUILD_MODE="prod"

# 编译参数校验
if [[ "$BUILD_MODE" != "dev" && "$BUILD_MODE" != "test" && "$BUILD_MODE" != "prod" ]]; then
    _error "Build mode must be 'dev', 'test' or 'prod'."
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