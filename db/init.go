package db

import (
	"context"
	"database/sql"
	_ "embed"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

//go:embed schema/games.sql
var gamesSchema string

type Database struct {
	Ctx     context.Context
	Queries *Queries
}

func Init() (Database, error) {
	ctx := context.Background()

	database, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return Database{}, err
	}

	_, err = database.ExecContext(ctx, gamesSchema)
	if err != nil {
		return Database{}, err
	}

	if os.Getenv("SEED_DATABASE") == "true" {
		seedDatabase()
	}

	queries := New(database)

	return Database{ctx, queries}, nil
}
