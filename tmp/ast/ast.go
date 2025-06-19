package ast

import (
	"github.com/xakepp35/aql/tmp/token"
	"strconv"
)

type Expr interface {
	Eval() int
}

type Number struct {
	Value int
}

func NewNumber(t interface{}) (*Number, error) {
	lit := t.(*token.Token).Lit
	v, err := strconv.Atoi(string(lit))
	if err != nil {
		return nil, err
	}
	return &Number{Value: v}, nil
}

func (n *Number) Eval() int {
	return n.Value
}

type Add struct {
	Left, Right Expr
}

func NewAdd(l, r interface{}) (*Add, error) {
	return &Add{
		Left:  l.(Expr),
		Right: r.(Expr),
	}, nil
}

func (a *Add) Eval() int {
	return a.Left.Eval() + a.Right.Eval()
}
