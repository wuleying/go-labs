package utils

import (
	"bytes"
	"encoding/gob"
	"github.com/go-clog/clog"
)

func GobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(data)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	return buff.Bytes()
}
