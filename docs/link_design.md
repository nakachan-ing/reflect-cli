# Reflect CLI - ノート間リンクの設計方針

このドキュメントでは、Reflect CLI におけるノート間リンク（linked_notes）の設計方針について整理する。特に、`fleeting - fleeting` 間のリンクを正当なものとして扱う根拠と活用方法を記載する。

---

## ✅ 想定されるノートリンクの種類

| リンクの種類 | 用途・意味 | 備考 |
|--------------|------------|------|
| `fleeting → permanent` | Reflectの出力元と昇華先 | `reflected_by:` や `linked_note:` で管理 |
| `permanent → permanent` | 知識の構造化・比較 | インデックスノート、構造ノートなど |
| `fleeting → fleeting` ✅ | 思考・調査の連鎖、関連ノートの追跡 | 新しい問いや調査が分岐する場合に自然に発生 |

---

## 🎯 `fleeting → fleeting` リンクを許容する理由

- 調査や問いが次のfleetingを生むことは自然な思考プロセス
- 1つのfleetingが、複数の問いや仮説を派生させるケースが多い
- Reflect CLIは「思考の記録 → 再訪 → 昇華」の連続を支援する設計なので、途中のつながりも記録すべき

---

## 🧩 Reflect CLIにおける実装方針

### Frontmatterでのリンク記述例

```yaml
linked_notes:
  - 20250402T1030_gc調査.md
  - 20250403T1215_heapregionsize_question.md
```

- `linked_notes` は **fleeting/permanent問わず使用可能**
- CLI側ではリンク先ノートの存在と `type` を内部的に確認するが、明示的な制限は設けない

---

## 💡 Reflect CLIの将来的な展望

- `zk link add <noteA> <noteB>` のような軽量リンク生成コマンド
- `zk structure` で思考のネットワークを可視化（グラフ構造表示）
- `zk list --linked-to <note>` による逆リンク探索
- Reflect済みノート間の関係性と、Reflect前の流れの可視化

---

## ✅ 結論

Reflect CLI は、「思考の流れそのものを記録し、再訪し、構造化する」ことを目的としたツールです。  
そのため、`fleeting → fleeting` のような **未昇華状態の思考同士のつながり** を記録・可視化することは非常に価値あり。

Reflect CLIでは `linked_notes:` によって、自由かつ明確にこれを実現する。
