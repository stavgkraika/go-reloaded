# Task 12: Handle (up,n), (low,n), (cap,n) Commands

## Description
Transform N previous words using commands with count parameters like (up,2), (low,3), (cap,4).

## Test First
**TestCaseTransformationMulti**: Test multiple word case transformations with lookback

## Implementation Goal
Extend case transformation functionality with lookback buffer to affect N previous words.

## Validation
Only the specified N words are affected by the transformation.

## Acceptance Criteria
- [ ] "hello world (up,2)" becomes "HELLO WORLD"
- [ ] "one two three (cap,2)" becomes "one Two Three"
- [ ] "a b c d (low,3)" becomes "a b c d" (last 3 become lowercase)
- [ ] Count larger than available words handled gracefully
- [ ] Count of 0 or negative handled appropriately
- [ ] Lookback buffer maintains word history
- [ ] Tests pass for various count values

## Phase
Phase 3 – Command State (Transformations)