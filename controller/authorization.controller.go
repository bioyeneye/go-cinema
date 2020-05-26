package controller

import (
	"github.com/bioyeneye/rest-gin-api/controller/interfaces"
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

func (controller *AuthorizationController) Token(ctx *gin.Context) string {
	isAuthenticated := true
	if isAuthenticated {
		return controller.jWtService.GenerateToken("bolaji", true)
	}
	return ""
}







