package ast

import "github.com/xakepp35/aql/pkg/vmi"

type (
	package expr

import (
	"github.com/xakepp35/aql/pkg/vm/op"
)

// Expr — корневой интерфейс всех выражений
type Expr interface {
	isExpr()
}



type (

	// Binary: + - * / % == != < <= > >= and or
	Binary struct {
			Op    op.Op
			Left  Expr
			Right Expr
		}

// Unary: - OVER
Unary struct {
	Op   op.Op
	Expr Expr
}

// Over: OVER x => ...
Over struct {
	Expr Expr
	Body Expr // если ARROW есть
}

// Call: f(a, b)
Call struct {
	Name []byte // без копирования
	Args []Expr
}

// Field Access: x.y
Field struct {
	Target Expr
	Name   []byte // без копирования
}

// Index1: arr[i]
Index1 struct {
	Target Expr
	Index  Expr
}

// Index2: arr[i:j]
Index2 struct {
	Target Expr
	Start  Expr
	End    Expr
}

// Pipe: a | b
Pipe struct {
	Left  Expr
	Right Expr
}

// Ident: переменная или встроенная функция
Ident struct {
	Name []byte
}

// Const: числа, строки, true, false, null
Const struct {
	Value any
}

// Dup: точка
Dup struct{}

// Group: (expr)
Group struct {
	Inner Expr
}


	// // комбинируем узлы
	// PipeExpr    struct{ Left, Right vmi.Node }
	// LogicalExpr struct {
	// 	Op          string
	// 	Left, Right vmi.Node
	// }
	// CompareExpr struct {
	// 	Op          string
	// 	Left, Right vmi.Node
	// }
	// BinaryExpr struct {
	// 	Op          string
	// 	Left, Right vmi.Node
	// }
	// UnaryExpr struct {
	// 	Op string
	// 	X  vmi.Node
	// }

	// // селекторы
	// FieldSel struct {
	// 	X    vmi.Node
	// 	Name []byte
	// }
	// IndexExpr struct{ X, I, J vmi.Node }
	// CallExpr  struct {
	// 	Fun  []byte
	// 	Args []vmi.Node
	// }

	// // stateful
	// OverExpr struct{ Seq, Scope vmi.Node }

	// // литералы
	// Number struct{ Text []byte }
	// String struct{ Text []byte }
	// Bool   struct{ Val bool }
	// Null   struct{}
	// Ident  struct{ Name []byte }
)
