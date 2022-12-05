package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          float64 `json:"id"`
	Description string  `json:"description"`
	Status      bool    `json:"status"`
}

var todoID float64 = 4
var todos = []Todo{
	{ID: 1, Description: "learn golang", Status: false},
	{ID: 2, Description: "go to gym", Status: false},
	{ID: 3, Description: "cook breakfast", Status: false},
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodosHandler)
	router.POST("/todos", postTodoHandler)

	router.Run(":7777")
}

func getTodosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func postTodoHandler(c *gin.Context) {
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	newTodo.ID = todoID + 1
	todos = append(todos, newTodo)
	c.JSON(http.StatusOK, todos)
}
