#!/usr/bin/env bash

# 根目录
ENV_ROOT_DIR="$(cd "$(dirname "$0")" && cd .. && pwd)"

# shellcheck disable=SC1091
source "$ENV_ROOT_DIR"/shell/common.sh

start() {
    # 编译
    "$ENV_ROOT_DIR"/shell/build.sh

    # 执行
    # nohup \
    "$ENV_BIN_DIR"/"$PROJECT_NAME"

    _info "$PROJECT_NAME started."
}

main(){
    start "$@"
}

main "$@"