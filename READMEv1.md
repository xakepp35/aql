# **AQL — Any Query Language**  
Stream-first. Stateful. Zero-alloc.

[![Go Report](https://goreportcard.com/badge/github.com/xakepp35/aql)](https://goreportcard.com/report/github.com/xakepp35/aql)
[![CI](https://github.com/xakepp35/aql/actions/workflows/ci.yml/badge.svg)](https://github.com/xakepp35/aql/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/xakepp35/aql)](LICENSE)

> **TL;DR** AQL объединяет лаконичный синтаксис `jq` и стриминговую,
> _stateful_ модель `zq`, но без аллокаций и копий.  
> **Пишем квери и фильтры к бд и очередям - на специальном языке под задачу

---

## ️⚡ Почему ещё один «q»?

|              | `jq` | `zq` | **AQL** |
|--------------|------|------|---------|
| Синтаксис C/JS-like | ✔ | ✖ | ✔ |
| Стриминговый ввод   | ✔ | ✔ | ✔ |
| **Stateful** агрегаты | ✖ (только `-s reduce`) | ✔ | **✔** |
| Zero-alloc / Zero-copy | ✖ | частично | **✔** |
| Unicode NFC сравнение | ✖ | ✔ | **✔** |
| Плагины Go / Wasm     | ✖ | ✖ | soon |

---

## ✨ Ключевые особенности

* **Stateful aggregate-функции**  
  `sum() avg() count() min() max() topk()` работают в потоке — без
  «слёрпа» всего файла в память.

* **`over … => (…)`** — smart-итератор для
  вложенных массивов, вдохновлённый `jq []`, но со
  стриминговыми агрегациями.

* **Zero-alloc runtime**  
  все значения — срезы `[]byte` поверх исходного буфера; любые
  функции, делающие `make/append`, помечены `// alloc`.

* **Полноценный парсер на goyacc**  
  → корректные приоритеты (`*` выше `+`, `&&` выше `||`,
  пайп `|` — самый низкий).

* **Модульная архитектура** (`/internal/lexer`, `/internal/parser`,
  `/internal/engine`) — легко расширять собственными операторами.

---

## 🚀 Быстрый старт

```bash
go install github.com/xakepp35/aql/cmd/aq@latest
echo '{"a":1}{"a":2}{"a":3}' \
  | aq 'sum(.a)'
# → {"sum":6}
````

```bash
# Flatten + aggregate в один проход
echo '[1,2,3] [4,5]' \
  | aq 'over this | sum(this)'
# → {"sum":15}
```

---

## 🧩 Грамматика (выжимка)

```ebnf
Expr        = PipeExpr ;
PipeExpr    = OrExpr { "|" OrExpr } ;
OrExpr      = AndExpr { "||" AndExpr } ;
AndExpr     = CmpExpr { "&&" CmpExpr } ;
CmpExpr     = AddExpr { ("=="|"!="|"<"|"<="|">"|">=") AddExpr } ;
AddExpr     = MulExpr { ("+"|"-") MulExpr } ;
MulExpr     = UnaryExpr { ("*"|"/"|"%") UnaryExpr } ;
UnaryExpr   = ["-"|"!"] Primary
           | "over" UnaryExpr
           | "over" UnaryExpr "=>" "(" Expr ")" ;
Primary     = Lit | Ident | Sel | Index | Call | "(" Expr ")" ;
```

Полный язык описан в [docs/spec.md](docs/spec.md).

---

## 🛠 Сборка из исходников

```bash
git clone https://github.com/xakepp35/aql
cd aql
go generate ./internal/parser   # генерация goyacc
go vet ./...
go test ./...
go run ./cmd/aq '1+2*3'
```

---

## 🏎️ Performance (Intel i7-12700, Go 1.22)

| Dataset                      | jq 1.6   | AQL (v0.1)   | Δ    |
| ---------------------------- | -------- | ------------ | ---- |
| 1 GB NDJSON, `sum(.value)`   | 520 MB/s | **1.7 GB/s** | ×3.3 |
| 10 GB ZNG, `count() by user` | –        | **5.4 GB/s** | –    |

Бенчмарки reproducible: `go test -bench ./bench`.

---

## 📅 Roadmap (v0.2 → v1)

* оконные агрегаты `window(size, step)`
* plug-in API (`go plugin` + WebAssembly)
* полноцветная ошибка с подсветкой фрагмента
* официальные пакеты `brew/apt/yay`
* интерактивный REPL (`aq -i`)

См. [projects/roadmap.md](projects/roadmap.md) для деталей.

---

## 🤝 Contributing

1. Fork → Branch → PR.
2. `go test ./...` должен быть зелёным.
3. Любой код, вызывающий `make/append/copy`, обязательно помечайте
   `// alloc` или `// copy`.

### Code of Conduct

Мы придерживаемся [Contributor Covenant](CODE_OF_CONDUCT.md).

---

## 📜 Лицензия

MIT © 2025 AQL Core Team

```

*Сделано с 💙 к stream-processing и zero-alloc-Go.*
```
