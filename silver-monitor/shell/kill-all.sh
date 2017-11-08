#!/usr/bin/env bash

# shellcheck disable=SC1091
source ./shell/common.sh

kill_all(){
    pgrep -f "silver-monitor-server" | xargs kill -9
    pgrep -f "silver-monitor-manager" | xargs kill -9

    _info "Is all stopped."
}

main(){
    kill_all "$@"
}

main "$@"