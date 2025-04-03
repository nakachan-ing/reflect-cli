# Zettelkasten Reflect CLI - マイルストーン：〜 4月上旬 設計方針まとめ

このドキュメントは、Zettelkasten Reflect CLI における 📍マイルストーン：〜 4月上旬 の各タスクに対する設計方針を整理したものである。

---

## ✅ コマンド構成のドラフト作成（new, reflect, structure）

- `zk new fleeting --type <type>`  
  → fleetingノート（調査・アイデア等）をテンプレート付きで作成

- `zk reflect <note-path> [--type <type>]`  
  → fleetingノートを元に対話式でPermanentノートへ昇華（Reflect）

- `zk structure [list|generate]`  
  → Permanentノート間のリンク構造を可視化・構造ノートを生成

---

## ✅ fleeting noteのタイプ定義

- Reflect対象：
  - `investigation`, `idea`, `question`, `literature`
- Reflect非対象：
  - `quote`, `log`, `reference`（参照用リンクには活用可）

- Reflectテンプレートには抽象化観点を定義（例：原則・因果・応用）
- 必須 / 推奨 frontmatter をタイプ別に定義（例：literatureはsource必須）

---

## ✅ ノートのディレクトリ構造と命名規則の決定

- ディレクトリ：
  - `fleeting/`, `permanent/`
- 命名規則：
  - `fleeting/20250401_idea.md`（--titleなし）
  - `fleeting/20250401_idea_gc-latency.md`（--titleあり：slug付き）
- titleやslugの入力バリデーション（英数・ひらがな・記号制限など）も設計に含む

---

## ✅ メタデータ管理方法の検討

- 初期は `zettel.json` による管理
- 将来的にSQLite移行も可能にするため、`NoteStore` インターフェースを抽象化設計
- `.zkconfig.json` で `backend: "json"` or `"sqlite"` を切り替え可能に

---

## ✅ Go構造体の定義

- `Note`, `Tag`, `NoteTags`, `NoteType`, `SubType` を構造体＋enumで定義
- Noteは `linked_notes`, `linked_issue` フィールドを含む
- タグは多対多対応（将来的にSQLiteで使用）

---

## ✅ NoteStoreインターフェース設計

- `GetAll()`, `FindByID()`, `Save()`, `Delete()` などの基本操作を定義
- タグ管理：`GetTags()`, `AddTag()` も含む
- Reflect CLIのCLI層はNoteStoreにのみ依存する設計で切り替え可能に

---

## ✅ プロジェクトノートを持たない方針の確定と連携設計

- Reflect CLIは「プロジェクトノート」は持たない
- 代わりに `linked_issue:` フィールドでLinearやGitHub Issueと連携
- `zk new` の `--linked` フラグで初期入力可能
- Reflect時に `linked_issue` はPermanentノートへも引き継がれる

---

## ✅ CLI内でReflectテンプレートを扱う構造の設計（ピン留め中）

- YAMLテンプレート方式（例：templates/reflect/ja/idea.yaml）
- 言語別（ja/en）にディレクトリを分け、`--lang` フラグ or configで切り替え
- Reflect時に subtype を判定してテンプレートをロード

※このタスクは現在ピン留め中。先に `zk new fleeting` の実装を優先。

---

## ✅ Permanentノート間リンクとSQLite設計

- frontmatterには `linked_notes` をリスト形式で保存
- SQLiteでは `note_links` テーブルを使って多対多で管理
- Reflect時はFleetingとの関係を `linked_note:` に記録

---

## ✅ linked_issue の設計

- 基本は **1ノート : 1 Issue** で設計
- フィールド名：`linked_issue`（将来 `linked_issues: []` に拡張可能）
- ノートの文脈（実務との接点）を残すためのリンクとして使用

---

