package controller

import (
	"github.com/bioyeneye/rest-gin-api/core/configs"
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiRoutes struct {
	Root          *gin.RouterGroup // ''
	ApiRoot           *gin.RouterGroup // 'api/v1'
	Authorization *gin.RouterGroup // 'api/v1/authorize'
	Video         *gin.RouterGroup // 'api/v1/videos'
}

//reference https://github.com/mattermost/mattermost-server/blob/master/api4/api.go
type APIRoutes struct {
	BaseRoutes          *ApiRoutes
}

func InitApplicationRoute(router *gin.Engine) {
	api := &APIRoutes{
		BaseRoutes:          &ApiRoutes{},
	}

	api.BaseRoutes.Root = router.Group("/")
	api.BaseRoutes.ApiRoot = router.Group(configs.API_URL_SUFFIX)

	api.BaseRoutes.Authorization = api.BaseRoutes.ApiRoot.Group("/authorize")
	api.BaseRoutes.Video = api.BaseRoutes.ApiRoot.Group("/videos", middlewares.AuthorizeMiddleware())

	api.InitAuthorizationRoutes()
	api.InitVideoRoutes()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "LOST-NOT-FOUND", "message": "I think you are lost, kindly re-route your request rightly MY-GUY."})
	})
}