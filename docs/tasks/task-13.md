# Task 13: Handle Malformed Commands

## Description
Skip invalid command syntax safely and return to Reading state without crashing or hanging.

## Test First
**TestInvalidCommandRecovery**: Test FSM recovery from malformed commands

## Implementation Goal
Add fast validation to detect and skip malformed commands like `(up,)`, `(hex,xyz)`, `(invalid)`.

## Validation
FSM recovers gracefully and continues processing without performance degradation.

## Acceptance Criteria
- [ ] `(up,)` - missing or invalid count parameter skipped
- [ ] `(hex,xyz)` - non-numeric parameter skipped  
- [ ] `(invalid)` - unknown command skipped
- [ ] `(up,0)` and `(up,-1)` - invalid counts handled
- [ ] Malformed commands don't affect surrounding text
- [ ] No infinite loops or hangs
- [ ] Performance remains fast with many malformed commands

## Performance Requirements
- Command validation must complete in O(1) time
- No regex or complex parsing for invalid commands
- Simple string checks only

## Phase
Phase 3 – Command State (Transformations)