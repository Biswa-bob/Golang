package middlewares

import (
	"net/http"

	"example.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "User is not Authorized",
		})

	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "User is not Authorized",
		})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
