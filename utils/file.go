package utils

import (
	"os"
)

type ErrDataIsNil struct {
	msg string
}

func (e ErrDataIsNil) Error() string {
	return "data is nil"
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
