package client_app

import "github.com/zawlinnnaing/oauth-golang/authorization-server/modules/user"

type Service struct {
	repo *Repository
}

func (s *Service) Register(body RegistrationBody, user user.User) (*ClientApp, error) {
	clientApp, err := s.repo.Register(body, user)
	return clientApp, err
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}
