package initializers

import (
	"todo/routers"

	"github.com/gin-gonic/gin"
)


func InitRoutes(router *gin.Engine) {
	routers.SetTodoRoutes(router)
}