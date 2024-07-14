package client_app

import (
	"time"

	"github.com/google/uuid"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/database"
)

type Repository struct{}

func (r *Repository) Register(body RegistrationBody) (*ClientApp, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	clientApp := &ClientApp{
		ID:          id.String(),
		Name:        body.Name,
		RedirectURI: body.RedirectURI,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	result := database.DB.Create(clientApp)
	if result.Error != nil {
		return nil, result.Error
	}
	return clientApp, nil
}

func NewRepository() *Repository {
	return &Repository{}
}
