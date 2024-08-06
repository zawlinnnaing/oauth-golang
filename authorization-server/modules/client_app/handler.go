package client_app

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	err := h.service.Register(body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}
