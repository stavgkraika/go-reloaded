# Task 6: Detect Transformation Commands

## Description
Recognize transformation command tokens like `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)` and transition to Command state.

## Test First
**TestDetectCommands**: Test command detection and state transition

## Implementation Goal
Implement detection logic that transitions FSM from Reading to Command state when commands are found.

## Validation
FSM correctly leaves Reading state when command detection occurs.

## Acceptance Criteria
- [ ] "(hex)" is detected as hex command
- [ ] "(bin)" is detected as binary command
- [ ] "(up)", "(low)", "(cap)" are detected as case commands
- [ ] Commands with parameters like "(up,2)" are detected
- [ ] Invalid command-like tokens are ignored
- [ ] State transition to Command occurs on detection
- [ ] Tests pass for all command types

## Phase
Phase 2 – Reading State (Base Text Logic)