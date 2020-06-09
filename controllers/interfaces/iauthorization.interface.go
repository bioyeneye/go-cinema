package interfaces

import (
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/bioyeneye/rest-gin-api/core/utilities"
	"github.com/gin-gonic/gin"
)

type IAuthorizationController interface {
	Token(ctx *gin.Context) (utilities.AccessTokenResponse, error)
	CurrentUser(ctx *gin.Context) middlewares.CurrentUser
}

