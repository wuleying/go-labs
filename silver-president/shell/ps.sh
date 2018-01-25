#!/usr/bin/env bash

# shellcheck disable=SC2009
# shellcheck disable=SC1091
source common.sh

silver_ps(){
    ps -ef | grep "$PROJECT_NAME" | grep -v grep
}

main(){
    silver_ps "$@"
}

main "$@"