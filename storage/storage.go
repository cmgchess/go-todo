package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cmgchess/gotodo/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	db *pgxpool.Pool
}

func NewPostgresStorage(db *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}

func (s *PostgresStorage) GetTodos(ctx context.Context) ([]models.Todo, error) {
	rows, err := s.db.Query(ctx, "SELECT id, name, description, completed, enabled, created_at, updated_at FROM todos")
	if err != nil {
		return nil, fmt.Errorf("failed to query todos: %v", err)
	}
	defer rows.Close()

	todos := make([]models.Todo, 0)
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.Completed, &todo.Enabled, &todo.CreatedAt, &todo.UpdatedAt); err == nil {
			todos = append(todos, todo)
		}
	}
	return todos, nil
}

func (s *PostgresStorage) GetTodoByID(ctx context.Context, id int) (*models.Todo, error) {
	var todo models.Todo
	err := s.db.QueryRow(ctx, "SELECT id, name, description, completed, enabled, created_at, updated_at FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Name, &todo.Description, &todo.Completed, &todo.Enabled, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("todo with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to query todo: %v", err)
	}
	return &todo, nil
}

func (s *PostgresStorage) AddTodo(ctx context.Context, todoRequest models.TodoRequest) (models.Todo, error) {
	todo := models.Todo{
		Name:        todoRequest.Name,
		Description: todoRequest.Description,
		Completed:   false,
		Enabled:     true,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	var id int
	err := s.db.QueryRow(ctx, "INSERT INTO todos (name, description, completed, enabled, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", todo.Name, todo.Description, todo.Completed, todo.Enabled, todo.CreatedAt, todo.UpdatedAt).Scan(&id)
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to insert todo: %v", err)
	}
	todo.ID = id
	return todo, nil
}

func (s *PostgresStorage) ChangeEnableStatus(ctx context.Context, id int, enabled bool) (*models.Todo, error) {
	var todo models.Todo
	var status = "enabled"
	if enabled {
		status = "disabled"
	}
	err := s.db.QueryRow(ctx, "UPDATE todos SET enabled = $1, updated_at = $2 WHERE id = $3 AND enabled = NOT $1 RETURNING id, name, description, completed, enabled, created_at, updated_at", enabled, time.Now().UTC(), id).Scan(&todo.ID, &todo.Name, &todo.Description, &todo.Completed, &todo.Enabled, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s todo with id %d not found", status, id)
		}
		return nil, fmt.Errorf("failed to change todo enable status: %v", err)
	}
	return &todo, nil
}

func (s *PostgresStorage) UpdateTodo(ctx context.Context, id int, todoRequest models.TodoRequest) (*models.Todo, error) {
	var todo models.Todo
	err := s.db.QueryRow(ctx, "UPDATE todos SET name = $1, description = $2, updated_at = $3 WHERE id = $4 RETURNING id, name, description, completed, enabled, created_at, updated_at", todoRequest.Name, todoRequest.Description, time.Now().UTC(), id).Scan(&todo.ID, &todo.Name, &todo.Description, &todo.Completed, &todo.Enabled, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("todo with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to update todo: %v", err)
	}
	return &todo, nil
}

func (s *PostgresStorage) DeleteTodo(ctx context.Context, id int) error {
	res, err := s.db.Exec(ctx, "DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %v", err)
	}
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("todo with id %d not found", id)
	}
	return nil
}
