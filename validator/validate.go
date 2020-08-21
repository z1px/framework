package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func New() (validate *validator.Validate, trans ut.Translator, err error) {
	//中文翻译器
	zh_ch := zh.New()
	uni := ut.New(zh_ch)
	trans, found := uni.GetTranslator("zh")
	if found {
		err = fmt.Errorf("验证器翻译失败")
		return
	}
	//验证器
	validate = validator.New()
	// 绑定验证器
	err = validate.RegisterValidation("mobile", IsMobile)
	if err != nil {
		return
	}
	err = validate.RegisterValidation("empty", IsEmpty)
	if err != nil {
		return
	}
	// 自定义mobile错误内容
	registerMobileFn := func(ut ut.Translator) (err error) {
		err = ut.Add("mobile", "手机号格式错误!", true) // see universal-translator for details
		return
	}
	translationMobileFn := func(ut ut.Translator, fe validator.FieldError) (t string) {
		t, _ = ut.T("mobile", fe.Field())
		return
	}
	if err = validate.RegisterTranslation("mobile", trans, registerMobileFn, translationMobileFn); err != nil {
		return
	}
	// 自定义empty错误内容
	registerEmptyFn := func(ut ut.Translator) (err error) {
		err = ut.Add("empty", "{0} 必须为空!", true) // see universal-translator for details
		return
	}
	translationEmptyFn := func(ut ut.Translator, fe validator.FieldError) (t string) {
		t, _ = ut.T("empty", fe.Field())
		return
	}
	if err = validate.RegisterTranslation("empty", trans, registerEmptyFn, translationEmptyFn); err != nil {
		return
	}
	//验证器注册翻译器
	if err = zhTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
		return
	}
	return
}

// 验证结构体
func Validate(params interface{}) (err error) {
	validate, trans, err := New()
	if err != nil {
		return
	}
	//查看是否符合验证
	if errs := validate.Struct(params); errs != nil {
		for _, validateErr := range errs.(validator.ValidationErrors) {
			err = fmt.Errorf(validateErr.Translate(trans))
			return
		}
	}
	return
}

// 验证单个字段
func IsValidate(value interface{}, tag string) (err error) {
	validate, trans, err := New()
	if err != nil {
		return
	}
	//查看是否符合验证
	if errs := validate.Var(value, tag); errs != nil {
		for _, validateErr := range errs.(validator.ValidationErrors) {
			err = fmt.Errorf(validateErr.Translate(trans))
			return
		}
	}
	return
}
