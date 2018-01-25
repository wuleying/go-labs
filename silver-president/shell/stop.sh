#!/usr/bin/env bash

# 根目录
ENV_ROOT_DIR="$(cd "$(dirname "$0")" && cd .. && pwd)"
# shell目录
ENV_SHELL_DIR="$ENV_ROOT_DIR"/shell

# shellcheck disable=SC1091
source "$ENV_SHELL_DIR"/common.sh

kill_process(){
    PID_FILE_PATH="$ENV_PID_DIR"/"$PROJECT_NAME".pid

    if [[ -f "$PID_FILE_PATH" ]]; then
        _info "PID_FILE_PATH:   $PID_FILE_PATH"

        # 杀死原进程
        kill -9 "$(sed -n 1p "$PID_FILE_PATH")"
        # 删除pid文件
        rm "$PID_FILE_PATH"

        _info "$PROJECT_NAME stopped."
    fi
}

kill_all(){
    pgrep -f "$PROJECT_NAME" | xargs kill -9

    _info "Is all stopped."
}

main(){
    kill_process "$@"
}

main "$@"