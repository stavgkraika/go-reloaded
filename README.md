# Go-Reloaded

A Go text processor that applies various transformations to input text including case conversions, number base conversions, article corrections, and punctuation formatting.

## Features

- **Number Base Conversion**: Convert hexadecimal and binary numbers to decimal
- **Case Transformations**: Apply uppercase, lowercase, and capitalization to words
- **Article Correction**: Automatically fix "a" vs "an" usage
- **Quote Formatting**: Clean up spacing around quotes
- **Punctuation Spacing**: Fix spacing around punctuation marks

## Usage

### Command Line
```bash
go run . input.txt output.txt
```

### Programmatic
```go
processor := NewProcessor("your text here")
result := processor.Process()
```

## Transformations

### Number Conversions
- `1E (hex)` → `30`
- `101 (bin)` → `5`

### Case Modifications
- `word (up)` → `WORD`
- `WORD (low)` → `word`
- `word (cap)` → `Word`
- `words here (up, 2)` → `WORDS HERE`

### Article Correction
- `A orange` → `An orange`
- `An lemon` → `A lemon`

### Quote & Punctuation Cleanup
- `' text '` → `'text'`
- `word ,text` → `word, text`

## Testing

Run all tests:
```bash
go run . golden
```

## Files

- `main.go` - Entry point and file I/O
- `processor.go` - Main text processing logic
- `parser.go` - Token parsing functions
- `commands.go` - Command application logic
- `utils.go` - Utility functions
- `tests/golden_tests.md` - Test cases