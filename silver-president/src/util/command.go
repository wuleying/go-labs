package util

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

// 执行命令
func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}

func Notification(title string) {
	command := "osascript"
	params := []string{"-e", fmt.Sprintf("display notification \"%s\" with title \"Silver president\" sound name \"Glass\"", title)}
	execCommand(command, params)
}
