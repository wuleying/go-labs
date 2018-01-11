#!/usr/bin/env bash

# shellcheck disable=SC1091
source ./shell/common.sh

kill_process(){
    PID_FILE_PATH="$ENV_PID_DIR"/"$PROJECT_NAME".pid

    if [[ -f "$PID_FILE_PATH" ]]; then
        _info "PID_FILE_PATH:   $PID_FILE_PATH"

        # 杀死原进程
        kill -9 "$(sed -n 1p "$PID_FILE_PATH")"
        # 删除pid文件
        rm "$PID_FILE_PATH"

        _info "silver-president-$PID_FILE_PATH stopped."
    fi
}

kill_all(){
    pgrep -f "silver-president" | xargs kill -9

    _info "Is all stopped."
}

main(){
    kill_process "$@"
}

main "$@"