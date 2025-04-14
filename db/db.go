package db

import (
	"context"
	"fmt"

	"github.com/cmgchess/gotodo/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgreSQLStorage(cfg configs.Config) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return dbpool, nil
}
