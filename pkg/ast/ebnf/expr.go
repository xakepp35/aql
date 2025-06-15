package ebnf

import (
	"errors"
)

func (c Choice) parse(rt *runtime) (any, error) {
	save := rt.lex.Pos()
	var results []any
	var lastErr error
	for _, alt := range c {
		n, err := alt.parse(rt)
		if err == nil {
			results = append(results, n)
			rt.lex.Restore(save)
			continue
		}
		rt.lex.Restore(save)
		lastErr = err
	}
	if len(results) > 0 {
		return rt.builder.BuildAlt(results)
	}
	return nil, lastErr
}

func (s Seq) parse(rt *runtime) (any, error) {
	var list []any
	for _, x := range s {
		n, err := x.parse(rt)
		if err != nil {
			return nil, err
		}
		list = append(list, n)
	}
	if len(list) == 0 {
		return rt.builder.BuildUnit()
	}
	if len(list) == 1 {
		return list[0], nil
	}
	return rt.builder.BuildSeq(list)
}

func (r Repeat) parse(rt *runtime) (any, error) {
	var list []any
	for {
		save := rt.lex.Pos()
		n, err := r.X.parse(rt)
		if err != nil {
			rt.lex.Restore(save)
			break
		}
		list = append(list, n)
	}
	if len(list) == 0 {
		return rt.builder.BuildUnit()
	}
	return rt.builder.BuildRep(rt.builder.BuildSeq(list))
}

func (o Option) parse(rt *runtime) (any, error) {
	save := rt.lex.Pos()
	n, err := o.X.parse(rt)
	if err != nil {
		rt.lex.Restore(save)
		return rt.builder.BuildUnit()
	}
	return rt.builder.BuildOpt(n)
}

func (t Terminal) parse(rt *runtime) (any, error) {
	if rt.lex.AcceptLit(t.Lit) {
		return rt.builder.BuildLit(t.Lit)
	}
	return nil, errors.New("token mismatch")
}

func (n NonTerm) parse(rt *runtime) (any, error) {
	r := rt.rules[string(n.Name)]
	if r == nil {
		return nil, errors.New("unknown rule")
	}
	body, err := r.Expr.parse(rt)
	if err != nil {
		return nil, err
	}
	return rt.builder.BuildRef(n.Name, body)
}
