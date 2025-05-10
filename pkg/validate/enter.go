package validate

import "github.com/go-playground/validator/v10"

type ValidatorPlugin interface {
	Tag() string
	Func() validator.Func
	Message() string
}

var Plugins []ValidatorPlugin

func Register(plugin ValidatorPlugin) {
	Plugins = append(Plugins, plugin)
}
