#!/usr/bin/env bash

# 根目录
ENV_ROOT_DIR="$(cd "$(dirname "$0")" && cd .. && pwd)"
# shell目录
ENV_SHELL_DIR="$ENV_ROOT_DIR"/shell

# shellcheck disable=SC2009
# shellcheck disable=SC1091
source "$ENV_SHELL_DIR"/common.sh

silver_ps(){
    ps -ef | grep "$PROJECT_NAME" | grep -v grep
}

main(){
    silver_ps "$@"
}

main "$@"