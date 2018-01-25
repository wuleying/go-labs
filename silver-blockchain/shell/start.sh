#!/usr/bin/env bash
# shellcheck disable=SC1091

# 根目录
ENV_ROOT_DIR="$(cd "$(dirname "$0")" && cd .. && pwd)"
source "$ENV_ROOT_DIR"/shell/common.sh

start() {
    # 编译
    "$ENV_ROOT_DIR"/shell/build.sh

    # 执行
    "$ENV_BIN_DIR"/"$PROJECT_NAME" w c
    "$ENV_BIN_DIR"/"$PROJECT_NAME" w c

    _info "silver-blockchain started."
}

clean_db() {
    if [[ -f "$ENV_DB_DIR"/silver-blockchain.db ]]; then
        rm "$ENV_DB_DIR"/silver-blockchain.db
    fi

    if [[ -f "$ENV_DB_DIR"/wallet.dat ]]; then
        rm "$ENV_DB_DIR"/wallet.dat
    fi
}

main(){
    clean_db "$@"
    start
}

main "$@"