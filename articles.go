package main

import (
	"strings"
	"unicode"
)

// fixArticles corrects a/an articles in processed text
func fixArticles(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		lower := strings.ToLower(words[i])
		if lower == "a" || lower == "an" {
			nextWord := words[i+1]
			if needsAn(nextWord) {
				if unicode.IsUpper(rune(words[i][0])) {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			} else {
				if unicode.IsUpper(rune(words[i][0])) {
					words[i] = "A"
				} else {
					words[i] = "a"
				}
			}
		}
	}
	return strings.Join(words, " ")
}