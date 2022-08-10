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

// ReplaceSlashPath 规范路径
func ReplaceSlashPath(path string, isBackslash bool) string {
	old := "\\"
	new := "/"
	if isBackslash == false {
		old = "/"
		new = "\\"
	}
	return strings.Replace(path, old, new, -1)
}

// ReplaceDirSlash 规范目录路径
func ReplaceDirSlash(path string, isBackslash bool) string {
	dirPath := ReplaceSlashPath(path, isBackslash)
	slash := "/"
	if isBackslash == false {
		slash = "\\"
	}
	// 判断最后一个字符是否是/
	if dirPath[len(dirPath)-1:] != slash {
		dirPath = dirPath + slash
	}
	return dirPath
}
