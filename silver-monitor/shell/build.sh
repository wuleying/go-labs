#!/usr/bin/env bash

# 根目录
ENV_ROOT_DIR="$(cd "$(dirname "$0")" && cd .. && pwd)"

# shellcheck disable=SC1091
source "$ENV_ROOT_DIR"/shell/common.sh

# 初始化
init(){
    _info "BUILD_MODE:      $BUILD_MODE"
    _info "ENV_ROOT_DIR:    $ENV_ROOT_DIR"
    _info "ENV_SRC_DIR:     $ENV_SRC_DIR"
    _info "ENV_BIN_DIR:     $ENV_BIN_DIR"
    _info "ENV_PID_DIR:     $ENV_PID_DIR"
}

# 清理工作
clean(){
    _info "start..."

    go clean

    if [[ -f "$ENV_BIN_DIR/$PROJECT_NAME-$BUILD_MODE" ]]; then
        _info "rm $ENV_BIN_DIR/$PROJECT_NAME-$BUILD_MODE"
        rm "$ENV_BIN_DIR/$PROJECT_NAME-$BUILD_MODE"
    fi

    _info "end..."
}

# 成生目录
mdir(){
    _info "start..."

    # 生成二进制文件目录
    if [[ ! -d "$ENV_BIN_DIR" ]]; then
        _info "mkdir $ENV_BIN_DIR"
        mkdir "$ENV_BIN_DIR"
    fi

    # 生成PID文件目录
    if [[ ! -d "$ENV_PID_DIR" ]]; then
        _info "mkdir $ENV_PID_DIR"
        mkdir "$ENV_PID_DIR"
    fi

    _info "end..."
}

# 编译
build() {
    _info "start..."

    if [[ "$BUILD_MODE" == "server" ]]; then
        go build -o "$ENV_BIN_DIR"/"$PROJECT_NAME"-"$SERVER_NAME" "$ENV_SRC_DIR"/app/"$SERVER_NAME"/*.go
        chmod +x "$ENV_BIN_DIR"/"$PROJECT_NAME"-"$SERVER_NAME"
    fi

    if [[ "$BUILD_MODE" == "manager" ]]; then
        go build -o "$ENV_BIN_DIR"/"$PROJECT_NAME"-"$MANAGER_NAME" "$ENV_SRC_DIR"/app/"$MANAGER_NAME"/*.go
        chmod +x "$ENV_BIN_DIR"/"$PROJECT_NAME"-"$MANAGER_NAME"
    fi

    _info "end..."
}

main(){
    init "$@"
    clean
    mdir
    build
}

main "$@"