# Zettelkasten Reflect CLI - ノートのディレクトリ構造と命名規則

このドキュメントは、Zettelkasten Reflect CLI におけるノートの保存ディレクトリ構造とファイル命名規則に関する設計をまとめたものである。

---

## 📁 ディレクトリ構造

```plaintext
zk-root/
├── fleeting/      # 素材・走り書きメモを格納（Reflect前）
└── permanent/     # Reflectによって抽象化された永久保存ノート
```

---

## 📛 ファイル名の命名規則

| 条件             | 命名規則例                          | 補足                          |
|------------------|-------------------------------------|-------------------------------|
| `--title` なし   | `20250401_idea.md`                  | タイプのみ付加                |
| `--title` あり   | `20250401_idea_gc-memory.md`        | title から生成された slug を付加 |
| Reflect出力      | `permanent/gc-memory-principles.md` | Reflect時に手動 or 自動命名   |

---

## 🏷️ --title オプションと文字制限について

### 🎯 設計方針

- `--title` は日本語、記号、全角文字など自由に入力可能
- ファイル名に使う slug は **ファイルシステムで安全・人間可読な形に変換**

---

### 🔤 slug化ルール（ファイル名用）

| 文字種         | slugへの変換          |
|----------------|------------------------|
| 英字・数字     | そのまま使用           |
| スペース       | ハイフン（`-`）に変換   |
| 全角文字       | 削除 or 無視           |
| 記号・絵文字   | 削除                   |
| 複数ハイフン   | 連続しないよう整理     |

### 🛠️ Goでの例実装

```go
import "regexp"
import "strings"

func Slugify(title string) string {
    re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
    slug := re.ReplaceAllString(strings.ToLower(title), "-")
    slug = strings.Trim(slug, "-")
    return slug
}
```

---

### 🧪 使用例

```bash
zk new fleeting --type idea --title "GCの挙動と最適化案"
# → ファイル名: 20250401_idea_gc.md

zk new fleeting --type question --title "JavaのGC戦略はなぜ分かりにくいのか？"
# → ファイル名: 20250401_question_java-gc.md
```

---

## ✅ 結論：設計サマリー

| 項目         | 方針内容 |
|--------------|----------|
| ディレクトリ構造 | `fleeting/` と `permanent/` の2階層でシンプルに運用 |
| タイトル      | 日本語OK、frontmatterに記録（title） |
| slug（ファイル名） | 半角英字＋数字＋ハイフンのみ。Goで整形処理 |
| Reflectノート | Reflect時に明示命名 or slug化による出力 |

---

## 📝 補足：今後の拡張余地

- Reflectノートの命名を `title` 入力必須にすることでユーザーの意図を強調
- `zk list` などでファイル名とタイトルを一覧表示
- slug生成をカスタマイズ可能にする設定ファイル対応（将来的に）

