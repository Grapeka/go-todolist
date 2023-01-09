package routers

import (
	"todo/controllers"

	"github.com/gin-gonic/gin"
)

func SetTodoRoutes(router *gin.Engine) {
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todo/:id", controllers.GetTodo)
	router.POST("/todo", controllers.AddTodo)
	router.PUT("/todo/:id", controllers.UpdateTodo)
	router.DELETE("/todo/:id", controllers.DeleteTodo)
}