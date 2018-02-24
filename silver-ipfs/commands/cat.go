package commands

// 读取文件信息
func CatFile(fileHash string) (error, string) {
	return execIPFSCommand("cat", fileHash)
}
