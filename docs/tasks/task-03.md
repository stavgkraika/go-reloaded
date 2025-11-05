# Task 3: Handle a → an Rule

## Description
Convert "a" to "an" if the next word starts with a vowel or h (with exceptions for words like "historic", "hotel").

## Test First
**TestHandleAtoAn**: Test conversion logic for various vowel and h-starting words

## Implementation Goal
Implement handleAtoAn() function in Reading state that correctly applies the a/an rule.

## Validation
All test cases pass including edge cases and exceptions.

## Acceptance Criteria
- [ ] "a apple" becomes "an apple"
- [ ] "a elephant" becomes "an elephant"
- [ ] "a historic" remains "a historic" (exception)
- [ ] "a hotel" remains "a hotel" (exception)
- [ ] "a house" becomes "an house"
- [ ] Case sensitivity handled correctly
- [ ] Tests pass for all vowel and h cases

## Phase
Phase 2 – Reading State (Base Text Logic)