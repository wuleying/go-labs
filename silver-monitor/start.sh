#!/usr/bin/env bash

# 杀死原进程
kill_process() {
    echo "kill process..."
    kill -9 "$(sed -n 1p ./pid)"
}

build() {
    echo "clean..."
    go clean

    echo "build begin..."
    go build -o silver-monitor main.go
    chmod +x ./silver-monitor
    echo "build end..."
}

begin_monitor() {
    nohup ./silver-monitor &
}

main(){
    kill_process "$@"
    build
    begin_monitor
}

main "$@"