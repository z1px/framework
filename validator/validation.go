package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/z1px/framework/util"
)

// 验证是否是手机号
func IsMobile(fl validator.FieldLevel) bool {
	return util.IsMobile(fl.Field().String())
}

// 验证是否是空字符串
func IsEmpty(fl validator.FieldLevel) bool {
	value := fl.Field().String() == ""
	return value
}
