package middlewares

import (
	"net/http"

	"deva.com/backend/v2/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}
	userId, err := utils.ValidateJWT(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	context.Set("userId", userId) // Store userId in context for later use
	context.Next()
}
