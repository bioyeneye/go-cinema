package controller

import (
	"github.com/bioyeneye/rest-gin-api/controller/interfaces"
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/bioyeneye/rest-gin-api/core/utilities"
	"github.com/gin-gonic/gin"
)

type AuthorizationController struct {
	jWtService   utilities.IJWTService
}




func NewAuthorizationController(jWtService utilities.IJWTService) interfaces.IAuthorizationController {
	return &AuthorizationController {
		jWtService:   jWtService,
	}
}

func (controller *AuthorizationController) Token(ctx *gin.Context) (utilities.AccessTokenResponse, error) {
	isAuthenticated := true
	if isAuthenticated {
		return controller.jWtService.GenerateToken("bolaji", true)
	}

	return utilities.AccessTokenResponse{}, nil
}

func (controller *AuthorizationController) CurrentUser(ctx *gin.Context) middlewares.CurrentUser {

	currentuser, exists := utilities.GetCurrentUser(ctx)
	if !exists {
		return middlewares.CurrentUser{}
	}

	return currentuser
}









