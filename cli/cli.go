package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/devxdh/task-cli/helper"
)

func EntryPoint() {
	reader := bufio.NewReader(os.Stdin)

	//Initial Message
	fmt.Println("=== Welcome to Your Task CLI ===")
	fmt.Println("Available commands: add | list | update | delete | exit")
	fmt.Println("Format: add [title] -d [description]  OR  add [title]")
	fmt.Println("--------------------------------")

	// Infinite loop for interactive use of CLI
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

		switch parts[0] {
		case "add":
		}
	}
}
