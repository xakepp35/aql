package ast

import (
	"github.com/xakepp35/aql/pkg/ast/asi"
	"github.com/xakepp35/aql/pkg/ast/expr"
	"github.com/xakepp35/aql/pkg/vm/op"
)

type Arena struct {
	binaries []expr.Binary
	literals []expr.Literal
	fields   []expr.Field
	calls    []expr.Call
	pipes    []expr.Pipe
	dups     []expr.Dup
	overs    []expr.Over
	unarys   []expr.Unary
	idents   []expr.Ident
	ternarys []expr.Ternary
}

// New возвращает свежую арену
func NewArena() *Arena {
	return &Arena{
		binaries: make([]expr.Binary, 0, 64),
		literals: make([]expr.Literal, 0, 64),
		fields:   make([]expr.Field, 0, 64),
		calls:    make([]expr.Call, 0, 64),
		pipes:    make([]expr.Pipe, 0, 64),
		dups:     make([]expr.Dup, 0, 64),
		overs:    make([]expr.Over, 0, 64),
		unarys:   make([]expr.Unary, 0, 64),
		idents:   make([]expr.Ident, 0, 64),
		ternarys: make([]expr.Ternary, 0, 64),
	}
}

func (a *Arena) Reset() {
	a.binaries = a.binaries[:0]
	a.literals = a.literals[:0]
	a.fields = a.fields[:0]
	a.calls = a.calls[:0]
	a.pipes = a.pipes[:0]
	a.dups = a.dups[:0]
	a.overs = a.overs[:0]
	a.unarys = a.unarys[:0]
	a.idents = a.idents[:0]
	a.ternarys = a.ternarys[:0]
}

// ==== Обёртки для каждого типа ====

func (a *Arena) Binary(lhs, rhs asi.AST, op op.Code) *expr.Binary {
	a.binaries = append(a.binaries, expr.Binary{
		Args: [2]asi.AST{lhs, rhs},
		Op:   op,
	})
	return &a.binaries[len(a.binaries)-1]
}

func (a *Arena) Unary(arg asi.AST, op op.Code) *expr.Unary {
	a.unarys = append(a.unarys, expr.Unary{Arg: arg, Op: op})
	return &a.unarys[len(a.unarys)-1]
}

func (a *Arena) Pipe(lhs, rhs asi.AST) *expr.Pipe {
	a.pipes = append(a.pipes, expr.Pipe{Args: [2]asi.AST{lhs, rhs}})
	return &a.pipes[len(a.pipes)-1]
}

func (a *Arena) Field(base asi.AST, name []byte) *expr.Field {
	a.fields = append(a.fields, expr.Field{Arg: base, Name: name})
	return &a.fields[len(a.fields)-1]
}

func (a *Arena) Literal(x any) *expr.Literal {
	a.literals = append(a.literals, expr.Literal{X: x})
	return &a.literals[len(a.literals)-1]
}

func (a *Arena) Ident(name []byte) *expr.Ident {
	a.idents = append(a.idents, expr.Ident{Name: name})
	return &a.idents[len(a.idents)-1]
}

func (a *Arena) Dup() *expr.Dup {
	a.dups = append(a.dups, expr.Dup{})
	return &a.dups[len(a.dups)-1]
}

func (a *Arena) Over(iter asi.AST, e asi.AST) *expr.Over {
	a.overs = append(a.overs, expr.Over{Iter: iter, Expr: e})
	return &a.overs[len(a.overs)-1]
}

func (a *Arena) Ternary(a1, a2, a3 asi.AST, op op.Code) *expr.Ternary {
	a.ternarys = append(a.ternarys, expr.Ternary{Args: [3]asi.AST{a1, a2, a3}, Op: op})
	return &a.ternarys[len(a.ternarys)-1]
}

func (a *Arena) Call(args []asi.AST, name []byte) *expr.Call {
	a.calls = append(a.calls, expr.Call{Name: name, Args: args})
	return &a.calls[len(a.calls)-1]
}
