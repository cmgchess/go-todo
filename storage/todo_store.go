package storage

import (
	"fmt"
	"time"

	"github.com/cmgchess/gotodo/models"
)

var todos = []models.Todo{
	{ID: 1, Name: "Todo 1", Description: "Description for Todo 1", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Enabled: true},
	{ID: 2, Name: "Todo 2", Description: "Description for Todo 2", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Enabled: true},
	{ID: 3, Name: "Todo 3", Description: "Description for Todo 3", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Enabled: true},
	{ID: 4, Name: "Todo 4", Description: "Description for Todo 4", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Enabled: true},
	{ID: 5, Name: "Todo 5", Description: "Description for Todo 5", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Enabled: true},
}
var currentID = len(todos)

func GetTodos() []models.Todo {
	return todos
}

func GetTodoByID(id int) (*models.Todo, error) {
	for i := range todos {
		if todos[i].ID == id {
			return &todos[i], nil
		}
	}
	return nil, fmt.Errorf("todo with id %d not found", id)
}

func AddTodo(todoRequest models.TodoRequest) models.Todo {
	currentID++
	todo := models.Todo{
		ID:          currentID,
		Name:        todoRequest.Name,
		Description: todoRequest.Description,
		Completed:   false,
		Enabled:     true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	todos = append(todos, todo)
	return todo
}

func ChangeEnableStatus(id int, enabled bool) (*models.Todo, error) {
	for i := range todos {
		if todos[i].ID == id && todos[i].Enabled != enabled {
			todos[i].Enabled = enabled
			todos[i].UpdatedAt = time.Now()
			return &todos[i], nil
		}
	}
	var status = "enabled"
	if enabled {
		status = "disabled"
	}
	return nil, fmt.Errorf("%s todo with id %d not found", status, id)
}

func UpdateTodo(id int, todoRequest models.TodoRequest) (*models.Todo, error) {
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Name = todoRequest.Name
			todos[i].Description = todoRequest.Description
			todos[i].UpdatedAt = time.Now()
			return &todos[i], nil
		}
	}
	return nil, fmt.Errorf("todo with id %d not found", id)
}

func DeleteTodo(id int) error {
	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo with id %d not found", id)
}
