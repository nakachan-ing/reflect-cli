# Zettelkasten Reflect CLI - メタデータ構造体定義（Go）

このドキュメントでは、Zettelkasten Reflect CLI におけるメタデータの管理に使用される Go構造体の定義をまとめる。

---

## 🧩 NoteType / SubType（カスタム型）

ノートの分類を `map[SubType]bool` で定義。

```go
type SubType string

var allowedSubType = map[subType]bool{
	"investigation": true,
	"idea":          true,
	"question":      true,
	"literature":    true,
	"quote":         true,
	"log":           true,
	"reference":     true,
}
```

---

## 📄 Note構造体

Reflect CLIで扱うノートの基本単位。

```go
type Note struct {
    ID           string     `json:"id"` // yyyymmddhhmmss
    Title        string     `json:"title"`
    NoteType     string   `json:"note_type"`   // fleeting / permanent
    SubType      SubType    `json:"sub_type"`    // idea, investigation など
    CreatedAt    *time.Time `json:"created_at"`
    UpdatedAt    *time.Time `json:"updated_at"`
    Archived     bool       `json:"archived"`
    Deleted      bool       `json:"deleted"`
    Reflected    bool       `json:"reflected"`
    FilePath     string     `json:"file_path"`   // Markdownファイルのパス
    Slug         string     `json:"slug"`        // タイトル由来のファイル名用slug
    LinkedIssue  []string    `json:"linked_issue"`
    LinkedNotes  []*Note    `json:"linked_notes,omitempty"` // 関連ノート（自己参照）
    Tags         []*Tag     `json:"tags,omitempty"`         // タグ一覧
}
```

---

## 🏷️ Tag構造体

ノートに紐づくタグ情報。

```go
type Tag struct {
    ID        string     `json:"id"` // 例: t-001
    Name      string     `json:"name"`
    CreatedAt *time.Time `json:"created_at"`
    Deleted   bool       `json:"deleted"`
}
```

---

## 🔗 NoteTags構造体（多対多の中間テーブル想定）

タグとノートの関連を表現する構造体（将来的にSQLite対応用）。

```go
type NoteTags struct {
    NoteID    string     `json:"note_id"`
    TagID     string     `json:"tag_id"`
    CreatedAt *time.Time `json:"created_at"`
    Deleted   bool       `json:"deleted"`
}
```

---

## 🔗 NoteLink構造体（多対多の中間テーブル想定）

ノートのリンクを表現する構造体（将来的にSQLite対応用）

```go
type NoteLink struct {
    FromNoteID string     `json:"from_note_id"`
    ToNoteID   string     `json:"to_note_id"`
    CreatedAt  *time.Time `json:"created_at"`
    Deleted    bool       `json:"deleted"`
}
```

---

## ✅ 今後の拡張候補

- `LinkedIssue`（LinearやGitHub IssueのURL）
- `Source`（文献・引用元）
- `Summary`（Reflect後の要約など）

この構造体定義により、CLIでのZettelkasten運用を効率よく、拡張性を保ちつつ実現する。
