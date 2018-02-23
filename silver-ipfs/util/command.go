package util

import (
	"bufio"
	"github.com/go-clog/clog"
	"io"
	"os/exec"
)

// 执行命令
func ExecCommand(commandName string, params []string) (err error, result string) {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	clog.Info("%s", cmd.Args)

	stdout, e := cmd.StdoutPipe()

	if e != nil {
		return e, result
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, e := reader.ReadString('\n')
		if e != nil || io.EOF == e {
			return e, result
		}
		clog.Info(line)
	}

	cmd.Wait()
	return nil, result
}
