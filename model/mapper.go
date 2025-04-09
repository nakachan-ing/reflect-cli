package model

import (
	"fmt"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

func MapNote(title, subType, slug, source, issue string, tags []*Tag) Note {
	t := time.Now()
	noteId := fmt.Sprintf("%d%02d%02dT%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	createdAt := t

	note := Note{
		ID:        noteId,
		Title:     title,
		NoteType:  "fleeting",
		SubType:   SubType(subType),
		CreatedAt: &createdAt,
		UpdatedAt: &createdAt,
		Archived:  false,
		Deleted:   false,
		Reflected: false,
		// FilePath:    "",
		Slug:        slug,
		Source:      source,
		LinkedIssue: issue,
		LinkedNotes: []*Note{},
		Tags:        tags,
	}
	return note
}

func MapTags(strTags []string) []*Tag {
	var tags []*Tag
	for _, strTag := range strTags {
		createAt := time.Now()
		tag := Tag{
			ID:        "T001",
			Name:      strTag,
			CreatedAt: &createAt,
			Deleted:   false,
		}
		tags = append(tags, &tag)
	}
	return tags
}

func MapTag(strTag string) Tag {
	createAt := time.Now()
	tag := Tag{
		ID:        "",
		Name:      strTag,
		CreatedAt: &createAt,
		Deleted:   false,
	}
	return tag
}

func MapFrontMatter(title, noteType, subType, source, issue string, tags []string, fileName string) ([]byte, error) {
	t := time.Now()
	formattedTime := t.Format("2006-01-02T15:04:05")

	frontMatter := FrontMatter{
		Title:       title,
		NoteType:    noteType,
		SubType:     subType,
		CreatedAt:   formattedTime,
		UpdatedAt:   formattedTime,
		Reflected:   false,
		Source:      source,
		LinkedIssue: issue,
		LinkedNotes: []string{},
		Tags:        tags,
	}
	if noteType == "permanent" {
		frontMatter.Reflected = true
		var linkedNotes []string
		linkedNotes = append(linkedNotes, fmt.Sprintf("[[%s]]", filepath.Join("..", "fleeting", fileName)))
		frontMatter.LinkedNotes = linkedNotes
	}

	frontMatterBytes, err := yaml.Marshal(frontMatter)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to YAML: %w", err)
	}
	return frontMatterBytes, nil
}

func MapReflectToPermanent(title, subType, slug, source, issue string, notes []*Note, tags []*Tag, responses []string, config *Config) Note {
	t := time.Now()
	id := fmt.Sprintf("%d%02d%02dT%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	return Note{
		ID:        id,
		Title:     title,
		NoteType:  "permanent",
		SubType:   SubType(subType),
		CreatedAt: &t,
		UpdatedAt: &t,
		Archived:  false,
		Deleted:   false,
		Reflected: true,
		// FilePath:    "",
		Slug:        slug,
		Source:      source,
		LinkedIssue: issue,
		LinkedNotes: []*Note{},
		Tags:        tags,
	}

}
