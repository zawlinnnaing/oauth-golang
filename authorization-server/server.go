package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/client_app"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/user"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/validators"
)

func createServer() *gin.Engine {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validators.PasswordValidator)
	}
	user.Router(router)
	client_app.Router(router)
	return router
}
