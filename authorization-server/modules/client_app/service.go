package client_app

type Service struct {
	repo *Repository
}

func (s *Service) Register(body RegistrationBody) error {
	// _, err := s.repo.Register(body)
	return nil
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}
