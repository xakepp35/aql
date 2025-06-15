// bnf/runtime.go
package ebnf

type Builder interface {
	BuildSeq(xs []any) (any, error)             // последовательность A B C
	BuildAlt(xs []any) (any, error)             // альтернатива A | B | C
	BuildOpt(x any) (any, error)                // [ X ]
	BuildRep(x any) (any, error)                // { X }
	BuildLit(lit []byte) (any, error)           // "true", "false", "=="
	BuildUnit() (any, error)                    // пустое значение
	BuildRef(name []byte, val any) (any, error) // вызов правила
	BuildGroup(inner any) (any, error)          // (A) → сохранить группировку
}

type runtime struct {
	lex     *lexer           // твой (!) zero-alloc лексер
	builder Builder          // динамический коллбек-строитель
	rules   map[string]*Rule // NonTerm  -> Rule
}
