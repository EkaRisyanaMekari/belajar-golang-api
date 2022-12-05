package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

var todoID int = 4
var todos = []Todo{
	{ID: 1, Description: "learn golang", Status: false},
	{ID: 2, Description: "go to gym", Status: false},
	{ID: 3, Description: "cook breakfast", Status: false},
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodosHandler)
	router.POST("/todos", postTodoHandler)
	router.DELETE("/todos/:id", deleteTodoHandler)

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

func deleteTodoHandler(c *gin.Context) {
	var deletedID = c.Param("id")
	filteredTodo := filter(todos, func(todo Todo) bool {
		id, err := strconv.Atoi(deletedID)
		if err != nil {
			panic(err)
		}
		return todo.ID != id
	})

	c.JSON(http.StatusOK, filteredTodo)
}

type filterFunc func(Todo) bool

func filter(todos []Todo, f filterFunc) []Todo {
	var filtered []Todo
	for _, todo := range todos {
		if f(todo) {
			filtered = append(filtered, todo)
		}
	}
	return filtered
}
