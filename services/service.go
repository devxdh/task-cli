package services

type Services struct{}
type TaskStatus string

const (
	StatusDue       TaskStatus = "due"
	StatusCompleted TaskStatus = "completed"
)

type Task struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}

func (s *Services) Add() error {
	return nil
}
