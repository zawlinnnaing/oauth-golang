package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/client_app"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository          *Repository
	clientAppRepository *client_app.Repository
}

const passwordCost = 14

var (
	ErrUserAlreadyExists = errors.New("user.already-exists")
	ErrUserNotFound      = errors.New("user.not-found")
	ErrInvalidPassword   = errors.New("user.invalid-password")
)

func (service *Service) SignUp(body *SignUpBody) (*User, error) {
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

func (service *Service) SignIn(body *SignInBody) (*SignInResponse, error) {
	user, err := service.repository.FindByEmail(body.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return nil, ErrInvalidPassword
	}
	tokenResult, err := SignToken(user.ID, AccessToken)
	if err != nil {
		return nil, err
	}
	refreshToken, err := SignToken(user.ID, RefreshToken)
	if err != nil {
		return nil, err
	}
	return &SignInResponse{
		RefreshToken: *refreshToken,
		Token:        *tokenResult,
	}, nil
}

func (service *Service) GrantAccessUI(context *gin.Context) {
	redirectURI := context.Query("redirect_uri")
	clientID := context.Query("client_id")
	if redirectURI == "" || clientID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "redirect_uri and client_id are required"})
		return
	}
	clientApp, err := service.clientAppRepository.FindByID(clientID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if clientApp.RedirectURI != redirectURI {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Redirect URI does not match with registered app.",
		})
	}
	context.HTML(http.StatusOK, "grant-access.tmpl", gin.H{
		"client_id":    clientApp.ID,
		"redirect_uri": clientApp.RedirectURI,
	})

}

func NewService() *Service {
	return &Service{
		repository:          NewRepository(),
		clientAppRepository: client_app.NewRepository(),
	}
}
