#!/usr/bin/env bash

# 根目录
ENV_ROOT_DIR=""
# bin目录
ENV_PID_DIR=""
# pid文件路径
PID_FILE_PATH=""

kill_process(){
    ENV_ROOT_DIR=$(cd "$(dirname "$1")" || exit; pwd)
    ENV_PID_DIR="$ENV_ROOT_DIR"/pid
    PID_FILE_PATH="$ENV_PID_DIR"/silver-monitor-server.pid

    #杀死原进程
    kill -9 "$(sed -n 1p "$PID_FILE_PATH")"

    if [[ -f "$PID_FILE_PATH" ]]; then
        rm "$PID_FILE_PATH"
    fi

    echo "silver-monitor-server stopped."
}

main(){
    kill_process "$@"
}

main "$@"