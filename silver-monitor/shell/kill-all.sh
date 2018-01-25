#!/usr/bin/env bash
# 滋道你在做什么不？
# 你这是在犯罪呀...

# 根目录
ENV_ROOT_DIR="$(cd "$(dirname "$0")" && cd .. && pwd)"

# shellcheck disable=SC1091
source "$ENV_ROOT_DIR"/shell/common.sh

kill_all(){
    pgrep -f "silver-monitor-server"    | xargs kill -9
    pgrep -f "silver-monitor-manager"   | xargs kill -9
    pgrep -f "build.sh manager"         | xargs kill -9
    pgrep -f "build.sh server"          | xargs kill -9

    _info "Is all stopped."
}

main(){
    kill_all "$@"
}

main "$@"