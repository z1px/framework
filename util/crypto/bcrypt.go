package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// HASH加密
func HashEncrypt(decrypted string) (encrypted string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(decrypted), bcrypt.DefaultCost)
	if err == nil {
		encrypted = string(hash)
	}
	return
}

// HASH比较
func HashCompare(decrypted string, encrypted string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(decrypted))
	if err != nil {
		return false
	} else {
		return true
	}
}
