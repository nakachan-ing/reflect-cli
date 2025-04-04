package model

import (
	"fmt"
	"time"
)

type Note struct {
	ID          string     `json:"id"` // yyyymmddhhmmss
	Title       string     `json:"title"`
	NoteType    string     `json:"note_type"` // fleeting / permanent
	SubType     SubType    `json:"sub_type"`  // idea, investigation など
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Archived    bool       `json:"archived"`
	Deleted     bool       `json:"deleted"`
	Reflected   bool       `json:"reflected"`
	FilePath    string     `json:"file_path"` // Markdownファイルのパス
	Slug        string     `json:"slug"`      // タイトル由来のファイル名用slug
	LinkedIssue string     `json:"linked_issue"`
	LinkedNotes []*Note    `json:"linked_notes,omitempty"` // 関連ノート（自己参照）
	// Tags        []*Tag     `json:"tags,omitempty"`         // タグ一覧
}

type SubType string

var AllowedSubType = map[SubType]bool{
	"investigation": true,
	"idea":          true,
	"question":      true,
	"literature":    true,
	"quote":         true,
	"log":           true,
	"reference":     true,
}

type subTypeInvalidError struct {
	Message string
}

func (e *subTypeInvalidError) Error() string {
	errorMsg := fmt.Sprintln("Type is invalid")
	errorMsg += fmt.Sprintln("Allowed types:")
	for subType := range AllowedSubType {
		errorMsg += fmt.Sprintf("  ・%v\n", subType)
	}
	return errorMsg
}

func IsSubType(subTypeInput string) (SubType, error) {
	subType := SubType(subTypeInput)

	if !AllowedSubType[subType] {
		// type is invalid の処理
		errorMsg := fmt.Sprintln("Type is invalid")
		errorMsg += fmt.Sprintln("Allowed types:")
		for subType := range AllowedSubType {
			errorMsg += fmt.Sprintf("  ・%v\n", subType)
		}
		return "", &subTypeInvalidError{
			Message: errorMsg,
		}
		// return "", fmt.Errorf("%s is invalid type\n\n%s", subTypeInput, AllowedSubTypesStringList()) // AllowedSubTypesStringListを使った場合、こちらをreturn

	} else {
		return subType, nil
	}
}

// 簡略化したい場合、こちらに切り替え
// func AllowedSubTypesStringList() string {
// 	var keys []string
// 	for k := range AllowedSubType {
// 		keys = append(keys, string(k))
// 	}
// 	sort.Strings(keys)

// 	msg := "Allowed types:\n"
// 	for _, k := range keys {
// 		msg += fmt.Sprintf("  ・%s\n", k)
// 	}
// 	return msg
// }
