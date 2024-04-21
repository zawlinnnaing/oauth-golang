package user

import (
	"errors"
)

type Service struct {
	repository *Repository
}

const passwordCost = 14

var (
	ErrUserAlreadyExists = errors.New("user.already-exists")
)

func (service *Service) SignUp(body *SignUpBody) (*User, error) {
	err := body.Validate()
	if err != nil {
		return nil, err
	}
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

func NewService() *Service {
	return &Service{
		repository: NewRepository(),
	}
}
