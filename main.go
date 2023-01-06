package main

import (
	"os"
	"todo/initializers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	initializers.LoadEnvVariables()
	initializers.InitRoutes(router) 

	port := os.Getenv("PORT")
	
	router.Run("localhost:" + port)
}