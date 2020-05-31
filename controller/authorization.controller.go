package controller

import (
	"github.com/bioyeneye/rest-gin-api/controller/interfaces"
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/bioyeneye/rest-gin-api/core/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthorizationController struct {
	jWtService utilities.IJWTService
}

func (api *APIRoutes) InitAuthorizationRoutes() {
	var jwtService = utilities.NewJWTService()
	var authorizationController = NewAuthorizationController(jwtService)

	api.BaseRoutes.Authorization.POST("/token", func(ctx *gin.Context) {
		token, _ := authorizationController.Token(ctx)
		if token != (utilities.AccessTokenResponse{}) {
			ctx.JSON(http.StatusOK, gin.H{
				"access_token":  token.AccessToken,
				"refresh_token": token.RefreshToken,
				"token_type":    token.TokenType,
				"expires_in":    token.ExpiresIn,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	api.BaseRoutes.Authorization.GET("/currentuser", func(ctx *gin.Context) {
		currentUser := authorizationController.CurrentUser(ctx)
		if currentUser != (middlewares.CurrentUser{}) {
			ctx.JSON(http.StatusOK, gin.H{
				"currentuser": currentUser,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

}

func NewAuthorizationController(jWtService utilities.IJWTService) interfaces.IAuthorizationController {
	return &AuthorizationController{
		jWtService: jWtService,
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
