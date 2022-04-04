package middleware

import (
	"github.com/gin-gonic/gin"
	"goGinDemo/config"
	"net/http"
	"strings"
)

func getToken(ctx *gin.Context) (string, bool) {
	authValue := ctx.GetHeader("Authorization")
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}
	authType := strings.Trim(arr[0], "\n\r\t")
	if strings.ToLower(authType) != strings.ToLower("Bearer") {
		return "", false
	}
	return strings.Trim(arr[1], "\n\t\r"), true
}

func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, ok := getToken(ctx)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		id, username, err := config.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		ctx.Set("id", id)
		ctx.Set("username", username)
		ctx.Writer.Header().Set("Authorization", "Bearer "+token)
		ctx.Next()
	}
}
