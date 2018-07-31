package main

import (
	"runtime"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(GetGoid())
}

func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)
	id, err := strconv.Atoi(idField[0])
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}
