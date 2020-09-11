package utils

import (
	"strconv"
)

// 字符串转uint
func Str2Int(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return num
}