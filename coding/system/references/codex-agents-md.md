# Custom instructions with AGENTS.md (Codex documentation)

- **Source:** OpenAI, "Custom instructions with AGENTS.md," Codex documentation. Living web page with no publication date or Codex version shown, read in the edition current on November 8, 2026. The old address on developers.openai.com redirects to learn.chatgpt.com.
- **Type:** official documentation
- **Where to find it:** https://learn.chatgpt.com/docs/agent-configuration/agents-md
- **Last reviewed:** 2026-11-08
- **Used by:** `coding/system/prompt.md`, through `bin/clauditory prompt coding` and the Codex install steps in `README.md`

## What the source covers

The page explains how Codex finds instruction files. It reads a global file from the Codex home folder, which is `~/.codex` unless `CODEX_HOME` says otherwise. At that level it uses `AGENTS.override.md` when that file exists and isn't empty, and `AGENTS.md` otherwise; only one global file is used. Codex then walks from the project root down to the working directory and takes at most one file from each folder, checking the override name first, then `AGENTS.md`, then any fallback names set in its configuration.

Codex joins the files it finds into one block, root first, separated by blank lines, and skips empty files. Files closer to the working directory come later in that block, so their guidance takes precedence. Discovery happens once per run or per interactive session. The combined text is capped by the `project_doc_max_bytes` setting, 32 KiB by default; once the total reaches the cap, Codex adds no more files. The page suggests raising the cap or splitting instructions across nested folders.

The page says nothing about plugins setting instructions.

## What the artifact applies

The README tells Codex users to append the coding prompt to `~/.codex/AGENTS.md`, the global file, so it applies in every project. Appending keeps whatever the user already has there. The 32 KiB cap covers the global file and every project file together, so the prompt stays at about 10 KB, which leaves room for a repository's own instructions. Because project files come later and take precedence, a repository's own writing rules override the prompt where they conflict, which is the behavior wanted.

## What the artifact omits

The install steps don't use `AGENTS.override.md`, because an override file would hide the user's existing global instructions instead of adding to them.

## Conflicts with other references

None. Claude Code receives the same prompt as an output style; the two mechanisms differ in delivery, not content.

## Review notes

Read the full page on November 8, 2026. The page's silence on plugins agrees with the earlier finding, recorded in the README, that Codex plugins can't set always-on instructions; that finding came from testing an install with the Codex CLI, not from this page.
