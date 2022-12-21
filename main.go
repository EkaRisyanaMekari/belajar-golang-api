package main

import (
	"belajar-golang-api/handler"
	"belajar-golang-api/todo"

	"github.com/gin-gonic/gin"
)

var TodoID int = 4

var Todos = []todo.Todo{
	{ID: 1, Description: "learn golang", Status: false},
	{ID: 2, Description: "go to gym", Status: true},
	{ID: 3, Description: "cook breakfast", Status: false},
}

func main() {
	handler.TodoID = TodoID
	handler.Todos = Todos

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/todos", handler.GetTodosHandler)
	v1.GET("/todos/outstanding", handler.GetTodosOutstandingHandler)
	v1.POST("/todos", handler.PostTodoHandler)
	v1.DELETE("/todos/:id", handler.DeleteTodoHandler)
	v1.PATCH("/todos/:id", handler.UpdateTodoHandler)

	router.Run(":7777")
}
