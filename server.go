package main

import (
	"github.com/bioyeneye/rest-gin-api/controller"
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/bioyeneye/rest-gin-api/entities"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"log"
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
	setupLogOutput()

	//configure db
	db := entities.SetupModels()

	server := gin.New()
	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.CORSMiddleware(),
		middlewares.ContentTypeMiddleware(),
		gindump.Dump())



	controller.InitApplication(server, db)

	//port := os.Getenv("PORT")

	//_ = server.Run(":" + port)
	_ = server.Run(port)
}
