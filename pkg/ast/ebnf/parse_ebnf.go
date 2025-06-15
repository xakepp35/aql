package ebnf

import (
	"bytes"
	"errors"
)

type (
	Rule struct {
		Name []byte
		Expr Expr
	}

	Expr interface {
		parse(*runtime) (any, error)
	}

	Choice   []Expr
	Seq      []Expr
	Repeat   struct{ X Expr }
	Option   struct{ X Expr }
	Terminal struct{ Lit []byte }
	NonTerm  struct{ Name []byte }
)

func ParseEBNF(src []byte) ([]Rule, error) {
	var rules []Rule
	lines := bytes.Split(src, []byte{'\n'})
	for _, raw := range lines {
		line := bytes.TrimSpace(raw)
		if len(line) == 0 || line[0] == ';' {
			continue
		}
		parts := bytes.SplitN(line, []byte("::="), 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid rule: missing ::= ")
		}
		name := bytes.TrimSpace(parts[0])
		rhs := bytes.TrimSpace(parts[1])

		expr, err := parseExpr(rhs)
		if err != nil {
			return nil, err
		}
		rules = append(rules, Rule{
			Name: name,
			Expr: expr,
		})
	}
	return rules, nil
}

func parseExpr(src []byte) (Expr, error) {
	var parts []Expr
	var part [][]byte
	level := 0
	start := 0

	for i := 0; i < len(src); i++ {
		c := src[i]
		switch c {
		case '|':
			if level == 0 {
				part = append(part, bytes.TrimSpace(src[start:i]))
				start = i + 1
			}
		case '[', '{':
			level++
		case ']', '}':
			level--
		}
	}
	part = append(part, bytes.TrimSpace(src[start:]))

	if len(part) > 1 {
		for _, p := range part {
			e, err := parseSeq(p)
			if err != nil {
				return nil, err
			}
			parts = append(parts, e)
		}
		return Choice(parts), nil
	}
	return parseSeq(part[0])
}

func parseSeq(src []byte) (Expr, error) {
	var parts []Expr
	tokens := tokenize(src)
	for _, tok := range tokens {
		switch {
		case len(tok) > 1 && tok[0] == '[' && tok[len(tok)-1] == ']':
			x, err := parseExpr(tok[1 : len(tok)-1])
			if err != nil {
				return nil, err
			}
			parts = append(parts, Option{X: x})
		case len(tok) > 1 && tok[0] == '{' && tok[len(tok)-1] == '}':
			x, err := parseExpr(tok[1 : len(tok)-1])
			if err != nil {
				return nil, err
			}
			parts = append(parts, Repeat{X: x})
		case len(tok) > 1 && tok[0] == '"' && tok[len(tok)-1] == '"':
			parts = append(parts, Terminal{Lit: tok[1 : len(tok)-1]})
		default:
			parts = append(parts, NonTerm{Name: tok})
		}
	}
	if len(parts) == 1 {
		return parts[0], nil
	}
	return Seq(parts), nil
}

func tokenize(src []byte) [][]byte {
	var out [][]byte
	var level int
	start := 0
	for i := 0; i < len(src); i++ {
		c := src[i]
		switch c {
		case ' ':
			if level == 0 && start < i {
				out = append(out, src[start:i])
				start = i + 1
			}
		case '[', '{':
			if level == 0 && start < i {
				out = append(out, src[start:i])
				start = i
			}
			level++
		case ']', '}':
			level--
			if level == 0 {
				out = append(out, src[start:i+1])
				start = i + 1
			}
		case '"':
			j := i + 1
			for j < len(src) && src[j] != '"' {
				j++
			}
			if j < len(src) {
				out = append(out, src[i:j+1])
				i = j
				start = i + 1
			}
		}
	}
	if start < len(src) {
		out = append(out, src[start:])
	}
	return out
}
