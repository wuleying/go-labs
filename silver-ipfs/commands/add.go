package commands

import (
	"github.com/juju/errors"
	"strings"
)

// 添加文件
func AddFile(filePath string) (string, error) {
	result, err := execIPFSCommand("add", filePath)
	if err != nil {
		return "", err
	}

	resultArr := strings.Fields(result)

	if len(resultArr[1]) < 1 {
		return "", errors.New("Result date error.")
	}

	return resultArr[1], nil

}
