package utils

import (
	"encoding/json"
)

// ToJsonBytes 转换成json byte数组
func ToJsonBytes(data interface{}) []byte {
	marshal, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil
	}
	return marshal
}
