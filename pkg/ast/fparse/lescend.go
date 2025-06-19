// Package fparse — быстрый LL-парсер AQL-выражений в fexpr.Compiler.
package fparse

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/ast/fexpr"
)

/*─────────────────────── PUBLIC ───────────────────────*/

// Parse возвращает лямбду-Compiler, готовую эмитить байткод.
func Parse(src []byte) (fexpr.Compiler, error) {
	p := parser{lex: newLexer(src)}
	p.next()
	return p.parsePipe()
}

/*────────────────────  LEXER  ────────────────────*/

type kind uint8

const (
	tEOF kind = iota
	tIdent
	tNum
	tStr
	tTrue
	tFalse
	tNull
	tOver
	tAndAnd
	tOrOr
	tEq
	tNeq
	tLe
	tGe
	tArrow
)

type token struct {
	k   kind
	lit []byte
}

type lexer struct {
	s []byte
	i int
}

func newLexer(b []byte) *lexer { return &lexer{s: b} }

func (l *lexer) next() token {
	// skip ws
	for l.i < len(l.s) && l.s[l.i] <= ' ' {
		l.i++
	}
	if l.i >= len(l.s) {
		return token{k: tEOF}
	}
	c := l.s[l.i]
	switch {
	case c >= '0' && c <= '9':
		st := l.i
		for l.i < len(l.s) && l.s[l.i] >= '0' && l.s[l.i] <= '9' {
			l.i++
		}
		return token{k: tNum, lit: l.s[st:l.i]}
	case c == '"' /*string*/ :
		st := l.i
		l.i++
		for l.i < len(l.s) && l.s[l.i] != '"' {
			l.i++
		}
		l.i++
		return token{k: tStr, lit: l.s[st:l.i]}
	case c == '&' && l.peek('&'):
		l.i += 2
		return token{k: tAndAnd}
	case c == '|' && l.peek('|'):
		l.i += 2
		return token{k: tOrOr}
	case c == '=' && l.peek('='):
		l.i += 2
		return token{k: tEq}
	case c == '!' && l.peek('='):
		l.i += 2
		return token{k: tNeq}
	case c == '<' && l.peek('='):
		l.i += 2
		return token{k: tLe}
	case c == '>' && l.peek('='):
		l.i += 2
		return token{k: tGe}
	case c == '=' && l.peek('>'):
		l.i += 2
		return token{k: tArrow}
	}
	// ident / keyword
	if isLetter(l.s[l.i]) || l.s[l.i] == '_' {
		st := l.i
		for l.i < len(l.s) && (isLetter(l.s[l.i]) || isDigit(l.s[l.i]) || l.s[l.i] == '_') {
			l.i++
		}
		word := l.s[st:l.i]
		switch string(word) {
		case "true":
			return token{k: tTrue}
		case "false":
			return token{k: tFalse}
		case "null":
			return token{k: tNull}
		case "OVER":
			return token{k: tOver}
		default:
			return token{k: tIdent, lit: word}
		}
	}
	// single
	l.i++
	return token{k: kind(c), lit: []byte{c}}
}
func (l *lexer) peek(b byte) bool { return l.i+1 < len(l.s) && l.s[l.i+1] == b }
func isLetter(c byte) bool        { return c|0x20 >= 'a' && c|0x20 <= 'z' }
func isDigit(c byte) bool         { return c >= '0' && c <= '9' }

/*────────────────────  PARSER  ────────────────────*/

type parser struct {
	lex *lexer
	tok token
}

func (p *parser) next() { p.tok = p.lex.next() }
func (p *parser) accept(k kind) bool {
	if p.tok.k == k {
		p.next()
		return true
	}
	return false
}
func (p *parser) expect(k kind) {
	if !p.accept(k) {
		panic("syntax")
	}
}

/*──── Grammar functions (LL, fastest path) ────*/

func (p *parser) parsePipe() (fexpr.Compiler, error) {
	left, _ := p.parseOr()
	for p.accept('|') {
		right, _ := p.parseOr()
		left = fexpr.Pipe(left, nil) // оп.Field переиспользуем (Pipe код совпадает)
		left = fexpr.Binary(left, right, op.Pipe)
	}
	return left, nil
}

func (p *parser) parseOr() (fexpr.Compiler, error) {
	left, _ := p.parseAnd()
	for p.tok.k == tOrOr {
		p.next()
		right, _ := p.parseAnd()
		left = fexpr.Binary(left, right, op.Or)
	}
	return left, nil
}

func (p *parser) parseAnd() (fexpr.Compiler, error) {
	left, _ := p.parseCmp()
	for p.tok.k == tAndAnd {
		p.next()
		right, _ := p.parseCmp()
		left = fexpr.Binary(left, right, op.And)
	}
	return left, nil
}

func (p *parser) parseCmp() (fexpr.Compiler, error) {
	left, _ := p.parseAdd()
	for {
		switch p.tok.k {
		case tEq, tNeq, '<', tLe, '>', tGe:
			opc := map[kind]op.Code{
				tEq: op.Eq, tNeq: op.Neq, '<': op.Lt,
				tLe: op.Le, '>': op.Gt, tGe: op.Ge,
			}[p.tok.k]
			p.next()
			right, _ := p.parseAdd()
			left = fexpr.Binary(left, right, opc)
		default:
			return left, nil
		}
	}
}

func (p *parser) parseAdd() (fexpr.Compiler, error) {
	left, _ := p.parseMul()
	for p.tok.k == '+' || p.tok.k == '-' {
		opc := op.Add
		if p.tok.k == '-' {
			opc = op.Sub
		}
		p.next()
		right, _ := p.parseMul()
		left = fexpr.Binary(left, right, opc)
	}
	return left, nil
}

func (p *parser) parseMul() (fexpr.Compiler, error) {
	left, _ := p.parseUnary()
	for p.tok.k == '*' || p.tok.k == '/' || p.tok.k == '%' {
		opc := map[kind]op.Code{'*': op.Mul, '/': op.Div, '%': op.Mod}[p.tok.k]
		p.next()
		right, _ := p.parseUnary()
		left = fexpr.Binary(left, right, opc)
	}
	return left, nil
}

func (p *parser) parseUnary() (fexpr.Compiler, error) {
	switch p.tok.k {
	case '-':
		p.next()
		r, _ := p.parseUnary()
		return fexpr.Unary(r, op.Not), nil
	case tOver:
		p.next()
		iter, _ := p.parseUnary()
		if p.accept(tArrow) {
			p.expect('(')
			body, _ := p.parsePipe()
			p.expect(')')
			return fexpr.Over(iter, body), nil
		}
		return fexpr.Over(iter, fexpr.Nop), nil
	default:
		return p.parsePost()
	}
}

func (p *parser) parsePost() (fexpr.Compiler, error) {
	node, _ := p.parseAtom()
loop:
	for {
		switch {
		case p.accept('.'):
			if p.tok.k != tIdent {
				return nil, errors.New("ident expected")
			}
			name := bytes.Clone(p.tok.lit)
			p.next()
			node = fexpr.Field(node, name)

		case p.accept('['):
			idx1, _ := p.parsePipe()
			if p.accept(']') { // single index
				node = fexpr.Binary(node, idx1, op.Index1)
				continue
			}
			p.expect(':')
			idx2, _ := p.parsePipe()
			p.expect(']')
			node = fexpr.Ternary(node, idx1, idx2, op.Index2)

		case p.tok.k == tIdent && p.lex.s[p.lex.i] == '(':
			name := bytes.Clone(p.tok.lit)
			p.next() // идентификатор
			p.expect('(')
			var args []fexpr.Compiler
			if !p.accept(')') {
				arg, _ := p.parsePipe()
				args = append(args, arg)
				for p.accept(',') {
					arg, _ = p.parsePipe()
					args = append(args, arg)
				}
				p.expect(')')
			}
			node = fexpr.Call(args, name)

		default:
			break loop
		}
	}
	return node, nil
}

func (p *parser) parseAtom() (fexpr.Compiler, error) {
	switch p.tok.k {
	case tIdent:
		name := bytes.Clone(p.tok.lit)
		p.next()
		return fexpr.Ident(name), nil
	case tNum:
		v, _ := strconv.ParseInt(string(p.tok.lit), 10, 64)
		p.next()
		return fexpr.I64(v), nil
	case tStr:
		s := p.tok.lit[1 : len(p.tok.lit)-1]
		p.next()
		return fexpr.StringBytes(s), nil
	case tTrue:
		p.next()
		return fexpr.True, nil
	case tFalse:
		p.next()
		return fexpr.False, nil
	case tNull:
		p.next()
		return fexpr.Nil, nil
	case '.':
		p.next()
		return fexpr.Dup, nil
	case '(':
		p.next()
		// n, _ := fexpr.Pipe()
		p.expect(')')
		return fexpr.Nop, nil
	default:
		return nil, errors.New("atom?")
	}
}
