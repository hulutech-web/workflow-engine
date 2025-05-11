package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type Domain struct{}

func (d Domain) Tag() string {
	return "domain"
}

func (d Domain) Func() validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		reg := `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}$`
		return regexp.MustCompile(reg).MatchString(value)
	}
}

func (d Domain) Message() string {
	return "{0}不是有效的域名"
}

func init() {
	Register(Domain{})
}
