# clauditory-lint

clauditory-lint checks prose against the rules in the writing-style and technical-writing skills that a program can test, and it fixes the problems that have a single mechanical fix. It runs on Vale, an open source prose linter, and adds two things Vale lacks: a ranking of the style guides behind each rule, and a fixer that applies Vale's suggested edits to the file.

## How the sources rank

Each rule belongs to one of three domains, and each domain ranks its sources:

| Domain | Primary | Secondary | Tertiary |
| --- | --- | --- | --- |
| Grammar and punctuation | Oxford (New Hart's Rules) | AP | Chicago |
| Word choice and formatting | NYT | Chicago | AP |
| Technical communication | Red Hat | Apple | Microsoft |

A secondary rule applies where no primary rule disagrees with it, and a tertiary rule applies where neither a primary nor a secondary rule does. A rule from a source the domain doesn't list, such as a Red Hat rule about word choice, ranks below the tertiary source. The owner's own choices from the writing-style skill, in the House style, outrank every source: no em dashes, the serial comma, and the patterns in the skill's revision pass.

The linter applies the ranking in two places:

- **Conflicting rules.** When two rules demand opposite things, the lower-ranked rule is turned off. AP's rule against the serial comma loses to Oxford, and Chicago's rule on spelling out numbers to one hundred loses to NYT's one-through-nine rule. Microsoft's `4 PM` loses to NYT's `4 p.m.` The registry also turns off rules that contradict the skills or duplicate a better-ranked rule, and it records a reason for each.
- **Overlapping alerts.** When several rules flag the same words for the same problem, the linter reports only the best-ranked one.

Run `clauditory-lint rules` to see every rule with its domain, tier, fix setting, and status, including the reason each disabled rule is off.

## Sources

- **Red Hat and Microsoft.** The Red Hat rules are Red Hat's official Vale style (vale-at-red-hat), vendored unchanged. The Microsoft rules are the community Vale style for the Microsoft Writing Style Guide, also unchanged. Both are MIT-licensed; their licenses sit in their folders under `styles/`.
- **Oxford, NYT, AP, Chicago, and Apple.** No Vale style exists for these guides, so the rules under `styles/` for them are written for this project from the guides themselves. Each rule links to the reference file that records what the guide says.
- **House.** These rules encode the writing-style skill's own choices.

`references/` records what the linter takes from each source and why.

## Install

The linter needs Go 1.24 or later to build. From a clone of the repository, run:

```
bin/clauditory install-lint
```

This command builds clauditory-lint, installs Vale v3.26.0 beside it if no Vale is on your PATH, and copies both into `~/.local/bin`. Pass `--dest DIR` to install elsewhere.

To run the linter without installing it, use `bin/clauditory lint`, which takes the same flags.

## Use

```
clauditory-lint README.md docs/*.md        # report problems
clauditory-lint --fix README.md            # apply the safe fixes in place
clauditory-lint --fix=all README.md        # also apply fixes that need a look
git log -1 --format=%B | clauditory-lint   # check standard input as Markdown
clauditory-lint --format=json notes.md     # machine-readable output
```

The `--min` flag sets the lowest severity to report, and `--fail` sets the lowest severity that makes the exit status 1 (the default is error). Status 2 means the linter itself failed.

A safe fix has one correct result, such as `utilize` to `use`, `50%` to `50 percent`, or `5 November 2026` to `November 5, 2026`. A fix that needs a look, such as spelling out a small number, is right in most sentences but wrong in some, so it runs only with `--fix=all`. Rules that need judgment, such as trailing participles or correction framing, report the problem and leave the rewrite to the writer.

Vale skips code spans and code blocks in Markdown. To turn the linter off for part of a file, wrap it in `<!-- vale off -->` and `<!-- vale on -->`.

## Claude Code hook

The Clauditory plugin includes a hook that runs after Claude writes or edits a Markdown, text, reStructuredText, or AsciiDoc file. If clauditory-lint is on the PATH, the hook lints the file and passes any warnings and errors back to Claude, which can then fix them. If the linter isn't installed, the hook does nothing. Set `CLAUDITORY_LINT_HOOK=off` to turn it off, or put the text `clauditory-lint: off` anywhere in a file to skip that file.

## Layout

```
cmd/clauditory-lint/   the command
registry.go            loads rules.yml and applies the ranking
alerts.go              runs Vale, filters overlapping alerts, and applies fixes
rules.yml              the domain, tier, fix setting, and status of each rule
styles/                Vale styles, one folder per source
hooks/hooks.json       the Claude Code hook, copied into the plugin by bin/clauditory build
references/            the sources behind the linter
testdata/              fixtures for go test
```

To add a rule, write a Vale rule file in the source's folder under `styles/`, add an entry to `rules.yml` when it needs a domain, topic, fix setting, or conflict, and run `go test ./...`. The tests need Vale on the PATH or in `CLAUDITORY_VALE`.
