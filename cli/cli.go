package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/devxdh/task-cli/helper"
	"github.com/devxdh/task-cli/services"
)

const (
	allowedCommands = "Allowed commands: add | list | complete | delete | exit"
)

func EntryPoint(svc *services.Services) {
	reader := bufio.NewReader(os.Stdin)

	//Initial Message
	fmt.Println("=== Welcome to Your Task CLI ===")
	fmt.Println(allowedCommands)
	fmt.Println("Format: add [title] -d [description] OR add [title] OR delete [id]")
	fmt.Println("--------------------------------")

	// Infinite loop for interactive CLI
	for {
		fmt.Print("\n$")

		// Reads everything the user types until they press Enter (\n)
		input, err := reader.ReadString('\n')

		if helper.HandleErr(err, "An error occured while reading input") {
			continue
		}

		// Removes New Line character from input(\n)
		input = strings.TrimSpace(input)

		// Splits input into parts (e.g. ["add", "Study", "-d", "Study", "on", "Monday"])
		parts := strings.Fields(input)

		// Re-prompts if input was empty
		if len(parts) == 0 {
			continue
		}

		command, title, description := parseInput(parts)

		if command == "" {
			fmt.Println(allowedCommands)
			continue
		}

		if command != "list" && command != "exit" {
			if !helper.IsValidTitle(title) {
				fmt.Println("Title cannot be empty")
				continue
			}
		}

		switch command {

		case "add":
			err := svc.Add(title, description)
			if helper.HandleErr(err, "Failed to save your new task") {
				continue // Jumps immediately back to the beginning of the for loop ($ prompt)
			}

			fmt.Println("Task added successfully!\n")

		case "list":
			list, err := svc.List()
			if helper.HandleErr(err, "Failed to List your tasks") {
				continue
			}

			if len(list) == 0 {
				fmt.Println("No tasks found! Your todo list is empty.")
				continue
			}

			fmt.Println("\n=== Your Tasks ===")
			for _, task := range list {
				timeStr := task.CreatedAt.Format("2006-01-02 15:04")
				fmt.Printf("[%d] %s (Status: %s)\n", task.Id, task.Title, task.Status)
				fmt.Printf("    Description: %s\n", task.Description)
				fmt.Printf("    Created At:  %s\n", timeStr)
				fmt.Println("--------------------------------")
			}

		case "delete":
			if title == "" {
				fmt.Println("Please provide a task ID to delete. Format: delete [id]")
				continue
			}

			err := svc.Delete(title)
			if err != nil {
				fmt.Printf("failed to delete task: %w\n", err)
				continue
			}

			fmt.Println("Task deleted successfully!")

		case "complete":
			if title == "" {
				fmt.Println("Please provide a task ID to complete. Format: complete [id]")
				continue
			}

			err := svc.Complete(title)
			if err != nil {
				fmt.Printf("failed to complete task: %w\n", err)
				continue
			}

			fmt.Println("Task completed successfully!")

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println(allowedCommands)
			continue
		}
	}
}
