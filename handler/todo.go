package handler

import (
	"net/http"
	"strconv"

	"belajar-golang-api/todo"

	"github.com/gin-gonic/gin"
)

var Todos = []todo.TodoInput{}
var TodoID int

func GetTodosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Todos)
}

func GetTodosOutstandingHandler(c *gin.Context) {
	outstandingTodo := Filter(Todos, func(todo todo.TodoInput) bool { return !todo.Status })
	c.JSON(http.StatusOK, outstandingTodo)
}

func PostTodoHandler(c *gin.Context) {
	var newTodo todo.TodoInput
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": BuildErrorMessages(err)})
		return
	}
	TodoID = TodoID + 1
	newTodo.ID = TodoID
	newTodo.Status = false
	Todos = append(Todos, newTodo)
	c.JSON(http.StatusOK, Todos)
}

func DeleteTodoHandler(c *gin.Context) {
	var deletedID = c.Param("id")
	filteredTodo := Filter(Todos, func(todo todo.TodoInput) bool {
		id, err := strconv.Atoi(deletedID)
		if err != nil {
			panic(err)
		}
		return todo.ID != id
	})

	Todos = filteredTodo
	c.JSON(http.StatusOK, Todos)
}

func UpdateTodoHandler(c *gin.Context) {
	updatedId, err1 := strconv.Atoi(c.Param("id"))
	if err1 != nil {
		panic(err1)
	}
	var updatedTodo todo.TodoInput
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": BuildErrorMessages(err)})
		return
	}

	for i := 0; i < len(Todos); i++ {
		var localTodo todo.TodoInput = Todos[i]
		if localTodo.ID == updatedId {
			Todos[i].Description = updatedTodo.Description
			Todos[i].Status = updatedTodo.Status
			break
		}
	}
	c.JSON(http.StatusOK, Todos)
}
