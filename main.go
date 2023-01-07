package main

import (
	"log"
	"os"
	"todo/db"
	"todo/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitRoutes(router *gin.Engine) {
	routers.SetTodoRoutes(router)
}

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()

	LoadEnvVariables()
	InitRoutes(router)
	db.ConnectToDataBase()
	db.SyncDB()

	port := os.Getenv("PORT")
	router.Run("localhost:" + port)
}