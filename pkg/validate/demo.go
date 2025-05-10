package validate

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Demo struct{}

func (d Demo) Tag() string {
	return "demo"
}

func (d Demo) Func() validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		if len(value) == 0 {
			return false
		}
		zap.S().Info("Validating demo field", zap.String("value", value))
		return true
	}
}

func (d Demo) Message() string {
	return "{0} is not valid"
}

func init() {
	Register(Demo{})
}
