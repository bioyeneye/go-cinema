package main

import (
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/bioyeneye/rest-gin-api/core/utilities"
	"github.com/bioyeneye/rest-gin-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"log"
	"net/http"
	"os"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	const port string = ":8000"

	//configure db
	dbconfig := utilities.NewDatabaseConfig()
	log.Println(dbconfig.Database)

	setupLogOutput()


	server := gin.New()
	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.CORSMiddleware(),
		middlewares.ContentTypeMiddleware(),
		//middlewares.BasicAuthentication(),
		gindump.Dump())

	server.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Work !",
		})
	})

	noAuthorizeGroup := server.Group("", middlewares.AuthorizeMiddleware())
	routes.SetAuthorizationRoutes(noAuthorizeGroup)

	//remove for routes that don't need to be authorized
	version1Api := server.Group("/v1/api", middlewares.AuthorizeMiddleware())
	routes.SetVideoRoutes(version1Api)

	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "LOST-NOT-FOUND", "message": "I think you are lost, kindly re-route your request rightly MY-GUY."})
	})

	//port := os.Getenv("PORT")

	//_ = server.Run(":" + port)
	_ = server.Run(port)
}
