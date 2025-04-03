# Reflect CLI - 開発段階におけるReflect的運用方針

このドキュメントでは、Reflect CLIがまだ未完成の段階においても、その思想（fleeting → reflect → permanent）を活かしながら開発を進めるための運用方針を整理する。

---

## 🎯 目的

Reflect CLIの完成を待たずに、開発そのものをReflect的に進めることで：

- CLI思想と実装がズレない
- 設計ログ・思考ノートが自然と蓄積される
- 将来Reflect CLIで扱える「元ノート群」を形成できる

---

## ✅ 現時点の前提

| ツール | 状態 | 用途 |
|--------|------|------|
| Reflect CLI | 未完成 | CLIでのノート管理・Reflect不可（構想中） |
| Linear | 利用中 | 設計・タスク・マイルストーン管理 |
| GitHub | 利用中 | 実装・コードレビュー・Issue管理 |
| Markdown/Obsidian | 利用可能 | 設計・思考の記録に使用（Reflect CLIのノート的立場） |

---

## 🛠️ CLI未完成でもReflect的に開発するステップ

### ① Linearで Reflectを模倣したチケット運用

- `fleeting` ラベル：走り書き、検討中の思考タスク
- `permanent` ラベル：設計がまとまり昇華されたもの
- カスタムフィールドで `type: idea`, `type: investigation` など追加
- コメントに GitHub PR や Obsidianノートのリンクを記録（手動 `linked_issue`）

---

### ② Markdownでノート（設計メモ）を蓄積

- CLIで作れない分、手動でノートファイルを残す
- 命名規則（例）：`20250403_idea_github連携.md`
- frontmatter例：

```yaml
---
type: fleeting
subtype: idea
linked_issue: "https://linear.app/reflect/CLI-101"
tags: [reflect, github]
---
```

---

### ③ Reflect CLI 完成後、蓄積ノートをReflect対象にする

- `zk reflect` 実装後に、今までの設計ノートを1つずつ抽象化
- Permanentノートに昇華し、Reflect CLIで再帰的に管理可能にする
- CLI自身の設計をReflect CLIでReflectするという自己言及的プロセス

---

## ✅ この運用のメリット

| 項目 | 内容 |
|------|------|
| 思考の継続 | Reflect CLI完成前からZettelkasten的な開発ができる |
| OSS哲学の体現 | 自分が使いたいものを「使いながら育てる」思想が明示される |
| 移行がスムーズ | CLIが完成しても、ノートやタスクの構造がすでに整っている |

---

## ✅ 結論

Reflect CLIの開発そのものをReflect CLI的に管理・思考することで：

- CLI完成前も、思想を使いながら構築できる
- CLI完成後のReflectフェーズも自然と迎えられる

