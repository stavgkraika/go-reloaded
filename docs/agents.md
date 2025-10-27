# AI and Developer Agents – Go Reloaded

## Overview
This document defines agents and roles used to assist in an Agile, TDD-driven workflow for the Go FSM Text Processor.

## Phase Mapping
- Test Writer Agent: All phases (creates test stubs).
- Implementation Agent: Phases 1–5.
- Refactor Agent: After each phase.
- Integration Agent: Phase 6.
- Reviewer Agent: Continuous.
- Builder Agent: Final phase (file I/O and CLI).

## Agents

### Test Writer Agent
Goal: Create initial failing tests for each FSM feature.
Responsibilities:
- Write failing tests before implementation.
- Use examples from the roadmap and golden tests.
- Validate test syntax with go test -c.

Outputs:
- _test.go files with clear, isolated cases.

### Implementation Agent
Goal: Implement minimal code to pass the current failing test.
Responsibilities:
- Modify specific modules (e.g., reading.go, command.go).
- Run go test and ensure green tests.
- Keep commits small and focused.

Outputs:
- Updated .go source files that pass tests.

### Refactor Agent
Goal: Optimize, clean, and document code after tests are green.
Responsibilities:
- Refactor for readability and maintainability.
- Add comments and GoDoc-style documentation.
- Ensure no test regressions.

Outputs:
- Cleaner code with unchanged behavior.

### Integration Agent
Goal: Verify full FSM interaction across all states.
Responsibilities:
- Run end-to-end tests.
- Validate state transitions and combined behavior.
- Use golden tests for final validation.

Outputs:
- Integration test reports.

### Reviewer Agent
Goal: Enforce code quality and style compliance.
Responsibilities:
- Run go fmt, golint, and go vet.
- Check commit messages follow TDD sequence (Test → Implement → Refactor).
- Flag untested code paths and missing docs.

Outputs:
- Review report and approval for merge.

### Builder Agent
Goal: Manage file I/O and CLI interface.
Responsibilities:
- Implement argument parsing (input.txt output.txt).
- Connect FSM to file system.
- Ensure file read/write tests pass.

Outputs:
- main.go and related wiring.

## Agent Handoff Protocol
1. Test Writer commits failing test.
2. Implementation Agent commits minimal passing code.
3. Refactor Agent cleans and documents code.
4. Reviewer Agent runs static analysis and approves.
5. Integration Agent validates full flow with golden tests.
6. Builder Agent finalizes CLI and file I/O.

## Trigger Rules
- If unit tests pass: notify Refactor Agent.
- If coverage is below target: return to Test Writer Agent.
- If lint or vet errors exist: notify Reviewer Agent.
- If integration fails: notify Implementation and Integration Agents.

## Task Ownership Table

| Agent               | Responsible For                                       | Output Artifacts            |
|---------------------|--------------------------------------------------------|-----------------------------|
| Test Writer         | Unit and integration tests                             | _test.go, reports           |
| Implementation      | Passing code for each feature                          | .go modules                 |
| Refactor            | Code cleanup and documentation                         | Updated .go files           |
| Integration         | End-to-end validation and golden tests                 | Integration reports         |
| Reviewer            | Formatting, linting, vetting, and policy checks        | Review summary              |
| Builder             | CLI entry point, file I/O integration                  | main.go, wiring             |
