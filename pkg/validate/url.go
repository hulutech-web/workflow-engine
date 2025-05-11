package validate

import (
	"github.com/go-playground/validator/v10"
	"net/url"
)

type Url struct{}

func (u Url) Tag() string {
	return "url"
}

func (u Url) Func() validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		// 添加URL验证逻辑
		parsedURL, err := url.Parse(value)
		if err != nil {
			return false
		}
		// 必须包含协议和域名
		if parsedURL.Scheme == "" || parsedURL.Host == "" {
			return false
		}
		// 可选：限制为http/https协议（按需添加）
		if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
			return false
		}
		return true
	}
}

func (u Url) Message() string {
	return "{0}不是有效的URL"
}

func init() {
	Register(Url{})
}
