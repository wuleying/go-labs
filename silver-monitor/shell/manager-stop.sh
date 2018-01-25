#!/usr/bin/env bash

# 根目录
ENV_ROOT_DIR="$(cd "$(dirname "$0")" && cd .. && pwd)"

# shellcheck disable=SC1091
source "$ENV_ROOT_DIR"/shell/common.sh

stop() {
    # 杀死原进程
    "$ENV_ROOT_DIR"/shell/kill.sh manager
}

main(){
    stop "$@"
}

main "$@"