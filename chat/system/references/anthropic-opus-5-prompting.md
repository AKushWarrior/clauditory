# Prompting Claude Opus 5 (Claude Platform documentation)

- **Source:** Anthropic, "Prompting Claude Opus 5," Claude Platform documentation, prompt engineering section. Living web page with no publication or revision date shown, read in the edition current on 2026-11-05.
- **Type:** official documentation
- **Where to find it:** https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/prompting-claude-opus-5
- **Last reviewed:** 2026-11-05
- **Used by:** chat/system/prompt.md

## What the source covers

### Purpose and form

The page is Anthropic's model-specific guide for Claude Opus 5. It sits under the general prompting best-practices page and lists only what differs from earlier Opus models and what to change in a prompt as a result. It says existing Opus 4.8 prompts work well on Opus 5 out of the box and that the patterns it lists are the ones that most often need tuning. Each section names a behavior, says which way the model leans and gives a short sample instruction to paste in.

The page has eight parts: capability improvements, response length, progress updates during agentic work, the length of written files, task scope and over-verification, subagent spawning, self-correction, and two sections on running with thinking disabled and on asking for reasoning in the response.

### Capability claims

The capability section says Opus 5 is strongest on hard, multifile coding work, finishes whole tasks without leaving stubs, and does best when given the full specification up front and left to run. It says the model reviews code with high precision and recall at a range of effort levels, and it warns that a review prompt asking for only high-severity findings may be followed literally, so the model reports less. Its advice is to ask for every finding and filter in a separate pass. Other claims cover strong results at low and medium effort, chart and diagram reading, a 1 million token context window with steady instruction following across it, spreadsheets and slide decks, and coordination of subagent teams with effective writer-verifier patterns and few cases of agents overwriting each other.

### Length and narration

The page says Opus 5's user-facing replies run longer than earlier Opus models' replies, and that lowering effort reduces thinking without reliably shortening the visible answer. Its fix is an explicit conciseness instruction, repeated as a short reminder near the end of a long system prompt. During agentic work the model narrates readily, announcing what it is about to do. The sample instruction asks for one sentence before the first tool call, updates only when something important turns up or the direction changes, and a final message that leads with the outcome. The page adds that positive examples of the wanted style work better than lists of things to avoid. A separate note says files the model writes to disk also run long, and gives a one-sentence length calibration.

### Over-verification and scope

This is the section that matters most for the prompt. The page says Opus 5 checks its own work without being told to. It names two kinds of instruction to remove: a required final verification step for non-trivial tasks, and an instruction to use a subagent to verify. It says such instructions cause over-verification and that removing them saves tokens with no loss in quality. Harness scaffolding that adds its own verification steps falls under the same advice.

The same section says the model can widen a task, adding steps nobody asked for or substituting its own view of what the task should be. The sample scope instruction asks the model to deliver what was asked at the intended scope, make routine judgment calls itself, check in only when different readings of the request would lead to materially different work, flag a mistaken request in a sentence and then carry on as asked, finish the whole task and stop short of actions clearly beyond the request.

### Subagent control

The page says Opus 5 delegates more readily than earlier models, and that delegation pays for genuinely independent, sizeable tracks of work but multiplies cost and time on small ones. It recommends either explicit guidance on when to delegate or a hard cap on spawning. The sample instruction limits delegation to large, independent, parallel work such as a wide multifile investigation, rules out delegating work the model can finish in a handful of tool calls, rules out subagents for verifying or double-checking its own work, and asks for one subagent in place of several where one will do. For Claude Code and the Agent SDK it names two environment variables that cap spawn depth and concurrency, plus an SDK spending cap. It also says Claude Code adds its own delegation instruction on Opus 5 only under its built-in system prompt preset, so a custom system prompt has to supply one.

### Self-correction

The page says Opus 5 catches and fixes its own mistakes without prompting, and that instructions to double-check or re-verify an answer compound with that behavior and add cost without improving results. It also says the model narrates corrections of its earlier statements more than earlier models did. The sample instruction limits correction notes to errors that would change the user's code, conclusions or decisions, and asks the model to fix trivial slips silently.

### Thinking disabled and reasoning in the response

Opus 5 thinks by default and accepts disabled thinking only at high effort or below. With thinking off, the model sometimes writes a tool call as plain text so it never runs, or leaks internal tags into the reply. The page's first remedy is to keep thinking on and lower effort, which it says beats disabled thinking at similar cost on most tasks. For integrations that must keep thinking off, it gives a combined instruction and warns that a rule telling the model not to think makes tag leakage worse. A last note says prompts asking the model to write out its reasoning verbatim or in a fixed format may be refused under a reasoning-extraction category, and suggests asking for a short explanation or a summary of actions in its place.

### Strength of the evidence

The page contains no numbers. It reports no task set, sample size, baseline or effect size for any claim, including the central one that removing verification instructions saves tokens with no loss in quality. Its authority comes from its authors, who train and test the model, and its claims should be read as Anthropic's internal observations stated as advice. They are probably reliable about the direction of a behavior (Opus 5 does lean toward longer replies, more delegation and more self-checking) and less reliable about the size of any effect on a given workload.

The advice also has a narrow target. It describes one model, addresses developers writing a whole system prompt for an API integration and assumes coding or agentic work. The prompt in this repository runs in a chat product, probably on whatever model the product serves. The successor guide for Opus 5.5 says the Opus 5 patterns remain a reasonable starting point on that model, which extends the page's reach by one generation but no further. The over-verification claim in particular concerns self-checks of the model's own work. It says nothing about checking external sources, which is a different activity with a different failure mode.

## What the prompt applies

**Subagent limits (Working loop, While working, third bullet).** The prompt tells the model to hand a goal to a subagent only when it is sizeable and independent of the rest, to finish work it can do in a handful of tool calls itself, and not to use subagents to recheck its own work. These three clauses follow the page's sample instruction nearly point for point, in shorter form. The page's warning that a custom system prompt has to supply its own delegation rule is a direct reason for the bullet to exist, since the prompt is written as instructions layered on a product, outside Claude Code's preset. The bullet's opening clause, which keeps the largest goal the model can work on now for itself, is the prompt's own and has no counterpart on the page. The page also asks for one subagent in place of several and for low spawn counts, which the prompt does not say.

**No separate verification step (Working loop, operational goals).** The revised prompt has no instruction to add a final verification pass and no instruction to verify through a subagent, which are the two forms the page names. On this point it follows the page. It keeps a completion bar instead: a goal counts as done only when evidence outside the model's own judgment shows it, such as passing tests, a file that contains the change or a page that renders, and the model's sense that it has finished does not count. This is a departure in substance, though not in the page's literal terms. A model could read the bar as a request to run extra checks, which is the cost the page warns about. The prompt's wording limits that risk, because each example is evidence the work usually produces anyway: code that was run, a file that was written, a page that was loaded. The page offers no evidence that would settle whether the bar costs tokens on Opus 5, and the external sources listed under conflicts give a reason to keep it.

**Self-checks before giving ground (Handling claims).** The page says instructions to double-check an answer add cost on Opus 5. The prompt asks for one targeted check: before agreeing with the user or giving ground, ask whether the same answer would follow if the user had asserted the opposite. This is a departure from the page's general advice, made for a specific failure (sycophancy under pushback) that the page does not address. The check fires only at the moment of a concession, so its cost is bounded.

**Full scope (Working loop, operational goals).** The instruction to cover the full scope, finish every part of the task and say when that isn't possible matches the page's request to finish the whole task. The prompt does not include the page's other half, which asks the model to stay inside the requested scope and to stop short of actions beyond it. Its goal list, judged by what the user is trying to accomplish, pulls the other way and could widen a narrow request, which is the tendency the page says Opus 5 already has. The confirmation rule for hard-to-undo actions limits the damage that widening can do, but it doesn't stop added work.

**Leading with the outcome (Composing answers).** The rule to lead with the bottom line matches the closing clause of the page's narration sample, which asks the final message to answer what happened or what was found before giving detail. The writing-style skill the prompt points to makes the same demand.

**Length (Composing answers).** The prompt hands length to the question and to the writing-style skill, which asks for length that follows content and for stopping when the content is done. This is a weaker form of the page's explicit conciseness instruction. The page says Opus 5 needs an explicit length rule because its default replies run long. The prompt's rule is explicit but conditional, and it sits in the last section of the prompt, which matches the page's advice to put a length reminder near the end.

**Explanations in place of reasoning dumps (Composing answers).** The prompt asks for the reasoning and evidence a reader needs to trust the answer, and for a report of what was verified and how. Both are forms the page permits: a short explanation of the answer and a summary of actions taken. Nothing in the prompt asks the model to reproduce its thinking verbatim, so it avoids the refusal category the page describes.

## What the prompt omits

**Mid-task progress updates.** The prompt sets no cadence for updates during agentic work. The page treats narration as a tuning point for agentic products. A chat prompt that runs on several products can't know how each one renders updates, so the omission is reasonable, though the page's one-sentence-before-starting pattern would cost little.

**Correction narration.** The prompt does not tell the model to skip notes about trivial self-corrections. Its rule to name the evidence that changed its mind applies to updates on the user's claims and leaves the model's own slips uncovered. Given the prompt's stress on calibration and on saying what changed, a rule that suppresses some corrections would sit awkwardly with the rest, which may explain its absence. The cost is some extra text on Opus 5.

**Deterministic caps.** The environment variables and spending cap are harness settings that a prompt can't set.

**Thinking-disabled mitigations, effort settings and reasoning-format rules.** These are API configuration. The prompt never tells the model not to think, which is the one prompt-level item the page warns against.

**Code review literalness and written deliverable length.** The prompt has no code review rules. Length of written files falls to the writing-style skill.

**Capability claims.** The vision, long-context and document claims describe the model and need no instruction.

## Conflicts with other references

Of the files cited below, `anthropic-prompting-best-practices.md`, `anthropic-opus-5-5-prompting.md` and `claude-code-subagents.md` exist in the folder and were read. The reference files for the harness design post, the progress mirage paper, FrontierChallenge, the coding-agent teams paper, the dual-control paper and the collaboration tax paper did not exist when this file was written, so the points below come from the local copies of those sources.

**Self-verification against external grounding.** The page says Opus 5 verifies its own work well and should not be told to. Anthropic's harness design post (Rajasekaran, March 2026) reports that agents grading their own work praise it even when it is mediocre, and that a separate, skeptical evaluator with tools to exercise the running application was a strong lever. That post tested Opus 4.5 and 4.6, before Opus 5, and found the evaluator worth its cost only for tasks at the edge of what the model handled alone. The progress mirage paper (Park and Choi, arXiv 2607.25152) found that an unnamed frontier agent claimed improvement in all 54 cycles of a loop while 56 percent of those cycles showed no gain or a loss, and that even a strong judge reading the full artifact accepted many regressions. Only a signal outside the agent's transcript closed the gap. FrontierChallenge (arXiv 2608.24979) found that 641 of 849 failed runs in the Claude Code scaffold, about 75.5 percent across several models including Opus 5, ended with language claiming completion. These results conflict with the page's confidence in self-verification, at least for long tasks and for completion claims. The prompt resolves the conflict by dropping the extra verification step and the verifying subagent, as the page asks, and keeping a bar that requires evidence from outside the model's judgment, as the papers support. The progress mirage result also backs the page's ban on subagent rechecks for a different reason: a second model reading the same transcript is still an in-band judge.

**Value of extra verifiers.** The coding-agent teams paper (arXiv 2610.17471) found that verifier workers stop adding value beyond about four, and that a lead that delegates only planning and verifies its own code matches a lead that delegates everything at half the cost. This agrees with the page's advice not to verify through subagents.

**Delegation in general.** The dual-control paper (arXiv 2610.27655) found a multi-agent system built from the same model scored worse than a single agent in 12 of 13 valid cells on a conversational benchmark with a live user, by a weighted mean of 28 points. The collaboration tax paper (arXiv 2608.22152) found that two agents with split information do worse than one agent with all of it, across 11 models. Both support the page's damping of delegation and the prompt's rule that work needing shared context stays with one agent. The page itself credits Opus 5 with strong team coordination, which these papers did not test on that model.

**Early stopping on the successor model.** The Opus 5.5 guide says Opus 5 patterns still apply but adds advice pushing the model to keep working in unattended runs. That pressure runs opposite to the Opus 5 page's worry about widened scope. The prompt's goal list as a stopping rule serves both: it defines when to stop, and it defines what remains to do.

**General self-check advice.** The best-practices page recommends asking the model to verify its answer against stated criteria, with Opus 5 named as the exception. The two Anthropic pages agree once the exception is read as part of the general advice.

## Review notes

I read the full text of the page from a local copy captured on 2026-11-05. The page shows no date, so I could not tell when it was first published or last changed; under the repository's rule for living documentation it counts from the version read. I did not open the pages it links to, including the Opus 5 model page, the recommended effort levels and the Agent SDK section on capping subagents. I confirmed the two Claude Code environment variables in the local copy of the Claude Code subagents page but did not check the SDK spending option. For the conflicts section I read the abstracts and targeted passages of the progress mirage, FrontierChallenge, coding-agent teams, dual-control and collaboration tax papers and searched the harness design post for its passages on evaluation, not the full texts. The claim that the prompt runs beneath a product's own system prompt is my inference from its first-person wording. I did not test any claim on a model, so I can't say whether the prompt's completion bar causes over-verification on Opus 5.
