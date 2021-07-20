package utils

import (
	"anew-server/pkg/common"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// 生成密码, 由于使用自适应hash算法, 不可逆
func GenPwd(str string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hash)
}

// 通过比较两个字符串hash判断是否出自同一个明文
// str 明文
// pwd 需要对比的密文
func ComparePwd(str string, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(str)); err != nil {
		return false
	}
	return true
}

func AesEncryptCBC2Hex(origData string) (string) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	key := []byte(common.Conf.System.Key)
	srcData := []byte(origData)
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	srcData = pkcs5Padding(srcData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted := make([]byte, len(srcData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, srcData)
	// 加密
	return hex.EncodeToString(encrypted)
}
func AesDecryptCBC2Hex(encrypted string) (string) {
	key := []byte(common.Conf.System.Key)
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	srcData,_ :=hex.DecodeString(encrypted)
	decrypted := make([]byte, len(srcData))                    // 创建数组
	blockMode.CryptBlocks(decrypted, srcData)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return string(decrypted)
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
