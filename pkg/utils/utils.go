package utils

import (
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"
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

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
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

// 判断文件或文件夹是否存在
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

// 字节的单位转换 保留两位小数
func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.0fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.0fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

//进行zlib压缩
func ZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

//进行zlib解压缩
func ZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}