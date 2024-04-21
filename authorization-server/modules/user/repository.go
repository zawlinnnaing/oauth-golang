package user

import (
	"errors"
	"github.com/google/uuid"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
}

func (repo *Repository) FindByEmail(email string) (*User, error) {
	user := &User{}
	result := database.DB.Where("email = ?", email).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func (repo *Repository) Create(body SignUpBody) (*User, error) {
	userID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), passwordCost)
	if err != nil {
		return nil, err
	}
	user := &User{
		ID:        userID.String(),
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  string(hashedPassword),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	result := database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func NewRepository() *Repository {
	return &Repository{}
}
