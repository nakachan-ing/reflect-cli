package noteio

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func ParseFrontMatter[T any](content string) (T, string, error) {
	var frontMatter T

	if !strings.HasPrefix(content, "---") {
		return frontMatter, content, fmt.Errorf("❌ Front matter not found")
	}

	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return frontMatter, content, fmt.Errorf("❌ Invalid front matter format")
	}

	frontMatterStr := strings.TrimSpace(parts[1])
	body := strings.TrimSpace(parts[2])

	// Parse YAML
	err := yaml.Unmarshal([]byte(frontMatterStr), &frontMatter)
	if err != nil {
		return frontMatter, content, fmt.Errorf("❌ Failed to parse front matter: %w", err)
	}

	return frontMatter, body, nil
}
