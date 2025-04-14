package router

import (
	"github.com/cmgchess/gotodo/handlers"
	"github.com/cmgchess/gotodo/storage"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	pingHandler := handlers.NewPingHandler()
	todoHandler := handlers.NewTodoHandler(storage.NewInMemoryStorage())

	r.HandleFunc("/ping", pingHandler.HealthHandler).Methods("GET")

	r.HandleFunc("/todos", todoHandler.GetTodosHandler).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.GetTodoByIDHandler).Methods("GET")
	r.HandleFunc("/todos", todoHandler.AddTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id}/enable", todoHandler.EnableTodoHandler).Methods("PATCH")
	r.HandleFunc("/todos/{id}/disable", todoHandler.DisableTodoHandler).Methods("PATCH")
	r.HandleFunc("/todos/{id}", todoHandler.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/todos/{id}", todoHandler.DeleteTodoHandler).Methods("DELETE")

	return r
}
