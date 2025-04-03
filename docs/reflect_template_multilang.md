# Zettelkasten Reflect CLI - Reflectテンプレートの構造と多言語対応設計

このドキュメントでは、Zettelkasten Reflect CLI における Reflectテンプレートのファイル構造と、将来のOSS展開を見据えた多言語対応の設計方針をまとめる。

---

## ✅ 目的

- 各 `subtype`（idea, question, investigation など）に対して Reflectプロンプトを用意
- CLI内でテンプレートを読み込み、対話形式で抽象化を支援
- 日本語・英語の **多言語対応** を前提としたテンプレート設計を行う

---

## 🗂️ ディレクトリ構造案（多言語対応）

```plaintext
templates/
└── reflect/
    ├── ja/
    │   ├── idea.yaml
    │   ├── investigation.yaml
    └── en/
        ├── idea.yaml
        ├── investigation.yaml
```

各言語ごとに `subtype.yaml` を配置することで、動的に言語切り替えが可能となる。

---

## 🧩 YAMLテンプレート構造（共通）

```yaml
type: idea
title: 気づきの抽象化テンプレート
abstract_dimensions:
  - 文脈
  - 洞察
  - 類似事例
prompts:
  - Q: このアイデアが生まれた背景や文脈は？
  - Q: 本質的な洞察や原則は何か？
  - Q: 他に似たような事例はあるか？
```

英語バージョン（例: `en/idea.yaml`）:

```yaml
type: idea
title: Template for abstracting ideas
abstract_dimensions:
  - Context
  - Insight
  - Analogous Cases
prompts:
  - Q: What context led to this idea?
  - Q: What is the core insight or principle?
  - Q: Are there similar known cases?
```

---

## ⚙️ Goでのテンプレート読み込み例

```go
type ReflectTemplate struct {
    Type               string   `yaml:"type"`
    Title              string   `yaml:"title"`
    AbstractDimensions []string `yaml:"abstract_dimensions"`
    Prompts            []Prompt `yaml:"prompts"`
}

type Prompt struct {
    Q string `yaml:"Q"`
}

func LoadTemplate(subtype string, lang string) (*ReflectTemplate, error) {
    if lang == "" {
        lang = "ja" // デフォルト言語
    }
    path := fmt.Sprintf("templates/reflect/%s/%s.yaml", lang, subtype)
    content, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var tmpl ReflectTemplate
    err = yaml.Unmarshal(content, &tmpl)
    return &tmpl, err
}
```

---

## 📋 CLIでの言語指定方法（例）

```bash
zk reflect note.md --lang en
```

または、`.zkconfig.json` に記録：

```json
{
  "lang": "ja",
  "backend": "json"
}
```

---

## ✅ この設計のメリット

| 項目           | 内容 |
|----------------|------|
| OSS向け展開     | 英語対応により国際的なコントリビューターも参加しやすい |
| 拡張性         | 他言語（例: fr, zh）の追加も容易に行える |
| ユーザー体験   | CLIを自分の言語で使えることで学習・思考がスムーズになる |
| 将来のGUI展開  | Reflectテンプレートをフロントエンドに取り込むことも簡単になる |

---

## ✅ 結論

- Reflectテンプレートは **YAMLファイル形式 + 言語ごとのディレクトリ構造** によって管理
- CLI実行時に `lang` を指定することで、**多言語プロンプトでのReflectが可能**
- OSSとしての将来性・ユーザー体験の両立を実現する構造

