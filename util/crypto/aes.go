package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// PKCS7 填充模式
func PKCS7Padding(decryptedByte []byte, blockSize int) []byte {
	padNum := blockSize - len(decryptedByte)%blockSize
	// Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padText := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(decryptedByte, padText...)
}

// 填充的反向操作，删除填充字符串
func PKCS7UnPadding(decryptedByte []byte) []byte {
	// 获取数据长度
	length := len(decryptedByte)
	// 获取填充字符串长度
	unPadNum := int(decryptedByte[length-1])
	// 截取切片，删除填充字节，并且返回明文
	return decryptedByte[:(length - unPadNum)]
}

// AES加密
func AesEncrypt(decrypted string, key string) (encrypted string, err error) {
	// 字符串转字节
	decryptedByte := []byte(decrypted)
	keyByte := []byte(Md5(key))

	// 创建加密算法实例
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return
	}
	// 获取块的大小
	blockSize := block.BlockSize()
	// 对数据进行填充，让数据长度满足需求
	decryptedByte = PKCS7Padding(decryptedByte, blockSize)
	// 采用AES加密方法中CBC加密模式
	blockMode := cipher.NewCBCEncrypter(block, keyByte[:blockSize])
	encryptedByte := make([]byte, len(decryptedByte))
	// 执行加密
	blockMode.CryptBlocks(encryptedByte, decryptedByte)
	// base64加密
	encrypted = base64.StdEncoding.EncodeToString(encryptedByte)
	return
}

// AES解密
func AesDecrypt(encrypted string, key string) (decrypted string, err error) {
	// base64解密
	encryptedByte, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return
	}
	// 字符串转字节
	keyByte := []byte(Md5(key))

	// 创建加密算法实例
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return
	}
	// 获取块大小
	blockSize := block.BlockSize()
	// 创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, keyByte[:blockSize])
	decryptedByte := make([]byte, len(encryptedByte))
	// 这个函数也可以用来解密
	blockMode.CryptBlocks(decryptedByte, encryptedByte)
	// 去除填充字符串
	decryptedByte = PKCS7UnPadding(decryptedByte)
	// 字节转字符串
	decrypted = string(decryptedByte)
	return
}
