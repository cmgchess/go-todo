package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cmgchess/gotodo/configs"
	"github.com/cmgchess/gotodo/db"
	"github.com/cmgchess/gotodo/router"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	db, err := db.NewPostgreSQLStorage(configs.Envs.DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	initStorage(ctx, db)

	r := router.SetupRouter(db)

	log.Printf("Server running on port %s", configs.Envs.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configs.Envs.Port), r))
}

func initStorage(ctx context.Context, db *pgxpool.Pool) {
	err := db.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
