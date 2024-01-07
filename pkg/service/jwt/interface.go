package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetUserId(token string) (string, error)
}
