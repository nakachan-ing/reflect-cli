package templateio

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoadFleetingTemplate(path string) (string, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return "", fmt.Errorf("failed to create template directory: %w", err)
	}
	templateContent, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read template file (%s): %w", path, err)
	}
	return string(templateContent), nil
}
