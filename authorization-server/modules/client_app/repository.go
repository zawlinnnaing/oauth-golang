package client_app

import (
	"time"

	"github.com/google/uuid"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/common/models"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/database"
)

type Repository struct{}

func (r *Repository) Register(body RegistrationBody, userId string) (*models.ClientApp, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	clientApp := &models.ClientApp{
		ID:          id.String(),
		Name:        body.Name,
		RedirectURI: *body.RedirectURI,
		UserID:      userId,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	result := database.DB.Create(clientApp)
	if result.Error != nil {
		return nil, result.Error
	}
	return clientApp, nil
}

func (r *Repository) FindByID(id string) (*models.ClientApp, error) {
	var clientApp models.ClientApp
	result := database.DB.Where("id = ?", id).First(&clientApp)
	if result.Error != nil {
		return nil, result.Error
	}
	return &clientApp, nil
}

func NewRepository() *Repository {
	return &Repository{}
}
