package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/devxdh/task-cli/helper"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Init() {
	env := helper.Env()
	connStr := env["DATABASE_URL"]

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var err error

	DB, err = pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	err = DB.Ping(ctx)
	if err != nil {
		log.Fatalf("Could not ping database: %v", err)
	}

	err = DDL_SEED()
	if err != nil {
		log.Fatal("DDL failed to seed DB: %v", err)
	}

	fmt.Println("DB connected and seeded Successfully!")
}

func DDL_SEED() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	queryCreateTable := `
		CREATE TABLE IF NOT EXISTS tasks (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			title VARCHAR(100) NOT NULL,
			description TEXT,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`
	_, err := DB.Exec(ctx, queryCreateTable)
	return err
}
