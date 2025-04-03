# Zettelkasten Reflect CLI - NoteStoreインターフェース仕様設計

このドキュメントでは、Zettelkasten Reflect CLI におけるメタデータ管理の中核となる `NoteStore` インターフェースの仕様と設計意図をまとめる。

---

## ✅ 目的

`NoteStore` は、CLI本体とストレージ実装（JSON or SQLite）の橋渡しを行う **データアクセス層の抽象化** を担う。  
CLIの処理はこのインターフェースを通してノートを取得・更新・削除することで、バックエンドを意識せずに柔軟に設計できる。

---

## 🎯 想定ユースケースと必要メソッド

| 操作例                     | 必要なメソッド                    |
|----------------------------|-----------------------------------|
| ノートの一覧を取得        | `GetAll()`                        |
| IDでノートを探す           | `FindByID(id string)`             |
| タイプで絞り込む           | `FindByType(noteType NoteType)`   |
| Reflect対象だけ取得        | `FindReflectable()`               |
| ノートを保存（新規・更新） | `Save(note *Note)`                |
| ノートを削除する           | `Delete(id string)`               |
| タグ一覧を取得             | `GetTags()`                       |
| ノートにタグを追加         | `AddTag(noteID string, tag *Tag)` |

---

## ✍️ NoteStoreインターフェース定義案（Go）

```go
type NoteStore interface {
    // ノート取得系
    GetAll() ([]*Note, error)
    FindByID(id string) (*Note, error)
    FindByType(noteType NoteType) ([]*Note, error)
    FindReflectable() ([]*Note, error)

    // ノート保存・削除
    Save(note *Note) error
    Delete(id string) error

    // タグ管理
    GetTags() ([]*Tag, error)
    AddTag(noteID string, tag *Tag) error
}
```

---

## 🔄 将来的に追加検討できるメソッド（v1.1〜）

| メソッド名            | 用途                                       |
|------------------------|--------------------------------------------|
| `FindBySubType()`      | サブタイプ（idea, questionなど）で絞り込み |
| `FindByTag(tag string)`| タグでノートを取得                         |
| `Search(query string)` | 検索語にマッチするノートを全文検索         |
| `ListLinkedNotes()`    | リンクされたノート一覧（構造ノート用）     |

---

## 💡 設計上の工夫

- `*Note` を使うことで、更新・保存時の差分管理がしやすい
- `[]*Note` を返すことで、JSONでもSQLでも共通の処理にできる
- `error` 戻り値でストレージエラーやバリデーションエラーにも対応可能
- CLIのコアロジックをデータ保存形式に依存させない

---

## ✅ 結論

- `NoteStore` を介してすべてのノート操作を抽象化することで、**CLI本体をクリーンに保ちつつ、将来的なJSON→SQLite移行も容易**になる。
- 初期実装は `JsonStore` から始め、将来的に `SqliteStore` を追加する構成が推奨される。

