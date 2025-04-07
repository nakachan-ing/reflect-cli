package model

type FrontMatter struct {
	Title       string   `yaml:"title"`
	NoteType    string   `yaml:"note_type"`
	SubType     string   `yaml:"sub_type"`
	CreatedAt   string   `yaml:"created_at"`
	UpdatedAt   string   `yaml:"updated_at"`
	Reflected   bool     `yaml:"reflected"`
	Source      string   `yaml:"source"`
	LinkedIssue string   `yaml:"linked_issue"`
	LinkedNotes []string `yaml:"linked_notes,omitempty"` // 関連ノート（自己参照）
	Tags        []string `yaml:"tags,omitempty"`         // タグ一覧
}
