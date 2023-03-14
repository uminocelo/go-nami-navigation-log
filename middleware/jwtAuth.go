package middleware

import (
	"nami_navigation_log/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helpers.ValidateJWToken(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}

		context.Next()
	}
}
