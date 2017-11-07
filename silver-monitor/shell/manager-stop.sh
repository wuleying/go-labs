#!/usr/bin/env bash

# shellcheck disable=SC1091
source ./shell/common.sh

stop() {
    # 杀死原进程
    "$ENV_SHELL_DIR"/kill.sh manager
}

main(){
    stop "$@"
}

main "$@"