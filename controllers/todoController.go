package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"todo/db"
	"todo/models"
	"todo/types"

	"github.com/gin-gonic/gin"
)


func GetTodos(context *gin.Context) {
	var todos []types.Todo
	db.DB.Find(&todos)
	context.IndentedJSON(http.StatusOK, todos)
}


func GetTodo(context *gin.Context) {
	id := context.Param("id")
	var todo types.Todo
	db.DB.First(&todo, id)
	context.IndentedJSON(http.StatusOK, todo)
}

func AddTodo(context *gin.Context) {
	var newTodo models.Todo
	if err := context.ShouldBindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Generate a 12-digit random number and set it as the id field
	newTodo.ID = fmt.Sprintf("%012d", rand.Intn(100000000000))

	db.DB.Create(&newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func UpdateTodo(context *gin.Context) {
	id := context.Param("id")
	var todo types.Todo
	db.DB.First(&todo, id)
	if todo.ID == "" {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	var updatedTodo models.Todo
	if err := context.ShouldBindJSON(&updatedTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTodo.ID = todo.ID
	db.DB.Save(&updatedTodo)
	context.JSON(http.StatusOK, gin.H{"data": updatedTodo})
}

func DeleteTodo(context *gin.Context) {
	id := context.Param("id")
	var todo types.Todo
	db.DB.First(&todo, id)
	if todo.ID == "" {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	db.DB.Delete(&todo)
	context.JSON(http.StatusOK, gin.H{"data": true})
}