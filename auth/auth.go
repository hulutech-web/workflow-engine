package auth

import (
	"errors"
	"time"
)

type AuthManager interface {
	Authenticate(token string) (User, error)
	Login(credentials Credentials) (*TokenPair, error)
	Refresh(refreshToken string) (*TokenPair, error)
	Logout(token string) error
}

type User interface {
	GetID() uint
	GetEmail() string
	Can(permission string) bool
}

type Credentials struct {
	Email    string
	Password string
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

type JwtConfig struct {
	Secret        string
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
	TokenRotation bool
}

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrUserNotFound = errors.New("user not found")
)
