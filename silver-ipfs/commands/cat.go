package commands

// 读取文件信息
func CatFile(fileHash string) (error, string) {
	command := "ipfs"
	params := []string{"cat", fileHash}
	err, out := execCommand(command, params)

	if err != nil {
		return err, ""
	}

	return nil, out
}
