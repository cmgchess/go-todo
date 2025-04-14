package router

import (
	"log"

	"github.com/cmgchess/gotodo/configs"
	"github.com/cmgchess/gotodo/db"
	"github.com/cmgchess/gotodo/handlers"
	"github.com/cmgchess/gotodo/storage"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	cfg := configs.Config{
		DBUser: configs.Envs.DBUser,
		DBPass: configs.Envs.DBPass,
		DBHost: configs.Envs.DBHost,
		DBName: configs.Envs.DBName,
	}
	db, err := db.NewPostgreSQLStorage(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	pingHandler := handlers.NewPingHandler()
	todoHandler := handlers.NewTodoHandler(storage.NewPostgresStorage(db))

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
