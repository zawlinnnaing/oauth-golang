package main

import (
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/user"
	"net/http"
)

func createServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/", user.Router)
	return mux
}
