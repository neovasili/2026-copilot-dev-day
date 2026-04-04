# Q&A bank for your session

## 1) "Can Copilot write idiomatic Go reliably?"

Answer:
- Often yes for common patterns, but quality varies by prompt context.
- Best results come from constraints, tests, and small reviewable edits.

## 2) "How do I avoid over-trusting generated code?"

Answer:
- Treat suggestions like junior teammate proposals.
- Ask for acceptance criteria first, then verify with tests and review.

## 3) "Should I prompt in English or Spanish?"

Answer:
- Both work; clarity matters more than language.
- For technical precision, many teams prefer English identifiers and constraints.

## 4) "When should I reject a suggestion?"

Answer:
- Reject when behavior is unclear, complexity increases, or style diverges from team rules.

## 5) "Can Copilot help with tests, not only implementation?"

Answer:
- Yes. Ask for table-driven tests, edge cases, and explicit failure scenarios.

## 6) "What is the minimum safe workflow?"

Answer:
- Prompt with constraints -> inspect diff -> run tests -> iterate.

## 7) "What if internet is unstable during demo?"

Answer:
- Keep a local backup path: pre-opened files, pre-run tests, and one static walkthrough.
