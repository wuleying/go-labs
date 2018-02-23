package util

import (
	"bytes"
	"encoding/binary"
	"github.com/go-clog/clog"
)

// 整数转十六进制
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		clog.Fatal(CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	return buff.Bytes()
}

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// 截断字符串
func StringSubstr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
