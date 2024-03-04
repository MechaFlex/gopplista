package db

import (
	"context"
	"database/sql"
	_ "embed"

	_ "github.com/glebarez/go-sqlite"
)

//go:embed schema/games.sql
var gamesSchema string

//go:embed seed.sql
var seed string

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

	// _, err = database.ExecContext(ctx, seed)
	// if err != nil {
	// 	return Database{}, err
	// }

	queries := New(database)

	return Database{ctx, queries}, nil
}
