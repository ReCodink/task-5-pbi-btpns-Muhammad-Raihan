package middlewares

import (
	"net/http"
	"strings"

	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/helpers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: Missing Authorization Header",
			})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is empty"})
			ctx.Abort()
			return
		}

		issuer, err := helpers.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("userID", issuer.Id)
		ctx.Next()
	}
}
