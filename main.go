package main

import (
	"github.com/devxdh/task-cli/cli"
	"github.com/devxdh/task-cli/database"
	"github.com/devxdh/task-cli/services"
)

func main() {
	pool := database.Init()
	svc := services.NewService(pool)
	defer pool.Close()

	cli.EntryPoint(svc)
}
