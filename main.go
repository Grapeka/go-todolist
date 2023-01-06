package main

import (
	"todo/initializers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	initializers.InitRoutes(router)
	router.Run("localhost:9090")
}