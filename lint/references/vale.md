# Vale

- **Source:** Joseph Kato and contributors, Vale, an open source prose linter, version 3.26.0, MIT License
- **Type:** Software
- **Where to find it:** https://github.com/vale-cli/vale and https://vale.sh
- **Last reviewed:** 2026-11-05
- **Used by:** clauditory-lint 1.0.0

## What the source covers

Vale checks prose against rules written in YAML. Each rule extends one of a dozen checks: existence (a pattern that shouldn't appear), substitution (a pattern with a preferred replacement), occurrence (a count limit, such as words per sentence), capitalization, conditional (a term that must be defined before use), repetition, sequence (a pattern over part-of-speech tags), readability, and spelling. Vale understands Markdown, reStructuredText, AsciiDoc, HTML, and plain text, and it skips code spans and code blocks. Rules sit in folders called styles, and a configuration file chooses the styles and turns single rules on or off. Vale reports each alert with its line, column span, matched text, and an optional action: replace, remove, or a regular-expression edit. Vale doesn't apply those actions itself; its language server offers them in editors as quick fixes.

## What the linter applies

The linter uses Vale as its engine. It embeds every style, writes a Vale configuration that turns off the rules the hierarchy disables, runs Vale with JSON output, and reads the alerts. It then does the two things Vale doesn't: it ranks overlapping alerts by the source hierarchy, and it applies the replace, remove, and edit actions to the file. For two rules whose fix Vale's actions can't express, writing times and spelling out small numbers, it adds its own fixers.

## What the linter omits

The linter doesn't use Vale's package manager, global configuration, or vocabularies. It passes `--no-global` so a user's own Vale setup can't change its results.

## Conflicts with other references

None. Vale is the engine, not a source of style rules.

## Review notes

I built Vale 3.26.0 from source with the Go toolchain, because the release downloads weren't reachable from the build environment. The module path changed from `github.com/errata-ai/vale` to `github.com/vale-cli/vale` in the 3.x series, and the install command uses the new path. Spans in Vale's JSON output count characters, not bytes, and the fixer counts the same way.
