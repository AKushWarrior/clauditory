# Claude Code output styles (Claude Code documentation)

- **Source:** Anthropic, "Output styles," Claude Code documentation. Living web page with no publication date shown, read in the edition current on November 8, 2026.
- **Type:** official documentation
- **Where to find it:** https://code.claude.com/docs/en/output-styles
- **Last reviewed:** 2026-11-08
- **Used by:** `coding/system/prompt.md`, through the output style that `bin/clauditory build` generates from it

## What the source covers

The page describes output styles, which are instructions Claude Code sends with every request in a session to set Claude's role, tone, and response format. Claude Code ships a default style and four built-in alternatives: Proactive, Concise, Explanatory, and Learning. A custom style is a Markdown file with frontmatter, stored at the user, project, or managed-policy level, and a plugin can ship styles in an `output-styles` folder.

Four frontmatter fields matter. The `name` and `description` fields label the style in the picker. The `keep-coding-instructions` field, off by default, keeps Claude Code's section of software engineering guidance, which covers how to scope changes, write comments, and verify work. Without it, a custom style replaces that section. The page notes that only the full system prompt includes that section, so the field has no effect on the shorter prompt that some models get. The `force-for-plugin` field applies a plugin's style whenever the plugin is enabled, overriding the user's own choice of style; when several enabled plugins set it, the first one loaded wins.

The page separates output styles from neighboring features. A style shapes every response in a session. Project knowledge belongs in CLAUDE.md, instructions for one kind of task in a skill, and anything that must happen without fail in a hook, because a style is only an instruction the model follows. Styles apply to the main conversation and to forks, but not to other subagents, which run their own system prompts. A style's text adds input tokens to every request, though prompt caching reduces that cost after the first request.

## What the artifact applies

The coding prompt ships as a forced plugin style with `keep-coding-instructions` set, so Claude Code keeps its own guidance on how to work and the style changes only how Claude writes. That split is why the coding prompt carries no working loop, research procedure, or subagent rules. The chat prompt's working loop overlapped with Claude Code's engineering instructions and, in places, contradicted them, such as its rule to hand operational goals to Sonnet subagents.

The page's advice on features shaped the rest of the plugin. The prompt covers every response, the skills hold the full writing rules for long documents, and the linter runs as a hook, because a hook runs whether or not the model remembers the instruction. Every style token is paid on every request, so the prompt condenses the two skills to about 10 KB instead of inlining both.

## What the artifact omits

The prompt doesn't set a role, since Claude Code's default role as a software engineer is the one wanted. It doesn't try to reach subagents, which the page says ignore output styles.

## Conflicts with other references

None. The Codex documentation describes a different mechanism, a global instructions file, and the same prompt text serves both.

## Review notes

Read the full page on November 8, 2026. The page doesn't say which models get the shorter system prompt, so on those models the coding instructions might drop out regardless of the field; the plugin can't control that.
