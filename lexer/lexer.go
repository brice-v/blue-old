// lexer will read a source input and stream tokens upon
// request of the blue programming language
package lexer

import (
	"blue/token"
	"fmt"
	"strings"
)

// Lexer is lex object type
// It contains the source input and positions to characters in the text
type Lexer struct {
	input   string
	pos     int  // current pos. in input (points to current char)
	readPos int  // current reading pos. in input (after current char)
	ch      rune // current char under examination
	prevCh  rune // previous char read

	filename string // filename is the name to print to the terminal for span
}

// New returns a pointer to the Lexer object
func New(input, filename string) *Lexer {
	l := &Lexer{input: input, filename: filename}
	l.readChar()
	return l
}

// NextToken matches against a byte and if it succeeds it will
// read the next char and return a token struct
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			// Note: cant use newToken here because it is not 1 byte long
			tok = l.makeTwoCharToken(token.EQ)
		} else if l.peekChar() == '>' {
			tok = l.makeTwoCharToken(token.RARROW)
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.pos)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.pos)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.pos)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.pos)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.pos)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.pos)
	case '[':
		tok = newToken(token.LBRACKET, l.ch, l.pos)
	case ']':
		tok = newToken(token.RBRACKET, l.ch, l.pos)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.pos)
	case '+':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.PLUSEQ)
		} else {
			tok = newToken(token.PLUS, l.ch, l.pos)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.NEQ)
		} else {
			tok = newToken(token.BANG, l.ch, l.pos)
		}
	case '-':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.MINUSEQ)
		} else {
			tok = newToken(token.MINUS, l.ch, l.pos)
		}
	case '/':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.DIVEQ)
		} else if l.peekChar() == '/' && l.peekNextChar() != '=' {
			tok = l.makeTwoCharToken(token.FDIV)
		} else if l.peekChar() == '/' && l.peekNextChar() == '=' {
			tok = l.makeThreeCharToken(token.FDIVEQ)
		} else {
			tok = newToken(token.FSLASH, l.ch, l.pos)
		}
	case '*':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.MULEQ)
		} else if l.peekChar() == '*' && l.peekNextChar() != '=' {
			tok = l.makeTwoCharToken(token.POW)
		} else if l.peekChar() == '*' && l.peekNextChar() == '=' {
			tok = l.makeThreeCharToken(token.POWEQ)
		} else {
			tok = newToken(token.STAR, l.ch, l.pos)
		}
	case '<':
		if l.peekChar() == '<' {
			if l.peekNextChar() == '=' {
				tok = l.makeThreeCharToken(token.LSHIFTEQ)
			} else {
				tok = l.makeTwoCharToken(token.LSHIFT)
			}
		} else if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.LTEQ)
		} else {
			tok = newToken(token.LT, l.ch, l.pos)
		}
	case '>':
		if l.peekChar() == '>' {
			if l.peekNextChar() == '=' {
				tok = l.makeThreeCharToken(token.RSHIFTEQ)
			} else {
				tok = l.makeTwoCharToken(token.RSHIFT)
			}
		} else if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.GTEQ)
		} else {
			tok = newToken(token.GT, l.ch, l.pos)
		}
	case '|':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.OREQ)
		} else {
			tok = newToken(token.PIPE, l.ch, l.pos)
		}
	case '&':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.ANDEQ)
		} else {
			tok = newToken(token.AMPERSAND, l.ch, l.pos)
		}
	case '^':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.XOREQ)
		} else {
			tok = newToken(token.HAT, l.ch, l.pos)
		}
	case '#':
		if l.peekChar() == '{' {
			tok = l.makeTwoCharToken(token.STRINGINTERP)
		} else {
			tok = newToken(token.HASH, l.ch, l.pos)
			l.readSingleLineComment()
		}
	case '%':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.PERCENTEQ)
		} else {
			tok = newToken(token.PERCENT, l.ch, l.pos)
		}
	case '.':
		if l.peekChar() == '.' {
			if l.peekNextChar() == '<' {
				tok = l.makeThreeCharToken(token.NONINCRANGE)
			} else {
				tok = l.makeTwoCharToken(token.RANGE)
			}
		} else {
			tok = newToken(token.DOT, l.ch, l.pos)
		}
	case '~':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.BINNOTEQ)
		} else {
			tok = newToken(token.TILDE, l.ch, l.pos)
		}
	case '`':
		tok.Type = token.BACKTICK
		tok.Literal = l.readExecString()
		return tok
	case ':':
		tok = newToken(token.COLON, l.ch, l.pos)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '"':
		if l.peekChar() == '"' && l.peekNextChar() == '"' {
			str := l.readRawString()
			tok.Type = token.RAW_STRING
			tok.Literal = str
		} else {
			str, err := l.readString()
			if err != nil {
				tok = newToken(token.ILLEGAL, l.prevCh, l.pos)
			} else {
				tok.Type = token.STRING
				tok.Literal = str
			}
		}
	default:
		start := l.pos
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			end := l.pos
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Span = token.Span{Start: start, End: end}
			return tok
		} else if isDigit(l.ch) {
			tok.Type, tok.Literal = l.readNumber()
			end := l.pos
			tok.Span = token.Span{Start: start, End: end}
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch, l.pos)
	}

	l.readChar()
	return tok
}

// GetSpanPrintable returns a printable string for the lexer with a message
// It prints the filename and the line:col so the link is clickable in
// most terminals
func (l *Lexer) GetSpanPrintable(span token.Span, msg string) string {
	lines := strings.Split(l.input, "\n")
	totCount := 0
	for lineno, line := range lines {
		lineno += 1
		linestart := totCount
		totCount += runeLen(line) + 1
		if span.End > totCount-1 {
			continue
		}
		fdata := fmt.Sprintf("%s:%d:%d", l.filename, lineno, span.Start-linestart+1)
		msgAndFdata := fmt.Sprintf("%s %s\n", fdata, msg)
		tildes := "~~~~~"
		padLen := (len(fdata) - len(tildes)) + 1
		lineMsg := fmt.Sprintf(strings.Repeat(" ", padLen+len(tildes)) + line + "\n")
		tildeMsg := fmt.Sprintf(strings.Repeat(" ", padLen+(span.Start-linestart)) + tildes + "^\n")
		return msgAndFdata + lineMsg + tildeMsg
	}
	return ""
}
