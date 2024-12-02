package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func (h *Handler) HandleSignUp(context *gin.Context) {
	var body SignUpBody
	if err := context.ShouldBind(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.Service.SignUp(&body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, user)
}

func (h *Handler) HandleSignIn(context *gin.Context) {
	var body SignInBody
	if err := context.ShouldBind(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Service.SignIn(&body)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, ErrUserNotFound) || errors.Is(err, ErrInvalidPassword) {
			status = http.StatusBadRequest
		}
		context.JSON(status, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, resp)
}

func (h *Handler) HandleGrantAccessUI(context *gin.Context) {
	h.Service.GrantAccessUI(context)
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}
