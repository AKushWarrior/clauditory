# Clauditory

Clauditory holds agent-facing prompts, skills, and hooks for chat, coding, and computer-use agents. Each artifact rests on published research or established references, and each one includes a `references` folder that records which sources shaped it and how.

## Write everything in the house register

Every piece of prose in this repository follows the writing-style skill at `chat/skills/writing-style/SKILL.md`, and documentation, code comments, and commit messages also follow the technical-writing skill at `chat/skills/technical-writing/SKILL.md`. That covers reference files, system prompts, skill text, READMEs, code comments meant for people, and every commit message. Read the skills before you write anything here, draft with them, and run their revision passes before you save or commit.

The reason is consistency. Agents that work in this repository read its files and its history, and they pick up the register they see. A repository written entirely in one clear, plain voice primes the next agent to write that way too. A single file or commit in default machine prose weakens that effect for everyone who reads it afterward.

## Layout

```
AGENTS.md                 this file
README.md                 what the repository holds and how to install it
bin/clauditory            command-line helper: list, install, package, build, prompt, lint
chat/
  system/
    prompt.md             system prompt for chat agents
    references/           one file per source behind the prompt
  skills/
    <skill-name>/
      SKILL.md            the skill, in Agent Skills format, with metadata.version
      CHANGELOG.md        what changed in each version
      references/         one file per source behind the skill
coding/                   (planned) artifacts for coding agents, same layout as chat/
lint/                     clauditory-lint, a Go prose linter built on Vale
  rules.yml               each rule's domain, tier, fix setting, and status
  styles/                 Vale styles, one folder per source
  references/             one file per source behind the linter
plugins/clauditory/       generated plugin for Claude Code, the Claude apps, and Codex
.claude-plugin/           Claude plugin marketplace that lists plugins/clauditory
.agents/plugins/          Codex plugin marketplace that lists plugins/clauditory
```

The plugin under `plugins/clauditory` is generated. Edit the sources under `chat/`, then run `bin/clauditory build` and commit the result; `bin/clauditory build --check` fails when the plugin is stale. Bump the version in both plugin manifests when a change should reach people who already installed it.

Every skill carries a version in its frontmatter, as `metadata.version`, and a `CHANGELOG.md` beside its `SKILL.md`. Use semantic versioning: raise the major version when a rule changes in a way that alters existing text, the minor version when you add or deepen rules, and the patch version for wording fixes that change no rule. Update the changelog in the same commit.

The writing-style skill sets the house register for any agent writing prose, including agents working in this repository, and the technical-writing skill adds rules for documentation, code comments, and commits. When their sources conflict, the news and book manuals rank first, in this order: The New York Times Manual of Style and Usage, the AP Stylebook, The Chicago Manual of Style, and the New Oxford Style Manual. Where Oxford's British conventions differ from American ones, the American convention wins. The technical guides from Google, Microsoft, Apple, IBM, and Red Hat and the git and Linux kernel patch guides come after them. The owner's explicit choices override every source: no em dashes, the serial comma, and the revision-pass bans.

The linter ranks sources by domain instead of in one list. For grammar and punctuation the order is Oxford, AP, and Chicago; for word choice and formatting, NYT, Chicago, and AP; and for technical communication, Red Hat, Apple, and Microsoft. The owner's House rules outrank all of them. A lower-ranked rule applies only where no higher-ranked rule disagrees, and `lint/README.md` explains how the linter settles each conflict. The linter carries a version in `lint/registry.go` and a `CHANGELOG.md` and follows the same versioning rule as the skills.

Run `bin/clauditory lint` over prose you write here before you commit it, and fix what it reports where the rule applies.

Use the most widely supported format for each kind of artifact. Skills follow the Agent Skills format: a folder holding a `SKILL.md` file with `name` and `description` frontmatter. System prompts are plain Markdown.

## Reference files

Every artifact keeps a `references/` folder beside it, with one Markdown file per source. Name each file with a short slug of the source, such as `nyt-manual-of-style.md` or `kobak-excess-vocabulary.md`. Each file uses these sections, in this order:

1. A title and a short metadata list: full citation, source type, where to find it, last-reviewed date, and the artifact that uses it.
2. What the source covers: an analysis of its contents, methods, findings, and principles, several pages long for substantial sources.
3. What the artifact applies, mapped to the specific sections and rules it shaped.
4. What the artifact omits from the source, and why.
5. Conflicts with other references in the same folder, and how the artifact resolved them.
6. Review notes: what was read, what could not be reached, and any doubts about the reading.

Sources that make claims about how AI systems perform must be recent, because model capability changes fast. A source of that kind must have been first published less than eight months before its last-reviewed date. Date it by its first publication, such as an arXiv paper's first version, not its latest revision. Living documentation, such as a vendor's current prompting guide, counts from the version read, if it covers current models. Studies of human readers and writers, style manuals, and other sources that make no claim about AI performance have no age limit. When a source ages out, remove its reference file at the next revision of the artifact, and either replace it with newer evidence or drop the rule it alone supported.

Paraphrase everything. Reference files contain no quotations from the source and no copies of it, its figures, or its tables. When you change an artifact, re-review every source behind the change and update the last-reviewed date in each file you touched.

## Commit messages

Commit messages follow section 8 of the technical-writing skill. In short, write an imperative subject line of about 50 characters with no closing period, such as "Add the Rhodes reference for the writing-style skill." Then leave a blank line and write a body wrapped at about 72 characters that explains the problem and why the change solves it. Name the files or behavior that changed, and skip anything the diff already makes obvious.

Commit as the repository owner. Don't add trailers that credit or link to an AI model or session, such as Co-Authored-By or Claude-Session lines.

## Using `bin/clauditory`

Run `bin/clauditory list` to see the available artifacts. `bin/clauditory install chat/skills/writing-style` copies a skill into `~/.claude/skills/` for Claude Code; pass `--dest` to install somewhere else. `bin/clauditory package chat/skills/writing-style` builds a zip you can upload as a skill in the Claude apps. Both commands omit the `references/` folder and `CHANGELOG.md`, since those document the work for people and would only add noise to an agent's context.
