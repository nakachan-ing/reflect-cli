# Reflect CLI - Linear × GitHub 開発連携ガイド

このドキュメントでは、Reflect CLI 開発において Linear と GitHub を連携させることで、タスク管理とPRレビューを効率化する方法をまとめる。

---

## ✅ 目的

- Linear で管理しているタスクを GitHub の Pull Request やブランチと **自動で関連づける**
- PRマージと同時に Linear 側の Issue ステータスを自動更新
- 一貫した命名ルールにより、開発フローの自動化・見通しを向上

---

## 🔗 1. Linear × GitHub 連携手順

### A. Linear 側で GitHub と接続

1. Linear 画面右上メニュー → **Settings**
2. **Integrations** → **GitHub**
3. 「Connect GitHub」 をクリック
4. 対象の GitHub リポジトリを選択
5. 以下のルールでリンクされる設定を確認：

- PRタイトル、ブランチ名、コミットメッセージに Linear Issue Key（例: `ZR-45`）を含めることで自動リンクされる

---

## 🧪 2. 命名ルール例（Reflect CLI用）

| 種別 | 命名例 | 説明 |
|------|--------|------|
| Linear Issue Key | `ZR-45` | Linear上で発行されるチームキー付き番号 |
| GitHub ブランチ | `feat/zk-new-type-ZR-45` | 機能ブランチ例 |
| PRタイトル | `✨ Add zk new --type option (ZR-45)` | タスクとの紐付けが可能 |
| コミットメッセージ（任意） | `fix: adjust slug generation (ZR-45)` | より詳細なトレースが可能に |

---

## 🔄 3. 自動リンク例

Linearに `✨ zk new --type 実装` というIssue（ZR-45）があるとき：

```bash
git checkout -b feat/zk-new-type-ZR-45
```

このブランチでPRを作成し、PRタイトルに `ZR-45` を含めると：

- PRが Linear に自動リンクされる
- PRのステータスで Linear 側の Issue 状態も更新できる（設定次第）

---

## 📌 Reflect CLI 開発におけるメリット

- Linear で設計管理 → GitHub で開発 → 双方向に状態が連携
- Pull Request から「どの設計タスクか」が一目瞭然
- 複数人での開発・レビュー時にもスムーズに把握できる

---

## ✅ 推奨ワークフロー

1. Linear にタスクを起票（例: `✨ zk new のtypeオプション設計`）
2. Linear Issue Key（例: `ZR-45`）を確認
3. ブランチ名・PR名・コミットに `ZR-45` を含める
4. PR作成 → Linearと自動リンク！

---

## 参考リンク

- [Linear × GitHub Integration Guide](https://linear.app/docs/github)
