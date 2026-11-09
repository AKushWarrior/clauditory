# Cited but Not Verified: Parsing and Evaluating Source Attribution in LLM Deep Research Agents

- **Source:** Hailey Onweller, Elias Lumer, Austin Huber, Pia Ramchandani, Vamse Kumar Subbiah, and Corey Feld (Commercial Technology and Innovation Office, PricewaterhouseCoopers U.S.), "Cited but Not Verified: Parsing and Evaluating Source Attribution in LLM Deep Research Agents," arXiv preprint, version 1 submitted May 7, 2026, arXiv 2605.06635
- **Type:** preprint (not peer-reviewed)
- **Where to find it:** https://arxiv.org/abs/2605.06635
- **Last reviewed:** 2026-11-05
- **Used by:** `chat/system/prompt.md`

## What the source covers

The paper measures whether the citations in reports written by deep research agents support the claims attached to them. Its answer is that they often don't, and that the usual quick checks miss it. For the frontier models tested, cited links almost always worked and almost always led to a page on the right topic, but the page backed the specific claim only 39 to 77 percent of the time. An ablation on two models found that letting the agent make more tool calls lowered that rate further.

### Method

The authors built a three-stage pipeline. An agent with web search writes a Markdown report with inline citations. A parser turns the report into a syntax tree, splits it into sentences, and pairs each sentence with its citations, without using a model. When a citation closes a passage, the parser assigns it to every uncited sentence before it in that passage, on the reasoning that models often cite once for several related claims. The pipeline then fetches every cited page and scores each claim and citation pair on three yes-or-no tests:

- **Link works:** the address returns content, as opposed to an error, a timeout, or a block. No model is involved.
- **Relevant content:** a model acting as judge decides whether the page covers the claim's topic, reading the first 5,000 characters of the page.
- **Fact check:** a judge model decides whether the facts, numbers, dates, and assertions in the claim are supported by the page. Anything contradicted, absent, or uncertain scores zero. The authors say they calibrated this judge by reviewing 50 to 100 of its decisions by hand.

The tests make a ladder. A citation can pass the first and fail the second, or pass the first two and fail the third.

The study covers 14 models. Four are from OpenAI (GPT-5.2, GPT-5.4, GPT-5 Mini, and Codex). Five are from Anthropic (Claude Sonnet 4.5, Sonnet 4.6, Opus 4.5, Opus 4.6, and Haiku 4.5). Two are from Google (Gemini 3.1 Pro and Gemini 3 Flash). The other three are open models (OSS-120B, Llama 4 Maverick, and Pixtral Large). The queries come from DeepResearch Bench and BrowseComp; the paper says 130 queries in total. Every model ran in the authors' own agent harness, with one system prompt that set the citation format and a minimum search depth. The runs are not the vendors' own deep research products. For the ablation, the authors capped the number of tool calls at 2, 10, 30, 50, 70, 100, or 150 for GPT-5.4 and Claude Opus 4.6.

### Findings

**Working, relevant links hide unsupported claims.** Twelve of the 14 models passed the link test on more than 94 percent of citations, and every frontier model passed the relevance test on more than 80 percent. Fact-check rates ranged from 24 percent (OSS-120B) to 77 percent (Claude Opus 4.5). Among frontier models the range was 39 percent (GPT-5 Mini) to 77 percent, with most between 45 and 60 percent. A reader who clicks a citation in one of these reports will almost always find a live page on the right subject. For the frontier models, that page fails to support the specific claim somewhere between a quarter and three-fifths of the time.

**Providers differ, and more citations went with lower accuracy.** OpenAI's models completed every query and produced the most citations (GPT-5 Mini made 1,272 attributions) but scored 39 to 59 percent on the fact check. Anthropic's models completed 83 to 97 percent of queries and scored highest. Opus 4.5 scored 77 percent and Haiku 4.5 scored 69 percent, against 52 to 59 percent for the Sonnet models and 54 percent for Opus 4.6. Google's two models scored 45 and 49 percent. The open models struggled with the task itself, completing only 17 to 40 percent of queries and making far fewer tool calls (7 to 60, against 73 to 211 for frontier models). The authors suggest that synthesizing from more passages raises the chance of misattributing or mixing up facts, and that selective citation may be the better strategy. They offer this as a hypothesis; the comparison across providers can't separate it from other differences between the models.

**More search, lower accuracy.** In the ablation, fact-check accuracy fell for both models as the tool-call cap rose, while the link and relevance tests stayed above 92 percent at every depth. GPT-5.4 fell from 79 percent at 2 calls to 17 percent at 150, with the largest drop between 2 and 10 calls (79 to 46 percent). Claude Opus 4.6 fell from 80 percent to 58 percent. The paper's headline figure, a drop of about 42 percent on average, is the mean of these two drops in percentage points (62 and 22). In relative terms the drops are about 79 percent and 28 percent. The decline is not strictly steady for either model; each has one step where accuracy rose slightly. The authors call this an information-overload effect: models keep finding live, relevant pages but synthesize their facts less accurately as the number of sources grows.

The error analysis on broken links found the expected causes (pages removed, paywalls and bot blocks, and timeouts). GPT-5.4 had one failed link in 2,159 checks, while two open models failed on 16 and 19 percent of links.

### Strength of the evidence

The central finding, that link and relevance checks pass while factual support often fails, is well grounded in what the pipeline measures, and the gap is large and consistent across 11 frontier models. The specific numbers are weaker than the paper's presentation suggests.

- **The judge is unnamed and loosely validated.** The paper doesn't say which model judged relevance and factual support, and it reports no agreement rate for the hand calibration of 50 to 100 decisions. Judge error affects every fact-check figure.
- **The scoring leans toward failure.** Absent or uncertain support counts as unsupported. If the fact-check judge saw only part of a long page, as the relevance judge did, claims supported further down would fail. The paper states the 5,000-character limit only for relevance. The parser's rule that a citation covers the uncited sentences before it can also charge a citation with claims its author never meant it to support. All three effects would push fact-check rates down, so the true rate of unsupported citations may be lower than reported. None changes the direction of the finding.
- **Sample sizes are small and partly unclear.** The success rates fall on multiples of one-thirtieth (16.7, 83.3, 93.3, and 96.7 percent), which suggests about 30 queries per model, not 130; the paper doesn't explain the difference. The ablation reports no number of attributions per depth and no confidence intervals, and some of its cells look small (one model's 92.3 percent link rate at 10 calls fits 12 of 13).
- **Selection may flatter some models.** Anthropic's models failed to complete more queries than OpenAI's. If the queries they dropped were the hardest, their fact-check rates are measured on an easier set.
- **The ablation confounds depth with volume.** A higher tool-call cap also means more citations and longer reports, so the ablation can't say whether accuracy falls because each synthesis has more sources or because the report makes more claims. Only two models were tested, and the cap is a maximum, not the number of calls actually made.
- **The harness is the authors'.** One system prompt that demanded a minimum search depth may have pushed models to cite pages they had skimmed. Vendor research products with their own citation pipelines might behave differently.

On transfer: every model tested predates Claude Opus 5 and Opus 5.5. Anthropic's Opus 5.5 prompting guide says that model is much less likely to cite the wrong source, without giving figures. The exact rates here should not be assumed for a current model. The qualitative points are likely to hold: a live, on-topic citation is weak evidence that the claim is supported, and reading more sources doesn't make a synthesis more accurate. Both describe the structure of the task more than the weaknesses of a particular model.

## What the prompt applies

**Opening a cited source before relying on it (Working loop, While working, informational goals, third bullet, second sentence).** The prompt tells the model to open a cited source before relying on it, including one a subagent supplied, and to check that it says what the citation claims. The rule is the paper's central lesson put into an instruction. The paper shows that the checks a hurried reader might run pass for nearly every citation. Those are that the link works and that the page is on topic. Support for the specific claim still fails in a large share of cases. Only reading the source against the claim catches that failure, which is exactly what the paper's fact-check test does. The paper measured citations in the agents' own reports, not citations an agent received from another agent. Its support for the subagent clause therefore rests on extension, because a subagent's report is the same kind of artifact the paper tested.

**Changing the angle when an approach stops producing new information (same section, fourth bullet).** The ablation is the closest evidence the prompt has for this rule. More tool calls lowered factual accuracy for both models while every surface measure held steady, so searching longer in the same way made the synthesis worse. That supports the prompt's premise that repeating an unproductive approach has a cost beyond time. The paper never tested changing strategy, so it supports stopping or redirecting a search that has gone flat, not any particular new angle.

**Stopping when the goals are met (Before starting, second bullet; While working, second bullet).** The prompt uses its goal list as a stopping rule and says speed comes partly from stopping on time. The ablation gives this rule a quality reason as well as a speed one, since extra research beyond what the answer needs lowered accuracy in these tests.

**Citing the sources the answer rests on (Composing answers, last bullet).** The prompt asks for links to the sources the answer depends on. The paper's finding that Claude Opus 4.5, which cited selectively, had the most accurate citations fits a practice of citing the sources an answer actually rests on rather than every page visited. The paper offers that interpretation only as a hypothesis.

**Saying what was verified (Composing answers, fourth bullet).** The prompt asks the model to say what it verified and how. The paper's three-level ladder gives that report a useful distinction. A citation whose link was checked is not a citation whose claim was checked.

The prompt's rule to confirm load-bearing claims with an independent second source (While working, informational goals, third bullet, first sentence) gets no support from this paper, which checks each citation against its own source only.

## What the prompt omits

**Separate link and relevance checks.** The prompt asks only for the strongest check, that the source says what is claimed. The paper's lower rungs pass almost always, so naming them would add steps that rarely catch anything.

**A cap on search depth.** The ablation could justify a numeric limit on tool calls. Its two-model result and the depth-volume confound are too thin to set a number, though, and a cap would conflict with the prompt's rule never to skip work the answer needs. The prompt relies on its stopping rule and its rule on changing the angle instead.

**Checking every citation.** The prompt limits its checks to load-bearing claims and citations it relies on. The paper's pipeline checks every citation, which is affordable for an offline evaluation but costly inside a conversation. The paper doesn't test whether checking a subset is enough.

**The evaluation machinery.** The parser, the judge models, and the scoring pipeline are tools for measuring citation quality across many reports. They don't translate into a behavior for a single agent.

**Provider comparisons and open-model results.** Provider comparisons and open-model results matter for choosing a model, not for writing a prompt.

## Conflicts with other references

Reference files for `zhu-misknow-agent.md` and `anthropic-prompting-best-practices.md` existed in the folder when this file was written and were read. Files for the Pradhan, Khurana, and Wineburg sources had not been written yet, so the notes on them come from the sources themselves; those files now exist and agree with the notes below.

- **Zhu and colleagues, MisKnow-Agent (`zhu-misknow-agent.md`).** The two papers cover opposite halves of the citation problem. Onweller and colleagues find citations that don't support their claims. Zhu and colleagues plant documents that do support the claim exactly but are false, and agents adopt them about half the time. A model that opened each Zhu document would find its citation accurate. Opening the source is therefore necessary but not sufficient, and the prompt pairs it with independent corroboration and judging sources by provenance. No conflict, but the pairing matters: either rule alone leaves one failure open.
- **Anthropic prompting guides (`anthropic-prompting-best-practices.md`, `anthropic-opus-5-prompting.md`, and `anthropic-opus-5-5-prompting.md`).** The Opus 5 guide says that model checks its own work unprompted and that explicit verification instructions make it over-verify. The prompt's citation check is an explicit verification instruction. The tension is narrower than it looks. The guide's warning targets rechecking one's own answer. This paper shows that checking cited sources against claims catches errors that surface checks miss. That is a different kind of check, and the models the paper tested did not perform it reliably on their own. The Opus 5.5 guide's claim that the model cites the wrong source much less often weakens the transfer of this paper's numbers, but it reports no rates, so it doesn't show the check is unnecessary. The prompt keeps the rule and scopes it to citations the model relies on, alongside second-source checks for load-bearing claims, which limits the cost. On the model's own work it follows the guide. An operational goal counts as done on evidence the work already produces, with no separate review pass, and subagents are not used to recheck the model's own work.
- **Pradhan and Goley, "Trust, but Don't Verify" (`pradhan-trust-dont-verify.md`, arXiv 2606.05403).** They find that models detect fabricated statistics in isolation but don't apply that ability while synthesizing several sources. That finding fits this paper's information-overload account: accuracy falls when synthesis involves many sources. Both suggest checking a claim against its source as a separate, focused step, which is what the prompt asks for.
- **Khurana and colleagues, AuthorityBench (`khurana-authoritybench.md`, arXiv 2606.13104).** They find that the mere presence of a citation, real or fabricated, raises hallucination rates. Combined with this paper, a citation is a weak signal in both directions: models trust it too much as readers and attach it too loosely as writers. No conflict.
- **Wineburg and McGrew, lateral reading (`wineburg-mcgrew-lateral-reading.md`, 2019).** Fact-checkers did well by leaving a site to see what others said about it. This paper's check goes the other way, into the cited source. The two checks answer different questions (is the source trustworthy, and does it say what is claimed), and the prompt asks for both.

## Review notes

I read the full text of version 1 from the local copy, including the methods, all three tables, the error analysis, and the limitations. Figure 1 and Figure 2 were not available as images; the numbers in Tables 2 and 3 cover what Figure 2 shows. I confirmed the submission date, May 7, 2026, from the local copy of the arXiv abstract page; that is within the recency window. The paper doesn't name its judge model, give a judge-human agreement rate, say whether the fact-check judge saw the whole page, or explain why success rates imply about 30 queries per model when the text says 130. The interpretation of the 42 percent figure as a mean of percentage-point drops is my calculation from the tables. The note on the Opus 5.5 guide rests on its one sentence about citing the wrong source. The notes on Pradhan and Khurana rest on their abstracts, and the Wineburg and McGrew note rests on the study's widely reported findings.
