package storage

import "github.com/cmgchess/gotodo/models"

type Storage interface {
	GetTodos() []models.Todo
	GetTodoByID(id int) (*models.Todo, error)
	AddTodo(todoRequest models.TodoRequest) models.Todo
	ChangeEnableStatus(id int, enabled bool) (*models.Todo, error)
	UpdateTodo(id int, todoRequest models.TodoRequest) (*models.Todo, error)
	DeleteTodo(id int) error
}
