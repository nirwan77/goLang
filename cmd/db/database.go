package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	dsn := "postgres://postgres:admin@localhost:5432/GoDatabase"

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return conn, nil
}
