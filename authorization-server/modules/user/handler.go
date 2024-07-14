package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/app_error"
)

func handleSignIn(w http.ResponseWriter, r *http.Request) {
	var body SignInBody
	err := json.NewDecoder(r.Body).Decode(&body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		app_error.HTTPError(w, "failed-to-parse-request", http.StatusBadRequest, err)
		return
	}
	err = body.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	resp, err := NewService().SignIn(&body)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, ErrUserNotFound) || errors.Is(err, ErrInvalidPassword) {
			status = http.StatusBadRequest
		}
		app_error.HTTPError(w, err.Error(), status, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	var body SignUpBody
	err := json.NewDecoder(r.Body).Decode(&body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		app_error.HTTPError(w, "failed-to-parse-request", http.StatusBadRequest, err)
		return
	}
	if err != nil {
		app_error.HTTPError(w, "internal-server-error", http.StatusInternalServerError, err)
		return
	}
	err = body.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			app_error.HTTPError(w, "internal-server-error", http.StatusInternalServerError, err)
		}
		return
	}
	user, err := NewService().SignUp(&body)
	if err != nil {
		app_error.HTTPError(w, err.Error(), http.StatusInternalServerError, err)
		return
	}
	resp, err := json.Marshal(user)
	if err != nil {
		app_error.HTTPError(w, "internal-server-error", http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func Router(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && r.URL.Path == "/users/sign-in" {
		handleSignIn(w, r)
		return
	}
	if r.Method == "POST" && r.URL.Path == "/users/sign-up" {
		handleSignUp(w, r)
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}
