package noteio

import (
	"fmt"
	"os"
)

func ParseNoteFile[T any](filePath string) (T, string, error) {
	var front T

	mdContent, err := os.ReadFile(filePath)
	if err != nil {
		return front, "", fmt.Errorf("failed to read Markdown file: %w", err)
	}

	front, body, err := ParseFrontMatter[T](string(mdContent))
	if err != nil {
		return front, "", fmt.Errorf("failed to parse front matter for %s: %v", filePath, err)
	}

	return front, body, nil
}
