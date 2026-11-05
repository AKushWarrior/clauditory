# Clauditory

Clauditory holds agent-facing prompts, skills and hooks for chat, coding and computer-use agents. Each artifact rests on published research or established references, and each one ships with a `references` folder that records which sources shaped it and how.

## Write everything in the house register

Every piece of prose in this repository follows the writing-style skill at `chat/skills/writing-style/SKILL.md`. That covers documentation, reference files, system prompts, skill text, code comments meant for people and every commit message. Read the skill before you write anything here, draft with it, and run its revision pass before you save or commit.

The reason is consistency. Agents that work in this repository read its files and its history, and they pick up the register they see. A repository written entirely in one clear, plain voice primes the next agent to write that way too. A single file or commit in default machine prose weakens that effect for everyone who reads it afterward.

## Layout

```
AGENTS.md                 this file
bin/clauditory            command-line helper for listing, installing and packaging artifacts
chat/
  system/                 (planned)
    prompt.md             system prompt for chat agents
    references/           one file per source behind the prompt
  skills/
    <skill-name>/
      SKILL.md            the skill, in Agent Skills format
      references/         one file per source behind the skill
coding/                   (planned) artifacts for coding agents, same layout as chat/
```

The first artifact is `chat/skills/writing-style`. It lives under `chat/` because it was written for chat agents, but its rules apply to any agent writing prose, including agents working in this repository. Its sources rank in this order when they conflict: The New York Times Manual of Style and Usage, the AP Stylebook, The Chicago Manual of Style, and then the technical guides (Google, Microsoft and Apple style guides and the git and Linux kernel patch guides), which fill in rules for documentation, code comments and commit messages.

Use the most widely supported format for each kind of artifact. Skills follow the Agent Skills format: a folder holding a `SKILL.md` file with `name` and `description` frontmatter. System prompts are plain Markdown.

## Reference files

Every artifact keeps a `references/` folder beside it, with one Markdown file per source. Name each file with a short slug of the source, such as `nyt-manual-of-style.md` or `kobak-excess-vocabulary.md`. Each file uses these sections, in this order:

1. A title and a short metadata list: full citation, source type, where to find it, last-reviewed date and the artifact that uses it.
2. What the source covers: an analysis of its contents, methods, findings and principles, several pages long for substantial sources.
3. What the artifact applies, mapped to the specific sections and rules it shaped.
4. What the artifact omits from the source, and why.
5. Conflicts with other references in the same folder, and how the artifact resolved them.
6. Review notes: what was read, what could not be reached and any doubts about the reading.

Paraphrase everything. Reference files contain no quotations from the source and no copies of it, its figures or its tables. When you change an artifact, re-review every source behind the change and update the last-reviewed date in each file you touched.

## Commit messages

Commit messages follow section 9 of the writing-style skill. In short: an imperative subject line of about 50 characters with no closing period, such as "Add the Rhodes reference for the writing-style skill", then a blank line and a body wrapped at about 72 characters that explains the problem and why the change solves it. Name the files or behavior that changed, and skip anything the diff already makes obvious.

Commit as the repository owner. Don't add co-author trailers that credit an AI model.

## Using bin/clauditory

Run `bin/clauditory list` to see the available artifacts. `bin/clauditory install chat/skills/writing-style` copies a skill into `~/.claude/skills/` for Claude Code; pass `--dest` to install somewhere else. `bin/clauditory package chat/skills/writing-style` builds a zip you can upload as a skill in the Claude apps. Both commands leave out the `references/` folder, since references document the work for people and would only add noise to an agent's context.
