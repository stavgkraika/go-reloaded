package main

import (
	"strings"
	"unicode"
)

// Processor handles single-pass text processing with transformations
type Processor struct {
	text   string   // Input text to process
	pos    int      // Current position in text
	result []string // Accumulated result words
}

// NewProcessor creates a new text processor instance
func NewProcessor(text string) *Processor {
	return &Processor{
		text:   text,
		pos:    0,
		result: make([]string, 0),
	}
}

// Process performs single-pass text processing with all transformations
func (p *Processor) Process() string {
	for p.pos < len(p.text) {
		// Skip whitespace
		if unicode.IsSpace(rune(p.text[p.pos])) {
			p.pos++
			continue
		}

		// Parse next token based on first character
		switch p.text[p.pos] {
		case '(':
			// Handle command tokens like (up), (hex), etc.
			if cmd := p.parseCommand(); cmd != "" {
				p.applyCommand(cmd)
			} else {
				// Not a valid command, treat as regular text
				token := p.parseWord()
				if token != "" {
					p.addWord(token)
				}
			}
		case '\'':
			// Handle quoted text
			quoted := p.parseQuoted()
			if quoted != "" {
				p.result = append(p.result, quoted)
			}
		case '.', ',', '!', '?', ':', ';':
			// Handle punctuation - attach to previous word
			punct := string(p.text[p.pos])
			p.pos++
			p.attachPunctuation(punct)
		default:
			// Handle regular words
			token := p.parseWord()
			if token != "" {
				p.addWord(token)
			}
		}
	}

	// Use strings.Builder for efficient concatenation
	var builder strings.Builder
	for i, word := range p.result {
		if i > 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	return builder.String()
}

// addWord adds a word to result, handling article correction (a <-> an)
func (p *Processor) addWord(word string) {
	// Quick check for articles without ToLower
	if len(word) <= 2 && (word == "a" || word == "A" || word == "an" || word == "An" || word == "AN") {
		// Look ahead to next word
		nextWord := p.peekNextWord()
		if nextWord != "" {
			isUpper := unicode.IsUpper(rune(word[0]))
			if needsAn(nextWord) {
				if isUpper {
					p.result = append(p.result, "An")
				} else {
					p.result = append(p.result, "an")
				}
			} else {
				if isUpper {
					p.result = append(p.result, "A")
				} else {
					p.result = append(p.result, "a")
				}
			}
			return
		}
	}

	p.result = append(p.result, word)
}

// peekNextWord looks ahead to find the next word without advancing position
func (p *Processor) peekNextWord() string {
	savedPos := p.pos
	defer func() { p.pos = savedPos }()

	// Skip whitespace
	for p.pos < len(p.text) && unicode.IsSpace(rune(p.text[p.pos])) {
		p.pos++
	}

	if p.pos >= len(p.text) {
		return ""
	}

	// Don't look ahead if next token is punctuation, command, or quote
	if isPunctChar(p.text[p.pos]) || p.text[p.pos] == '(' || p.text[p.pos] == '\'' {
		return ""
	}

	return p.parseWord()
}

// attachPunctuation attaches punctuation to the previous word
func (p *Processor) attachPunctuation(punct string) {
	if len(p.result) > 0 {
		last := len(p.result) - 1
		p.result[last] = p.result[last] + punct
	} else {
		// No previous word, add punctuation as standalone
		p.result = append(p.result, punct)
	}
}









// needsAn checks if a word should be preceded by "an" instead of "a"
func needsAn(word string) bool {
	if len(word) == 0 {
		return false
	}

	first := unicode.ToLower(rune(word[0]))
	if first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' {
		return true
	}

	// Special case for 'h' words that sound like vowels
	if first == 'h' && len(word) >= 4 {
		lower := strings.ToLower(word)
		if strings.HasPrefix(lower, "hour") || strings.HasPrefix(lower, "honest") || strings.HasPrefix(lower, "heir") {
			return true
		}
	}

	return false
}

