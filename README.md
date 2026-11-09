# Clauditory

Clauditory is a set of agent-facing prompts and skills, each built from published research and established style guides. Each artifact has a `references` folder that records which sources shaped it, what it took from each, and where the sources disagree.

The plugin contains:

- **A chat system prompt** (`chat/system/prompt.md`) that sets how an agent handles claims, plans and checks its work, runs independent work in parallel, and writes up answers. It is for the Claude apps.
- **A coding writing prompt** (`coding/system/prompt.md`), a condensed blend of the two writing skills for Claude Code and Codex. It covers replies, code comments, commits, pull requests, and docs, and leaves the agent's own coding instructions in charge of how it works.
- **The writing-style skill** (`chat/skills/writing-style`), which sets a plain newspaper register for everything an agent writes. It draws on recent research into how readers and AI-text detectors tell model prose from human prose.
- **The technical-writing skill** (`chat/skills/technical-writing`), which adds rules for documentation, instructions, code comments, error messages, commit messages, and pull requests.
- **The spawn-research-subagent skill** (`chat/skills/spawn-research-subagent`), which splits research into parallel subagents and checks what they bring back.

The writing-style skill also includes **clauditory-lint** (`lint/`), a prose linter with Vale compiled in. It checks text against the rules a program can test: the skills' revision passes, sentence rhythm, a catalog of machine-writing tells, and the word lists of the major style guides. It also fixes the mechanical problems. In a source file it checks only the comments and docstrings. The skill tells the agent to run it before sending or saving substantial text, and the plugin also runs it after Claude writes a prose or source file.

## Install

### Claude Code

Run these two commands in Claude Code:

```
/plugin marketplace add AKushWarrior/clauditory
/plugin install clauditory@clauditory
```

The plugin adds the three skills and applies the coding writing prompt as an output style while the plugin is enabled. The style keeps Claude Code's built-in coding instructions. To turn it off, disable the plugin.

The plugin also adds a hook that lints each Markdown or text file Claude writes, and the comments in each source file, and passes problems back to Claude. The linter comes inside the plugin, so there's nothing else to install. See `lint/README.md` to run it by hand.

### Claude apps (claude.ai and desktop)

1. Open Customize > Plugins, choose Add > Add marketplace, and enter `AKushWarrior/clauditory`.
2. Install Clauditory from that marketplace.
3. Paste the contents of `chat/system/prompt.md` into your profile preferences in Settings, or into a project's instructions.

Chat loads the plugin's skills but not its output style, so the system prompt has to go in by hand.

### Codex

```
codex plugin marketplace add AKushWarrior/clauditory
codex plugin add clauditory@clauditory
```

This installs the three skills. Codex plugins can't set always-on instructions, so add the coding writing prompt to your global instructions file:

```
git clone https://github.com/AKushWarrior/clauditory
clauditory/bin/clauditory prompt coding >> ~/.codex/AGENTS.md
```

## Repository layout

Source artifacts live under `chat/` and `coding/`. The installable plugin in `plugins/clauditory` is generated from them by `bin/clauditory build`, which copies each skill without its `references` folder or changelog and turns the coding writing prompt into an output style. It also builds the linter for every platform into the writing-style skill and copies the linter's hook from `lint/hooks`. Run `bin/clauditory build --check` to confirm the plugin matches its sources. `AGENTS.md` describes the layout, the reference-file format, the versioning rule, and the house rules for writing and committing.

## License

Mozilla Public License 2.0. See `LICENSE`.
