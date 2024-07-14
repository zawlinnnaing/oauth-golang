package client_app

import "net/http"

func Router(w http.ResponseWriter, r *http.Request) {
	handler := NewHandler(NewService(NewRepository()))
	if r.Method == "POST" && r.URL.Path == "/client-apps/register" {
		handler.handleRegister(w, r)
		return
	}
}
