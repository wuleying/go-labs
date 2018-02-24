package commands

// 添加文件
func AddFile(filePath string) (error, string) {
	return execIPFSCommand("add", filePath)
}
