package validator

import (
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	valiRules "github.com/hulutech-web/workflow-engine/pkg/validate"
	"go.uber.org/zap"
	"reflect"
)

type Service struct {
	instance *validator.Validate
	trans    unTrans.Translator
}

func NewService() *Service {
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		zap.S().Error("register default translation failed", zap.Error(err))
		return nil
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	srv := &Service{
		instance: validate,
		trans:    trans,
	}

	// 加载验证插件
	for _, plugin := range valiRules.Plugins {
		if err := srv.RegisterValidation(plugin.Tag(), plugin.Func(), plugin.Message()); err != nil {
			zap.S().Errorf("注册验证插件失败: %s", plugin.Tag())
		}
	}

	return srv
}

// 新增自定义验证注册方法
func (s *Service) RegisterValidation(tag string, fn validator.Func, translation string) error {
	// 注册验证函数
	if err := s.instance.RegisterValidation(tag, fn); err != nil {
		return err
	}

	// 注册验证翻译
	return s.instance.RegisterTranslation(tag, s.trans, func(ut unTrans.Translator) error {
		return ut.Add(tag, translation, true)
	}, func(ut unTrans.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
}

// 验证对象
func (s *Service) Validate(obj interface{}) (msg []string) {
	err := s.instance.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg = append(msg, err.Translate(s.trans))
		}
	}
	return msg
}
