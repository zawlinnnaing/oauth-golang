package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/config"
)

var signingMethod = jwt.SigningMethodHS256
var jwtSecret = []byte(config.JWT_SECRET)

type AuthService struct{}

func (a *AuthService) ValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method != signingMethod {
			return nil, fmt.Errorf("unexpected Signing method: %v", token.Method.Alg())
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func NewAuthService() *AuthService {
	return &AuthService{}
}
