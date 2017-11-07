#!/usr/bin/env bash

# shellcheck disable=SC1091
source ./shell/common.sh

kill_process(){
    PID_FILE_PATH="$ENV_PID_DIR"/silver-monitor-server.pid

    _info "PID_FILE_PATH:   $PID_FILE_PATH"

    #杀死原进程
    kill -9 "$(sed -n 1p "$PID_FILE_PATH")"

    if [[ -f "$PID_FILE_PATH" ]]; then
        rm "$PID_FILE_PATH"
    fi

    _info "silver-monitor-server stopped."
}

main(){
    kill_process "$@"
}

main "$@"