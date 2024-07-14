package main

import (
	"net/http"

	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/client_app"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/user"
)

func createServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/", user.Router)
	mux.HandleFunc("/client-apps/", client_app.Router)
	return mux
}
