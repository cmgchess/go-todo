package storage

import (
	"fmt"
	"time"

	"github.com/cmgchess/gotodo/models"
)

type InMemoryStorage struct {
	todos     []models.Todo
	currentID int
}

func NewInMemoryStorage() *InMemoryStorage {
	todos := []models.Todo{
		{ID: 1, Name: "Todo 1", Description: "Description for Todo 1", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Enabled: true},
		{ID: 2, Name: "Todo 2", Description: "Description for Todo 2", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Enabled: true},
		{ID: 3, Name: "Todo 3", Description: "Description for Todo 3", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Enabled: true},
		{ID: 4, Name: "Todo 4", Description: "Description for Todo 4", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Enabled: true},
		{ID: 5, Name: "Todo 5", Description: "Description for Todo 5", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Enabled: true},
	}
	return &InMemoryStorage{
		todos:     todos,
		currentID: len(todos),
	}
}

func (s *InMemoryStorage) GetTodos() []models.Todo {
	return s.todos
}

func (s *InMemoryStorage) GetTodoByID(id int) (*models.Todo, error) {
	for i := range s.todos {
		if s.todos[i].ID == id {
			return &s.todos[i], nil
		}
	}
	return nil, fmt.Errorf("todo with id %d not found", id)
}

func (s *InMemoryStorage) AddTodo(todoRequest models.TodoRequest) models.Todo {
	s.currentID++
	todo := models.Todo{
		ID:          s.currentID,
		Name:        todoRequest.Name,
		Description: todoRequest.Description,
		Completed:   false,
		Enabled:     true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	s.todos = append(s.todos, todo)
	return todo
}

func (s *InMemoryStorage) ChangeEnableStatus(id int, enabled bool) (*models.Todo, error) {
	for i := range s.todos {
		if s.todos[i].ID == id && s.todos[i].Enabled != enabled {
			s.todos[i].Enabled = enabled
			s.todos[i].UpdatedAt = time.Now()
			return &s.todos[i], nil
		}
	}
	var status = "enabled"
	if enabled {
		status = "disabled"
	}
	return nil, fmt.Errorf("%s todo with id %d not found", status, id)
}

func (s *InMemoryStorage) UpdateTodo(id int, todoRequest models.TodoRequest) (*models.Todo, error) {
	for i := range s.todos {
		if s.todos[i].ID == id {
			s.todos[i].Name = todoRequest.Name
			s.todos[i].Description = todoRequest.Description
			s.todos[i].UpdatedAt = time.Now()
			return &s.todos[i], nil
		}
	}
	return nil, fmt.Errorf("todo with id %d not found", id)
}

func (s *InMemoryStorage) DeleteTodo(id int) error {
	for i := range s.todos {
		if s.todos[i].ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo with id %d not found", id)
}
