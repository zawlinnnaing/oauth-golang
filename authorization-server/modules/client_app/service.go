package client_app

import "github.com/zawlinnnaing/oauth-golang/authorization-server/modules/common/models"

type Service struct {
	repo *Repository
}

func (s *Service) Register(body RegistrationBody, userId string) (*models.ClientApp, error) {
	clientApp, err := s.repo.Register(body, userId)
	return clientApp, err
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}
