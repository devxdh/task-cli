package main

import (
	"fmt"

	"github.com/devxdh/task-cli/database"
)

func main() {
	fmt.Println("Hello, World!")
	database.Init()

	defer database.DB.Close()
}
