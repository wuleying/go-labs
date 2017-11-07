#!/usr/bin/env bash

# shellcheck disable=SC1091
source ./shell/common.sh

# 初始化
init(){
    # 编译参数
    if [ -n "$1" ]; then
        BUILD_MODE=$(echo "$1" | tr '[:upper:]' '[:lower:]')
    fi

    if [[ "$BUILD_MODE" != "server" && "$BUILD_MODE" != "manager" ]]; then
        _error "Build mode must be 'server' or 'manager'."
    fi

    _info "BUILD_MODE:      $BUILD_MODE"
    _info "ENV_ROOT_DIR:    $ENV_ROOT_DIR"
    _info "ENV_SRC_DIR:     $ENV_SRC_DIR"
    _info "ENV_BIN_DIR:     $ENV_BIN_DIR"
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

# 成生编译目录
mdir(){
    _info "start..."

    if [[ ! -d "$ENV_BIN_DIR" ]]; then
        _info "mkdir $ENV_BIN_DIR"
        mkdir "$ENV_BIN_DIR"
    fi

    _info "end..."
}

# 编译
build() {
    _info "start..."

    if [[ "$BUILD_MODE" == "server" ]]; then
        go build -o "$ENV_BIN_DIR"/"$PROJECT_NAME"-"$SERVER_NAME" "$ENV_SRC_DIR"/"$SERVER_NAME"/*.go
        chmod +x "$ENV_BIN_DIR"/"$PROJECT_NAME"-"$SERVER_NAME"
    fi

    if [[ "$BUILD_MODE" == "manager" ]]; then
        go build -o "$ENV_BIN_DIR"/"$PROJECT_NAME"-"$MANAGER_NAME" "$ENV_SRC_DIR"/"$MANAGER_NAME"/*.go
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