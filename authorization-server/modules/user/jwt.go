package user

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/config"
)

const (
	AccessToken = iota
	RefreshToken
)

var (
	ErrInvalidTokenType = errors.New("invalid token type")
)

type TokenOptions struct {
	expiresIn int64
}

type TokenResponse struct {
	Token string `json:"token"`
	TTL   int64  `json:"ttl"`
}

var signingMethod = jwt.SigningMethodHS256
var jwtSecret = []byte(config.JWT_SECRET)

func SignToken(userId string, tokenType int) (*TokenResponse, error) {
	var expiresIn time.Time
	switch tokenType {
	case AccessToken:
		expiresIn = time.Now().Add(time.Hour * 24)
	case RefreshToken:
		expiresIn = time.Now().Add(time.Hour * 24 * 7)
	default:
		return nil, ErrInvalidTokenType
	}
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiresIn),
		Subject:   userId,
		Issuer:    "simple-oauth",
	}
	token := jwt.NewWithClaims(signingMethod, claims)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}
	return &TokenResponse{
		Token: tokenStr,
		TTL:   expiresIn.Unix(),
	}, nil
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
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
