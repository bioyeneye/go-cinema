package routes

import (
	"github.com/bioyeneye/rest-gin-api/controller"
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/bioyeneye/rest-gin-api/core/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

//routes
func SetAuthorizationRoutes(router *gin.RouterGroup) []gin.IRoutes  {

	var jwtService = utilities.NewJWTService()
	var authorizationController = controller.NewAuthorizationController(jwtService)

	return []gin.IRoutes {
		router.POST("/token", func(ctx *gin.Context) {
			token, _ := authorizationController.Token(ctx)
			if token != (utilities.AccessTokenResponse{}) {
				ctx.JSON(http.StatusOK, gin.H{
					"access_token": token.AccessToken,
					"refresh_token": token.RefreshToken,
					"token_type": token.TokenType,
					"expires_in": token.ExpiresIn,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, nil)
			}
		}),
		router.GET("/currentuser", func(ctx *gin.Context) {
			currentUser := authorizationController.CurrentUser(ctx)
			if currentUser != (middlewares.CurrentUser{}) {
				ctx.JSON(http.StatusOK, gin.H{
					"currentuser": currentUser,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, nil)
			}
		}),
	}
}