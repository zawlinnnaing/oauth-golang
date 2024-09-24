package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/user"
)

func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRepo := user.NewRepository()
		authorization := ctx.GetHeader("Authorization")
		authorization = strings.ReplaceAll(authorization, "Bearer ", "")
		authorization = strings.TrimSpace(authorization)
		fmt.Println("Auth Token:", authorization)
		token, err := user.ValidateToken(authorization)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		userId, err := token.Claims.GetSubject()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to parse userId",
			})
			ctx.Abort()
			return
		}
		user, err := userRepo.FindByID(userId)
		if err != nil {
			ctx.JSON(401, gin.H{
				"message": "User not found",
			})
			return
		}
		ctx.Set("user", *user)
		ctx.Next()
	}
}
