#!/usr/bin/env bash

# 根目录
ENV_ROOT_DIR="$(cd "$(dirname "$0")" && cd .. && pwd)"
# shell目录
ENV_SHELL_DIR="$ENV_ROOT_DIR"/shell

# shellcheck disable=SC1091
source "$ENV_SHELL_DIR"/common.sh

start() {
    # 杀死原进程
    "$ENV_SHELL_DIR"/stop.sh

    # 编译
    "$ENV_SHELL_DIR"/build.sh

    # 执行
    nohup \
    "$ENV_BIN_DIR"/"$PROJECT_NAME" \
    --config="$ENV_CONFIG_DIR"/config-dev.ini &

    _info "$PROJECT_NAME started."
}

main(){
    start "$@"
}

main "$@"