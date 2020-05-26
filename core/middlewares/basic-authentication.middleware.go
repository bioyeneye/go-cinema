package middlewares

import "github.com/gin-gonic/gin"

func BasicAuthentication() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"pragmatic":"reviews",
	})
}