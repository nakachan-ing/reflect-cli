# Zettelkasten Reflect CLI - プロジェクト構造とモジュール設計方針

このドキュメントでは、Zettelkasten Reflect CLI における「JSON / SQLite 両対応の実装」を見据えた、プロジェクトディレクトリ構成とモジュール設計方針を整理する。

---

## ✅ 目的

- バックエンド（JSON / SQLite）を切り替え可能にする
- NoteStore インターフェースを中心に、依存逆転の原則に基づく柔軟な構成を実現する
- 保守性・拡張性・テスト性を高める

## 🗂️ 推奨ディレクトリ構造

```plaintext
zk-reflect-cli/
├── cmd/                  # CLIエントリポイント（cobra等）
│   └── root.go
├── store/                # データアクセス層（抽象＋実装）
│   ├── store.go          # NoteStoreインターフェース定義
│   ├── jsonstore/        # JSONによる実装
│   │   ├── json_store.go
│   │   └── zettel.json
│   └── sqlitestore/      # SQLiteによる実装
│       ├── sqlite_store.go
│       └── zk.db
├── model/                # 構造体定義（Note, Tag など）
│   └── note.go
├── config/               # 設定ファイル読み込み
│   └── config.go         # .zkconfig.json を扱う
├── templates/            # Reflectテンプレート（YAML）
│   └── idea.yaml
├── utils/                # 汎用処理（slug生成など）
│   └── slug.go
├── docs/                # 設計ドキュメントなど
│   └── project_structure.md
├── go.mod
└── main.go
```

## 🔁 実装切り替えの例（Go）
```go
func LoadStore(config Config) NoteStore {
    switch config.Backend {
    case "sqlite":
        return sqlitestore.New(config)
    case "json":
        return jsonstore.New(config)
    default:
        return jsonstore.New(config)
    }
}
```

---

## 🔃 依存関係の方向性（依存逆転）

```plaintext
CLIコマンド（cmd/）
   ↓
NoteStore インターフェース（store.Store）
   ↓
jsonstore/ or sqlitestore/（実装モジュール）
```

→ CLI本体は保存形式に依存しない構造となり、将来的な移行や差し替えが容易になります。

---

## ✅ モジュール分離のメリット
|項目|内容|
|--|--|
|柔軟性|JSON, SQLite を`config.Backend` で切り替え可能|
|テスト性|モックを差し替えたCLIユニットテストが可能|
|拡張性|将来的に REST API やクラウドストアにも対応しやすい|


📌 今後の拡張候補

- `store/reststore/`（REST APIベースの同期）
- `store/cloudstore/`（S3やGCSへの保存）
- `store/memstore/`（テスト専用のインメモリ実装）

---

## ✅ 結論

- **バックエンドごとの実装は `store` 以下にモジュール化**
- **CLIは `NoteStore` を介して抽象的に操作する**
- **JSONからSQLiteへの移行・共存にも対応できる柔軟な構成を目指す**
