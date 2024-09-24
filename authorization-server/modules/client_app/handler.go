package client_app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/user"
)

type Handler struct {
	service *Service
}

func (h *Handler) handleRegister(context *gin.Context) {
	var body RegistrationBody
	if err := context.ShouldBind(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authUser, exists := context.Get("user")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "user.does-not-exist"})
		return
	}
	clientApp, err := h.service.Register(body, authUser.(user.User))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": clientApp})
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}
