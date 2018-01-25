#!/usr/bin/env bash

# shellcheck disable=SC1091
source common.sh

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

    if [[ -f "$ENV_BIN_DIR/$PROJECT_NAME" ]]; then
        _info "rm $ENV_BIN_DIR/$PROJECT_NAME"
        rm "$ENV_BIN_DIR/$PROJECT_NAME"
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

    go build -o "$ENV_BIN_DIR"/"$PROJECT_NAME" "$ENV_SRC_DIR"/*.go
    chmod +x "$ENV_BIN_DIR"/"$PROJECT_NAME"

    _info "end..."
}

main(){
    init "$@"
    clean
    mdir
    build
}

main "$@"