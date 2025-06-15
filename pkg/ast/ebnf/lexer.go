package ebnf

type kind uint8

const (
	_ kind = iota
	tEOF
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
	off, end int
	k        kind
}

func NewLexer(src []byte) *lexer {
	return &lexer{s: src}
}

/*──────────────────── LEXER ────────────────────*/
type lexer struct {
	s   []byte
	i   int
	tok token
}

func (l *lexer) Next() {
	// skip ws
	for l.i < len(l.s) && l.s[l.i] <= ' ' {
		l.i++
	}
	if l.i >= len(l.s) {
		l.tok = token{k: tEOF}
		return
	}
	c := l.s[l.i]
	switch {
	case c >= '0' && c <= '9': // number
		st := l.i
		for l.i < len(l.s) && l.s[l.i] >= '0' && l.s[l.i] <= '9' {
			l.i++
		}
		l.tok = token{k: tNum, off: st, end: l.i}
		return
	case c == '"': // string
		st := l.i
		l.i++
		for l.i < len(l.s) && l.s[l.i] != '"' {
			l.i++
		}
		l.i++
		l.tok = token{k: tStr, off: st, end: l.i}
		return
	case c == '&' && l.s[l.i+1] == '&':
		l.i += 2
		l.tok.k = tAndAnd
		return
	case c == '|' && l.s[l.i+1] == '|':
		l.i += 2
		l.tok.k = tOrOr
		return
	case c == '=' && l.s[l.i+1] == '=':
		l.i += 2
		l.tok.k = tEq
		return
	case c == '!' && l.s[l.i+1] == '=':
		l.i += 2
		l.tok.k = tNeq
		return
	case c == '<' && l.s[l.i+1] == '=':
		l.i += 2
		l.tok.k = tLe
		return
	case c == '>' && l.s[l.i+1] == '=':
		l.i += 2
		l.tok.k = tGe
		return
	case c == '=' && l.s[l.i+1] == '>':
		l.i += 2
		l.tok.k = tArrow
		return
	}
	// ident / keyword
	if isLetter(c) || c == '_' {
		st := l.i
		for l.i < len(l.s) && (isLetter(l.s[l.i]) || isDigit(l.s[l.i]) || l.s[l.i] == '_') {
			l.i++
		}
		word := l.s[st:l.i]
		switch string(word) {
		case "true":
			l.tok.k = tTrue
		case "false":
			l.tok.k = tFalse
		case "null":
			l.tok.k = tNull
		case "over":
			l.tok.k = tOver
		default:
			l.tok.k = tIdent
			l.tok.off, l.tok.end = st, l.i
		}
		return
	}
	// single-char
	l.i++
	l.tok.k = kind(c)
	l.tok.off, l.tok.end = l.i-1, l.i
}

func isLetter(c byte) bool { return c|0x20 >= 'a' && c|0x20 <= 'z' }
func isDigit(c byte) bool  { return c >= '0' && c <= '9' }
