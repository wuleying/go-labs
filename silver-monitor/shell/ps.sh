#!/usr/bin/env bash

# shellcheck disable=SC2009

silver_ps(){
    ps -ef | grep silver-monitor | grep -v grep
}

main(){
    silver_ps "$@"
}

main "$@"