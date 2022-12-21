package handler

import (
	"belajar-golang-api/todo"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func BuildErrorMessages(err error) []string {
	var errMessages = []string{}
	for _, e := range err.(validator.ValidationErrors) {
		message := fmt.Sprintf("Error field %s [%s].", e.Field(), e.ActualTag())
		errMessages = append(errMessages, message)
	}
	return errMessages
}

type filterFunc func(todo.Todo) bool

func Filter(todos []todo.Todo, f filterFunc) []todo.Todo {
	var filtered []todo.Todo
	for _, todo := range todos {
		if f(todo) {
			filtered = append(filtered, todo)
		}
	}
	return filtered
}
