package db

import (
	"Taskie/cfg"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func NewClient(cfg *cfg.Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.DBPassword,
		cfg.DB.DBHost,
		cfg.DB.DBPort,
		cfg.DB.DBName)

	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}
