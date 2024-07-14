package client_app

import (
	"net/http"

	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/common"
)

type Handler struct {
	service *Service
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var body RegistrationBody
	err := common.DecodeAndValidate(w, r, &body)
	if err != nil {
		return
	}
	err = h.service.Register(w, body)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}
