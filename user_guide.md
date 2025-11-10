# Go-Reloaded User Guide

## Usage
```bash
go run . <input.txt> <output.txt>
go run . test
go run . golden
```

## Commands

### Case Conversion
- `(up)` - Convert to uppercase: `hello (up)` → `HELLO`
- `(low)` - Convert to lowercase: `HELLO (low)` → `hello`
- `(cap)` - Capitalize: `hello world (cap)` → `Hello World`

### Number Conversion
- `(hex)` - Hex to decimal: `1E (hex)` → `30`
- `(bin)` - Binary to decimal: `101 (bin)` → `5`

### Multiple Words
Use count to affect multiple previous words:
```
hello world (up, 2) → HELLO WORLD
one two three (cap, 2) → one Two Three
```

## Automatic Corrections

### Articles
- `A orange` → `An orange`
- `A honest` → `An honest`

### Punctuation
- `Hello , world !` → `Hello, world!`
- `Wait ... what ?` → `Wait... what?`

### Quotes
Commands work inside single quotes:
```
He said ' hello (up) ' → He said 'HELLO'
```

## Examples
```
I added 1E (hex) files ,and 10 (bin) were removed (low, 2) .
→ I added 30 files, and 2 were removed.

He shouted ' stop (up) right now (cap, 2) '
→ He shouted 'STOP Right Now'
```