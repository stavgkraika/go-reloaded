# Task 5: Handle Quotes

## Description
Format quoted text by removing spaces inside quotes: `' awesome '` → `'awesome'`, handle both single and multiple words.

## Test First
**TestHandleQuotes**: Test quote formatting for single and multi-word cases

## Implementation Goal
Implement quote detection and formatting (may transition to Quote state).

## Validation
Single and multi-word quote cases pass formatting tests.

## Acceptance Criteria
- [ ] "' word '" becomes "'word'"
- [ ] "' multiple words '" becomes "'multiple words'"
- [ ] Leading and trailing spaces inside quotes are removed
- [ ] Spaces between words inside quotes are preserved
- [ ] Nested or malformed quotes handled gracefully
- [ ] Tests pass for various quote scenarios

## Phase
Phase 2 – Reading State (Base Text Logic)