package main

import (
	"fmt"
	"strings"
)

// UserConfirm asks the user to confirm an action.
func UserConfirm(msg string) bool {
	if config.Confirm {
		return true
	}
	var input string
	fmt.Print(msg)
	fmt.Scan(&input)
	input = strings.ToLower(input)
	input = strings.TrimSpace(input)
	if input == "y" || input == "yes" {
		return true
	}
	return false
}
