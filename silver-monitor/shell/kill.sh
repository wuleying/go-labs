#!/usr/bin/env bash

# 根目录
ENV_ROOT_DIR="$(cd "$(dirname "$0")" && cd .. && pwd)"

# shellcheck disable=SC1091
source "$ENV_ROOT_DIR"/shell/common.sh

kill_process(){
    if [[ "$BUILD_MODE" == "server" ]]; then
        PID_FILE_PATH="$ENV_PID_DIR"/"$FILE_NAME_SERVER_PID"
    fi

    if [[ "$BUILD_MODE" == "manager" ]]; then
        PID_FILE_PATH="$ENV_PID_DIR"/"$FILE_NAME_MANAGER_PID"
    fi

    if [[ -f "$PID_FILE_PATH" ]]; then
        _info "PID_FILE_PATH:   $PID_FILE_PATH"

        # 杀死原进程
        kill -9 "$(sed -n 1p "$PID_FILE_PATH")"
        # 删除pid文件
        rm "$PID_FILE_PATH"

        _info "silver-monitor-$PID_FILE_PATH stopped."
    fi
}

main(){
    kill_process "$@"
}

main "$@"