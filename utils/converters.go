package utils

import (
	"regexp"
	"strings"
)

func ConvertNameToFilePath(name string) string {
	// Define a regular expression to match all non-alphanumeric characters
	var nonAlphanumeric = regexp.MustCompile(`[^a-zA-Z0-9]+`)

	// Replace all non-alphanumeric characters with dashes
	cleanedName := nonAlphanumeric.ReplaceAllString(name, "-")

	// Convert the result to lowercase
	filePath := strings.ToLower(cleanedName)

	// Trim leading and trailing dashes (in case the original name starts or ends with non-alphanumeric characters)
	filePath = strings.Trim(filePath, "-")

	return filePath
}
