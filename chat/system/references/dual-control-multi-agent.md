# Does Multi-Agent Collaboration Pay When Agents Serve Users? Testing Agent Scaling in Dual-Control Settings

- **Source:** Ethan Yang and Xiangcheng Shi (University of Pennsylvania), "Does Multi-Agent Collaboration Pay When Agents Serve Users? Testing Agent Scaling in Dual-Control Settings," arXiv preprint arXiv:2610.27655, first version 20 October 2026. Pre-registered, with code, prompts and trajectories released.
- **Type:** Preprint (not peer-reviewed), academic
- **Where to find it:** https://arxiv.org/abs/2610.27655
- **Last reviewed:** 2026-11-05
- **Used by:** `chat/system/prompt.md`

## What the source covers

The paper tests whether a published rule for choosing between one agent and several carries over to agents that talk with a user. The rule, from Kim and colleagues (2026), says coordination helps while the single agent is weak and hurts once it is strong, with the crossover near 45 percent single-agent accuracy. That rule was fitted on autonomous benchmarks with no user in the loop. This paper reruns the comparison on a customer-service benchmark and finds no case where the multi-agent system helped: it lost in 12 of 13 valid cells and tied in the 13th, including the two cells where the rule predicts a gain.

### Method

The benchmark is tau2-bench, in which an agent serves a simulated customer under a written policy and is scored on the final state of a database or device. It has three domains: airline and retail, where only the agent has tools, and telecom, where the customer also acts on the device. Each cell pairs one model with one domain and runs 50 tasks three times for each system, 150 episodes per system.

The single-agent system is the benchmark's stock agent. The multi-agent system is a centralized design built from the same model: an orchestrator that alone talks to the customer and plans, a database agent that alone calls tools, and a policy advisor that answers questions with a fresh context. Each orchestrator turn costs at least two calls where the single agent makes one. The authors chose the centralized design because the conversation has one point of contact with the user, which the other common designs would have to break.

Compute is matched on completion tokens. The single agent runs uncapped first; its 75th-percentile spend sets a per-episode budget, and the multi-agent system gets four times that from one shared ledger for all its agents. An episode that runs out ends with a budget-exhausted label, which is scored as a failure and kept separate from plain inaction. A model plays the customer (gpt-5.2 at low effort); the benchmark's default simulator is used in one robustness check.

Six models from three families ran, two tiers each: gpt-5-nano and gpt-5-mini, gemini-3.1-flash-lite and gemini-3-flash, and two small open Qwen 3.5 models (4 billion and 9 billion parameters). Sixteen of 18 cells ran on the registered setup; about 7,950 conversations ran in total. The primary test regresses the gain on a capability index measured on the model's other two domains, which avoids a statistical artifact the authors prove: regressing a gain on its own baseline produces a negative slope even when none exists.

### Findings

**No cell shows a benefit.** Across the 12 cells in the primary fit, the weighted mean gain was minus 28 points (95 percent interval minus 44 to minus 12). In the 13 cells that met the mechanical validity checks, the multi-agent system was worse in 12, each with an interval excluding zero, and tied in one (gpt-5-nano on airline, 42.7 percent single-agent accuracy). Losses ranged from 11 points to 81 points; the largest fell where the single agent was strongest, such as gemini-3-flash on telecom (91.3 percent alone, 10.7 percent as a team). Per task, the team did better on 43 of 650 task-cell pairs and worse on 330.

**The rule's distinctive prediction failed.** Only two cells had single-agent accuracy clearly below 45 percent (both telecom, with the low-tier gpt-5-nano and gemini-3.1-flash-lite). The rule predicts a gain there. The team lost by 13 and 22 points. With no budget cap at all, it drew level on one of the two at 5.5 times the single agent's median spend and changed nothing on the other. The capability index also showed that those two cells were low because telecom is hard: on their other domains both models scored above 45 percent.

**Whether the gain shrinks with capability is unresolved.** The fitted slope was minus 0.41 points of gain per point of capability, with an interval from minus 1.51 to plus 0.68. Cells scattered far more than sampling noise allows, because model-by-domain effects dominated. What the data fix is the level: the gain is negative across the observed range.

**Most of the loss runs through the budget.** About 74 percent of the mean loss came from episodes that ran out of budget; 18 to 68 percent of team episodes ended that way. The authors located the overhead in the orchestrator's planning calls: raising the database agent's read limit from 4 to 25 calls left gains unchanged.

**The mechanism a team is supposed to supply did not appear.** The usual argument for a checker in this setting is to stop wrongful database writes. Pooled over 13 cells, wrongful writes moved from 55 to 43, a change of minus 0.6 points per episode with an interval covering zero, while budget exhaustion rose by 34 points and inaction by 7. The team converted some failures into others. The policy advisor in this design only advised; a version with veto power was outlined but not run.

**A strong single agent isolates the cause.** On three cells the authors ran one agent with the same plan-and-review cycle in one context, under the same budget. On both cells below 45 percent, it did no worse than the stock agent while the team lost to it, which places the loss in splitting the work across agents. One confound applies: the single-context planner saw the tool definitions and the team's orchestrator did not.

**Robustness and cost.** Capping the single agent too cost it 4 to 7 points and left the team still behind. Swapping in the benchmark's default customer simulator produced the study's only positive gain, 10 points on the tied cell, with a change between simulators whose interval covered zero. In dollars the team cost 1.4 to 3.1 times as much per cell, and the single agent was cheaper and more accurate in 11 of 12 priced cells.

### Strength of the evidence

The design is careful. It was pre-registered with a dated change log, budgets are enforced per call and checked against provider bills to within 1 percent, every team cell was checked from logs for broken behavior, and the authors report a deviation instead of hiding it. That deviation is the main weakness. The orchestrator planned without seeing the tools' schemas, so only about 25 percent of its write orders ran exactly as issued, against a pre-registered bar of 90 percent. A fix raised that share to 72 percent in a small smoke test (five tasks on each of three cells), but the grid was not rerun. The paper therefore measures one architecture with a handicapped planner, and a better-built team might do better. Other limits: only two cells test the rule's distinctive prediction; the customer is simulated, and simulators are more cooperative than people; the models may have seen the benchmark in training (the 4-billion-parameter model scored above every closed model on telecom, which the authors flag); and the official score misses violations that never reach a write.

### Transfer to a chat agent following a system prompt

The setting is close to a chat agent's: one conversation, one user who supplies information turn by turn, and actions under rules. The models are small and older than current frontier models, and no Claude model ran, so the size of the losses shouldn't be carried over. The direction is more likely to transfer, because the authors' explanation is structural: a conversation can't be decomposed in advance, every tool call has to pass through the agent facing the user, and coordination spends budget the conversation needs. The paper tests splitting a single conversation's turn-level work across agents. It doesn't test a chat agent handing a separate, sizeable job to a subagent while it keeps the conversation, which is what the prompt allows.

## What the prompt applies

The paper supports the restrictive half of the delegation rules in the "While working" subsection.

- **Keep work that needs shared context or a chain of linked decisions with a single agent.** This is the paper's strongest bearing. The team split turn-by-turn work that depends on what the customer just said, and the split cost 11 to 81 points. The single-context agent running the same plan-and-review cycle did not lose, which ties the cost to dividing tightly coupled work. The prompt's rule names exactly this kind of work as unsuitable for delegation.
- **Hand a goal to a subagent only when it is sizeable and independent.** The paper offers negative evidence: when work was neither independent nor separable, the team lost at every capability level tested. It says nothing positive about delegating independent work, since it never tested that.
- **Don't use subagents to recheck your own work.** The policy advisor was a second agent consulted on the main agent's plans, and it did not detectably reduce wrongful writes while it consumed budget. This is consistent with the rule. The advisor had no veto, so the paper doesn't show that an empowered checker fails.
- **Count a goal as done only when something outside your own judgment shows it.** The paper scores outcomes by database state and found that a team can look busy (planning, consulting, ordering writes) while the state goes unchanged. That supports judging by outside evidence, though the paper doesn't test any instruction about it.

The paper also bears on what the prompt leaves unsaid about capability. It tested cheap models (gpt-5-nano, gemini-3.1-flash-lite and the small Qwen models), and teams built from them lost as badly as teams built from stronger ones. It never mixed models within a team, so it gives no evidence about cheap models in particular roles. Its one statement about weak models is that same-model teams didn't rescue them in this setting.

## What the prompt omits

- **A default of no delegation for user-facing work.** The authors conclude that a single agent should be the default for agents that serve people, with the burden of proof on multi-agent designs. The prompt is permissive by comparison: it allows delegation of sizeable independent goals. The two can coexist because the prompt keeps the conversation and any linked decisions with the main agent.
- **Budget accounting.** The paper shows that coordination overhead can exhaust a fixed budget before the task is done. The prompt mentions cost only through confirmation before spending money, and it leaves token budgets to the harness.
- **The capability-saturation rule itself.** The prompt doesn't adopt the 45 percent rule or any capability threshold, which fits this paper's finding that the rule failed to transfer.
- **Sharing tool definitions with planners.** The fix the authors tested, showing the planner the tool schemas, is a harness design point a system prompt can't set.

## Conflicts with other references

When this file was drafted, the other new reference files did not exist yet, so the points below were first worked from the sources themselves. Before finishing, I read the relevant passages of `anthropic-opus-5-prompting.md`, `anthropic-opus-5-5-prompting.md`, `anthropic-harness-design.md` and `progress-mirage.md`, which other reviewers had written in the meantime, and checked that the points below agree with them.

- **Teams of coding agents (`coding-agent-teams.md`, arXiv 2610.17471).** That paper found teams beat solo agents on long autonomous coding tasks, while this one found teams lost in every valid cell. The difference is the task: probing a black-box program splits naturally across workers, and a customer conversation does not. The prompt's independence test is the rule that reconciles the two.
- **Collaboration tax (`collaboration-tax.md`, arXiv 2608.22152).** Both papers find a cost to dividing work that needs shared context. The collaboration-tax paper says the cost falls as models get stronger; this paper found the largest losses on the strongest cells, though its slope test couldn't settle the question. The two measured different things (peers pooling private views versus an orchestrator with role agents), so they agree that the cost exists and disagree only on how it scales.
- **Anthropic, Prompting Claude Opus 5 (`anthropic-opus-5-prompting.md`).** The guide says Opus 5 coordinates subagent teams well and needs damping because it delegates too readily. This paper had no Claude models, so it can't speak to Opus 5's coordination skill. It agrees with the guide's damping advice.
- **Anthropic harness design post (`anthropic-harness-design.md`).** Anthropic's harness kept separate planner and evaluator agents for long autonomous builds and found them worth the cost. That setting had no user in the loop and work that splits into sprints, which is the regime where this paper's authors expect teams to have a chance.

- **Anthropic, Prompting Claude Opus 5.5 (`anthropic-opus-5-5-prompting.md`).** The guide reports that small Opus 5.5 teams on research tasks matched a single agent's quality and finished sooner when given a time budget. This paper found the opposite direction for teams in a customer conversation. The tasks differ (research splits into separate sub-questions; a live conversation does not), and the guide gives no numbers. The prompt's rule follows both: delegate separable research, keep the conversation with one agent.

## Review notes

I read the main text in full, from the abstract through the conclusion, including the validity, budget, failure-taxonomy and extension sections and their tables. I did not read appendices A to C (the full hypothesis table, the notes on the Qwen cells and the reporting notes) beyond what the main text summarizes. The paper's model names and dates are as reported; I did not check them against provider documentation. The account of the Kim and colleagues rule comes from this paper's summary, since I did not read that paper. The view that a chat agent's delegation of separate jobs differs from this paper's split of turn-level work is my inference.
