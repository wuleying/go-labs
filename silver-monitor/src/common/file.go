package common

import (
    "strconv"
    "os"
    "log"
    "path/filepath"
    "strings"
)

// 保存pid
func SavePid(pid_path string) {
    pid_path = PIDS_DIR + "/" + pid_path

    file, err := os.OpenFile(pid_path, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)

    if err != nil {
        log.Fatal("Open file failed.", err.Error())
    }

    defer file.Close()

    pid := os.Getpid()

    log.Printf("Pid:%d", pid)

    file.WriteString(strconv.Itoa(pid))
}

// 获取当前目录
func GetCurrentDirectory() string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal("Get current directory failed.", err.Error())
    }

    return strings.Replace(dir, "\\", "/", -1)
}

// 获取上级目录
func GetParentDirectory(dirctory string) string {
    return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}