package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ErrDataIsNil struct {
	msg string
}

func (e ErrDataIsNil) Error() string {
	return "data is nil"
}

// CheckPathExist 检测文件路径是否存在
func CheckPathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// CreateDir 创建目录
func CreateDir(fullPath string) (bool, error) {
	parentPath, _ := filepath.Split(fullPath)
	if ok := CheckPathExist(parentPath); ok {
		return true, nil
	} else {
		if err := os.MkdirAll(parentPath, os.ModePerm); err != nil {
			return false, err
		}
		return true, nil
	}
}

// RemoveFile 删除文件
func RemoveFile(path string) error {
	_err := os.Remove(path)
	return _err
}

// RemoveDir 删除文件夹
func RemoveDir(path string) error {
	_err := os.RemoveAll(path)
	return _err
}

// ReadAllFile 一次性读取文件
func ReadAllFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// SaveFile 保存文件
func SaveFile(filePath string, data []byte) (bool, error) {

	// 如果data为空时, 返回错误信息
	if data == nil {
		return false, ErrDataIsNil{}
	}

	// 创建文件
	create, err := os.Create(filePath)
	if err != nil {
		return false, err
	}
	// 用后关闭
	defer create.Close()

	// 写入文件
	_, err = create.Write(data)
	if err != nil {
		return false, err
	}
	return true, nil
}

// SaveJsonFile 保存成json文件
func SaveJsonFile(filePath string, data interface{}) (bool, error) {
	return SaveFile(filePath, ToJsonBytes(data))
}

// GetFileMd5 计算文件md5
func GetFileMd5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	md5Hash := hex.EncodeToString(hash.Sum(nil))
	return md5Hash, nil
}
