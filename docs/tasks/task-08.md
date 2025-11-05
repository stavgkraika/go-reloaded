# Task 8: Parse Command Tokens

## Description
Identify command type and optional parameters from tokens like `(up,2)`, extracting both command and count values.

## Test First
**TestParseCommandToken**: Test command parsing accuracy for various formats

## Implementation Goal
Implement parseCommand() function that returns command type and count parameters.

## Validation
Command parsing is accurate for all supported command formats.

## Acceptance Criteria
- [ ] "(hex)" parsed as command="hex", count=1
- [ ] "(up,2)" parsed as command="up", count=2
- [ ] "(low,5)" parsed as command="low", count=5
- [ ] Invalid formats return appropriate error/default values
- [ ] Whitespace in commands handled correctly
- [ ] Edge cases like "(up,)" handled gracefully
- [ ] Tests pass for all parsing scenarios

## Phase
Phase 3 – Command State (Transformations)