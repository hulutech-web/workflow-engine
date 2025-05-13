package validate

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type TemplateForm struct{}

func (d TemplateForm) Tag() string {
	return "template_form"
}

func (d TemplateForm) Func() validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		if len(value) == 0 {
			return false
		}
		zap.S().Info("Validating TemplateForm field", zap.String("value", value))
		return true
	}
}

func (d TemplateForm) Message() string {
	return "{0} is not valid"
}

func init() {
	Register(TemplateForm{})
}
