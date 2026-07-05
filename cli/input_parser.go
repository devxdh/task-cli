package cli

import (
	"fmt"
	"slices"
	"strings"
)

func parseInput(parts []string) (command, title, description string) {
	if len(parts) == 0 {
		return "", "", ""
	}

	command = parts[0]
	allowed := []string{"add", "complete", "delete", "list", "exit"}

	if !slices.Contains(allowed, command) {
		fmt.Printf("%s is not allowed!\n", command)
		return "", "", ""
	}
	var titleParts []string
	var descriptionParts []string
	foundDescriptionFlag := false

	for i := 1; i < len(parts); i++ {
		if parts[i] == "-d" {
			foundDescriptionFlag = true
			continue
		}

		if !foundDescriptionFlag {
			titleParts = append(titleParts, parts[i])
		} else {
			descriptionParts = append(descriptionParts, parts[i])
		}
	}

	title = strings.TrimSpace(strings.Join(titleParts, " "))
	description = strings.TrimSpace(strings.Join(descriptionParts, " "))

	return command, title, description
}
