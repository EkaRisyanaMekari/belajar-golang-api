package main

import (
	"belajar-golang-api/handler"
	"belajar-golang-api/todo"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var TodoID int = 4

var Todos = []todo.TodoInput{
	{ID: 1, Description: "learn golang", Status: false},
	{ID: 2, Description: "go to gym", Status: true},
	{ID: 3, Description: "cook breakfast", Status: false},
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/belajar-golang-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error")
	}
	fmt.Println("Database connected !")

	db.AutoMigrate(&todo.Todo{})

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
