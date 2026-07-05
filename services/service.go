package services

import (
	"context"
	"fmt"
	"time"

	"github.com/devxdh/task-cli/helper"
	"github.com/jackc/pgx/v5"
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
		description = "No description provided at creation"
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

func (s *Services) List() ([]Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	queryListTasks := `
		SELECT id, title, description, created_at, updated_at FROM tasks;
	`

	rows, err := s.DB.Query(ctx, queryListTasks)
	if err != nil {
		return nil, fmt.Errorf("failed to execute list query: %w", err)
	}
	defer rows.Close()

	tasks, err := pgx.CollectRows(rows, pgx.RowToStructByName[Task])
	if err != nil {
		return nil, fmt.Errorf("error collecting database rows %w\n", err)
	}

	return tasks, nil
}
