package service

import (
	"errors"
	"fmt"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/resp"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/core/config"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(loginReq *req.AuthLoginReq) (*resp.AuthLoginResp, error)
	Info(token string) (*resp.AuthLoginInfoResp, error)
	Logout(token string) error
	RefreshToken(token string) (string, error)
}

type authService struct {
	db  *gorm.DB
	cfg *config.Config
}

func (a authService) Login(loginReq *req.AuthLoginReq) (*resp.AuthLoginResp, error) {
	var user models.User
	if err := a.db.Where("username = ?", loginReq.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户名或密码错误")
		}
	}
	pwd := util.ToolsUtil.MakeMd5(loginReq.Password)
	if pwd != loginReq.Password {
		return nil, fmt.Errorf("用户名或密码错误")
	}
	token, err := util.JwtUtil.GenerateToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("生成token失败")
	}
	var res resp.AuthLoginResp
	res.Token = token
	var info resp.AuthLoginInfoResp
	response.Copy(&info, user)
	res.Info = info
	return &res, nil
}

func (a authService) Logout(token string) error {
	return nil
}

func (a authService) RefreshToken(token string) (string, error) {
	claims, err := util.JwtUtil.ParseToken(token)
	if err != nil {
		return "", fmt.Errorf("token解析失败")
	}
	userId := claims.ID
	newToken, err := util.JwtUtil.GenerateToken(userId)
	if err != nil {
		return "", fmt.Errorf("生成token失败")
	}
	return newToken, nil
}

func (a authService) Info(token string) (*resp.AuthLoginInfoResp, error) {
	claims, err := util.JwtUtil.ParseToken(token)
	if err != nil {
		return nil, fmt.Errorf("token解析失败")
	}
	userId := claims.ID
	var user models.User
	if err := a.db.First(&user, userId).Error; err != nil {
		return nil, fmt.Errorf("用户不存在")
	}
	var res resp.AuthLoginInfoResp
	response.Copy(&res, user)
	return &res, nil
}

func NewAuthService(db *gorm.DB, cfg *config.Config) AuthService {
	return &authService{db: db, cfg: cfg}
}
