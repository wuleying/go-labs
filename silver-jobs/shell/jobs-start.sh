#!/usr/bin/env bash

# shellcheck disable=SC1091
source ./shell/common.sh

start() {
    # 编译
    "$ENV_SHELL_DIR"/build.sh

    # 执行
    #nohup \
    "$ENV_BIN_DIR"/"$PROJECT_NAME"

    _info "silver-credit started."
}

main(){
    start "$@"
}

main "$@"