package noteio

import (
	"fmt"
	"log"
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

func UpdateFrontMatter[T any](frontMatter T, body string) string {
	// Convert to YAML
	frontMatterBytes, err := yaml.Marshal(frontMatter)
	if err != nil {
		log.Printf("❌ Failed to convert front matter to YAML: %v", err)
		return body
	}

	// Preserve `---` and merge YAML with body
	return fmt.Sprintf("---\n%s---\n\n%s", string(frontMatterBytes), body)
}
