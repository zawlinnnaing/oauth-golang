package client_app

import (
	"encoding/json"
	"net/http"
)

type Service struct {
	repo *Repository
}

func (s *Service) Register(w http.ResponseWriter, body RegistrationBody) error {
	err := body.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return err
	}
	_, err = s.repo.Register(body)
	return err
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}
