# Zettelkasten Reflect CLI - ノートのリンクと外部Issue連携設計

このドキュメントでは、Zettelkasten Reflect CLI における以下の設計方針を整理する：

- Permanentノート間のリンク構造
- Reflect元ノートとの関係
- 外部Issue（Linear/GitHub）との連携
- SQLiteとの整合性

---

## ✅ 1. Permanentノート間リンクの基本方針

| リンク種別            | 内容                                               |
|------------------------|----------------------------------------------------|
| Fleeting → Permanent   | Reflect時に自動的にリンクされる（linked_note）     |
| Permanent ↔ Permanent  | 手動リンクまたは `zk link` コマンドで管理予定     |
| 構造ノート生成         | これらのリンク情報を元に `zk structure generate` で可視化 |

---

## 📍 リンクの保存場所（frontmatter vs body）

| 保存先       | 用途・特徴                                          |
|--------------|------------------------------------------------------|
| frontmatter  | `linked_notes: [...]` としてCLI側が扱いやすい形式   |
| body         | `[[note-id]]` 形式で記述。Neovim/Obsidian等と相性◎  |

Reflect CLIの標準では **frontmatterを正式管理場所** とし、bodyはユーザーの自由記述とする。

---

## ✅ 2. SQLiteでのノートリンク管理

Permanentノート間の多対多リンクを扱うために、`note_links` 中間テーブルを用意する：

```sql
CREATE TABLE note_links (
    from_note_id TEXT NOT NULL,
    to_note_id   TEXT NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted      BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (from_note_id, to_note_id),
    FOREIGN KEY (from_note_id) REFERENCES notes(id),
    FOREIGN KEY (to_note_id)   REFERENCES notes(id)
);
```

Go構造体：

```go
type NoteLink struct {
    FromNoteID string     `json:"from_note_id"`
    ToNoteID   string     `json:"to_note_id"`
    CreatedAt  *time.Time `json:"created_at"`
    Deleted    bool       `json:"deleted"`
}
```

---

## ✅ 3. 外部Issue（Linear/GitHubなど）とのリンク戦略

### フィールド名：`linked_issue`

Reflect CLIでは、ノートが生まれた文脈として「外部のプロジェクトやタスクへのリンク」を記録する。

例（frontmatter）：

```yaml
linked_issue: "https://linear.app/you/project/ISSUE-123"
```

### 関係性：**1ノート : 1 Issue（基本設計）**

| 理由                              | 内容 |
|-----------------------------------|------|
| シンプルで分かりやすい            | タスク or プロジェクトと1対1で結びつける構造が大半 |
| 実装が簡単                        | フロントマターで1フィールドで済む |
| 将来的にリスト化も可能            | `linked_issues: [...]` に拡張可能な構造を意識しておくと◎ |

---

## ✅ まとめ

| 項目 | Reflect CLIとしての設計 |
|------|--------------------------|
| ノートリンク | frontmatterに `linked_notes` を記録（Permanent間） |
| Reflect元リンク | `linked_note` でFleeting→Permanent関係を記録 |
| 外部Issueリンク | `linked_issue`（1対1）として記録。拡張可能 |
| SQLite対応     | `note_links` 中間テーブルでPermanent間リンクを管理 |

