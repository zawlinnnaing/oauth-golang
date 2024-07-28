package user

import "github.com/gin-gonic/gin"

func Router(router *gin.Engine) {
	userRouter := router.Group("/users")
	handler := NewHandler(NewService())
	userRouter.POST("/sign-up", handler.HandleSignUp)
	userRouter.POST("/sign-in", handler.HandleSignIn)
}
