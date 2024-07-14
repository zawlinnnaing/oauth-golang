package common

import (
	"encoding/json"
	"net/http"
)

type RequestBody interface {
	Validate() error
}

func DecodeAndValidate(w http.ResponseWriter, r *http.Request, body RequestBody) error {
	err := json.NewDecoder(r.Body).Decode(body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "failed-to-parse-request", http.StatusBadRequest)
		return err
	}
	err = body.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return err
	}
	return nil
}
