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

var allowedSubType = map[SubType]bool{
	"investigation": true,
	"idea":          true,
	"question":      true,
	"literature":    true,
	"quote":         true,
	"log":           true,
	"reference":     true,
}

func IsSubType(subTypeInput string) (SubType, error) {
	subType := SubType(subTypeInput)

	if !allowedSubType[subType] {
		// type is invalid の処理
		return "", fmt.Errorf("[%v] is invalid type", subTypeInput)
	} else {
		return subType, nil
	}
}
