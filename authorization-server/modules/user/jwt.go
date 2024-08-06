package user

import (
	"errors"
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

var signingMethod = jwt.SigningMethodHS512

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
	token := jwt.NewWithClaims(signingMethod, jwt.MapClaims{
		"sub": userId,
		"iss": "simple-oauth",
		"exp": expiresIn,
	})
	tokenStr, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return nil, err
	}
	return &TokenResponse{
		Token: tokenStr,
		TTL:   expiresIn,
	}, nil
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// fmt.Println("token", token.Method)
		// if token.Method != signingMethod {
		// 	return nil, fmt.Errorf("unexpected Signing method: %v", token.Method.Alg())
		// }
		return config.JWT_SECRET, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
