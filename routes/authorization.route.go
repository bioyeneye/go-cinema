package routes

import (
	"github.com/bioyeneye/rest-gin-api/controller"
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
			token := authorizationController.Token(ctx)
			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, nil)
			}
		}),
	}
}
