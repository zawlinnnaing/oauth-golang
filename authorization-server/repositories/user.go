package repositories

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/database"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/dtos"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const passwordCost = 14

type UserRepository struct {
}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	result := database.DB.Where("email = ?", email).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) Create(body dtos.SignUpBody) (*models.User, error) {
	userID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), passwordCost)
	if err != nil {
		return nil, err
	}
	user := &models.User{
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

func (repo *UserRepository) FindByID(id string) (*models.User, error) {
	user := &models.User{}
	result := database.DB.Where("id = ?", id).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
