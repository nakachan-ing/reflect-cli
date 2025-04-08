package jsonstore

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nakachan-ing/reflect-cli/model"
)

func InsertNoteToJson(newNote model.Note, config *model.Config) error {
	notes, err := LoadNotes(*config)
	if err != nil {
		return err
	}

	notes = append(notes, newNote)

	notesJsonBytes, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to convert to JSON: %w", err)
	}

	err = os.WriteFile(config.ZettelJsonPath, notesJsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}

func SaveUpdatedJson[T any](v []T, jsonPath string) error {
	updatedJson, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("❌ Failed to convert to JSON: %w", err)
	}

	err = os.WriteFile(jsonPath, updatedJson, 0644)
	if err != nil {
		return fmt.Errorf("❌ Failed to write JSON file: %w", err)
	}

	log.Printf("✅ Successfully updated JSON file: %s", jsonPath)
	return nil
}
