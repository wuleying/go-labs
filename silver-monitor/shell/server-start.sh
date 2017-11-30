#!/usr/bin/env bash

# shellcheck disable=SC1091
source ./shell/common.sh

start() {
    # 杀死原进程
    "$ENV_SHELL_DIR"/kill.sh server

    # 编译
    "$ENV_SHELL_DIR"/build.sh server

    # 执行
    nohup \
    "$ENV_BIN_DIR"/"$PROJECT_NAME"-"$SERVER_NAME" \
    --config="$ENV_CONFIG_DIR"/config-dev.ini &

    _info "silver-server started."
}

main(){
    start "$@"
}

main "$@"