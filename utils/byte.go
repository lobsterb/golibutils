package utils

import (
	"encoding/json"
)

// ToJsonBytes 转换成json byte数组
func ToJsonBytes(data interface{}) []byte {
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return marshal
}
