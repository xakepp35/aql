package ast

import (
	"fmt"
	"strconv"

	"github.com/xakepp35/aql/pkg/vmi"
)

// Eval плашет значения через стек, берет аргументы из state и пушит результат.

func (f *FieldSel) Eval(this vmi.State) error {
	// X.Eval пушит объект
	if err := f.X.Eval(this); err != nil {
		return err
	}
	obj := this.Pop()
	m, ok := obj.(map[string]any)
	if !ok {
		return fmt.Errorf("expected object for field access")
	}
	val := m[string(f.Name)]
	this.Push(val)
	return nil
}

func (i *IndexExpr) Eval(this vmi.State) error {
	// X, I, (optional J)
	if err := i.X.Eval(this); err != nil {
		return err
	}
	if err := i.I.Eval(this); err != nil {
		return err
	}
	var jval any
	if i.J != nil {
		if err := i.J.Eval(this); err != nil {
			return err
		}
		jval = this.Pop()
	}

	idxAny := this.Pop()
	seqAny := this.Pop()

	a, ok := seqAny.([]any)
	if !ok {
		return fmt.Errorf("expected array for index access")
	}

	ixf, ok := idxAny.(float64)
	if !ok {
		return fmt.Errorf("expected numeric index")
	}
	idx := int(ixf)
	if i.J == nil {
		if idx < 0 || idx >= len(a) {
			return fmt.Errorf("index out of bounds")
		}
		this.Push(a[idx])
	} else {
		jxf, ok := jval.(float64)
		if !ok {
			return fmt.Errorf("expected numeric end index")
		}
		jdx := int(jxf)
		// bounds check omitted
		slice := a[idx:jdx]
		this.Push(slice)
	}
	return nil
}

func (p *PipeExpr) Eval(this vmi.State) error {
	// eval left, pop result, set _
	if err := p.Left.Eval(this); err != nil {
		return err
	}
	left := this.Pop()
	prev := this.Get("_")
	this.Set("_", left)
	// eval right pushes result
	if err := p.Right.Eval(this); err != nil {
		this.Set("_", prev)
		return err
	}
	// restore prev
	this.Set("_", prev)
	return nil
}

func (b *BinaryExpr) Eval(this vmi.State) error {
	// left, right
	if err := b.Left.Eval(this); err != nil {
		return err
	}
	if err := b.Right.Eval(this); err != nil {
		return err
	}
	right := this.Pop()
	left := this.Pop()

	lf, lok := left.(float64)
	rf, rok := right.(float64)
	if !lok || !rok {
		return fmt.Errorf("binary op '%s' expects numbers", b.Op)
	}
	var res float64
	switch b.Op {
	case "+":
		res = lf + rf
	case "-":
		res = lf - rf
	case "*":
		res = lf * rf
	case "/":
		res = lf / rf
	case "%":
		res = float64(int64(lf) % int64(rf))
	default:
		return fmt.Errorf("unknown op: %s", b.Op)
	}
	this.Push(res)
	return nil
}

func (c *CompareExpr) Eval(this vmi.State) error {
	if err := c.Left.Eval(this); err != nil {
		return err
	}
	if err := c.Right.Eval(this); err != nil {
		return err
	}
	right := this.Pop()
	left := this.Pop()
	var out bool
	switch c.Op {
	case "==":
		out = left == right
	case "!=":
		out = left != right
	case "<", "<=", ">", ">=":
		lf, lok := left.(float64)
		rf, rok := right.(float64)
		if !lok || !rok {
			return fmt.Errorf("comparison expects numbers")
		}
		switch c.Op {
		case "<":
			out = lf < rf
		case "<=":
			out = lf <= rf
		case ">":
			out = lf > rf
		case ">=":
			out = lf >= rf
		}
	default:
		return fmt.Errorf("unknown compare op: %s", c.Op)
	}
	this.Push(out)
	return nil
}

func (l *LogicalExpr) Eval(this vmi.State) error {
	if err := l.Left.Eval(this); err != nil {
		return err
	}
	lv := this.Pop()
	lb, lok := lv.(bool)
	if !lok {
		return fmt.Errorf("logical left side not bool")
	}
	if l.Op == "||" && lb {
		this.Push(true)
		return nil
	}
	if l.Op == "&&" && !lb {
		this.Push(false)
		return nil
	}
	if err := l.Right.Eval(this); err != nil {
		return err
	}
	rv := this.Pop()
	rb, rok := rv.(bool)
	if !rok {
		return fmt.Errorf("logical right side not bool")
	}
	switch l.Op {
	case "&&":
		this.Push(lb && rb)
	case "||":
		this.Push(lb || rb)
	default:
		return fmt.Errorf("unknown logical op: %s", l.Op)
	}
	return nil
}

func (c *CallExpr) Eval(this vmi.State) error {
	// args in order
	args := make([]any, 0, len(c.Args))
	for _, a := range c.Args {
		if err := a.Eval(this); err != nil {
			return err
		}
		args = append(args, this.Pop())
	}
	// reverse args
	for i, j := 0, len(args)-1; i < j; i, j = i+1, j-1 {
		args[i], args[j] = args[j], args[i]
	}
	for i := range args {
		this.Push(args[i])
	}
	this.Call(string(c.Fun))
	return nil
}

func (u *UnaryExpr) Eval(this vmi.State) error {
	if err := u.X.Eval(this); err != nil {
		return err
	}
	v := this.Pop()
	switch u.Op {
	case "-":
		f, ok := v.(float64)
		if !ok {
			return fmt.Errorf("unary - expects number")
		}
		this.Push(-f)
		return nil
	default:
		return fmt.Errorf("unsupported unary op: %s", u.Op)
	}
}

func (i *Ident) Eval(this vmi.State) error {
	val := this.Get(string(i.Name))
	if val == nil {
		return fmt.Errorf("undefined variable: %s", i.Name)
	}
	this.Push(val)
	return nil
}

func (n *Number) Eval(this vmi.State) error {
	f, err := strconv.ParseFloat(string(n.Text), 64)
	if err != nil {
		return err
	}
	this.Push(f)
	return nil
}

func (s *String) Eval(this vmi.State) error {
	this.Push(string(s.Text))
	return nil
}

func (b *Bool) Eval(this vmi.State) error {
	this.Push(b.Val)
	return nil
}

func (n *Null) Eval(this vmi.State) error {
	this.Push(nil)
	return nil
}

func (o *OverExpr) Eval(this vmi.State) error {
	// просто возвращаем Seq
	if err := o.Seq.Eval(this); err != nil {
		return err
	}
	val := this.Pop()
	this.Push(val)
	return nil
}
