# Changelog

Versions follow semantic versioning. A major version turns rules on or off, or changes what they report, in a way that changes existing results, a minor version adds rules or fixes, and a patch version corrects a rule's pattern or wording.

## 2.0.0 (2026-11-05)

- Compile Vale in, so the linter is one program with no other install, and ship it xz-compressed for Linux, macOS, and Windows inside the writing-style skill.
- Add the rules extracted page by page from the NYT, AP, Chicago, Oxford, IBM, Apple, Google, Microsoft, and Red Hat guides, matched by a word-list lookup and a prefiltered pattern pass that are much faster than Vale's matcher on lists this size.
- Add the vale-ai-tells style for machine-writing tells, and a commit-message mode, `--commit`.
- Catch every dash and every run of hyphens standing in for a dash.
- Show each alert with its rule and message only.

## 1.0.0 (2026-11-05)

- First version: Vale styles for House, Oxford, NYT, AP, Chicago, and Apple rules, the vendored Red Hat and Microsoft styles, conflict handling between them, a fixer for Vale's replace and edit actions, and a Claude Code hook.
