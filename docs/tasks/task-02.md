# Task 2: Tokenizer and Input Reading

## Description
Split input text into tokens (words, punctuation, quotes, commands) while preserving all symbols and boundaries.

## Test First
**TestTokenizer**: Validate token boundaries for commands, quotes, punctuation

## Implementation Goal
Implement tokenizer that preserves symbols and correctly segments different types of tokens.

## Validation
Proper token segmentation confirmed for all token types.

## Acceptance Criteria
- [ ] Words are correctly tokenized
- [ ] Punctuation marks are separate tokens
- [ ] Quote marks are identified as tokens
- [ ] Command patterns like (hex), (up,2) are recognized as single tokens
- [ ] Whitespace handling preserves token boundaries
- [ ] Tests pass for various input combinations

## Phase
Phase 1 – Core FSM Setup (Foundation)