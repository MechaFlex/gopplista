package db

import (
	"context"
	"database/sql"
	_ "embed"

	_ "github.com/glebarez/go-sqlite"
)

//go:embed seed.sql
var seed string

func seedDatabase() error {
	ctx := context.Background()

	database, err := sql.Open("sqlite", "sqlite/db.sqlite")
	if err != nil {
		return err
	}

	_, err = database.ExecContext(ctx, seed)
	if err != nil {
		return err
	}

	return nil
}
