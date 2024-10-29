package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/event-booking/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized request"})
		return
	}

	customerId, error := utils.VerifyToken(token)
	if error != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized request"})
		return
	}

	context.Set("customerId", customerId)
	context.Next()
}
