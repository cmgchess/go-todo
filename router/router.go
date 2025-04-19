package router

import (
	"net/http"

	"github.com/cmgchess/gotodo/handlers"
	"github.com/cmgchess/gotodo/middleware"
	"github.com/cmgchess/gotodo/storage"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool) *mux.Router {
	r := mux.NewRouter()
	sr := r.PathPrefix("/api/v1").Subrouter()

	sr.Use(middleware.LoggingMiddleware)

	pingHandler := handlers.NewPingHandler()
	todoHandler := handlers.NewTodoHandler(storage.NewPostgresStorage(db))

	r.Handle("/ping", middleware.LoggingMiddleware(http.HandlerFunc(pingHandler.HealthHandler))).Methods(http.MethodGet)

	sr.HandleFunc("/todos", todoHandler.GetTodosHandler).Methods(http.MethodGet)
	sr.HandleFunc("/todos/{id}", todoHandler.GetTodoByIDHandler).Methods(http.MethodGet)
	sr.HandleFunc("/todos", todoHandler.AddTodoHandler).Methods(http.MethodPost)
	sr.HandleFunc("/todos/{id}/enable", todoHandler.EnableTodoHandler).Methods(http.MethodPatch)
	sr.HandleFunc("/todos/{id}/disable", todoHandler.DisableTodoHandler).Methods(http.MethodPatch)
	sr.HandleFunc("/todos/{id}", todoHandler.UpdateTodoHandler).Methods(http.MethodPut)
	sr.HandleFunc("/todos/{id}", todoHandler.DeleteTodoHandler).Methods(http.MethodDelete)

	return r
}
