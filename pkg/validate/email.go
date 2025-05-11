package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type Email struct{}

func (e Email) Tag() string {
	return "email"
}

func (e Email) Func() validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		reg := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		return regexp.MustCompile(reg).MatchString(value)
	}
}

func (e Email) Message() string {
	return "{0}不是有效的邮箱地址"
}

func init() {
	Register(Email{})
}
