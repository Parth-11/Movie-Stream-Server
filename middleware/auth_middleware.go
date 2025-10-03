package middleware

import (
	"net/http"

	"github.com/Parth-11/Movie-Stream-Server/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token, err := utils.GetAccessToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
		}

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token is provided"})
			ctx.Abort()
		}

	}
}
