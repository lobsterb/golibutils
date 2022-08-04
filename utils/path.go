package utils

import (
	"os"
	"strings"
)

// GetWorkDir 获取工作目录
func GetWorkDir() (string, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return workDir, nil
}

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

// NormativePath 规范路径
func NormativePath(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}

// NormativeDirPath 规范目录路径
func NormativeDirPath(path string) string {
	dirPath := NormativePath(path)
	// 判断最后一个字符是否是/
	if dirPath[len(dirPath)-1:] != "/" {
		dirPath = dirPath + "/"
	}
	return dirPath
}
