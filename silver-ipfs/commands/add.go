package commands

// 添加文件
func AddFile(filePath string) (error, string) {
	command := "ipfs"
	params := []string{"add", filePath}
	err, out := execCommand(command, params)

	if err != nil {
		return err, ""
	}

	return nil, out
}
