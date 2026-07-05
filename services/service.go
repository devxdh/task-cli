package services

import (
	"context"
	"fmt"
	"time"

	"github.com/devxdh/task-cli/helper"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskStatus string

type Task struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

const (
	StatusDue       TaskStatus = "due"
	StatusCompleted TaskStatus = "completed"
)

type Services struct {
	DB *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Services {
	return &Services{DB: pool}
}

func (s *Services) Add(title, description string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if !helper.IsValidDescription(description) {

	}

	queryInsertTask := `
		INSERT INTO tasks (title, description) VALUES ($1, $2);
	`

	_, err := s.DB.Exec(ctx, queryInsertTask, title, description)
	if err != nil {
		return fmt.Errorf("Database execution error: %v", err)
	}

	return nil
}
