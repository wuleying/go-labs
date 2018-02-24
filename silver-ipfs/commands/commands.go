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

	/*
		cmd.Start()

		reader := bufio.NewReader(stdout)

		//实时循环读取输出流中的一行内容
		for {
			line, e := reader.ReadString('\n')
			if e != nil || io.EOF == e {
				return e, result
			}
			clog.Trace(line)
		}

		cmd.Wait()
	*/

	return nil, string(out)
}
