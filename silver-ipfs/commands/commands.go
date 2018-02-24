package commands

import (
	"github.com/go-clog/clog"
	"os/exec"
)

// 执行命令
func execCommand(commandName string, params []string) (string, error) {
	cmd := exec.Command(commandName, params...)

	clog.Trace("%s", cmd.Args)

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func execIPFSCommand(commandType string, param string) (string, error) {
	command := "ipfs"
	params := []string{commandType, param}
	out, err := execCommand(command, params)

	if err != nil {
		return "", err
	}

	return out, nil
}
