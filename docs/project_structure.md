# Reflect CLI - プロジェクトディレクトリ構造（v0.1）

このドキュメントでは、Reflect CLI プロジェクトにおける推奨ディレクトリ構成と各ディレクトリの役割について説明する。

---

## 📁 ディレクトリ構成

```plaintext
zk-reflect-cli/
├── cmd/                        # CLIエントリポイント（cobra用）
│   ├── root.go
│   ├── new.go
│   ├── reflect.go
│   └── structure.go
│
├── config/                     # 設定ファイル管理
│   └── config.go               # .zkconfig.json など
│
├── internal/                   # Reflect CLI固有の内部処理
│   ├── reflect/                # Reflectの中核処理（テンプレロード〜出力まで）
│   │   └── engine.go
│   ├── parser/                 # frontmatter / YAML / JSONパース
│   │   └── frontmatter.go
│   ├── render/                 # ノート構造のMarkdown出力ビルダ
│   │   └── builder.go
│   ├── validation/             # slug/title/tagsのチェック系
│   │   └── validator.go
│   └── fileutil/               # ファイルI/Oユーティリティ
│       └── fs.go
│
├── model/                      # データ構造定義（Note, Tag など）
│   ├── note.go
│   ├── tag.go
│   └── note_link.go
│
├── store/                      # 抽象的なデータストアと実装
│   ├── store.go                # NoteStore インターフェース定義
│   ├── jsonstore/              # JSONベースの永続化
│   │   ├── json_store.go
│   │   └── zettel.json
│   └── sqlitestore/            # SQLiteベースの永続化
│       ├── sqlite_store.go
│       └── zk.db
│
├── templates/                  # Reflectテンプレート（YAML定義）
│   ├── ja/
│   │   └── idea.yaml
│   └── en/
│       └── idea.yaml
│
├── utils/                      # 汎用関数（slug生成・日付整形など）
│   └── slug.go
│
├── docs/                       # 設計ドキュメント・README補足
│   ├── architecture.md
│   ├── naming_strategy.md
│   ├── reflect_link_direction.md
│   ├── dev_strategy.md
│   └── milestone_april_early.md
│
├── go.mod
└── main.go                     # CLIの起動ポイント
```

---

## ✅ 設計方針のポイント

| ディレクトリ | 役割と狙い |
|--------------|------------|
| `cmd/`       | CLIのエントリーポイントとコマンド実装を分離しやすく保守性が高い構成 |
| `internal/`  | Reflect CLI専用ロジックを外部から隠蔽し、安全に管理する（Goのinternal構造） |
| `model/`     | 全体で使うデータ構造（Note, Tagなど）を一元管理 |
| `store/`     | 永続化層の抽象化（NoteStore）と具体実装（JSON, SQLite）を分離 |
| `templates/` | Reflect時のプロンプトテンプレートをローカライズ対応で分離 |
| `config/`    | CLI設定情報の読み込み・切り替え管理 |
| `utils/`     | 汎用処理の再利用（slug生成、ID発行など） |
| `docs/`      | OSS開発やコントリビューター向けの設計ドキュメント群 |

---

## 📌 補足

- この構成は `v0.1` の段階でのベース設計です。
- 将来的に `test/`, `examples/`, `internal/api/` などの追加が考えられる。
- コマンドが増えても `cmd/` 内に追加しやすく、拡張しやすい構造。
