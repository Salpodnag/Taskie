package db

import (
	"Taskie/cfg"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewClient(cfg *cfg.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.DBPassword,
		cfg.DB.DBHost,
		cfg.DB.DBPort,
		cfg.DB.DBName)

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
