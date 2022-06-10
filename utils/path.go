package utils

import "os"

// CheckFilePathExist 检测文件路径是否存在
func CheckFilePathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
