package client_app

import (
	"time"

	"github.com/google/uuid"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/database"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/user"
)

type Repository struct{}

func (r *Repository) Register(body RegistrationBody, u user.User) (*ClientApp, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	clientApp := &ClientApp{
		ID:          id.String(),
		Name:        body.Name,
		RedirectURI: body.RedirectURI,
		User:        u,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	result := database.DB.Create(clientApp)
	if result.Error != nil {
		return nil, result.Error
	}
	return clientApp, nil
}

func (r *Repository) FindByID(id string) (*ClientApp, error) {
	var clientApp ClientApp
	result := database.DB.Where("id = ?", id).First(&clientApp)
	if result.Error != nil {
		return nil, result.Error
	}
	return &clientApp, nil
}

func NewRepository() *Repository {
	return &Repository{}
}
