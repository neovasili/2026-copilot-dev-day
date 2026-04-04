# Why `copilot-instructions.md` and `AGENTS.md` both matter

This is a short talk track you can use before the live prompt.

## The short version

- `AGENTS.md` defines your engineering contract: scope discipline, coding style, testing rules, and stop conditions.
- `.github/copilot-instructions.md` is the **Copilot integration point** that makes repository instructions automatically available to Copilot requests.
- Best setup: keep durable policy in `AGENTS.md`, and use `copilot-instructions.md` to tell Copilot to follow it every time.

## Purpose of `.github/copilot-instructions.md`

Use it to provide repository-wide custom instructions that Copilot can automatically include in prompts for this repo.

Why this helps:

- Less prompt repetition in every chat.
- More consistent outputs across edits.
- Easier onboarding: team rules live in version control.

## Purpose of `AGENTS.md`

Use it as a tool-agnostic behavior contract for AI agents and humans:

- smallest-diff policy
- quality-gate loop
- style and testability constraints
- expected output format

`AGENTS.md` is where you define **how work should be done**, not just coding style.

## Key difference

- `.github/copilot-instructions.md` answers: "How does Copilot receive repository instructions automatically?"
- `AGENTS.md` answers: "What exact behavior/rules should an agent follow while changing code?"

They are complementary, not redundant.

## What happens if we do not use them

Without `.github/copilot-instructions.md`:

- Copilot may miss repository conventions unless you restate them manually each time.
- Responses become less consistent across prompts/sessions.

Without `AGENTS.md`:

- You lose a clear, reusable workflow contract (minimal diff, run checks, iterate to green).
- Guidance tends to drift into ad hoc prompting.

Without both:

- Higher prompt overhead.
- Higher chance of larger-than-needed diffs and inconsistent quality discipline.

## 30-second on-stage summary

"`AGENTS.md` defines the rules. `copilot-instructions.md` makes Copilot actually apply those rules by default. If we skip either one, we get more prompt friction and less consistency."

## Official references

- GitHub Docs - Adding repository custom instructions:
  <https://docs.github.com/en/copilot/how-tos/configure-custom-instructions/add-repository-instructions>
- GitHub Docs - Copilot CLI best practices (supported instruction locations, including `AGENTS.md`):
  <https://docs.github.com/en/copilot/how-tos/copilot-cli/cli-best-practices>
