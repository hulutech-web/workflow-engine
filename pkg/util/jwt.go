package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hulutech-web/workflow-engine/core/config"
	"time"
)

type jwtUtil struct{}

var (
	JwtUtil   = jwtUtil{}
	conf      = config.NewConfig()
	jwtSecret = []byte(conf.Jwt.Secret)
)

type Claims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func (jwtUtil) GenerateToken(userId uint) (string, error) {
	claims := Claims{
		userId,
		jwt.RegisteredClaims{
			Issuer:    "tinyGW",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (jwtUtil) ParseToken(tokenString string) (*Claims, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("解析错误！")
	}
}
