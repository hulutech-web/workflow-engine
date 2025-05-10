package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type Phone struct{}

func (v Phone) Func() validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return regexp.MustCompile(`^1[3-9]\d{9}$`).MatchString(value)
	}
}

func (v Phone) Message() string {
	return "{0}必须是有效的手机号码"
}

func (v Phone) Tag() string {
	return "phone"
}

func init() {
	Register(Phone{})
}
