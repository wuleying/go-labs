package commands

// 添加文件
func AddFile(filePath string) (string, error) {
	return execIPFSCommand("add", filePath)
}
