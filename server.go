package main

import (
	"github.com/bioyeneye/rest-gin-api/controller"
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/bioyeneye/rest-gin-api/core/utilities"
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
		gindump.Dump())

	server.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Work !",
		})
	})

	controller.InitApplicationRoute(server)

	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "LOST-NOT-FOUND", "message": "I think you are lost, kindly re-route your request rightly MY-GUY."})
	})

	//port := os.Getenv("PORT")

	//_ = server.Run(":" + port)
	_ = server.Run(port)
}
