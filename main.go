// Text processing program that handles various text transformations in a single pass
// Supports: case conversion, number base conversion, article correction, punctuation normalization
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Processor handles single-pass text processing with transformations
type Processor struct {
	text   string // Input text to process
	pos    int    // Current position in text
	result []string // Accumulated result words
}

// main function handles command line arguments and file I/O
func main() {
	// Check for test mode
	if len(os.Args) == 2 && os.Args[1] == "test" {
		runAllTests()
		return
	}

	// Validate command line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		fmt.Println("       go run . test")
		return
	}

	// Read input file
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Process text and write output
	processor := NewProcessor(string(input))
	output := processor.Process()

	err = os.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
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
				p.addWord(token)
			}
		case '\'':
			// Handle quoted text
			quoted := p.parseQuoted()
			p.result = append(p.result, quoted)
		case '.', ',', '!', '?', ':', ';':
			// Handle punctuation - attach to previous word
			punct := string(p.text[p.pos])
			p.pos++
			p.attachPunctuation(punct)
		default:
			// Handle regular words
			token := p.parseWord()
			p.addWord(token)
		}
	}

	return strings.Join(p.result, " ")
}

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

	// Invalid command, reset position
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

// addWord adds a word to result, handling article correction (a -> an)
func (p *Processor) addWord(word string) {
	// Handle article correction: "a" -> "an" before vowel sounds
	if strings.ToLower(word) == "a" {
		// Look ahead to next word
		nextWord := p.peekNextWord()
		if nextWord != "" && needsAn(nextWord) {
			p.result = append(p.result, "An")
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

	// Don't look ahead if next token is punctuation or command
	if isPunctChar(p.text[p.pos]) || p.text[p.pos] == '(' {
		return ""
	}

	return p.parseWord()
}

// attachPunctuation attaches punctuation to the previous word
func (p *Processor) attachPunctuation(punct string) {
	if len(p.result) > 0 {
		last := len(p.result) - 1
		p.result[last] += punct
	} else {
		// No previous word, add punctuation as standalone
		p.result = append(p.result, punct)
	}
}

// applyCommand applies transformation commands to previous words
func (p *Processor) applyCommand(cmdText string) {
	cmd, count := parseCommandText(cmdText)
	if cmd == "" {
		return
	}

	// Special handling for cap command without count - apply to current phrase
	if cmd == "cap" && count == 1 && len(p.result) >= 2 {
		// Find start of current phrase (consecutive non-punctuation words)
		start := len(p.result) - 1
		for start > 0 && !hasPunctuation(p.result[start-1]) {
			start--
		}
		// Capitalize all words in phrase
		for i := start; i < len(p.result); i++ {
			if !hasPunctuation(p.result[i]) {
				p.result[i] = capitalize(p.result[i])
			}
		}
		return
	}

	// Apply command to specified number of previous words
	if len(p.result) < count {
		count = len(p.result)
	}

	start := len(p.result) - count
	for i := start; i < len(p.result); i++ {
		switch cmd {
		case "up":
			p.result[i] = strings.ToUpper(p.result[i])
		case "low":
			p.result[i] = strings.ToLower(p.result[i])
		case "cap":
			p.result[i] = capitalize(p.result[i])
		case "hex":
			// Convert hexadecimal to decimal
			if val, err := strconv.ParseInt(p.result[i], 16, 64); err == nil {
				p.result[i] = strconv.FormatInt(val, 10)
			}
		case "bin":
			// Convert binary to decimal
			if val, err := strconv.ParseInt(p.result[i], 2, 64); err == nil {
				p.result[i] = strconv.FormatInt(val, 10)
			}
		}
	}
}

// isPunctChar checks if character is punctuation
func isPunctChar(ch byte) bool {
	return ch == '.' || ch == ',' || ch == '!' || ch == '?' || ch == ':' || ch == ';'
}

// hasPunctuation checks if word contains punctuation
func hasPunctuation(word string) bool {
	for i := 0; i < len(word); i++ {
		if isPunctChar(word[i]) {
			return true
		}
	}
	return false
}

// isValidCommand validates command format and content
func isValidCommand(cmdText string) bool {
	if len(cmdText) < 3 || cmdText[0] != '(' || cmdText[len(cmdText)-1] != ')' {
		return false
	}

	cmd, _ := parseCommandText(cmdText)
	return cmd != ""
}

// parseCommandText extracts command name and count from command text
func parseCommandText(cmdText string) (string, int) {
	if len(cmdText) < 3 {
		return "", 0
	}

	// Remove parentheses
	content := strings.TrimSpace(cmdText[1 : len(cmdText)-1])

	// Split by comma or space
	var parts []string
	if strings.Contains(content, ",") {
		parts = strings.Split(content, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
	} else {
		parts = strings.Fields(content)
	}

	if len(parts) == 0 {
		return "", 0
	}

	cmd := parts[0]
	count := 1

	// Parse count if provided
	if len(parts) == 2 {
		if c, err := strconv.Atoi(parts[1]); err == nil && c > 0 {
			count = c
		} else {
			return "", 0
		}
	}

	// Validate command
	validCmds := map[string]bool{"up": true, "low": true, "cap": true, "hex": true, "bin": true}
	if !validCmds[cmd] {
		return "", 0
	}

	return cmd, count
}

// needsAn determines if "a" should be changed to "an" before this word
func needsAn(word string) bool {
	if len(word) == 0 {
		return false
	}

	first := strings.ToLower(string(word[0]))

	// Check for vowel sounds
	if strings.Contains("aeiou", first) {
		return true
	}

	// Special case for silent 'h' words
	if first == "h" {
		lower := strings.ToLower(word)
		silentH := []string{"hour", "honest", "honor", "heir"}
		for _, silent := range silentH {
			if strings.HasPrefix(lower, silent) {
				return true
			}
		}
	}

	return false
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

// runTest executes a single test case and reports result
func runTest(testNum int, input, expected string) {
	processor := NewProcessor(input)
	actual := processor.Process()

	if strings.TrimSpace(actual) == strings.TrimSpace(expected) {
		fmt.Printf("Test %d: PASS\n", testNum)
	} else {
		fmt.Printf("Test %d: FAIL\n", testNum)
		fmt.Printf("  Input:    %q\n", input)
		fmt.Printf("  Expected: %q\n", expected)
		fmt.Printf("  Actual:   %q\n", actual)
	}
}

// runAllTests executes all test cases from golden_tests.md
func runAllTests() {
	tests := []struct {
		input    string
		expected string
	}{
		{"1E (hex)", "30"},
		{"101 (bin)", "5"},
		{"Go (up)", "GO"},
		{"HELLO WORLD (low, 2)", "hello world"},
		{"brooklyn bridge (cap, 2)", "Brooklyn Bridge"},
		{"A orange tree.", "An orange tree."},
		{"He said: ' this is fine ' .", "He said: 'this is fine'."},
		{"Wait ,what ?!", "Wait, what?!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"10 files (up, 6)", "10 FILES"},
		{"1f (hex) files (up)", "31 FILES"},
		{"He shouted ' stop (up) '", "He shouted 'STOP'"},
		{"new york (cap) !", "New York!"},
		{"one two three (up, 2) four (low)", "one TWO THREE four"},
		{"101 (bin)!", "5!"},
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"He said: ''", "He said: ''"},
		{"A honest mistake.", "An honest mistake."},
		{"wow (up), amazing (low, 1)!", "WOW, amazing!"},
		{"I added 1E (hex) files ,and 10 (bin) were removed (low, 2) .", "I added 30 files, and 2 were removed."},
	}

	fmt.Println("Running all tests...")
	for i, test := range tests {
		runTest(i+1, test.input, test.expected)
	}
}
