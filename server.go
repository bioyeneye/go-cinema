package main

import (
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/bioyeneye/rest-gin-api/routes"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"net/http"
	"os"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	const port string = ":8000"

	setupLogOutput()
	server := gin.New()
	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.CORSMiddleware(),
		middlewares.BasicAuthentication(),
		gindump.Dump())

	server.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Work !",
		})
	})

	noAuthorizeGroup := server.Group("")
	routes.SetAuthorizationRoutes(noAuthorizeGroup)

	//remove for routes that don't need to be authorized
	version1Api := server.Group("/v1/api", middlewares.AuthorizeMiddleware())
	routes.SetVideoRoutes(version1Api)

	//port := os.Getenv("PORT")

	//_ = server.Run(":" + port)
	_ = server.Run(port)
}
