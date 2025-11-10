package main

import (
	"unicode"
)

// parseCommand extracts and validates command tokens like (up), (hex, 2)
func (p *Processor) parseCommand() string {
	start := p.pos
	if p.pos >= len(p.text) || p.text[p.pos] != '(' {
		return ""
	}

	// Find closing parenthesis
	p.pos++
	for p.pos < len(p.text) && p.text[p.pos] != ')' {
		p.pos++
	}

	if p.pos >= len(p.text) {
		// No closing paren found, reset and return empty
		p.pos = start
		return ""
	}

	// Extract command content
	cmdText := p.text[start : p.pos+1]
	p.pos++

	// Validate command format
	if isValidCommand(cmdText) {
		return cmdText
	}

	// Invalid command, reset position and return empty
	p.pos = start
	return ""
}

// parseQuoted extracts quoted text and processes commands within quotes
func (p *Processor) parseQuoted() string {
	if p.pos >= len(p.text) || p.text[p.pos] != '\'' {
		return ""
	}

	start := p.pos
	p.pos++ // Skip opening quote

	// Find closing quote
	for p.pos < len(p.text) && p.text[p.pos] != '\'' {
		p.pos++
	}

	if p.pos >= len(p.text) {
		// No closing quote, return as is
		p.pos = start + 1
		return "'"
	}

	// Extract content between quotes
	content := p.text[start+1 : p.pos]
	p.pos++ // Skip closing quote

	// Process commands inside quotes
	if content == "" {
		return "''"
	}

	processor := NewProcessor(content)
	processed := processor.Process()
	
	// Fix article corrections after processing
	processed = fixArticles(processed)
	
	return "'" + processed + "'"
}

// parseWord extracts a single word (non-whitespace, non-punctuation sequence)
func (p *Processor) parseWord() string {
	start := p.pos

	// Read until whitespace or punctuation
	for p.pos < len(p.text) {
		ch := p.text[p.pos]
		if unicode.IsSpace(rune(ch)) || isPunctChar(ch) {
			break
		}
		p.pos++
	}

	return p.text[start:p.pos]
}