package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

// 整数转十六进制
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)

	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
