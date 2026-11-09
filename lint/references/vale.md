# Vale

- **Source:** Joseph Kato and contributors, Vale, an open source prose linter, version 3.26.0, MIT License
- **Type:** Software
- **Where to find it:** https://github.com/vale-cli/vale and https://vale.sh
- **Last reviewed:** 2026-11-08
- **Used by:** clauditory-lint 2.1.0 (`lint/internal/vale`)

## What the source covers

Vale checks prose against rules written in YAML. Each rule extends one of a dozen checks. They are existence (a pattern that shouldn't appear), substitution (a pattern with a preferred replacement), occurrence (a count limit, such as words per sentence), and capitalization. The others are conditional (a term that must be defined before use), repetition, sequence (a pattern over part-of-speech tags), readability, and spelling. Vale reads Markdown, reStructuredText, AsciiDoc, HTML, and plain text, and it skips code spans and code blocks. Rules sit in folders called styles, and a configuration file chooses the styles and turns single rules on or off. Vale reports each alert with its line, column span, matched text, and an optional action: replace, remove, or a regular-expression edit. Vale doesn't apply those actions itself; its language server offers them in editors as quick fixes.

## What the linter applies

The linter compiles Vale in. Vale's internal packages are copied into `lint/internal/vale` with their import paths rewritten, so one Go program carries the engine and every style, and the agent running it needs nothing else installed. The linter writes a Vale configuration for the styles, runs Vale on the text in memory, and then does what Vale doesn't: it settles conflicts between overlapping alerts and applies the replace, remove, and edit actions to the file. For two rules whose fix Vale's actions can't express, writing times and spelling out small numbers, it adds its own fixers.

## What the linter omits

- **Code parsers.** Vale parses source code with tree-sitter grammars written in C, which would tie each build to a C toolchain for every platform. The copy drops them, along with the XML and DITA readers that call external programs, and uses the regex-based comment reader Vale ships for its WebAssembly build. The linter checks prose files, so it loses nothing it uses.
- **Package manager and global configuration.** The linter ignores a user's own Vale setup, so its results don't depend on the machine.

## Conflicts with other references

None. Vale is the engine, not a source of style rules.

## Review notes

I copied Vale 3.26.0 from the Go module proxy, because the release downloads weren't reachable from the build environment. The module path changed from `github.com/errata-ai/vale` to `github.com/vale-cli/vale` in the 3.x series. Vale's matcher is a backtracking regular-expression engine that tries every alternative of a rule at every position, so the linter matches the large extracted word lists itself (see `style-guide-rules.md`) and leaves Vale the hand-written and vendored styles. Spans in Vale's output count characters, not bytes, and the fixer counts the same way.
