# FSM Text Processor – TDD Roadmap (Go)

This document defines the Agile and TDD development plan for the Finite State Machine (FSM)-based text transformation tool written in Go.

Each task:
- Starts with writing a failing test (TDD-first).
- Proceeds with incremental implementation.
- Ends with validation through passing tests.
- Is designed for entry-level developers using AI-assisted coding.

---

## Overview

### FSM States
- Reading: Base state, normal text scanning, rule application.
- Command: Handles transformation commands like `(hex)`, `(bin)`, `(up,n)`.
- Quote: Handles quoted text `' ... '`.
- Punctuation: Handles punctuation formatting and spacing.

---

## Phase 1 – Core FSM Setup (Foundation)

| # | Task | Description | Test First | Implementation Goal | Validation |
|---|------|-------------|------------|---------------------|------------|
| 1 | Define FSM states and transitions | Create enum or constants for Reading, Command, Quote, Punctuation | TestFSMInitialization: FSM starts in Reading, transitions compile | Implement FSM structure with state switch logic | Transitions trigger as expected |
| 2 | Tokenizer and input reading | Split input text into tokens (words, punctuation, quotes, commands) | TestTokenizer: validate token boundaries for commands, quotes, punctuation | Implement tokenizer preserving symbols | Proper token segmentation confirmed |

---

## Phase 2 – Reading State (Base Text Logic)

| # | Task | Description | Test First | Implementation Goal | Validation |
|---|------|-------------|------------|---------------------|------------|
| 3 | Handle a → an rule | Convert “a” to “an” if next word starts with vowel or h (with exceptions) | TestHandleAtoAn | Implement handleAtoAn() in Reading | All cases pass |
| 4 | Normalize punctuation spacing | Ensure . , ! ? : ; stick to previous word, spaced from next | TestNormalizePunctuation | Implement normalizePunctuation() in Reading | Spacing fixed |
| 5 | Handle quotes | Format `' awesome '` → `'awesome'`, handle multiple words | TestHandleQuotes | Implement quote detection and formatting (may call Quote state) | Single and multi-word cases pass |
| 6 | Detect transformation commands | Recognize `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)` tokens | TestDetectCommands | Implement detection logic, transition to Command | FSM leaves Reading on command detection |
| 7 | Integration test – Reading | Combine above rules in one paragraph | TestReadingIntegration | Ensure full Reading pipeline executes | Complex sentences pass |

---

## Phase 3 – Command State (Transformations)

| # | Task | Description | Test First | Implementation Goal | Validation |
|---|------|-------------|------------|---------------------|------------|
| 8 | Parse command tokens | Identify command and optional parameter `(up,2)` | TestParseCommandToken | Implement parseCommand() returning command type and count | Command parsing accurate |
| 9 | Execute (hex) | Convert previous hex word to decimal | TestHexConversion | Implement applyHexConversion() | Numeric result valid |
| 10 | Execute (bin) | Convert previous binary word to decimal | TestBinConversion | Implement applyBinConversion() | Numeric result valid |
| 11 | Execute (up), (low), (cap) | Modify previous word’s case | TestCaseTransformationSingle | Implement applyCaseTransformation() with count=1 | Correct capitalization |
| 12 | Handle (up,n), (low,n), (cap,n) | Transform N previous words | TestCaseTransformationMulti | Extend case transformation with lookback buffer | Only N words affected |
| 13 | Handle malformed commands | Skip invalid syntax safely (e.g., `(up,)`, `(hex,xyz)`) | TestInvalidCommandRecovery | Add validation and safe return to Reading | FSM recovers without crash |
| 14 | Integration test – Command | Multiple commands in sequence | TestCommandIntegration | Confirm transitions back to Reading after each | Seamless flow |

---

## Phase 4 – Quote State (Text Within Quotes)

| # | Task | Description | Test First | Implementation Goal | Validation |
|---|------|-------------|------------|---------------------|------------|
| 15 | Detect opening quote | Transition Reading → Quote on first `'` | TestDetectQuoteStart | Implement state change and start buffer | Opening quote recognized |
| 16 | Accumulate text in Quote | Collect text verbatim until closing `'` | TestAccumulateQuotedText | Append words inside quotes, ignore commands | Text preserved |
| 17 | Detect closing quote | Exit Quote, trim spaces, emit `'text'` | TestDetectQuoteEnd | Implement detection and output | Formatting correct |
| 18 | Handle unclosed quotes | Graceful EOF handling with unmatched `'` | TestUnclosedQuote | Add recovery logic | No crash; output consistent |

---

## Phase 5 – Punctuation State (Formatting Rules)

| # | Task | Description | Test First | Implementation Goal | Validation |
|---|------|-------------|------------|---------------------|------------|
| 19 | Detect punctuation | Enter Punctuation on . , ! ? : ; | TestDetectPunctuation | Transition Reading → Punctuation | Detection works |
| 20 | Normalize single punctuation | Attach to previous word, add space after | TestNormalizeSinglePunct | Implement normalization | Output correct |
| 21 | Handle grouped punctuation | Support ... and !? without spacing issues | TestGroupedPunctuation | Merge consecutive punctuation marks | Special cases handled |
| 22 | Trim redundant spaces | Remove extra spaces before punctuation | TestTrimPunctuationSpaces | Implement cleanup | No double spaces remain |
| 23 | Integration test – Punctuation | Mix punctuation and words | TestPunctuationIntegration | Ensure transitions back to Reading | Output matches rules |

---

## Phase 6 – System Integration and Validation

| # | Task | Description | Test First | Implementation Goal | Validation |
|---|------|-------------|------------|---------------------|------------|
| 24 | Full FSM integration test | End-to-end test with all rules | TestFullFSMIntegration | Connect all states and transitions | All transformations applied correctly |
| 25 | File I/O layer | Read input file, write output file | TestFileIO | Implement CLI argument handling and file operations | Input/output processed correctly |
| 26 | Performance and edge cases | Stress test with large input, odd spacing, invalid commands | TestPerformanceAndEdgeCases | Optimize token handling and error recovery | Stable output and runtime |
| 27 | Golden test validation | Validate against project golden tests | TestGoldenSet | Verify outputs match expected golden results | Functional compliance confirmed |
