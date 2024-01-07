package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func (j *jwtService) GetUserId(token string) (string, error) {
	aToken, err := j.ValidateToken(token)
	if err != nil {
		return "", err
	}

	claims := aToken.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["userId"])

	return userId, nil
}
