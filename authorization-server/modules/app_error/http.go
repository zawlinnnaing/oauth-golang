package app_error

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func HTTPError(w http.ResponseWriter, key string, status int, error error) {
	errorMsg := Message{Message: key, Error: error.Error()}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(errorMsg)
	if err != nil {
		http.Error(w, "Write JSON Error", http.StatusInternalServerError)
		return
	}
}
