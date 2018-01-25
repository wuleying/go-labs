#!/usr/bin/env bash

# shellcheck disable=SC1091
source common.sh

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