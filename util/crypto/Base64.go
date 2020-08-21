package crypto

import (
	"encoding/base64"
)

// BASE64加密
func Base64Encrypt(decrypted string) (encrypted string) {
	// 字符串转字节
	decryptedByte := []byte(decrypted)
	// base64编码
	encrypted = base64.StdEncoding.EncodeToString(decryptedByte)
	return
}

// BASE64解密
func Base64Decrypt(encrypted string) (decrypted string, err error) {
	// base64解码
	decryptedByte, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return
	}
	// 字节转字符串
	decrypted = string(decryptedByte)
	return
}
