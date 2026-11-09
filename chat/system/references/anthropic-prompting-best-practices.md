# Prompting best practices (Claude Platform documentation)

- **Source:** Anthropic, "Prompting best practices," Claude Platform documentation, prompt engineering section. Living web page covering Claude models from Claude Haiku 4.5 through Claude Fable 5.1 and Claude Opus 5.5, read in the edition current on November 5, 2026.
- **Type:** official documentation
- **Where to find it:** https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/claude-4-best-practices
- **Last reviewed:** 2026-11-05
- **Used by:** `chat/system/prompt.md`

## What the source covers

### Purpose and form

The page is Anthropic's main reference for developers who write prompts for current Claude models. It has three parts. A short model-specific part points to separate guides for each model and lists what differs in each one. A long middle part gives techniques for all current models: general principles, output and formatting, tool use, thinking, agentic systems, and a few capability tips. A last part covers migration from earlier model generations. Most advice comes as a short claim followed by a sample prompt block that a developer can paste in.

The page is practitioner guidance from the model's maker. It reports almost no method. Two numbers appear in the whole page: parallel tool calling can be pushed to about 100 percent with a prompt, and putting the query after long documents can improve quality by up to 30 percent on complex multi-document inputs. Neither includes a task set, sample size, baseline, or model name. The page also says that a technique naming a specific model was measured on that model and should be rechecked on others. That sentence is the clearest statement of how to read the whole page: much of it is tuning advice tied to model versions, and the advice changes from release to release, sometimes reversing.

The rest of this section covers the parts that bear on a chat system prompt that runs a working loop. A separate reference for the writing-style skill (`chat/skills/writing-style/references/anthropic-prompting-docs.md`) covers the same page's advice on examples, roles, tags, output formatting, and prose, so those topics appear here only where they touch the system prompt. The model-specific guides for Claude Opus 5 and Claude Opus 5.5 have their own reference files in this folder (`anthropic-opus-5-prompting.md` and `anthropic-opus-5-5-prompting.md`), and this file draws on them only where the general page points to them.

### Explicit instructions

The page's first principle is that Claude follows clear, explicit instructions well and does not reliably infer ambitions that the prompt leaves unstated. It frames the model as a capable new employee with no knowledge of local norms, and offers a test: if a colleague with little context would be confused by the prompt, the model will be too. It asks for explicit output formats and constraints, and for numbered or bulleted steps when order or completeness matters. The migration section repeats the point: newer models take instructions literally, so a developer who wants extra effort or particular features must ask for them.

### Positive framing and calm wording

In the formatting section the page advises telling the model what to do in place of what to avoid. The advice is stated only for formatting, though the page's own sample blocks often mix positive and negative wording. A related theme runs through the tool-use and migration sections: newer models respond strongly to the system prompt, so capitalized, absolute phrasing written to fix undertriggering in older models now causes overtriggering. The page recommends plain conditional wording that says when to do something.

### Giving reasons

The page claims that a rule with its motivation attached works better than a bare rule, because the model can generalize from the reason. Its example replaces an all-capitals ban on ellipses with a sentence explaining that the output goes to a text-to-speech engine. The claim is plausible and matches how the page writes its own samples, but no measurement supports it.

### Parallel tool calling

The page says current models already run independent tool calls in parallel and that a prompt can raise the rate to about 100 percent. The sample block asks the model to send all independent calls together, to call dependent tools in sequence, and never to guess a missing parameter or use a placeholder. For one model in long agent loops, the page says the instruction works better as a system message repeated after each round of tool results. The 100 percent figure has no stated test behind it.

### State and progress tracking

The agentic section says current models track state well over long tasks by making steady progress on a few things at a time. Most of its advice concerns work that spans several context windows. It says to use the first window to set up tests and scripts and to keep test status in a structured file. It also says to keep free-form progress notes in text, use git as a log and checkpoint, and tell the model to track its progress. It offers a sample telling the model that context will be compacted, that it should save state before the refresh, and that it should not stop early because of its token budget.

### Early stopping and overcorrection

The page approaches early stopping from two sides. One side pushes persistence. The compaction sample tells the model to keep working through budget limits, and a second sample encourages it to use its whole output budget on a long task. The model table lists early stopping in long agent prompts as a known issue for Claude Haiku 5.5 and finishing long tasks as a topic for Claude Fable 5.1. The other side warns against overcorrection. The migration section says anti-laziness prompts written for older models should be reduced, because Claude 4.6 models and later take more initiative and overtrigger on them. The overeagerness section says some models overengineer, adding files, abstractions, and flexibility nobody asked for, and offers a sample that limits changes to what was requested or clearly needed. The general page gives no rule for telling these cases apart beyond testing on the model in use. The Opus 5.5 guide adds the most specific advice on the persistence side, for unattended runs only.

### Verification

The page recommends asking the model to check its answer against stated criteria before finishing, and says this catches errors reliably for coding and math. It then names an exception: Claude Opus 5 verifies its work well without being told, and verification instructions carried over from older prompts make it over-verify, adding tokens and latency. On that model the page says to remove such instructions, not rewrite them. For long autonomous tasks it recommends giving the model tools to check its own work, such as browser or computer-use tools for interfaces. For research it recommends clear success criteria and cross-checking information across several sources. Its chaining section describes a draft, review, and refine cycle run as separate calls, which it presents as a pipeline a developer builds outside the model's turn. The model table lists verification on coding tasks as a topic in the Sonnet 5.5 and Haiku 5.5 guides, which suggests the right amount of verification prompting differs by model.

### Action, suggestion and risky actions

The tool-use section says current models follow the verb in a request closely. A request to suggest changes tends to get suggestions, while a direct instruction gets the change. The page offers two opposite samples, one setting a default of acting and one setting a default of restraint. A separate sample on autonomy and safety, introduced for Claude Opus 4.6, tells the model to weigh how reversible and how visible an action is. It says to act freely on local, reversible steps such as editing files or running tests, and to ask before actions that are destructive, hard to reverse, or visible to others. Its examples include deleting files or branches, force-pushing, pushing code, commenting on issues, and sending messages. It also says not to bypass safety checks or discard unfamiliar files to get around an obstacle.

### Subagents

The page says current models delegate on their own when subagent tools are well described, and warns that some overuse delegation. It names Claude Opus 4.6, which may spawn a subagent where a direct search would do, and Claude Opus 5, which delegates more readily than earlier models, and it points to the Opus 5 guide for a damping sample. Its own sample limits subagents to parallel, isolated, or independent work, and keeps simple tasks, sequential steps, single-file edits, and work that needs shared context with one agent.

### Time and speed

The general page says nothing about time budgets or elapsed time. Its only speed advice is parallel tool calling and lower effort. The time-budget guidance for multi-agent harnesses appears only in the Opus 5.5 guide.

### Thinking

The page says a general instruction to think thoroughly often beats a hand-written step-by-step plan, because the model's own reasoning frequently exceeds what a person would prescribe. It recommends committing to an approach and revisiting it only on new contradicting evidence, to curb overthinking. It also notes that a model which thinks more often than wanted can be steered with a sentence saying thinking should be used only when it improves the answer.

### Strength of the evidence

The page is authoritative about Claude's behavior in the sense that its authors build and test the models. It is weak as evidence in the usual sense. It cites no experiments, gives two unexplained numbers, and attaches several claims to single model versions. Its advice is also internally conditional: verification helps except on a model where it hurts, persistence prompts help except where they overtrigger, thoroughness helps except where it inflates thinking. For a system prompt that will run on whatever Claude model a product serves, the general principles (be explicit, give reasons, phrase calmly, parallelize independent calls, ask before risky actions, limit delegation to independent work) should transfer well. The version-tuned notes should be treated as hypotheses to recheck on the model in use.

One more limit applies. The page addresses developers who write the whole system prompt through the API. The prompt in this repository is written in the first person to "me," which suggests it sits as user-level instructions beneath a product's own system prompt (this is an inference from its wording). In that setting it competes with other instructions, and the page says nothing about how layered instructions interact.

## What the prompt applies

The prompt was revised on November 5, 2026. The revision changed its verification and delegation rules, so this section describes the revised text.

**Explicit goals and a stopping rule (Working loop, Before starting).** The prompt asks the model to list the goals a complete answer must meet, judged by what the user is trying to do, and to treat that list as the definition of done. The rule follows the page's advice to state success criteria explicitly and not to rely on the model to infer them. It goes further than the page by asking the model to write its own criteria at the start of each task, which the page neither tests nor describes.

**Positive, calm phrasing (whole prompt).** Nearly every bullet says what to do, and none uses capitals or threats. The phrasing follows the page's advice on positive framing and on softening forceful language. A few negatives remain (don't apologize for disagreeing, never skip work the answer needs, don't use subagents to recheck your own work). They are short and specific, so they are unlikely to overtrigger in the way the page describes.

**Reasons attached to some rules.** Several rules carry their reason: manufactured disagreement counts as a failure on par with flattery, pushback and the model's own earlier statements carry no evidence, and polish and authority cues are weak signals. Those reasons follow the page's claim that the model generalizes from a stated motive. The rules on delegation, completion, and confirmation are bare instructions, so the prompt applies this principle only in part.

**Parallel tool calls (While working, first two bullets).** The prompt asks for independent tool calls and searches to go out together and for only dependent steps to wait. The request follows the core of the page's parallel-calling sample. The added line that speed comes from parallelism and from stopping when the goals are met, and never from skipping needed work, is the prompt's own. It shares an aim with the Opus 5.5 guide's one-sentence time signal but goes beyond anything on the general page.

**Early stopping (Before starting; operational goals).** The stopping rule says work is done only when every goal is met or shown to be unmeetable, and the operational bullets ask the model to deliver the full scope and say when it can't. Both address the early-stopping problem with a completion criterion in place of the page's persistence wording, which follows the page's warning against blanket anti-laziness prompts. The scaling line at the top of the loop (a quick question gets a quick answer) also follows that warning. The prompt also follows the page's overeagerness advice for operational work: it asks for delivery at the scope the user intended, with a sentence of pushback where the request looks mistaken and no quiet narrowing or widening. Judging goals by what the user is trying to accomplish could still widen an informational request, which no rule limits.

**Verification (operational and informational goals).** The revision changed this rule most. The earlier text asked the model to check every result before marking a goal done. The revised text sets a completion bar instead. A goal counts as done only when something outside the model's own judgment shows it, such as passing tests, the saved file, or the rendered page. The model should use the evidence the work already produces instead of adding a separate review pass. It also tells the model not to use subagents to recheck its own work. The page's general advice favors a self-check against criteria, and its Opus 5 exception says to remove verification instructions on that model. The revised prompt follows the exception in dropping an explicit verification step and in banning subagent rechecks. It departs from the exception in keeping a completion bar at all, though the bar points to evidence the work already produces and rules out a separate review pass, so it should rarely add checks. It follows the page's research advice in asking the model to confirm load-bearing claims with an independent second source and to open a cited source before relying on it, including one a subagent supplied. The page's over-verification warning concerns self-checks of the model's own work, so it doesn't reach these source checks.

**Confirmation for risky actions (operational goals).** The prompt asks the model to confirm before an action that is hard to undo, spends the user's money, or will be seen by other people, unless the user has already authorized it. The rule follows the page's autonomy-and-safety sample on hard-to-reverse actions, and the Opus 5.5 guide's instruction to keep a confirmation step even when pushing an agent to keep working. Its third trigger matches the page's category of actions visible to others, such as posting or sending messages. It departs in two ways. It adds spending money, which the page does not list, and it leaves out the page's rule against bypassing safety checks.

**Subagent limits (While working, third and fourth bullets).** The revised prompt tells the model to keep the largest goal it can work on now for itself and to hand off a goal only when it is sizeable and independent of the rest. It also tells the model to finish work it can do in a handful of tool calls itself, and to keep work that needs shared context or linked decisions with one agent. The limits follow the page's damping sample and the Opus 5 guide's sample almost point for point. The earlier version pushed delegation more, which the page would have called unnecessary on recent models. Two operational bullets still direct delegation: sizeable informational goals go to the research-subagent skill and sizeable operational goals go to general-purpose subagents on Sonnet. The page says nothing about which model a subagent should use, so the routing is the repository owner's choice.

**Time budgets.** The prompt has no time budget, and the general page offers none. The Opus 5.5 guide covers the topic; its reference file discusses how the prompt relates to it.

**Reporting (Composing answers).** The request to say what was verified and how, what couldn't be checked, and which goals remain open serves the same purpose as the page's sample asking for a summary after tool use. The prompt makes the report about verification status, which the page's sample does not specify.

## What the prompt omits

**Persistent state files.** The prompt asks the model to keep its goal list in a to-do tool or a file on longer tasks and to update it as goals close. That covers part of the page's state advice and protects the list from compaction. It never asks for free-form progress notes, a structured test-status file, or git checkpoints. The page's fuller state advice targets work across several context windows, and a chat agent usually finishes inside one. The closing report on open goals is the hand-off.

**Compaction and context-budget wording.** The prompt does not tell the model that context will be compacted or that it should not wind down near the limit. Whether that applies depends on the product running the prompt, which the prompt can't know.

**A missing-parameter rule.** The page's parallel-calling sample tells the model never to guess a missing parameter. The prompt leaves this out. Its rules on treating assertions as hypotheses and on external evidence of completion discourage guessing in general but don't address tool arguments.

**Scope limits for informational work.** The page's overeagerness sample and the Opus 5 guide's scope sample both limit work to what was asked. The prompt has such a rule for operational goals only, as noted above.

**Turn-scoped reminders.** The page says one model follows the parallel-call instruction better in long loops when it is repeated after each round of tool results. A static system prompt can't do this.

**Formatting, examples, roles, and tags.** The prompt contains none of these: a role line, worked examples, tags, or a formatting block. It hands register and format to the writing-style skill, with the technical-writing skill added for documentation, code comments, and commits; the writing-style skill's own reference covers this part of the page.

**Model-specific tuning.** The prompt ignores effort settings, thinking configuration, prefill migration, and other API-level controls, since none can be set from inside a prompt.

**Research scaffolding.** The page's research sample asks for competing hypotheses, confidence tracking, and a notes file. The prompt keeps calibration and cross-checking but leaves out the hypothesis tree and notes file, because chat tasks rarely need persisted research state.

## Conflicts with other references

The model-specific references `anthropic-opus-5-prompting.md` and `anthropic-opus-5-5-prompting.md` and the subagent reference `claude-code-subagents.md` existed in this folder when this section was written and were read. Reference files for the harness design post (`anthropic-harness-design.md`), the progress mirage paper (`progress-mirage.md`), FrontierChallenge (`frontierchallenge.md`), the coding-agent teams paper (`coding-agent-teams.md`), the dual-control paper (`dual-control-multi-agent.md`), and the collaboration tax paper (`collaboration-tax.md`) came later, so the points below come from the local copies of those sources. Sources first published before March 5, 2026, that earlier versions of this file cited (older engineering posts, essays, and papers on multi-agent failure, planning prompts, and sycophancy) are no longer used here, under the repository's recency rule.

**Self-checks against external grounding.** The page recommends self-checks in general and the Opus 5 guide says to drop them on that model. The progress mirage paper (Park and Choi, arXiv 2607.25152) found an agent's own claims of progress uninformative. It claimed improvement in all 54 cycles while 56 percent showed no gain or a loss, and even a strong judge reading the full artifact accepted many regressions. Only a signal outside the transcript closed the gap. FrontierChallenge (arXiv 2608.24979) found that about 75.5 percent of failed runs in the Claude Code scaffold, across several models including Opus 5, ended with language claiming completion. Anthropic's harness design post (March 2026) found that agents grading their own work praise it, and that a separate evaluator with tools to exercise the application helped on tasks at the edge of the model's ability. The prompt resolves the tension by replacing a self-check step with a bar that requires outside evidence. That keeps the part of verification the papers support and drops the part the Opus 5 guide calls wasteful.

**Delegation.** The page says current models delegate on their own and some overuse it. The dual-control paper (arXiv 2610.27655) found a multi-agent system built from the same model worse than a single agent in 12 of 13 valid cells on a benchmark with a live user. The collaboration tax paper (arXiv 2608.22152) found two agents with split information do worse than one agent with all of it. The coding-agent teams paper (arXiv 2610.17471) found teams beat solo agents on difficult coding tasks at three to seven times the cost, and that extra verifier workers stop helping beyond about four. All three agree with the page's damping sample and with the prompt's revised limits. The Opus 5.5 guide reports that agent teams with a time budget finished research tasks sooner at comparable quality, which pulls toward more parallel delegation for research. The prompt routes research through a separate skill and doesn't take a position on team size.

**Prescribed reasoning steps.** The page says a general instruction to think often beats a hand-written reasoning plan. The prompt prescribes specific steps: treat each assertion as a hypothesis, run a counterfactual check before giving ground, and list goals before starting. The claim-checking steps are targeted fixes for sycophancy, a failure the page does not discuss, and the SWAY study (Bhalla and Gligorić, arXiv 2604.02423, see `bhalla-sway.md`) supports the counterfactual check specifically. The goal-listing step on simple tasks has no recent support beyond the page's own advice on success criteria, and the prompt's scaling line is its only limit.

**Early stopping.** The page warns against anti-laziness prompts, while the Opus 5.5 guide gives a strong keep-going addition for unattended agents and says to leave it out where a person is present. The prompt follows both by using a completion criterion and no keep-going language.

**State tracking.** The page and the harness design post agree on structured artifacts for handing off context in long runs. The prompt adopts only the lightest form, a goal list kept in a to-do tool or file on longer tasks, for the reasons in the omissions section. The difference is one of scope.

**Prompt format and output format.** The page says the formatting of a prompt tends to carry into the output. The prompt is written in headings and bullets while asking for writing-style output, which defaults to paragraphs. The writing-style skill's reference notes the same tension for that skill.

## Review notes

I read the full text of the page from the local copy of the claude-4-best-practices URL, which the site served under the title "Prompting best practices." The writing-style reference cites the same content at a different URL, so the two addresses probably lead to one page, but I did not confirm the redirect. For this revision I also read the full Opus 5 and Opus 5.5 guides, which answered part of a doubt in the earlier version. The Opus 5.5 guide says the Opus 5 patterns remain a reasonable starting point but does not restate the over-verification advice, so whether over-verification persists on Opus 5.5 is still not stated anywhere I read. I did not read the Fable, Sonnet, Haiku, or Opus 4.8 guides, so model-specific claims about early stopping and verification on those models rest on this page's summary table. The 100 percent parallel-calling figure and the 30 percent long-context figure are Anthropic's own and the page gives no test behind them. For the conflicts section I read abstracts and targeted passages of the cited papers and searched the harness design post, not their full texts. The reading of the prompt as user-level instructions layered under a product system prompt is my inference from its wording. I did not test any claim on a model.
