package lexer

import (
	"unicode/utf8"
)

// TokenKind represents a lexical token category.
// NOTE: keep order in sync with tokenNames array for performance.
//
//go:generate stringer -type=TokenKind
type TokenKind int

const (
	// Special
	TIllegal TokenKind = iota
	TEOF

	// Identifiers + literals
	TIdentifier // foo, _bar, x1
	TNumber     // 123, 0x2a
	TFloat      // 3.14, 1e-9
	TString     // "hello"
	TNull       // null
	TTrue       // true
	TFalse      // false
	TOver       // over

	// Operators & delimiters
	// single‑rune first to leverage fast table
	TPlus     // +
	TMinus    // -
	TStar     // *
	TSlash    // /
	TPercent  // %
	TDot      // .
	TComma    // ,
	TColon    // :
	TSemi     // ;
	TLParen   // (
	TRParen   // )
	TLBracket // [
	TRBracket // ]
	TLBrace   // {
	TRBrace   // }
	TPipe     // |

	// multi‑rune
	TArrow // =>
	TEq    // ==
	TNeq   // !=
	TLt    // <
	TLte   // <=
	TGt    // >
	TGte   // >=
	TAnd   // &&
	TOr    // ||
)

// Token represents a single lexical token.
type Token struct {
	Kind TokenKind
	Lit  []byte // slice of original input (zero‑copy)
	Pos  int    // byte offset in source
}

// Lexer implements zero‑alloc scanning over UTF‑8 source.
type Lexer struct {
	src []byte
	pos int // current reading offset
	wid int // width of last rune in bytes
	tok Token
}

// New returns a new lexer over src.
func New(src []byte) *Lexer {
	return &Lexer{src: src, pos: 0}
}

// Next returns the next token (reuse internal struct to avoid allocs).
func (lx *Lexer) Next() Token {
	lx.skipWS()
	if lx.pos >= len(lx.src) {
		return lx.make(TEOF, nil)
	}

	ch, w := lx.peekRune()

	// Fast single‑byte tokens first (common ASCII path)
	switch ch {
	case '+':
		return lx.take1(TPlus)
	case '-':
		return lx.take1(TMinus)
	case '*':
		return lx.take1(TStar)
	case '/':
		return lx.take1(TSlash)
	case '%':
		return lx.take1(TPercent)
	case '.':
		return lx.take1(TDot)
	case ',':
		return lx.take1(TComma)
	case ':':
		return lx.take1(TColon)
	case ';':
		return lx.take1(TSemi)
	case '(':
		return lx.take1(TLParen)
	case ')':
		return lx.take1(TRParen)
	case '[':
		return lx.take1(TLBracket)
	case ']':
		return lx.take1(TRBracket)
	case '{':
		return lx.take1(TLBrace)
	case '}':
		return lx.take1(TRBrace)
	case '|':
		if lx.peek2('|') {
			return lx.take2(TOr)
		}
		return lx.take1(TPipe)
	case '&':
		if lx.peek2('&') {
			return lx.take2(TAnd)
		}
	case '=':
		if lx.peek2('=') {
			return lx.take2(TEq)
		}
	case '!':
		if lx.peek2('=') {
			return lx.take2(TNeq)
		}
	case '<':
		if lx.peek2('=') {
			return lx.take2(TLte)
		}
		return lx.take1(TLt)
	case '>':
		if lx.peek2('=') {
			return lx.take2(TGte)
		}
		return lx.take1(TGt)
	}

	// multi‑rune arrow =>
	if ch == '=' && lx.peekAhead(1) == '>' {
		return lx.take2(TArrow)
	}

	// Ident / keyword
	if isIdentStart(ch) {
		return lx.scanIdent()
	}

	// Number
	if isDigit(ch) {
		return lx.scanNumber()
	}

	// String literal
	if ch == '"' {
		return lx.scanString()
	}

	// Unknown byte → illegal token
	lx.advance(w)
	return lx.make(TIllegal, lx.src[lx.tok.Pos:lx.pos])
}

// ---- scanning helpers ----

func (lx *Lexer) make(kind TokenKind, lit []byte) Token {
	lx.tok.Kind = kind
	lx.tok.Lit = lit
	return lx.tok
}

func (lx *Lexer) take1(kind TokenKind) Token {
	start := lx.pos
	lx.advance(1)
	return lx.make(kind, lx.src[start:lx.pos])
}

func (lx *Lexer) take2(kind TokenKind) Token {
	start := lx.pos
	lx.advance(2)
	return lx.make(kind, lx.src[start:lx.pos])
}

func (lx *Lexer) advance(n int) { lx.pos += n }

func (lx *Lexer) peekRune() (rune, int) {
	if lx.pos >= len(lx.src) {
		return 0, 0
	}
	ch, w := utf8.DecodeRune(lx.src[lx.pos:])
	lx.wid = w
	return ch, w
}

func (lx *Lexer) peekAhead(off int) byte {
	idx := lx.pos + off
	if idx >= len(lx.src) {
		return 0
	}
	return lx.src[idx]
}

func (lx *Lexer) peek2(expected byte) bool {
	return lx.pos+1 < len(lx.src) && lx.src[lx.pos+1] == expected
}

func (lx *Lexer) skipWS() {
	for lx.pos < len(lx.src) {
		switch lx.src[lx.pos] {
		case ' ', '\t', '\n', '\r':
			lx.pos++
		default:
			return
		}
	}
}

func isIdentStart(r rune) bool {
	return r == '_' || r == '$' || ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z') || r >= 0x80
}

func isIdentPart(r rune) bool {
	return isIdentStart(r) || ('0' <= r && r <= '9')
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func (lx *Lexer) scanIdent() Token {
	start := lx.pos
	for {
		ch, w := lx.peekRune()
		if !isIdentPart(ch) {
			break
		}
		lx.advance(w)
	}
	lit := lx.src[start:lx.pos]

	switch string(lit) {
	case "true":
		return lx.make(TTrue, lit)
	case "false":
		return lx.make(TFalse, lit)
	case "null":
		return lx.make(TNull, lit)
	case "over": // ✨ ДОБАВЛЕНО
		return lx.make(TOver, lit)
	default:
		return lx.make(TIdentifier, lit)
	}
}

func (lx *Lexer) scanNumber() Token {
	start := lx.pos
	isFloat := false

	// Integer part or 0x
	if lx.src[lx.pos] == '0' && (lx.peekAhead(1) == 'x' || lx.peekAhead(1) == 'X') {
		lx.advance(2)
		for lx.pos < len(lx.src) && isHexDigit(lx.src[lx.pos]) {
			lx.pos++
		}
		return lx.make(TNumber, lx.src[start:lx.pos])
	}

	for lx.pos < len(lx.src) && isDigit(rune(lx.src[lx.pos])) {
		lx.pos++
	}

	// Fractional
	if lx.pos < len(lx.src) && lx.src[lx.pos] == '.' {
		isFloat = true
		lx.pos++
		for lx.pos < len(lx.src) && isDigit(rune(lx.src[lx.pos])) {
			lx.pos++
		}
	}

	// Exponent
	if lx.pos < len(lx.src) && (lx.src[lx.pos] == 'e' || lx.src[lx.pos] == 'E') {
		isFloat = true
		lx.pos++
		if lx.pos < len(lx.src) && (lx.src[lx.pos] == '+' || lx.src[lx.pos] == '-') {
			lx.pos++
		}
		for lx.pos < len(lx.src) && isDigit(rune(lx.src[lx.pos])) {
			lx.pos++
		}
	}
	kind := TNumber
	if isFloat {
		kind = TFloat
	}
	return lx.make(kind, lx.src[start:lx.pos])
}

func isHexDigit(b byte) bool {
	return ('0' <= b && b <= '9') || ('a' <= b && b <= 'f') || ('A' <= b && b <= 'F')
}

func (lx *Lexer) scanString() Token {
	start := lx.pos
	lx.advance(1) // skip opening quote

	escaped := false
	for lx.pos < len(lx.src) {
		c := lx.src[lx.pos]
		if c == '\\' {
			escaped = true
			lx.pos += 2 // skip escape sequence – may go past end; kept simple
			continue
		}
		if c == '"' {
			break
		}
		lx.pos++
	}
	if lx.pos >= len(lx.src) { // unterminated
		return lx.make(TIllegal, lx.src[start:lx.pos])
	}
	lx.advance(1) // consume closing quote

	lit := lx.src[start:lx.pos]

	if escaped {
		// alloc: we must unescape → produces new slice
		// alloc
		unquoted, _ := unescapeString(lx.src[start+1 : lx.pos-1]) // implement simple unescape
		return lx.make(TString, unquoted)
	}
	return lx.make(TString, lit)
}

// unescapeString handles minimal set (\", \\, \n, \t). Returns freshly allocated slice. // alloc
func unescapeString(in []byte) ([]byte, error) { // alloc
	out := make([]byte, 0, len(in)) // alloc unavoidable
	for i := 0; i < len(in); i++ {
		if in[i] == '\\' && i+1 < len(in) {
			i++
			switch in[i] {
			case 'n':
				out = append(out, '\n')
			case 't':
				out = append(out, '\t')
			case 'r':
				out = append(out, '\r')
			case '\\':
				out = append(out, '\\')
			case '"':
				out = append(out, '"')
			default:
				out = append(out, '\\', in[i])
			}
		} else {
			out = append(out, in[i])
		}
	}
	return out, nil
}
