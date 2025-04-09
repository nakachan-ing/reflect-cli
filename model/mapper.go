package model

import (
	"fmt"
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

func MapFrontMatter(title, subType, source, issue string, tags []string) ([]byte, error) {
	t := time.Now()
	formattedTime := t.Format("2006-01-02T15:04:05")

	frontMatter := FrontMatter{
		Title:       title,
		NoteType:    "fleeting",
		SubType:     subType,
		CreatedAt:   formattedTime,
		UpdatedAt:   formattedTime,
		Reflected:   false,
		Source:      source,
		LinkedIssue: issue,
		LinkedNotes: []string{},
		Tags:        tags,
	}
	frontMatterBytes, err := yaml.Marshal(frontMatter)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to YAML: %w", err)
	}
	return frontMatterBytes, nil
}
