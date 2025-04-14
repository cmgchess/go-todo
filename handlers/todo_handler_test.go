package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cmgchess/gotodo/models"
	"github.com/gorilla/mux"
)

func TestTodoHandlers(t *testing.T) {
	t.Run("should return 200 if todos return successfully", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			GetTodosFunc: func(ctx context.Context) ([]models.Todo, error) {
				return []models.Todo{}, nil
			},
		})
		req, err := http.NewRequest(http.MethodGet, "/todos", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos", todoHandler.GetTodosHandler).Methods(http.MethodGet)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code 200, got %d", rr.Code)
		}
	})

	t.Run("should return 500 if internal error occurs when fetching todos", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			GetTodosFunc: func(ctx context.Context) ([]models.Todo, error) {
				return nil, errors.New("internal server error")
			},
		})
		req, err := http.NewRequest(http.MethodGet, "/todos", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		
		router.HandleFunc("/todos", todoHandler.GetTodosHandler).Methods(http.MethodGet)
		router.ServeHTTP(rr, req)
	
		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status code 500, got %d", rr.Code)
		}
	})

	t.Run("should return 200 if todo by ID return successfully", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			GetTodoByIDFunc: func(ctx context.Context, id int) (*models.Todo, error) {
				return &models.Todo{ID: id, Name: "Test Todo"}, nil
			},
		})
		req, err := http.NewRequest(http.MethodGet, "/todos/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", todoHandler.GetTodoByIDHandler).Methods(http.MethodGet)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code 200, got %d", rr.Code)
		}
	})

	t.Run("should return 400 if invalid id passed when get todo by ID", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{})
		req, err := http.NewRequest(http.MethodGet, "/todos/bla", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", todoHandler.GetTodoByIDHandler).Methods(http.MethodGet)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", rr.Code)
		}
	})

	t.Run("should return 404 if todo not found when get todo by ID", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			GetTodoByIDFunc: func(ctx context.Context, id int) (*models.Todo, error) {
				return nil, errors.New("todo not found")
			},
		})
		req, err := http.NewRequest(http.MethodGet, "/todos/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", todoHandler.GetTodoByIDHandler).Methods(http.MethodGet)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status code 404, got %d", rr.Code)
		}
	})

	t.Run("should return 201 if todo added successfully", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			AddTodoFunc: func(ctx context.Context, todoRequest models.TodoRequest) (models.Todo, error) {
				return models.Todo{ID: 1, Name: todoRequest.Name}, nil
			},
		})
		body := strings.NewReader(`{"name": "Test Todo", "description": "Testing add"}`)
		req, err := http.NewRequest(http.MethodPost, "/todos", body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos", todoHandler.AddTodoHandler).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code 201, got %d", rr.Code)
		}
	})

	t.Run("should return 400 if invalid payload when adding todo", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{})
		body := strings.NewReader("hello")
		req, err := http.NewRequest(http.MethodPost, "/todos", body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos", todoHandler.AddTodoHandler).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", rr.Code)
		}
	})

	t.Run("should return 400 if model validation failed when adding todo", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{})
		body := strings.NewReader(`{"description": ""}`)
		req, err := http.NewRequest(http.MethodPost, "/todos", body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos", todoHandler.AddTodoHandler).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", rr.Code)
		}
	})

	t.Run("should return 500 if internal error occurs when adding todo", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			AddTodoFunc: func(ctx context.Context, todoRequest models.TodoRequest) (models.Todo, error) {
				return models.Todo{}, errors.New("internal server error")
			},
		})
		body := strings.NewReader(`{"name": "Test Todo", "description": "Testing add"}`)
		req, err := http.NewRequest(http.MethodPost, "/todos", body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos", todoHandler.AddTodoHandler).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status code 500, got %d", rr.Code)
		}
	})

	t.Run("should return 200 if todo enabled successfully", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			ChangeEnableStatusFunc: func(ctx context.Context, id int, enabled bool) (*models.Todo, error) {
				return &models.Todo{ID: id, Name: "Test Todo", Enabled: enabled}, nil
			},
		})
		req, err := http.NewRequest(http.MethodPatch, "/todos/1/enable", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}/enable", todoHandler.EnableTodoHandler).Methods(http.MethodPatch)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code 200, got %d", rr.Code)
		}
	})

	t.Run("should return 400 if invalid id passed when enable todo", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{})
		req, err := http.NewRequest(http.MethodPatch, "/todos/bla/enable", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}/enable", todoHandler.EnableTodoHandler).Methods(http.MethodPatch)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", rr.Code)
		}
	})

	t.Run("should return 404 if todo not found when enable", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			ChangeEnableStatusFunc: func(ctx context.Context, id int, enabled bool) (*models.Todo, error) {
				return nil, errors.New("todo not found")
			},
		})
		req, err := http.NewRequest(http.MethodPatch, "/todos/1/enable", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}/enable", todoHandler.EnableTodoHandler).Methods(http.MethodPatch)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status code 404, got %d", rr.Code)
		}
	})

	t.Run("should return 200 if todo disabled successfully", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			ChangeEnableStatusFunc: func(ctx context.Context, id int, enabled bool) (*models.Todo, error) {
				return &models.Todo{ID: id, Name: "Test Todo", Enabled: enabled}, nil
			},
		})
		req, err := http.NewRequest(http.MethodPatch, "/todos/1/disable", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}/disable", todoHandler.DisableTodoHandler).Methods(http.MethodPatch)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code 200, got %d", rr.Code)
		}
	})

	t.Run("should return 400 if invalid id passed when disable todo", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{})
		req, err := http.NewRequest(http.MethodPatch, "/todos/bla/disable", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}/disable", todoHandler.DisableTodoHandler).Methods(http.MethodPatch)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", rr.Code)
		}
	})

	t.Run("should return 404 if todo not found when disable", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			ChangeEnableStatusFunc: func(ctx context.Context, id int, enabled bool) (*models.Todo, error) {
				return nil, errors.New("todo not found")
			},
		})
		req, err := http.NewRequest(http.MethodPatch, "/todos/1/disable", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}/disable", todoHandler.DisableTodoHandler).Methods(http.MethodPatch)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status code 404, got %d", rr.Code)
		}
	})

	t.Run("should return 200 if todo updated successfully", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			UpdateTodoFunc: func(ctx context.Context, id int, todoRequest models.TodoRequest) (*models.Todo, error) {
				return &models.Todo{ID: id, Name: todoRequest.Name}, nil
			},
		})
		body := strings.NewReader(`{"name": "Updated Todo", "description": "Testing update"}`)
		req, err := http.NewRequest(http.MethodPut, "/todos/1", body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", todoHandler.UpdateTodoHandler).Methods(http.MethodPut)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code 200, got %d", rr.Code)
		}
	})

	t.Run("should return 400 if invalid id passed when update todo", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{})
		body := strings.NewReader(`{"name": "Updated Todo", "description": "Testing update"}`)
		req, err := http.NewRequest(http.MethodPut, "/todos/bla", body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", todoHandler.UpdateTodoHandler).Methods(http.MethodPut)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", rr.Code)
		}
	})

	t.Run("should return 400 if invalid payload when updating todo", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{})
		body := strings.NewReader("hello")
		req, err := http.NewRequest(http.MethodPut, "/todos/1", body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", todoHandler.UpdateTodoHandler).Methods(http.MethodPut)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", rr.Code)
		}
	})

	t.Run("should return 400 if model validation failed when updating todo", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{})
		body := strings.NewReader(`{"description": ""}`)
		req, err := http.NewRequest(http.MethodPut, "/todos/1", body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", todoHandler.UpdateTodoHandler).Methods(http.MethodPut)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", rr.Code)
		}
	})

	t.Run("should return 404 if todo not found when update", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			UpdateTodoFunc: func(ctx context.Context, id int, todoRequest models.TodoRequest) (*models.Todo, error) {
				return nil, errors.New("todo not found")
			},
		})
		body := strings.NewReader(`{"name": "Updated Todo", "description": "Testing update"}`)
		req, err := http.NewRequest(http.MethodPut, "/todos/1", body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", todoHandler.UpdateTodoHandler).Methods(http.MethodPut)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status code 404, got %d", rr.Code)
		}
	})

	t.Run("should return 200 if todo deleted successfully", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			DeleteTodoFunc: func(ctx context.Context, id int) error { return nil },
		})
		req, err := http.NewRequest(http.MethodDelete, "/todos/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", todoHandler.DeleteTodoHandler).Methods(http.MethodDelete)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Errorf("expected status code 204, got %d", rr.Code)
		}
	})

	t.Run("should return 400 if invalid id passed when delete todo", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{})
		req, err := http.NewRequest(http.MethodDelete, "/todos/bla", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", todoHandler.DeleteTodoHandler).Methods(http.MethodDelete)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", rr.Code)
		}
	})

	t.Run("should return 404 if todo not found when delete", func(t *testing.T) {
		todoHandler := NewTodoHandler(&mockStore{
			DeleteTodoFunc: func(ctx context.Context, id int) error { return errors.New("todo not found") },
		})
		req, err := http.NewRequest(http.MethodDelete, "/todos/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/todos/{id}", todoHandler.DeleteTodoHandler).Methods(http.MethodDelete)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status code 404, got %d", rr.Code)
		}
	})
}

type mockStore struct {
	GetTodosFunc           func(ctx context.Context) ([]models.Todo, error)
	GetTodoByIDFunc        func(ctx context.Context, id int) (*models.Todo, error)
	AddTodoFunc            func(ctx context.Context, todoRequest models.TodoRequest) (models.Todo, error)
	ChangeEnableStatusFunc func(ctx context.Context, id int, enabled bool) (*models.Todo, error)
	UpdateTodoFunc         func(ctx context.Context, id int, todoRequest models.TodoRequest) (*models.Todo, error)
	DeleteTodoFunc         func(ctx context.Context, id int) error
}

func (m *mockStore) GetTodos(ctx context.Context) ([]models.Todo, error) {
	return m.GetTodosFunc(ctx)
}

func (m *mockStore) GetTodoByID(ctx context.Context, id int) (*models.Todo, error) {
	return m.GetTodoByIDFunc(ctx, id)
}

func (m *mockStore) AddTodo(ctx context.Context, todoRequest models.TodoRequest) (models.Todo, error) {
	return m.AddTodoFunc(ctx, todoRequest)
}

func (m *mockStore) ChangeEnableStatus(ctx context.Context, id int, enabled bool) (*models.Todo, error) {
	return m.ChangeEnableStatusFunc(ctx, id, enabled)
}

func (m *mockStore) UpdateTodo(ctx context.Context, id int, todoRequest models.TodoRequest) (*models.Todo, error) {
	return m.UpdateTodoFunc(ctx, id, todoRequest)
}

func (m *mockStore) DeleteTodo(ctx context.Context, id int) error {
	return m.DeleteTodoFunc(ctx, id)
}
