package commands

import (
	"github.com/go-clog/clog"
	"os/exec"
)

// 执行命令
func execCommand(commandName string, params []string) (error, string) {
	cmd := exec.Command(commandName, params...)

	clog.Trace("%s", cmd.Args)

	out, err := cmd.Output()

	if err != nil {
		return err, ""
	}

	return nil, string(out)
}

func execIPFSCommand(commandType string, param string) (error, string) {
	command := "ipfs"
	params := []string{commandType, param}
	err, out := execCommand(command, params)

	if err != nil {
		return err, ""
	}

	return nil, out
}
