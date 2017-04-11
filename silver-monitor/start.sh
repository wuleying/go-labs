#!/usr/bin/env bash

#杀死原进程
kill -9 `sed -n 1p  ./pid`

go clean

echo "build begin..."
go build -o silver-monitor main.go

echo "run..."
chmod +x ./silver-monitor
#./silver_monitor
nohup ./silver-monitor &