# go-scrape

---
## 仕様
- ブラウザの自動操作アプリである
- JSON形式のプロファイルに沿ってブラウザの操作を行う
- 実行時間の指定が可能
- 実行回数（1...n）の指定が可能
- Headlessブラウザ、ブラウザのデータの指定が可能

## プロファイルについて
### 形式
```json
[
  {
    "name": "一意となる名前",
    "repeat": "1",
    "args": {
      "arg1": "arg1",
      "arg2": "arg2"
    },
    "operation": {
      "wakeUp": {
        "date": "YYYY/MM/DD",
        "time": "hh:mm"
      },
      "url": "エントリーポイントのURL",
      "control": [
        { "操作": "操作内容" }
      ]
    }
  }
]
```

### `control`で指定できるアクションと操作内容一覧
| アクション名 |    操作内容    |
|:------:|:----------:|
| click  |  CSSセレクタ   |
| doubleClick  |  CSSセレクタ   |
| input  |  入力する文字列   |
| select  | 選択するオプション名 |
| sendKey  | 送信したいキー入力  |
| wait  | 待機時間[msec] |
| screenShot  |    保存名     |
| to  |   遷移先URL   |
| reload  |     -      |

## TODO
1. profileの充実
   - repeatのオプションを追加
     - 指定された回数のリピート
     - 無指定だと1回
     - 数値または、"n"を指定できる
       - 数値→その回数分の繰り返し
       - n→無限ループ
   - varのオプションを追加
     - オプションの内容
       - {args1}、{ARGS1}等で置換可能
       - "var" [{ "variable": "variable content" }, { "ARGS1": "ARGS1 content" }]
     - repeatの繰り返し回数ごと({N})等で痴漢できるようにする
     - "assign"でセレクタの結果を代入
       - **operation内の書き方を再考しないと使いやすさに影響有り**
       - `"assign": { "${variableName}": "CSS selector" }` 
   - 条件分岐の考慮
2. アクションの追加
   - get
     - 指定したセレクタの文言を取得する
     - セレクタを指定？保存先は？

3. 制御内の構文を考え直し
- 案1 
```json
{
  "control": [
    {
      "#selector1": [
        {"do1": ""},
        {"do2": ""},
        {"do3": ""}
      ]
    },
    {
      "#selector2": [
        {"do4": ""},
        {"do5": ""},
        {"do6": ""}
      ]
    }
  ]
}
```

```go
type Control []Selector

type Selector struct {
    selector string
	actions []Action
}

type Action struct {
	name string
	contents string
}

```

- 案2
```json
{
  "control": [
    {
      "action": {
        "target": "",
        "contents": ""
      }
    },
    {
      "action2": {
        "target": "",
        "contents": ""
      }
    }
  ]
}
```

```go
type Control []

type Selector struct {
    selector string
	actions []Action
}

type Action struct {
	name string
	contents string
}

```

- 案3
```json
{
  "control": [
    {"click": "CSS Selector"},
    {"doubleClick": "CSS Selector"},
    {"input": {
      "target": "CSS Selector",
      "text": "input text"
    }},
    {"select": {
      "target": "CSS Selector",
      "text": "select target text"
    }},
    {"sendKey": "want send key"},
    {"wait": "wait time"},
    {"screenShot": "save path"},
    {"reload": ""},
    
    <!-- ↓新規追加分 -->
    {"exit": ""},
    {"cmd": ""},
    {"assign-text": { 
      "variableName": "CSS Selector"
    }},
    {"assign-value": { 
      "variableName": "CSS Selector"
    }},
    {"if": {
        "condition1": [
          { "actions": "action contents" }
        ]
      }
    }
  ]
}
```