package db

import (
	"app/api"
)

type Database interface {
	AddTodo(todo api.Todo) error
	AllTodos() ([]api.Todo, error)
}
