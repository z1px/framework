package util

import (
	"regexp"
)

// 检查手机格式
func IsMobile(mobile string) (matched bool) {
	var err error
	matched, err = regexp.MatchString(`^1[3456789]\d{9}$`, mobile)
	if err != nil {
		matched = false
	}
	return
}

// 检查邮箱格式
func IsEmail(email string) (matched bool) {
	var err error
	matched, err = regexp.MatchString(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`, email)
	if err != nil {
		matched = false
	}
	return
}
