# Reflect CLI - 日本語タイトルにおける slug 生成戦略

本ドキュメントでは、`zk new fleeting --title` における日本語タイトル入力時の slug 生成に関する課題と対応方針をまとめる。

---

## 🎯 背景

Reflect CLI では、タイトルが指定された場合、識別しやすいファイル名のために `slug` を生成し、Markdownファイル名に含める。

```plaintext
fleeting/20250402_idea_g1gc-tuning.md
```

ただし、日本語タイトルを扱う際に、使用中の `github.com/gosimple/slug` ライブラリでは次のような課題あり。

---

## ❌ 問題点：日本語が中国語の拼音読みとして変換されてしまう

| 入力タイトル | 期待されるslug | slugパッケージの出力 |
|--------------|----------------|------------------------|
| 巨大オブジェクト | kyodai-obujekuto | ju-da |
| 調査メモ         | chousa-memo        | diao-cha |
| 日本語だけのタイトル | slugなし or ローマ字 | ピンイン表記になる |

---

## ✅ Reflect CLI の対応方針

### 1. 英語が含まれている場合

- `slug.Make(title)` を使ってそのまま slug 生成

### 2. ひらがな・カタカナのみの場合

- 自作 or 軽量ライブラリでローマ字に変換して slug に使用
    - 例: `カスタム辞書変換`, `github.com/ebc-2in2crc/kana` など

### 3. 漢字が含まれる場合（かつ英語がない）

- 現時点では slug 生成をスキップ
- ファイル名は `yyyymmdd_subtype.md` の形式とする

---

## 🛠 実装例（Go）

```go
func GenerateSlug(title string) string {
    if hasASCII(title) {
        return slug.Make(title)
    }

    if isKanaOnly(title) {
        return kana.ToRomaji(title)
    }

    return "" // slugなし
}
```

---

## 💡 補足：将来的な拡張案

- `kagome` などを使って形態素解析し、意味ベースで変換する
- よく使われる単語だけに特化した `手作業マッピング辞書` を導入する

---

## ✅ 結論

Reflect CLI では slug の自動生成を行うが、日本語に関しては以下のような柔軟な戦略を取る：

- 英語があればそれを優先して slug に使う
- 日本語だけのときは、ローマ字化を試みる or slugなしにする
- slugなしでも CLI は問題なく動作する設計にする

この方針により、多言語ユーザーでも自然に CLI を運用できるようになる。
