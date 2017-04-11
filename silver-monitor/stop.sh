#!/usr/bin/env bash

#杀死原进程
kill -9 `sed -n 1p  ./pid`