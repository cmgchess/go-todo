package storage

import (
	"context"

	"github.com/cmgchess/gotodo/models"
)

type Storage interface {
	GetTodos(ctx context.Context) ([]models.Todo, error)
	GetTodoByID(ctx context.Context, id int) (*models.Todo, error)
	AddTodo(ctx context.Context, todoRequest models.TodoRequest) (models.Todo, error)
	ChangeEnableStatus(ctx context.Context, id int, enabled bool) (*models.Todo, error)
	UpdateTodo(ctx context.Context, id int, todoRequest models.TodoRequest) (*models.Todo, error)
	DeleteTodo(ctx context.Context, id int) error
}
