package common

import (
    "strconv"
    "os"
    "log"
)

// 保存pid
func SavePid(pid_path string) {
    file, err := os.OpenFile(pid_path, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)

    if err != nil {
        log.Fatal("open file failed.", err.Error())
    }

    defer file.Close()

    pid := os.Getpid()

    log.Printf("pid:%d", pid)

    file.WriteString(strconv.Itoa(pid))
}
