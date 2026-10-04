# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Mangrove Auth is a self-hosted authentication server written in Go. It serves as the identity layer for the Open Admin ecosystem (per README.md).

## Labels

Issues and pull requests use these labels:

- `bug`: something isn't working as expected
- `enhancement`: a new feature or an improvement request
- `documentation`: improvements or additions to the docs
- `test`: adding or improving tests
- `refactor`: code changes that neither fix a bug nor add a feature
- `chore`: maintenance and tooling work

## Issue template

Issues follow the general template in `.github/ISSUE_TEMPLATE/general.md`, with these sections in order: Summary, Acceptance criteria (a `- [ ]` checklist), Additional notes (a `-` list), and References (a `-` list written as `- [Label](URL)`). The template sets no default labels or assignees, so pick the label when creating the issue.

## Branching strategy

The project follows GitHub Flow: `main` is always stable and is never committed to directly. Each change is made on a short-lived branch created from `main`, merged back through a pull request, and the branch is deleted afterwards.

Branches are named `<type>/<issue-number>-<short-description>`, e.g. `chore/0-project-repo-setup`. `<type>` is the Conventional Commits type, the same one used in commit messages, and maps to the issue label: `enhancement` → `feat`, `bug` → `fix`, `documentation` → `docs`; `test`, `refactor` and `chore` are unchanged.
