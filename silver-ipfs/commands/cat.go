package commands

// 读取文件信息
func CatFile(fileHash string) (string, error) {
	return execIPFSCommand("cat", fileHash)
}
