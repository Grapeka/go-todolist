package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          string    `json:"id"`
	Title       string `json:"title"`
	Completed   bool   `json:"complete"`
	TimeCreated string `json:"time_created"`
}

var todos = [] Todo {
	{ID: "1", Title: "First Todo", Completed: false, TimeCreated: "2019-01-01"},
	{ID: "2", Title: "Second Todo", Completed: false, TimeCreated: "2019-01-01"},
	{ID: "3", Title: "Third Todo", Completed: false, TimeCreated: "2019-01-01"},
	{ID: "4", Title: "Fourth Todo", Completed: false, TimeCreated: "2019-01-01"},
}

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")
	for _, todo := range todos {
		if todo.ID == id {
			context.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func AddTodo(context *gin.Context) {
	var newTodo Todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func DeleteTodo(context *gin.Context) {
	id := context.Param("id")
	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted"})
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}