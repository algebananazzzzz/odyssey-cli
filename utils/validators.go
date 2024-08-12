package utils

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

// ValidateNotEmpty checks if the input is not empty.
func ValidateAlphanumeric(input string) error {
	pattern := `^[a-zA-Z0-9_][\w-]*[a-zA-Z0-9_]$`

	matched, err := regexp.MatchString(pattern, input)
	if err != nil {
		// Handle any regex errors
		fmt.Println("Regex error:", err)
		return errors.New("error in regular expression matching")
	}

	if !matched {
		return errors.New("invalid input: please only use letters, numbers, underscores, and dashes")
	}

	return nil
}

// ValidateFilePath checks if the input is a valid file path (no spaces, no weird characters).
func ValidateFilePath(input string) error {
	// Define a regular expression for a valid file path (no spaces, no weird characters)
	var validPathPattern = `^[a-zA-Z0-9_\-/\.]+$`
	matched, err := regexp.MatchString(validPathPattern, input)

	if err != nil {
		return errors.New("error in regular expression matching")
	}

	if !matched {
		return errors.New("invalid file path: no spaces or special characters allowed")
	}

	if _, err := os.Stat(input); err == nil {
		return errors.New("the specified path already exists")
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("unable to check file path: %v", err)
	}

	return nil
}
