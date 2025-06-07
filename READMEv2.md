# AQ MVP Specification v0.1

## 1 Purpose & Scope

`aq` is a stream‑oriented query/transform CLI that combines the *stateful* data‑flow model of **zq** with the familiar C/JS‑style syntax of **jq** while targeting **zero‑alloc / zero‑copy** performance in Go. This document defines the **minimum viable product (MVP)** behaviour, grammar, type‑system, operator set and implementation contracts.

---

## 2 Data Model

### 2.1 Scalar Types (one‑to‑one with Go primitives)

| Category    | Go kind                    | Literal examples       |
| ----------- | -------------------------- | ---------------------- |
| Boolean     | `bool`                     | `true`, `false`        |
| Signed ints | `int`, `int8` … `int64`    | `-42`, `0x2A`          |
| Unsigned    | `uint`, `uint8` … `uint64` | `23u`                  |
| Floating    | `float32`, `float64`       | `3.14`, `1e-9`         |
| Complex     | `complex64`, `complex128`  | `1+2i`                 |
| Byte        | `byte` (`uint8`)           | `'A'`, `0x41`          |
| Rune        | `rune` (`int32`)           | `'λ'`                  |
| String      | UTF‑8 `string`             | `"hello"`              |
| Time        | `time.Time` RFC‑3339       | `2025-06-06T12:34:56Z` |
| Nil         | Go `nil`                   | `null`                 |

### 2.2 Composite Types

* **Array** `[N]T` – fixed length
* **Slice** `[]T` – variable
* **Map** `map[string]any`
* **Struct** anonymous record `{x:1, y:"bar"}`
* **Interface** `any` – carried as **tagged union** in runtime `Value`

> **Zero‑alloc rule:** All composite values are *views* over the original buffer unless an operation is explicitly marked `// alloc` or `// copy`.

---

## 3 Syntax (EBNF excerpt)

```
Expr        = PipeExpr ;
PipeExpr    = OrExpr { "|" OrExpr } ;
OrExpr      = AndExpr { "||" AndExpr } ;
AndExpr     = CmpExpr { "&&" CmpExpr } ;
CmpExpr     = AddExpr { ( "==" | "!=" | "<" | "<=" | ">" | ">=" ) AddExpr } ;
AddExpr     = MulExpr { ( "+" | "-" ) MulExpr } ;
MulExpr     = UnaryExpr { ( "*" | "/" | "%" ) UnaryExpr } ;
UnaryExpr   = [ "!" | "-" | "~" ] Primary ;
Primary     = Lit | Ident | PathSel | Index | Call | Group ;
PathSel     = Primary "." Ident ;
Index       = Primary "[" Expr [":" Expr ] "]" ;
Call        = Ident "(" [ ArgList ] ")" ;
Group       = "(" Expr ")" ;
Lit         = IntLit | FloatLit | StringLit | BoolLit | NullLit ;
```

Operator precedence matches Go (highest → lowest): **index**, unary, `* / %`, `+ -`, comparisons, `&&`, `||`, pipe `|` (lowest but left‑assoc).

---

## 4 Operator Semantics

| Syntax            | Opcode    | Description                 | `//alloc` notes              |            |   |
| ----------------- | --------- | --------------------------- | ---------------------------- | ---------- | - |
| \`                | \`        | 0x01                        | Pipe current stream into RHS | none       |   |
| `.`               | 0x02      | Field/attr access           | none                         |            |   |
| `[]`              | 0x03      | Index/slice/iterate         | none (slice view)            |            |   |
| `+`               | 0x10      | Addition / concat           | may alloc string             |            |   |
| `-`               | 0x11      | Subtraction                 | –                            |            |   |
| `*`               | 0x12      | Multiplication              | –                            |            |   |
| `/`               | 0x13      | Division                    | –                            |            |   |
| `%`               | 0x14      | Modulo                      | –                            |            |   |
| `== != < > <= >=` | 0x20–0x25 | Comparisons (Unicode NFC)   | may alloc for norm           |            |   |
| `&&`              | 0x30      | Logical AND (short‑circuit) | –                            |            |   |
| \`                |           | \`                          | 0x31                         | Logical OR | – |
| `over`            | 0x40      | Stateful iterator           | –                            |            |   |
| `=>`              | 0x41      | Lateral scope begin/end     | –                            |            |   |

> **Static dispatch:** Each opcode maps to an entry in `var OpFuncs [256]OpFunc`.  *No* `switch` in hot path.

---

## 5 Aggregate & Window Functions

* `sum(expr)`
* `avg(expr)`
* `count([expr])`
* `min(expr)` / `max(expr)`
* `stddev(expr)`
* `percentile(expr, p)`
* `group_by(expr)` → keyed buckets
* `distinct(expr)` / `distinct_count(expr)`
* `topk(expr, k)`
* `window(size, step)` modifiers (tumbling / sliding)

All aggregates expose incremental `Step` and `Result` without extra alloc.  Bucket storage uses arena pages; only `distinct`/`topk` may allocate.

---

## 6 Execution Model

1. **Stream** of decoded `Value`s enters root pipeline.
2. Parser → AST → compiled **plan** of opcode array + constants table.
3. Runtime has

   ```go
   type State struct {
       Acc   [MaxOps]uintptr // per‑op scratch ptr
       Arena *bytearena.Arena
   }
   ```
4. Engine fetches opcode, calls `OpFuncs[op](&state, v)`; function may emit 0..N outputs (generator pattern via closure or channel).  No reflection.
5. **Yield contract**: A function must *not* allocate unless flagged.  If allocation is unavoidable, implementation **must** carry `// alloc` or `// copy` on the first line.

---

## 7 Memory & Performance Contracts

* All paths use borrowed `[]byte` slices; strings created via `unsafe.String` when safe (Go 1.22).
* `Arena` provides page‑sized bump allocator reused per input chunk.
* CI has linter that scans for stray `make(` / `append(` / `copy(` without explicit comment.
* Target throughput: **>400 MiB/s** NDJSON on Intel i7‑12700.

---

## 8 CLI Behaviour

```
Usage: aq [-j|-z] [-o json|text|table] [--pretty] 'QUERY' [file …]
  -j   treat input as NDJSON (default)
  -z   treat input as ZNG
  -s   slurp entire input into single array value (compat)
  -n   start with null input, rely on generators (compat with jq)
```

Non‑flag arguments are files, `-` is stdin. Output goes to stdout line‑delimited JSON by default.

---

## 9 Error Semantics

* **Syntax** → abort before execution, exit 64.
* **Type** → emit diagnostic with byte‑offset + snippet, exit 65.
* **Runtime** → pipeline stops, prints first error, exit 66.
* Panic in `OpFunc` ⇒ recover, wrap, exit 67.

---

## 10 Gherkin Features

### Feature: Basic arithmetic

```
Scenario: Add numbers
  Given stream "1 2 3"
  When I run aq " . + 1 "
  Then the output lines are "2", "3", "4"
```

### Feature: Stateful sum across stream

```
Scenario: Sum NDJSON numbers
  Given file numbers.jsonl with lines "1", "2", "3"
  When I run aq "sum(this)"
  Then stdout is "{\"sum\":6}"
```

### Feature: Over + lateral scope on arrays

```
Scenario: Flatten array of arrays and sum
  Given stream "[1,2,3] [4,5,6]"
  When I run aq "over this | sum(this)"
  Then stdout is "{\"sum\":21}"
```

### Feature: Unicode‑aware group\_by

```
Scenario: Equivalent Unicode keys group together
  Given stream "{\"name\":\"Władysław Stępniak\"} {\"name\":\"Władysław Stępniak\"}"
  When I run aq "count() by name"
  Then stdout has one record with "count":2
```

### Feature: Zero‑alloc audit

```
Scenario: Engine has no hidden allocations
  Given benchmark "1..1e6"
  When I profile mem
  Then peak allocations ≤ 10 MiB
```

---

## 11 Open Items (out‑of‑scope for MVP)

* Binary formats (CBOR, MessagePack)
* Plugin ABI
* Parallel execution
* WASM build

---

© 2025 Bro‑Lang core team.
