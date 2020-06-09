package controllers

import (
	"github.com/bioyeneye/rest-gin-api/core/configs"
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type ApiRoutes struct {
	Root          *gin.RouterGroup // ''
	ApiRoot       *gin.RouterGroup // 'api/v1'
	Authorization *gin.RouterGroup // 'api/v1/authorize'
	Video         *gin.RouterGroup // 'api/v1/videos'
}

//reference https://github.com/mattermost/mattermost-server/blob/master/api4/api.go
type APIRoutes struct {
	BaseRoutes *ApiRoutes
	DB         *gorm.DB
}

func InitApplication(router *gin.Engine, db *gorm.DB) {
	api := &APIRoutes{
		BaseRoutes: &ApiRoutes{},
		DB:         db,
	}

	api.BaseRoutes.Root = router.Group("/")
	api.BaseRoutes.ApiRoot = router.Group(configs.API_URL_SUFFIX)

	api.BaseRoutes.Authorization = api.BaseRoutes.ApiRoot.Group("/authorize")
	api.BaseRoutes.Video = api.BaseRoutes.ApiRoot.Group("/videos", middlewares.AuthorizeMiddleware())

	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Work !",
		})
	})

	//routes
	api.InitAuthorizationRoutes()
	api.InitVideoRoutes()

	//services

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "LOST-NOT-FOUND", "message": "I think you are lost, kindly re-route your request rightly MY-GUY."})
	})
}
