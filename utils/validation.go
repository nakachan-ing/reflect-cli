package utils

import (
	"fmt"
	"regexp"
)

func ValidateIssueURL(issue string) (string, string) {
	re := regexp.MustCompile(`^https?://[\w./%-]+$`)
	if !re.MatchString(issue) {
		return "", fmt.Sprintf("Warning: Issue URL is invalid and will be skipped: %v", issue)

	}

	return issue, ""
}
