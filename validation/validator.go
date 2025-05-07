package validation

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type Validator interface {
	Validate(interface{}) error
	RegisterValidation(tag string, fn validator.Func) error
}

type CustomValidator struct {
	validate *validator.Validate
}

func NewValidator() *CustomValidator {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &CustomValidator{validate: v}
}

func (cv *CustomValidator) Validate(s interface{}) error {
	return cv.validate.Struct(s)
}

func (cv *CustomValidator) RegisterValidation(tag string, fn validator.Func) error {
	return cv.validate.RegisterValidation(tag, fn)
}

// 自定义验证规则
func RegisterCustomRules(v Validator) {
	cv := v.(*CustomValidator)

	_ = cv.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		// 至少8个字符，包含数字和字母
		if len(password) < 8 {
			return false
		}

		hasLetter := false
		hasNumber := false
		for _, c := range password {
			switch {
			case c >= 'a' && c <= 'z':
				hasLetter = true
			case c >= 'A' && c <= 'Z':
				hasLetter = true
			case c >= '0' && c <= '9':
				hasNumber = true
			}
		}

		return hasLetter && hasNumber
	})
}
