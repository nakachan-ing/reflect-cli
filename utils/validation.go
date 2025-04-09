package utils

import (
	"fmt"
	"net/url"
)

// ValidateIssueURL checks if the URL is valid (RFC3986) and warns if non-ASCII chars are found.
// Returns: (URL or "", warning or "")
func ValidateIssueURL(issue string) (string, string) {
	if issue == "" {
		return "", ""
	}

	parsed, err := url.Parse(issue)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return "", fmt.Sprintf("Invalid URL format: %v", issue)
	}

	// Check for non-ASCII characters (e.g. Japanese)
	if containsNonASCII(issue) {
		return issue, fmt.Sprintf(
			"URL contains non-ASCII characters and may not be portable:\n   %s\nConsider using percent-encoded format if this causes issues.",
			issue,
		)
	}

	return issue, ""
}

// containsNonASCII returns true if string contains any character outside ASCII range
func containsNonASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return true
		}
	}
	return false
}
