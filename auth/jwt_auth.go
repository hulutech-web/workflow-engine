package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis"
	"time"
)

type JwtAuth struct {
	redis    *redis.Client
	config   JwtConfig
	userRepo UserRepository
}

type UserRepository interface {
	FindByEmail(email string) (User, error)
	FindByID(id uint) (User, error)
}

func NewJwtAuth(redis *redis.Client, config JwtConfig, userRepo UserRepository) *JwtAuth {
	return &JwtAuth{
		redis:    redis,
		config:   config,
		userRepo: userRepo,
	}
}

func (a *JwtAuth) Authenticate(tokenString string) (User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 检查Redis黑名单
		if exists, _ := a.redis.Exists(ctx, tokenString).Result(); exists > 0 {
			return nil, ErrInvalidToken
		}

		userID := uint(claims["sub"].(float64))
		return a.userRepo.FindByID(userID)
	}

	return nil, ErrInvalidToken
}

func (a *JwtAuth) Login(creds Credentials) (*TokenPair, error) {
	user, err := a.userRepo.FindByEmail(creds.Email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// 验证密码...

	accessToken, err := a.createToken(user.GetID(), a.config.AccessExpiry)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.createToken(user.GetID(), a.config.RefreshExpiry)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(a.config.AccessExpiry),
	}, nil
}

func (a *JwtAuth) createToken(userID uint, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(expiry).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.config.Secret))
}

// 其他方法实现...
