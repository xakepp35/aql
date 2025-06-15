package pratt

/*────────────────────────  TOKENIZER ─────────────────────*/

type tokKind int

const (
	tEOF tokKind = iota
	tIdent
	tNumber
	tString

	// односимвольные берём как rune, а мультирунные ниже:
	tTrue
	tFalse
	tNull
	tOver
	tAndAnd // &&
	tOrOr   // ||
	tEq     // ==
	tNeq    // !=
	tLe     // <=
	tGe     // >=
	tArrow  // =>
)

type token struct {
	kind tokKind
	lit  []byte
}

type lexer struct {
	src []byte
	i   int
}

func newLexer(src []byte) *lexer { return &lexer{src: src} }

func (l *lexer) nextToken() token {
	// skip ws
	for l.i < len(l.src) && (l.src[l.i] == ' ' || l.src[l.i] == '\t' ||
		l.src[l.i] == '\n' || l.src[l.i] == '\r') {
		l.i++
	}
	if l.i >= len(l.src) {
		return token{kind: tEOF}
	}
	ch := l.src[l.i]
	switch {
	case ch >= '0' && ch <= '9': // number
		start := l.i
		for l.i < len(l.src) && l.src[l.i] >= '0' && l.src[l.i] <= '9' {
			l.i++
		}
		return token{kind: tNumber, lit: l.src[start:l.i]}

	case ch == '"' /*string*/ :
		start := l.i
		l.i++
		for l.i < len(l.src) && l.src[l.i] != '"' {
			l.i++
		}
		l.i++ // consume "
		return token{kind: tString, lit: l.src[start:l.i]}

	case ch == '&' && l.peek('&'):
		l.i += 2
		return token{kind: tAndAnd, lit: []byte("&&")}
	case ch == '|' && l.peek('|'):
		l.i += 2
		return token{kind: tOrOr, lit: []byte("||")}
	case ch == '=' && l.peek('='):
		l.i += 2
		return token{kind: tEq, lit: []byte("==")}
	case ch == '!' && l.peek('='):
		l.i += 2
		return token{kind: tNeq, lit: []byte("!=")}
	case ch == '<' && l.peek('='):
		l.i += 2
		return token{kind: tLe, lit: []byte("<=")}
	case ch == '>' && l.peek('='):
		l.i += 2
		return token{kind: tGe, lit: []byte(">=")}
	case ch == '=' && l.peek('>'):
		l.i += 2
		return token{kind: tArrow, lit: []byte("=>")}
	}

	// identifiers / keywords
	if isLetter(l.src[l.i]) || l.src[l.i] == '_' {
		start := l.i
		for l.i < len(l.src) && (isLetter(l.src[l.i]) || isDigit(l.src[l.i]) || l.src[l.i] == '_') {
			l.i++
		}
		word := string(l.src[start:l.i])
		switch word {
		case "true":
			return token{kind: tTrue, lit: l.src[start:l.i]}
		case "false":
			return token{kind: tFalse, lit: l.src[start:l.i]}
		case "null":
			return token{kind: tNull, lit: l.src[start:l.i]}
		case "over":
			return token{kind: tOver, lit: l.src[start:l.i]}
		default:
			return token{kind: tIdent, lit: l.src[start:l.i]}
		}
	}

	// single char token
	l.i++
	return token{kind: tokKind(rune(ch)), lit: []byte{ch}}
}

func (l *lexer) peek(b byte) bool { return l.i+1 < len(l.src) && l.src[l.i+1] == b }
func isLetter(c byte) bool        { return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' }
func isDigit(c byte) bool         { return c >= '0' && c <= '9' }
