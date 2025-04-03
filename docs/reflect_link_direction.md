# Zettelkasten Reflect CLI - Reflectリンクの方向性と設計方針

このドキュメントでは、Zettelkasten Reflect CLI における Fleeting ノートと Permanent ノート間のリンク（Reflect）に関する設計方針を整理する。

---

## ✅ 背景：Reflectにおけるリンクの方向性

Reflect CLIでは、fleeting ノートを元に permanent ノートを生成（Reflect）する。  
このとき、**どちらのノートにどのようなリンク情報を持たせるか** が設計上の論点である。

---

## 🔄 想定されるリンク方向

| ノート種類 | リンク項目 | 説明 |
|------------|------------|------|
| Permanent  | `linked_note` | Reflect元となった Fleetingノートへのリンク（必須） |
| Fleeting   | `reflected_by` | Reflectで生成された Permanentノートへのリンク（任意） |

---

## ✅ `reflected_by` をFleetingに記録するメリット

- `zk list --reflectable` などで Reflect済みかどうかを判定しやすくなる
- どの Permanent に反映されたか後からたどりやすい
- Reflect済みかどうかを明示するフラグとしても使える

---

## 💭 しかし記録しない場合の利点もある

- 1つの Fleeting から複数の Permanent が生まれるケース（split）に柔軟に対応
- Reflect後の Permanent が削除・再生成されたときの整合性を考えなくて済む
- CLIの処理がシンプルになる（双方向リンクの更新が不要）

---

## 🧭 Reflect CLIとしての推奨設計（v0.1時点）

| 方針項目 | 内容 |
|----------|------|
| Reflectリンクの方向性 | Permanentノートにだけ `linked_note` を記録 |
| Fleeting側のリンク | `reflected_by` は記録しない（将来的な拡張候補） |
| Reflect済み判定方法 | Permanentノートの `linked_note` を解析して判断 |
| 将来的な拡張 | `zk reflect --update-source` で Fleeting 側にも記録可能にする余地を残す |

---

## ✅ 結論

Reflect CLI は、「思考の昇華」に集中したシンプルな設計を目指す。  
そのため、**Reflectリンクは片方向（Permanent → Fleeting）に限定**することで、初期実装の柔軟性・簡潔性を保つ。

- フル双方向リンクは将来的にニーズが明確になった段階で対応を検討する。
- Reflectの非線形性（1:n, 再Reflectなど）にも柔軟に対応できる構造を維持する。

