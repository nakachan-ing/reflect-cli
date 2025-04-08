package jsonstore

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nakachan-ing/reflect-cli/internal/noteio"
	"github.com/nakachan-ing/reflect-cli/model"
)

func UpdateNotes(filePath, noteID string, config *model.Config) ([]model.Note, error) {
	mdContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read Markdown file: %w", err)
	}

	frontMatter, _, err := noteio.ParseFrontMatter[model.FrontMatter](string(mdContent))
	if err != nil {
		return nil, fmt.Errorf("failed to parse front matter for %s: %v", filePath, err)
		// body = string(mdContent) // フロントマターの解析に失敗した場合、全文をセット
	}

	updatedNotes, err := LoadNotes(*config)
	if err != nil {
		log.Printf("Error loading notes from JSON: %v", err)
		os.Exit(1)
	}

	found := false
	for i := range updatedNotes {
		if updatedNotes[i].ID == noteID {
			updatedAt := time.Now()
			updatedNotes[i].Title = frontMatter.Title
			updatedNotes[i].NoteType = frontMatter.NoteType
			updatedNotes[i].SubType = model.SubType(frontMatter.SubType)
			updatedNotes[i].UpdatedAt = &updatedAt // 更新日時も更新
			updatedNotes[i].Reflected = frontMatter.Reflected
			updatedNotes[i].Source = frontMatter.Source
			updatedNotes[i].LinkedIssue = frontMatter.LinkedIssue
			// updatedNotes[i].LinkedNote = "ここはあとで実装する(変換実装がまだできていない)"
			// updatedNotes[i].Tags = "ここもあとで実装する(ノートに直更新でよいのか？)"
			found = true
			break
		}
	}

	if !found {
		log.Printf("Note with ID %s not found", noteID)
	}

	return updatedNotes, nil
}
