package noteio

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nakachan-ing/reflect-cli/model"
)

func WriteNoteFile(note model.Note, frontMatter string, body string, config model.Config) (string, error) {
	content := fmt.Sprintf("---\n%s---\n\n%s", frontMatter, body)

	var dir string
	if note.NoteType == "fleeting" {
		dir = filepath.Join(config.BaseDir, "fleeting")
	}
	if note.NoteType == "permanent" {
		dir = filepath.Join(config.BaseDir, "permanent")
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	var fileName string
	if note.Slug != "" {
		fileName = fmt.Sprintf("%s_%s_%s.md", note.ID, note.SubType, note.Slug)
	} else {
		fileName = fmt.Sprintf("%s_%s.md", note.ID, note.SubType)
	}

	filePath := filepath.Join(dir, fileName)
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to write note file: %w", err)
	}

	return filePath, nil
}
