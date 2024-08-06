package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func handleSignIn(w http.ResponseWriter, r *http.Request) {
// 	var body SignInBody
// 	err := json.NewDecoder(r.Body).Decode(&body)
// 	w.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		app_error.HTTPError(w, "failed-to-parse-request", http.StatusBadRequest, err)
// 		return
// 	}
// 	err = body.Validate()
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}
// 	resp, err := NewService().SignIn(&body)
// 	if err != nil {
// 		status := http.StatusInternalServerError
// 		if errors.Is(err, ErrUserNotFound) || errors.Is(err, ErrInvalidPassword) {
// 			status = http.StatusBadRequest
// 		}
// 		app_error.HTTPError(w, err.Error(), status, err)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(resp)
// }

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

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}
