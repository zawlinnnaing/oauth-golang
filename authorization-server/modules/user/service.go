package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
}

const passwordCost = 14

var (
	ErrUserAlreadyExists = errors.New("user.already-exists")
	ErrUserNotFound      = errors.New("user.not-found")
	ErrInvalidPassword   = errors.New("user.invalid-password")
)

func (service *Service) SignUp(body *SignUpBody) (*User, error) {
	existingUser, err := service.repository.FindByEmail(body.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}
	user, err := service.repository.Create(*body)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *Service) SignIn(body *SignInBody) (*SignInResponse, error) {
	user, err := service.repository.FindByEmail(body.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return nil, ErrInvalidPassword
	}
	tokenResult, err := SignToken(user.ID, AccessToken)
	if err != nil {
		return nil, err
	}
	refreshToken, err := SignToken(user.ID, RefreshToken)
	if err != nil {
		return nil, err
	}
	return &SignInResponse{
		RefreshToken: *refreshToken,
		Token:        *tokenResult,
	}, nil
}

func NewService() *Service {
	return &Service{
		repository: NewRepository(),
	}
}
