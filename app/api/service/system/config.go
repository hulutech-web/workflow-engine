package system

import (
	"errors"
	"fmt"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"gorm.io/gorm"
)

type ConfigService interface {
	Get(tenantId uint, cnfType string, names ...string) (data map[string]string, err error)
	GetVal(tenantId uint, cnfType string, name string, defaultValue ...string) (data string, err error)
	GetMap(tenantId uint, cnfType string, name string) (data map[string]string, err error)
	Set(tenantId uint, cnfType string, name, value string) (err error)
}

type configServiceImpl struct {
	db *gorm.DB
}

func (c configServiceImpl) Get(tenantId uint, cnfType string, names ...string) (data map[string]string, err error) {
	chain := c.db.Model(&models.Config{}).Where("type = ?", cnfType)
	if tenantId > 0 {
		chain = chain.Where("tenant_id = ?", tenantId)
	}
	if len(names) > 0 {
		chain = chain.Where("name in (?)", names)
	}
	var configs []models.Config
	if err = chain.Find(&configs).Error; err != nil {
		return nil, fmt.Errorf("获取配置失败:%s", err.Error())
	}
	data = make(map[string]string)
	for i := 0; i < len(configs); i++ {
		data[configs[i].Name] = configs[i].Value
	}
	return data, nil
}

func (c configServiceImpl) GetVal(tenantId uint, cnfType string, name string, defaultValue ...string) (data string, err error) {
	conf, err := c.Get(tenantId, cnfType, name)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return "", err
	}
	data, ok := conf[name]
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return "", fmt.Errorf("未找到配置项:%s", name)
	}
	return data, nil
}

func (c configServiceImpl) GetMap(tenantId uint, cnfType string, name string) (data map[string]string, err error) {
	conf, err := c.GetVal(tenantId, cnfType, name, "")
	if err != nil {
		return nil, err
	}
	if conf == "" {
		return map[string]string{}, nil
	}
	err = util.ToolsUtil.JsonToObj(conf, &data)
	if err != nil {
		return nil, fmt.Errorf("配置项:%s 格式错误:%s", name, err.Error())
	}
	return data, nil
}

func (c configServiceImpl) Set(tenantId uint, cnfType string, name, value string) (err error) {
	var conf models.Config
	chain := c.db.Model(&models.Config{}).Where("type = ?", cnfType).Where("name = ?", name)
	if tenantId > 0 {
		chain = chain.Where("tenant_id = ?", tenantId)
	}
	err = chain.First(&conf).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if err = c.db.Create(&conf).Error; err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	if err = c.db.Model(&conf).Update("value", value).Error; err != nil {
		return err
	}
	return nil
}

func NewConfigService(db *gorm.DB) ConfigService {
	return &configServiceImpl{db: db}
}
