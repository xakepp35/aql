package expr

import (
	"unsafe"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/ast/asi"
)

type ExprRaw = [64]byte

type Arena []ExprRaw

// New возвращает свежую арену
func NewArena(capHist uint) *Arena {
	raw := make(Arena, 0, capHist)
	return &raw
}

//go:inline
func (a *Arena) Reset() {
	*a = (*a)[:0]
}

//go:inline
func (a *Arena) alloc() *ExprRaw {
	*a = append(*a, ExprRaw{})
	return &(*a)[len(*a)-1]
}

// ==== Обёртки для каждого типа ====

func (a *Arena) Binary(lhs, rhs asi.AST, op op.Code) *Binary {
	res := As[Binary](a.alloc())
	*res = Binary{
		Args: [2]asi.AST{lhs, rhs},
		Op:   op,
	}
	return res
}

func (a *Arena) Unary(arg asi.AST, op op.Code) *Unary {
	res := As[Unary](a.alloc())
	*res = Unary{Arg: arg, Op: op}
	return res
}

func (a *Arena) Pipe(lhs, rhs asi.AST) *Pipe {
	res := As[Pipe](a.alloc())
	*res = Pipe{Args: [2]asi.AST{lhs, rhs}}
	return res
}

func (a *Arena) Field(base asi.AST, name []byte) *Field {
	res := As[Field](a.alloc())
	*res = Field{Arg: base, Name: name}
	return res
}

func (a *Arena) Literal(x any) *Literal {
	res := As[Literal](a.alloc())
	*res = Literal{X: x}
	return res
}

func (a *Arena) Ident(name []byte) *Ident {
	res := As[Ident](a.alloc())
	*res = Ident{Name: name}
	return res
}

func (a *Arena) Dup() *Dup {
	res := As[Dup](a.alloc())
	*res = Dup{}
	return res
}

func (a *Arena) Over(iter asi.AST, e asi.AST) *Over {
	res := As[Over](a.alloc())
	*res = Over{Iter: iter, Expr: e}
	return res
}

func (a *Arena) Ternary(a1, a2, a3 asi.AST, op op.Code) *Ternary {
	res := As[Ternary](a.alloc())
	*res = Ternary{Args: [3]asi.AST{a1, a2, a3}, Op: op}
	return res
}

func (a *Arena) Call(args []asi.AST, name []byte) *Call {
	res := As[Call](a.alloc())
	*res = Call{Name: name, Args: args}
	return res
}

func As[T any](b *ExprRaw) *T {
	return (*T)(unsafe.Pointer(b))
}
