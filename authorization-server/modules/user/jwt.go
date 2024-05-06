package user

import (
	"errors"
	jwt "github.com/golang-jwt/jwt/v5"
	"time"
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

func SignToken(userId string, tokenType int) (*TokenResponse, error) {
	var expiresIn int64
	switch tokenType {
	case AccessToken:
		expiresIn = time.Now().Add(time.Hour * 24).Unix()
	case RefreshToken:
		expiresIn = time.Now().Add(time.Hour * 24 * 7).Unix()
	default:
		return nil, ErrInvalidTokenType
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{
		"userId": userId,
		"iss":    "simple-oauth",
		"exp":    expiresIn,
	})
	tokenStr, err := token.SigningString()
	if err != nil {
		return nil, err
	}
	return &TokenResponse{
		Token: tokenStr,
		TTL:   expiresIn,
	}, nil
}
