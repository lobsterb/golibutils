package utils

import "strings"

// ToLower 字符串转小写
func ToLower(str string) string {
	return strings.ToLower(str)
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
