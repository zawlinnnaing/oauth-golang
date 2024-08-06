package client_app

import (
	"github.com/gin-gonic/gin"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/middlewares"
)

func Router(router *gin.Engine) {
	clientAppRouter := router.Group("/client-apps")
	clientAppRouter.Use(middlewares.Authenticated())
	handler := NewHandler(NewService(NewRepository()))
	clientAppRouter.POST("/register", handler.handleRegister)
}
