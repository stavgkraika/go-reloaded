# Go-Reloaded User Guide

## Overview
Go-Reloaded is a text processing tool that applies various transformations to text files including case conversion, number base conversion, article correction, and punctuation normalization.

## Installation & Usage

### Basic Usage
```bash
go run . <input.txt> <output.txt>
```

### Test Mode
```bash
go run . test
```

## Features

### 1. Case Conversion Commands

#### Uppercase `(up)`
Converts text to uppercase.
```
Input:  hello (up)
Output: HELLO
```

#### Lowercase `(low)`
Converts text to lowercase.
```
Input:  HELLO (low)
Output: hello
```

#### Capitalize `(cap)`
Capitalizes the first letter of each word.
```
Input:  hello world (cap)
Output: Hello World
```

### 2. Number Base Conversion

#### Hexadecimal to Decimal `(hex)`
Converts hexadecimal numbers to decimal.
```
Input:  1E (hex)
Output: 30
```

#### Binary to Decimal `(bin)`
Converts binary numbers to decimal.
```
Input:  101 (bin)
Output: 5
```

### 3. Command Modifiers

#### Applying to Multiple Words
Use a number after the command to specify how many previous words to affect.
```
Input:  hello world (up, 2)
Output: HELLO WORLD

Input:  one two three (cap, 2)
Output: one Two Three
```

### 4. Article Correction
Automatically corrects "a" to "an" before vowel sounds.
```
Input:  A orange tree
Output: An orange tree

Input:  A honest mistake
Output: An honest mistake
```

### 5. Punctuation Normalization
Removes spaces before punctuation marks.
```
Input:  Hello , world !
Output: Hello, world!

Input:  Wait ... what ?
Output: Wait... what?
```

### 6. Quote Processing
Processes commands inside single quotes.
```
Input:  He said ' hello (up) '
Output: He said 'HELLO'

Input:  She whispered ' stop (low) now '
Output: She whispered 'stop now'
```

## Command Syntax

### Basic Format
- Commands are enclosed in parentheses: `(command)`
- Commands can include a count: `(command, count)`
- Spaces around commas are optional: `(up,2)` or `(up, 2)`

### Valid Commands
- `up` - Convert to uppercase
- `low` - Convert to lowercase  
- `cap` - Capitalize first letter
- `hex` - Convert hexadecimal to decimal
- `bin` - Convert binary to decimal

### Count Parameter
- Default count is 1 (affects previous word)
- Count specifies how many previous words to modify
- If count exceeds available words, all available words are affected

## Examples

### Complex Text Processing
```
Input:  I added 1E (hex) files ,and 10 (bin) were removed (low, 2) .
Output: I added 30 files, and 2 were removed.
```

### Mixed Commands
```
Input:  one two three (up, 2) four (low)
Output: one TWO THREE four
```

### Quotes with Commands
```
Input:  He shouted ' stop (up) right now (cap, 2) '
Output: He shouted 'STOP Right Now'
```

## File Format

### Input Files
- Plain text files (.txt)
- Any encoding supported by Go
- Commands embedded within the text

### Output Files
- Processed text with all transformations applied
- Same encoding as input
- Maintains original structure with corrected spacing

## Error Handling

### Invalid Commands
Invalid commands are treated as regular text:
```
Input:  hello (invalid) world
Output: hello (invalid) world
```

### Malformed Numbers
Invalid hex/binary numbers remain unchanged:
```
Input:  XYZ (hex)
Output: XYZ (hex)
```

### Missing Files
The program will display an error message if input file doesn't exist or output file cannot be created.

## Testing
Run the built-in test suite to verify functionality:
```bash
go run . test
```

This runs 20 comprehensive tests covering all features and edge cases.