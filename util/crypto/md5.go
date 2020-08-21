package crypto

import (
	"crypto/md5"
	"fmt"
)

// MD5加密
func Md5(str string) (md5str string) {
	strByte := []byte(str)
	has := md5.Sum(strByte)
	md5str = fmt.Sprintf("%x", has) // 将[]byte转成16进制
	return
}
