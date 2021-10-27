package lexer

import (
	"blue/token"
	"unicode"
	"unicode/utf8"
)

// runeLen will return the number of unicode characters in a string
func runeLen(input string) int {
	return utf8.RuneCountInString(input)
}

// toRunes wraps converting a string to a slice of runes
func toRunes(input string) []rune {
	return []rune(input)
}

// newToken will return a Token object with a type and literal
func newToken(tokenType token.Type, ch rune, pos int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Span: token.Span{Start: pos, End: pos}}
}

// isLetter will return true if the rune given matches the pattern below
func isLetter(ch rune) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_' || ch == '?'
}

// isDigit will return true if the rune matches the below pattern
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' || unicode.IsNumber(ch)
}

// isHexChar will return true if the rune given is a hex character
func isHexChar(ch rune) bool {
	return 'a' <= ch && ch <= 'f' || 'A' <= ch && ch <= 'F' || '0' <= ch && ch <= '9'
}

// isOctalChar will return true if the rune given is an octal character
func isOctalChar(ch rune) bool {
	return '0' <= ch && ch <= '7'
}

// isBinaryChar will return true if the rune given is a binary character
func isBinaryChar(ch rune) bool {
	return '0' == ch || '1' == ch
}
