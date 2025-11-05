# Task 4: Normalize Punctuation Spacing

## Description
Ensure punctuation marks (. , ! ? : ;) stick to the previous word and are properly spaced from the next word.

## Test First
**TestNormalizePunctuation**: Test spacing rules for all punctuation types

## Implementation Goal
Implement normalizePunctuation() function in Reading state that fixes spacing around punctuation.

## Validation
Spacing is fixed according to standard punctuation rules.

## Acceptance Criteria
- [ ] "word ." becomes "word."
- [ ] "word." becomes "word. "
- [ ] "word , word" becomes "word, word"
- [ ] Multiple spaces before punctuation are removed
- [ ] Proper spacing after punctuation is ensured
- [ ] All punctuation types handled: . , ! ? : ;
- [ ] Tests pass for various spacing scenarios

## Phase
Phase 2 – Reading State (Base Text Logic)