package util

import (
	"os/exec"
)

// 执行命令
func ExecCommand(commandName string, params []string) (error, string) {
	cmd := exec.Command(commandName, params...)
	out, err := cmd.CombinedOutput()

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
