package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//随机字符串
func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// 字符串转int
func Str2Int(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return num
}

// 字符串转uint
func Str2Uint(str string) uint {
	num, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint(num)
}

// 字符串转uint数组, 默认逗号分割
func Str2UintArr(str string) (ids []uint) {
	idArr := strings.Split(str, ",")
	for _, v := range idArr {
		ids = append(ids, Str2Uint(v))
	}
	return
}
// 判断文件是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// 获取文件的MD5
func GetFileMd5(filename string) string {
	file, _ := ioutil.ReadFile(filename)
	return fmt.Sprintf("%x", md5.Sum(file))
}

// 判断uint数组是否包含item元素
func ContainsUint(arr []uint, item uint) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

// TCP端口测试
func Tcping(ip string, port string) bool {
	var conn net.Conn
	var err error

	if conn, err = net.DialTimeout("tcp", ip+":"+port, 2*time.Second); err != nil {
		return false
	}
	defer conn.Close()
	return true
}
