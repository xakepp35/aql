package ast

import "github.com/xakepp35/aql/pkg/vmi"

type (
	// комбинируем узлы
	PipeExpr    struct{ Left, Right vmi.Node }
	LogicalExpr struct {
		Op          string
		Left, Right vmi.Node
	}
	CompareExpr struct {
		Op          string
		Left, Right vmi.Node
	}
	BinaryExpr struct {
		Op          string
		Left, Right vmi.Node
	}
	UnaryExpr struct {
		Op string
		X  vmi.Node
	}

	// селекторы
	FieldSel struct {
		X    vmi.Node
		Name []byte
	}
	IndexExpr struct{ X, I, J vmi.Node }
	CallExpr  struct {
		Fun  []byte
		Args []vmi.Node
	}

	// stateful
	OverExpr struct{ Seq, Scope vmi.Node }

	// литералы
	Number struct{ Text []byte }
	String struct{ Text []byte }
	Bool   struct{ Val bool }
	Null   struct{}
	Ident  struct{ Name []byte }
)
