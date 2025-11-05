# FSM Design Overview (No-Regex Single Pass)

## 1. Design Philosophy
The FSM (Finite State Machine) implementation of Go Reloaded operates in a single forward pass, processing the input text sequentially and transforming it according to defined rules.
Regular expressions are intentionally avoided to improve clarity, performance, and maintainability.

## 2. Tokenization Strategy
The tokenizer reads the input rune by rune, grouping characters into tokens such as:
- Words
- Punctuation marks
- Command markers (e.g., (hex), (up,2))
- Quotation marks (')

### Tokenization Algorithm (Simplified)
```go
for each rune in text:
    if whitespace:
        flush current token
    else if punctuation or quote:
        flush current token
        emit punctuation as its own token
    else:
        accumulate rune in current token
```

Each token is appended to a list and later processed by the FSM, maintaining context awareness throughout.

## 3. FSM Flow
Each token is processed once in a single-pass loop:

| State | Description | Transitions |
|--------|--------------|-------------|
| Reading | Default mode, scans normal text | May transition to Command, Quote, or Punctuation |
| Command | Handles commands like (hex), (bin), (up,n) | Returns to Reading |
| Quote | Processes text inside '...' | Returns to Reading on closing quote |
| Punctuation | Handles punctuation formatting | Returns to Reading |

### Example Flow
```
Reading → Command → Reading → Quote → Reading → Punctuation → Reading
```

## 4. Command Parsing (Without Regex)
Command detection is performed using simple string logic:

```go
func parseCommandToken(tok string) (cmd string, count int, valid bool) {
	if len(tok) < 4 || tok[0] != '(' || tok[len(tok)-1] != ')' {
		return "", 0, false
	}
	body := tok[1 : len(tok)-1]
	parts := strings.Split(body, ",")
	cmd = strings.TrimSpace(parts[0])
	if len(parts) == 2 {
		c, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err == nil {
			count = c
		}
	}
	return cmd, count, true
}
```

This method:
- Avoids regex
- Works deterministically
- Keeps parsing context within FSM control

## 5. Advantages of No-Regex FSM
| Aspect | FSM Approach | Regex Approach |
|---------|---------------|----------------|
| Performance | O(n), single pass | O(n²) in complex regex cases |
| Control | Fine-grained | Limited by regex engine |
| Debuggability | High | Low |
| Context Awareness | Strong | Weak |
| Readability | Clear | Dense |

## 6. Implementation Guidelines
- Use strings and unicode packages for token handling.
- Avoid regexp entirely.
- Keep FSM transitions explicit and deterministic.
- Maintain state machine logic separate from transformation logic.
