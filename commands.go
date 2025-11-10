package main

import (
	"strconv"
	"strings"
)

// applyCommand applies transformation commands to previous words
func (p *Processor) applyCommand(cmdText string) {
	cmd, count := parseCommandText(cmdText)
	if cmd == "" {
		return
	}

	// Special handling for cap command without explicit count
	if cmd == "cap" && count == 1 {
		// Find start of current phrase (consecutive non-punctuation words)
		start := len(p.result) - 1
		for start > 0 && !isPunctChar(p.result[start-1][len(p.result[start-1])-1]) {
			start--
		}
		// Capitalize all words in phrase
		for i := start; i < len(p.result); i++ {
			p.result[i] = capitalize(p.result[i])
		}
		return
	}

	// Apply command to specified number of previous words
	if len(p.result) < count {
		count = len(p.result)
	}
	if count == 0 {
		return
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
			if val, err := strconv.ParseInt(p.result[i], 16, 64); err == nil {
				p.result[i] = strconv.FormatInt(val, 10)
			}
		case "bin":
			if val, err := strconv.ParseInt(p.result[i], 2, 64); err == nil {
				p.result[i] = strconv.FormatInt(val, 10)
			}
		}
	}
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
	content := strings.TrimSpace(cmdText[1 : len(cmdText)-1])
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
	if len(parts) == 2 {
		if c, err := strconv.Atoi(parts[1]); err == nil && c > 0 {
			count = c
		} else {
			return "", 0
		}
	}
	validCmds := map[string]bool{"up": true, "low": true, "cap": true, "hex": true, "bin": true}
	if !validCmds[cmd] {
		return "", 0
	}
	return cmd, count
}