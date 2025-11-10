package main

import (
	"strings"
	"unicode"
)

// isPunctChar checks if character is punctuation
func isPunctChar(ch byte) bool {
	return ch == '.' || ch == ',' || ch == '!' || ch == '?' || ch == ':' || ch == ';'
}

// capitalize converts first letter to uppercase, rest to lowercase
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}