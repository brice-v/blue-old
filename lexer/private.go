package lexer

import (
	"blue/token"
	"encoding/hex"
	"strings"
)

// readChar gives us the next character and advances out position
// in the input string
func (l *Lexer) readChar() {
	l.prevCh = l.ch
	if l.readPos >= runeLen(l.input) {
		l.ch = 0
	} else {
		l.ch = toRunes(l.input)[l.readPos]
	}
	l.pos = l.readPos
	l.readPos += runeLen(string(l.ch))
}

// peekChar will return the rune that is in the readPosition without consuming any input
func (l *Lexer) peekChar() rune {
	if l.readPos >= runeLen(l.input) {
		return 0
	}
	return toRunes(l.input)[l.readPos]
}

// peekNextChar will return the rune right after the readPosition without consuming any input
func (l *Lexer) peekNextChar() rune {
	if l.readPos >= runeLen(l.input) || l.readPos+1 >= runeLen(l.input) {
		return 0
	}
	return toRunes(l.input)[l.readPos+1]
}

// readNumber will keep consuming valid digits of the input according to `isDigit`
// and return the string
// TODO: readNumber can be refactored to be cleaner
func (l *Lexer) readNumber() (token.Type, string) {
	position := l.pos
	if l.ch == '0' {
		if l.peekChar() == 'x' && isHexChar(l.peekNextChar()) {
			// consume the 0 and x and continue to the number
			l.readChar()
			l.readChar()
			for isHexChar(l.ch) || (l.ch == '_' && isHexChar(l.peekChar())) {
				l.readChar()
			}
			return token.HEX, string(toRunes(l.input)[position:l.pos])
		} else if l.peekChar() == 'o' && isOctalChar(l.peekNextChar()) {
			// consume the 0 and the o and continue to the number
			l.readChar()
			l.readChar()
			for isOctalChar(l.ch) || (l.ch == '_' && isOctalChar(l.peekChar())) {
				l.readChar()
			}
			return token.OCTAL, string(toRunes(l.input)[position:l.pos])
		} else if l.peekChar() == 'b' && isBinaryChar(l.peekNextChar()) {
			// consume the 0 and the b and continue to the number
			l.readChar()
			l.readChar()
			for isBinaryChar(l.ch) || (l.ch == '_' && isBinaryChar(l.peekChar())) {
				l.readChar()
			}
			return token.BINARY, string(toRunes(l.input)[position:l.pos])
		}
	}
	dotFlag := false
	for isDigit(l.ch) || (l.ch == '_' && isDigit(l.peekChar())) {
		if l.peekChar() == '.' && !dotFlag && l.peekNextChar() != '.' {
			dotFlag = true
			l.readChar()
			l.readChar()
		}
		l.readChar()
	}
	if dotFlag {
		return token.FLOAT, string(toRunes(l.input)[position:l.pos])
	}
	return token.INT, string(toRunes(l.input)[position:l.pos])
}

// readIdentifier will keep consuming valid letters out of the input according to `isLetter`
// and return the string
func (l *Lexer) readIdentifier() string {
	position := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return string(toRunes(l.input)[position:l.pos])
}

// readSingleLineComment will continue to consume input until the EOL is reached
func (l *Lexer) readSingleLineComment() {
	for l.ch != 0 {
		if l.ch == 0 {
			break
		}
		if l.ch == '#' {
			l.readChar()
			for l.ch != '\n' {
				l.readChar()
				if l.ch == 0 {
					break
				}
			}
			break
		}
		l.readChar()
	}
}

func (l *Lexer) readExecString() string {
	b := strings.Builder{}
	for {
		l.readChar()
		if l.ch == '`' || l.ch == 0 {
			l.readChar()
			break
		}
		b.WriteRune(l.ch)
	}
	return b.String()
}

func (l *Lexer) readRawString() string {
	b := &strings.Builder{}
	// Skip the first 2 " chars
	l.readChar()
	l.readChar()
	for {
		l.readChar()
		if (l.ch == '"' && l.peekChar() == '"' && l.peekNextChar() == '"') || l.ch == 0 {
			l.readChar()
			l.readChar()
			l.readChar()
			break
		}
		b.WriteRune(l.ch)
	}
	// Skip the final part of the raw string token
	// l.readChar()
	return b.String()
}

// readString will consume tokens until the string is fully read
func (l *Lexer) readString() (string, error) {
	b := &strings.Builder{}
	for {
		l.readChar()

		// Support some basic escapes like \"
		if l.ch == '\\' {
			switch l.peekChar() {
			case '"':
				b.WriteByte('"')
			case 'n':
				b.WriteByte('\n')
			case 'r':
				b.WriteByte('\r')
			case 't':
				b.WriteByte('\t')
			case '\\':
				b.WriteByte('\\')
			case 'x':
				// Skip over the the '\\', 'x' and the next two bytes (hex)
				l.readChar()
				l.readChar()
				l.readChar()
				src := string([]rune{l.prevCh, l.ch})
				dst, err := hex.DecodeString(src)
				if err != nil {
					return "", err
				}
				b.Write(dst)
				continue
			}

			// Skip over the '\\' and the matched single escape char
			l.readChar()
			continue
		} else {
			if l.ch == '"' || l.ch == 0 {
				break
			}
		}

		b.WriteRune(l.ch)
	}

	return b.String(), nil
}

// skipWhitespace will continue to advance if the current byte is considered
// a whitespace character such as ' ', '\t', '\n', '\r'
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// makeTwoCharToken takes a tokens type and returns the new token
// while advancing the readPosition and current char
func (l *Lexer) makeTwoCharToken(typ token.Type) token.Token {
	ch := l.ch
	// consume next char because we know it is an =
	l.readChar()
	return token.Token{Type: typ, Literal: string(ch) + string(l.ch)}
}

// makeThreeCharToken takes a tokens type and returns the new token
// while advancing the readPosition and current char to the proper position
func (l *Lexer) makeThreeCharToken(typ token.Type) token.Token {
	ch := l.ch
	l.readChar()
	ch1 := l.ch
	l.readChar()
	return token.Token{Type: typ, Literal: string(ch) + string(ch1) + string(l.ch)}
}
