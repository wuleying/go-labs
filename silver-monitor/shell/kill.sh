#!/usr/bin/env bash

# shellcheck disable=SC1091
source ./shell/common.sh

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

        _info "silver-monitor-server stopped."
    fi
}

main(){
    kill_process "$@"
}

main "$@"