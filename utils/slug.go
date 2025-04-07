package utils

import (
	"fmt"
	"regexp"
	"strings"
)

type slugGenerateError struct {
	Message string
}

func (e *slugGenerateError) Error() string {
	errorMsg := fmt.Sprintln("Failed to create slug because it does not contain alphabetic characters. Create file without slug.")
	return errorMsg
}

func Slugify(title string) (string, error) {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	slug := re.ReplaceAllString(strings.ToLower(title), "-")
	slug = strings.Trim(slug, "-")
	if slug == "" {
		errorMsg := fmt.Sprintln("Failed to create slug because it does not contain alphabetic characters. Create file without slug.")
		return "", &slugGenerateError{
			Message: errorMsg,
		}
	}

	return slug, nil
}
