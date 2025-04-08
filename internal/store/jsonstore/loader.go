package jsonstore

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nakachan-ing/reflect-cli/model"
)

type ZettelJsonNotExistError struct {
	Message string
}

func (e *ZettelJsonNotExistError) Error() string {
	return e.Message
}

type ZettelJsonReadError struct {
	Message string
}

func (e *ZettelJsonReadError) Error() string {
	return e.Message
}

type ZettelJsonParseError struct {
	Message string
}

func (e *ZettelJsonParseError) Error() string {
	return e.Message
}

func LoadJson[T any](filePath string, v *[]T) error {
	// var multiErr zettelJsonLoadError
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// ファイルが存在しない場合は空のスライスを返す
		*v = []T{}
		return nil
	} else if err != nil {
		zettelJsonNotExistErrorMsg := &ZettelJsonNotExistError{
			Message: fmt.Sprintf("failed to check JSON file: %v", err),
		}
		// multiErr.zettelJsonNotExistError = zettelJsonNotExistErrorMsg
		return zettelJsonNotExistErrorMsg
	}

	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		zettelJsonReadErrorMsg := &ZettelJsonReadError{
			Message: fmt.Sprintf("failed to read JSON file: %v", err),
		}
		// multiErr.zettelJsonReadError = zettelJsonReadErrorMsg
		return zettelJsonReadErrorMsg
	}

	if len(jsonBytes) > 0 {
		err = json.Unmarshal(jsonBytes, v)
		if err != nil {
			zettelJsonParseErrorMsg := &ZettelJsonParseError{
				Message: fmt.Sprintf("failed to parse JSON: %v", err),
			}
			// multiErr.ZettelJsonParseError = zettelJsonParseErrorMsg
			return zettelJsonParseErrorMsg
		}
	}

	return nil
}

func LoadNotes(config model.Config) ([]model.Note, error) {
	zettelJsonPath := config.ZettelJsonPath

	// ディレクトリがない場合は作成
	if err := os.MkdirAll(filepath.Dir(zettelJsonPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	// notes.json が存在しない場合、空の JSON 配列 `[]` で初期化
	if _, err := os.Stat(zettelJsonPath); os.IsNotExist(err) {
		if err := os.WriteFile(zettelJsonPath, []byte("[]"), 0644); err != nil {
			return nil, fmt.Errorf("failed to create notes.json file: %w", err)
		}
	} else if err != nil {
		// ファイルの存在確認時の別のエラー（例: 権限エラー）
		return nil, fmt.Errorf("failed to check notes.json: %w", err)
	}

	// JSON をロード
	var notes []model.Note
	if err := LoadJson(zettelJsonPath, &notes); err != nil {
		return nil, fmt.Errorf("error loading notes from JSON: %w", err)
	}

	return notes, nil
}

func LoadTags(config model.Config) ([]model.Tag, error) {
	tagsJsonPath := config.TagsJsonPath

	// ディレクトリがない場合は作成
	if err := os.MkdirAll(filepath.Dir(tagsJsonPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	// notes.json が存在しない場合、空の JSON 配列 `[]` で初期化
	if _, err := os.Stat(tagsJsonPath); os.IsNotExist(err) {
		if err := os.WriteFile(tagsJsonPath, []byte("[]"), 0644); err != nil {
			return nil, fmt.Errorf("failed to create tags.json file: %w", err)
		}
	} else if err != nil {
		// ファイルの存在確認時の別のエラー（例: 権限エラー）
		return nil, fmt.Errorf("failed to check tags.json: %w", err)
	}

	// JSON をロード
	var tags []model.Tag
	if err := LoadJson(tagsJsonPath, &tags); err != nil {
		return nil, fmt.Errorf("error loading tags from JSON: %w", err)
	}

	return tags, nil
}
