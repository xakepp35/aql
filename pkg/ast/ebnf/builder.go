package ebnf

import (
	"github.com/xakepp35/aql/pkg/ast/asi"
	"github.com/xakepp35/aql/pkg/ast/expr"
)

type ASTBuilder struct {
	arena *expr.Arena
}

func NewASTBuilder(a *expr.Arena) *ASTBuilder {
	return &ASTBuilder{arena: a}
}

func (b *ASTBuilder) BuildSeq(xs []any) (any, error) {
	if len(xs) == 0 {
		return b.arena.Literal(nil), nil
	}
	if len(xs) == 1 {
		return xs[0], nil
	}
	cur := xs[0].(asi.AST)
	for _, x := range xs[1:] {
		next := x.(asi.AST)
		cur = b.arena.Binary(cur, next, "_seq")
	}
	return cur, nil
}

func (b *ASTBuilder) BuildAlt(xs []any) (any, error) {
	if len(xs) == 0 {
		return b.arena.Literal(nil), nil
	}
	cur := xs[0].(asi.AST)
	for _, x := range xs[1:] {
		next := x.(asi.AST)
		cur = b.arena.Binary(cur, next, "_alt")
	}
	return cur, nil
}

func (b *ASTBuilder) BuildOpt(x any) (any, error) {
	// wrap: x | unit
	unit, _ := b.arena.Literal(nil)
	return b.BuildAlt([]any{x, unit})
}

func (b *ASTBuilder) BuildRep(x any) (any, error) {
	// wrap: { x } => повторение
	return b.arena.Unary(x.(asi.AST), "_rep")
}

func (b *ASTBuilder) BuildLit(lit []byte) (any, error) {
	return b.arena.Literal(lit)
}

func (b *ASTBuilder) BuildUnit() (any, error) {
	return b.arena.Literal(nil)
}

func (b *ASTBuilder) BuildRef(name []byte, val any) (any, error) {
	return b.arena.Unary(val.(asi.AST), name)
}
