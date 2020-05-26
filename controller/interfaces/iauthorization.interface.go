package interfaces

import "github.com/gin-gonic/gin"

type IAuthorizationController interface {
	Token(ctx *gin.Context) string
}

