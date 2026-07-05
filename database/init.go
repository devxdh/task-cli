package database

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/devxdh/task-cli/helper"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Init() *pgxpool.Pool {
	env := helper.Env()
	connStr := env["DATABASE_URL"]

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Could not ping database: %v", err)
	}

	err = DDL_SEED(pool)
	if err != nil {
		log.Fatal("DDL failed to seed DB: %v", err)
	}

	log.Println("DB connected and seeded Successfully!")

	return pool
}

//go:embed sql/001-init.sql
var initSchema string

func DDL_SEED(pool *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Injecting database schema from sql/001-init.sql...")

	_, err := pool.Exec(ctx, initSchema)
	if err != nil {
		return err
	}

	log.Println("Database schema injected successfully!")

	return nil
}
