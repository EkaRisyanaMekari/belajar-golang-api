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
	{ID: 2, Description: "go to gym", Status: true},
	{ID: 3, Description: "cook breakfast", Status: false},
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodosHandler)
	router.GET("/todos/outstanding", getTodosOutstandingHandler)
	router.POST("/todos", postTodoHandler)
	router.DELETE("/todos/:id", deleteTodoHandler)
	router.PATCH("/todos/:id", updateTodoHandler)

	router.Run(":7777")
}

func getTodosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func getTodosOutstandingHandler(c *gin.Context) {
	outstandingTodo := filter(todos, func(todo Todo) bool { return !todo.Status })
	c.JSON(http.StatusOK, outstandingTodo)
}

func postTodoHandler(c *gin.Context) {
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todoID = todoID + 1
	newTodo.ID = todoID
	newTodo.Status = false
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

	todos = filteredTodo
	c.JSON(http.StatusOK, todos)
}

func updateTodoHandler(c *gin.Context) {
	updatedId, err1 := strconv.Atoi(c.Param("id"))
	if err1 != nil {
		panic(err1)
	}
	var updatedTodo Todo
	if err := c.BindJSON(&updatedTodo); err != nil {
		return
	}

	for i := 0; i < len(todos); i++ {
		var localTodo Todo = todos[i]
		if localTodo.ID == updatedId {
			todos[i].Description = updatedTodo.Description
			todos[i].Status = updatedTodo.Status
			break
		}
	}
	c.JSON(http.StatusOK, todos)
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
