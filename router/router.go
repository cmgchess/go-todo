package router

import (
	"github.com/cmgchess/gotodo/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ping", handlers.PingHandler).Methods("GET")

	r.HandleFunc("/todos", handlers.GetTodosHandler).Methods("GET")
	r.HandleFunc("/todos/{id}", handlers.GetTodoByIDHandler).Methods("GET")
	r.HandleFunc("/todos", handlers.AddTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id}/enable", handlers.EnableTodoHandler).Methods("PATCH")
	r.HandleFunc("/todos/{id}/disable", handlers.DisableTodoHandler).Methods("PATCH")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodoHandler).Methods("DELETE")

	return r
}
