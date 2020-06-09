package main

import (
	"fmt"
	"github.com/bioyeneye/rest-gin-api/controllers"
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/bioyeneye/rest-gin-api/core/utilities"
	"github.com/bioyeneye/rest-gin-api/db"
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
	//note: use utiilites NewDBConfig
	dbConfig := utilities.NewDBConfigFromEnv()
	dbConString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Name, dbConfig.Password)
	dbInstance, err := db.SetupDbModels("postgres", dbConString)
	if err != nil {
		panic("Failed to connect to database!")
	}

	server := gin.New()
	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.CORSMiddleware(),
		middlewares.ContentTypeMiddleware(),
		gindump.Dump())

	controllers.InitApplication(server, dbInstance)
	_ = server.Run(port)
}
