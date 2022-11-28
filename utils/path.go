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
func ReplaceSlashPath(path string, useBackslash bool) string {
	slash := "\\"
	newSlash := "/"
	if useBackslash == true {
		slash = "/"
		newSlash = "\\"
	}

	return strings.Replace(path, slash, newSlash, -1)
}

// NormalizationDirSlash 规范目录路径
func NormalizationDirSlash(path string, useBackslash bool) string {
	dirPath := ReplaceSlashPath(path, useBackslash)
	slash := "/"
	if useBackslash == true {
		slash = "\\"
	}
	// 判断最后一个字符是否是/
	if dirPath[len(dirPath)-1:] != slash {
		dirPath = dirPath + slash
	}

	// 判断是否有多个/获取\\
	doubleSlash := slash + slash
	dirPath = strings.Replace(dirPath, doubleSlash, slash, -1)

	return dirPath
}
