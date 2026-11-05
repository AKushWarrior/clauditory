# The Last Fingerprint: How Markdown Training Shapes LLM Prose

- **Source:** E. M. Freeburg (independent researcher), "The Last Fingerprint: How Markdown Training Shapes LLM Prose," arXiv preprint arXiv:2603.27006v1 [cs.CL], submitted March 27, 2026.
- **Type:** preprint (not peer-reviewed)
- **Where to find it:** https://arxiv.org/abs/2603.27006 (PDF at https://arxiv.org/pdf/2603.27006)
- **Last reviewed:** 2026-11-05
- **Used by:** chat/skills/writing-style/SKILL.md

## What the source covers

### Purpose and scope

The paper tries to explain why some large language models use so many em dashes. By early 2026 the em dash had become the most discussed punctuation tell of machine-written prose, yet the author found no formal account of why that one mark, out of all the marks available, would be overrepresented. The paper's proposal is that the em dash is a leftover of markdown. Models trained on corpora full of markdown learn that dashes in various forms (horizontal rules, list markers, front-matter delimiters) mark boundaries between units of text. When a model is told to stop formatting, headers, bullets, and bold are withheld, but the em dash is also ordinary prose punctuation, so the instruction does not register it as formatting and it slips through. The author calls it markdown's last fingerprint and describes it as the smallest unit of markdown-style thinking that survives compression into prose.

The paper is short (about 14 pages with references) and has three parts: a conceptual argument, three small experiments, and a discussion with limitations. The author is an independent researcher with no stated institutional affiliation, and the paper has not been through peer review.

### The five-step genealogy

The argument runs in steps. First, training corpora are saturated with markdown, because GitHub, Stack Exchange, Reddit, and documentation sites use it natively, and quality filtering favors well-structured sources. Second, models internalize the structural orientation of that text, which shows up as unrequested headings, bullets, and bold in default output. Third, the em dash has a dual status: in markdown, dashes act as architecture, while in prose the em dash is legitimate punctuation. The author is careful to say the em dash character itself has no role in markdown syntax; the claim concerns the dash family generally. Fourth, post-training (the paper uses RLHF as shorthand for the whole stack of fine-tuning methods) amplifies or suppresses the latent tendency, partly because human raters may read dash-rich prose as articulate. Fifth, the author argues that the markdown account and the fine-tuning account complement each other: fine-tuning explains that the em dash was selected, and markdown training explains why it was available to be selected.

The author also names two rival explanations and treats both as compatible. One holds that nineteenth-century literature, where em dash use peaked, is overrepresented in training data. The other holds that the em dash works as a prestige marker that raters reward. The paper says the markdown account adds the one thing the rivals lack, an explanation for why the em dash specifically resists formatting suppression.

### Method

The experiments used twelve instruction-tuned models from five providers. Anthropic was represented by Claude Opus 4.6, Claude Sonnet 4, and Claude Haiku 3.5. OpenAI had GPT-4.1, GPT-4o, GPT-4o Mini, and GPT-5.4. Meta had Llama 3.1 8B Instruct and Llama 3.3 70B Instruct. Google had Gemini 2.5 Pro and Gemini 2.5 Flash, and DeepSeek had DeepSeek V3.

Each model wrote essays on ten everyday topics, asked for at 1,000 words each, under two conditions. In the unconstrained condition the prompt only asked for the essay. In the prose-constrained condition the prompt added an instruction to write flowing paragraphs and avoid markdown, headers, bullets, bold, and lists. That yields roughly 10,000 words per model per condition and about 240,000 words overall. The author counted the Unicode em dash character and four overt markdown features (headings, bullets, bold, and numbered lists), normalized to counts per 1,000 words. Calls used provider-default sampling settings, no system prompt, and a cap of 2,048 output tokens. One model, Gemini 2.5 Pro, completed only three topics per condition because of rate limits. All runs took place in February and March 2026.

The human baseline came from eight published essays across literary criticism, journalism, and technical writing, totaling 57,232 words.

### Main findings

Markdown suppression worked almost perfectly. Under the prose-constrained prompt, every model reached zero markdown features except Claude Haiku 3.5, which kept a small residual of 0.9 per 1,000 words.

Em dashes behaved very differently. Unconstrained rates per 1,000 words were 10.62 for GPT-4.1, 9.09 for Claude Opus 4.6, 8.29 for Claude Sonnet 4, 7.51 for Claude Haiku 3.5, 6.95 for DeepSeek V3, 4.16 for GPT-4o Mini, 4.12 for GPT-4o, 3.53 for Gemini 2.5 Pro, 1.43 for GPT-5.4, 1.28 for Gemini 2.5 Flash, and zero for both Llama models. Under the markdown-suppression prompt, GPT-4.1 only fell to 9.10 (a 14 percent drop) and DeepSeek V3 to 5.41 (22 percent). Claude Opus 4.6 fell to 0.19 (98 percent), Haiku 3.5 to 0.18, Sonnet 4 to 1.31, and Gemini 2.5 Pro to zero. GPT-4o fell from 4.12 to 2.68, GPT-5.4 from 1.43 to 0.29. Two low-rate models, GPT-4o Mini and Gemini 2.5 Flash, rose slightly, which the author attributes to sampling noise. The author checked the Llama output for substitute marks such as double hyphens or en dashes and found none.

The human mean was 3.23 per 1,000 words, with a median of 3.83 and a range from 0.33 to 17.12. That is a fifty-fold spread across eight essays. By my count, eight of the twelve models exceeded the human mean unconstrained, but only three (GPT-4.1, DeepSeek V3, and GPT-4o Mini) did so under the markdown-suppression prompt. The author draws the careful conclusion that models do not uniformly overuse the em dash relative to human writing. The defensible claims are that the em dash resists formatting suppression, that rates differ systematically by provider, and that some models stay well above the human mean even when told to avoid markdown.

Study 2 added a third condition on four models: markdown suppression plus an explicit ban on em dashes. GPT-4.1 went from 11.51 unconstrained to 8.20 under markdown suppression to 3.86 under the explicit ban. DeepSeek V3 went from 8.66 to 4.75 to 1.57. Claude Opus 4.6 went from 8.46 to 0.19 to zero. GPT-5.4 went from 0.75 to 0.15 to zero. So naming the mark helps a great deal, but two of the four models still produced it after being told directly not to.

A supplementary long-form test asked four models for 5,000-word essays on three topics. GPT-4.1 reached 14.03 per 1,000 words unconstrained and still produced 6.97 under the explicit ban (over 9,331 words). Claude Opus 4.6 reached 12.91 unconstrained and dropped to zero under either suppression condition. The author reads this as support for the idea that longer text offers more transition points for the structural habit to surface.

Study 3 compared Llama 3.1 8B base and instruct versions, run locally with the same quantization on the same hardware, on ten essay topics. The base model produced 4 em dashes in 8,186 words (0.49 per 1,000) and 28 markdown features. The instruct model produced no em dashes in 9,241 words and 5 markdown features. The author reads the base result as evidence that a latent tendency exists before fine-tuning, and the instruct result as evidence that fine-tuning can push the tendency to zero as well as amplify it.

### Strengths

The suppression design is clean for what it tests. The same prompt reliably removes one class of marks and leaves another, and the three-condition gradient shows that the persistence is not a simple matter of the instruction failing to mention the mark. The paper reports raw rates for every model, states its sample sizes, and admits where differences fall inside noise. It also reports a human baseline and explicitly declines to claim blanket overuse. The practical finding for anyone steering model output is useful and holds up: an instruction about formatting does not control punctuation, and for some models even a direct prohibition leaves a residue.

### Weaknesses and limits

The evidence for the causal story is thin. Nothing in the experiments tests the markdown mechanism directly. The Llama base result shows a small number of em dashes (four instances) but cannot show they come from markdown rather than from literary text, which the author concedes is a live alternative. The base-versus-instruct test covers one family, one small model, and the one family where fine-tuning removes the mark, so it cannot show amplification. The author lists this and the falsifiability problem among the limitations: the genealogy predicts elevated dashes when fine-tuning amplifies the tendency and low dashes when fine-tuning suppresses it, which leaves little that could count against it. The author proposes a corpus analysis of dash positions and a model trained on low-markdown text as future tests, neither of which was run.

The sample is small. Ten topics per condition, one prompt template, and default sampling mean the per-model rates carry real uncertainty, and the paper reports no confidence intervals or significance tests. Run-to-run variation is visible: Claude Opus 4.6 scored 9.09 in the first study and 8.46 unconstrained in the second. The human baseline has only eight essays chosen by the author, and the weighted mean sits below the median, so the baseline is a rough anchor. The paper also mixes API runs for the main study with local quantized runs for the base comparison.

Some of the paper's own numbers do not line up. The abstract gives the top of the range as 9.1, which is GPT-4.1 under suppression, even though the table shows 10.62 unconstrained. Section 3.4 and Section 5.4 refer to a figure of 14.03 for OpenAI as though it came from the main table, but that figure comes from the long-form test. The claim that no prior mechanistic account exists rests on the author's own search, which the paper does not describe. Model names and the use of RLHF as a catch-all also blur which stage of training is responsible. The finding about Altman is cited to a social media post, with no stated detail beyond a summary that em dash frequency was adjusted. I treat the post-training attribution as a hypothesis, not a result.

Finally, the models tested are snapshots from early 2026, and the paper itself shows rapid change: GPT-5.4 sits at 1.43 against 10.62 for GPT-4.1. Rates for any given model will drift with each release.

### Principles the paper implies for writers

Em dash use is a property of the writer and the training process, not a mark of quality or of authorship. Human rates vary widely, so a single em dash proves nothing, but a model that is not corrected will often produce them at rates above the typical human level. A general instruction to write plainly does not catch the habit, and a named prohibition catches more of it. For an agent writing prose, the safe procedure is to name the mark in an explicit check after drafting and to rewrite each hit.

## What the skill applies

The skill uses this source for one rule and for the design of the process around it.

The rule is item 1 of the revision pass in section 10 of SKILL.md, which tells the writer to replace every em dash and every spaced dash used as one with a comma, colon, parentheses, or a period, and to write ranges with the word to. The paper supplies the reason the rule exists as a separate, explicit item and the reason it covers more than the em dash character. Its suppression results show that the general instruction to avoid markdown leaves the mark in place for many models, and that only a direct prohibition reduces it much. That is why the skill does not rely on its section 6 formatting rules to handle dashes and instead lists dashes first in the revision pass. The paper does not discuss replacement punctuation, so the list of replacements in that item comes from the style manuals and the user's request, not from Freeburg.

The paper also shaped the placement of the rule. Because naming the mark in an explicit instruction measurably reduced it, even when it did not fully eliminate it, the skill asks for a check after drafting, and the closing instruction of section 10 asks the writer to confirm the fixes did not create new hits. The long-form result, where GPT-4.1 produced more dashes as essays grew longer, supports applying the check to long responses and reports, not only to short chat replies.

Section 3 of SKILL.md, the rule that asides and breaks are marked with commas, colons, parentheses, or a new sentence and that a mark that does not clarify should be dropped, is consistent with the paper's picture of the dash as a structural joint inserted at transitions. The paper did not produce that wording, but its description of the em dash as a transition marker explains why the skill's replacements are chosen to make the relation between the parts explicit.

The opening paragraph of SKILL.md says readers recognize default machine prose on sight. The popular recognition of the em dash that the paper documents is one example of that, and the paper's own account of the two online discourses is the best sourced statement in the folder of how widespread the em dash tell is.

## What the skill omits

The skill leaves out the genealogy, the markdown-leakage theory, and the fine-tuning attribution argument. They explain why models produce the mark but do not change what a writer should do about it, and the theory is only weakly supported by the paper's own evidence. The skill gives no explanation of cause, because a rule that depends on a contested mechanism is more fragile than a rule that depends on the observed behavior.

It omits all per-model rates and the provider ranking. Those numbers describe early-2026 snapshots that the paper itself shows are changing quickly, and the skill is meant to work across models and releases. It also does not tell the writer to check how a particular model behaves, since the rule applies to everything.

It omits the detection and model-attribution discussion. Identifying which provider wrote a text is not a writing task.

It does not adopt the paper's human baseline as a tolerance. A human mean of about 3 per 1,000 words, and the finding that human use varies fifty-fold, might suggest permitting an occasional em dash. The skill does not do this because the user explicitly asked for em dashes to be banned, and because the skill's concern is the reader's reaction to machine prose, which a human-range rate does not remove. Section 10 item 1 is therefore a flat replacement rule with no allowance.

It omits the finding that the base model showed a small latent rate, the Llama-specific results, and the long-form amplification claim as stated findings. The base comparison rests on four instances in one small model and is too thin to cite.

## Conflicts with other references

At the time of this review the references folder at chat/skills/writing-style/references/ contained no other files, so the comparisons below rest on the descriptions in the project brief for those files and not on text I read. They should be rechecked once the other files exist.

Style manuals. The paper shows that skilled human writers use the em dash, at rates from nearly zero to over 17 per 1,000 words in its sample. That matches nyt-manual-of-style.md, which allows the mark with limits, ap-stylebook.md, which prescribes spaced dashes, and chicago-manual-of-style.md, which prescribes the unspaced em dash and uses en dashes for ranges. The paper therefore gives no support to treating the em dash as an error in general. The skill is stricter than all three manuals by user choice: it bans the mark outright, uses the word to in ranges, and does not follow Chicago's en dash. When the manuals and the paper agree that humans use the mark and the user has asked for a ban, the user's explicit choice sets the rule. Precedence among the manuals (NYT, then AP, then Chicago) is not in play here, since the skill departs from all of them on this point.

One Token Away from Collapse and the Taboo study (one-token-away-from-collapse.md, taboo-negative-constraints.md). The skill keeps its pattern lists, the em dash included, in the post-draft revision pass because those studies found that constraints given while drafting make a model plan a thinner answer. An earlier version of this entry described a conflict with a priming claim, that naming a banned word makes a model produce it; that claim was dropped with its source. Freeburg's gradient agrees with the Taboo study's compliance results: an explicit prohibition lowered the em dash rate in all four models tested, although it did not always reach zero. Freeburg did not measure content, so the papers do not conflict. The skill names the em dash only in the revision pass, where the check works on a finished draft, and treats the prohibition as a reduction, not a guarantee.

Anthropic prompting docs (anthropic-prompting-docs.md). Per the brief, those docs favor positive instructions over prohibitions. Freeburg tested only prohibitions, and a positive instruction (use commas and periods) was never compared against them. The skill pairs the ban with replacements in the same item, which keeps a positive alternative in view.

Kobak, Juzek and Ward, Liang, and Wikipedia lists. Those sources concern vocabulary, and Freeburg concerns one punctuation mark with raw frequency counts, so their measures are not comparable. I found no direct conflict. Where Wikipedia-style lists of machine tells include the em dash, Freeburg is the only source in the set that measured it.

Boggia (boggia-artificial-epanorthosis.md) and Reinhart (reinhart-llm-grammatical-style.md). No conflict, since they concern different features.

## Review notes

I read the full text of the PDF from arXiv (about 6,600 words including references), covering the abstract, all seven sections, the three tables, and the reference list. I also retrieved the arXiv abstract page. I did not read the cited sources, the author's raw data, prompts beyond the ones quoted in the paper, or any code, and the paper links no data or code repository. I could not check the Altman post or the Goedecke essay it cites.

My reading of the tables is direct, but the counts I give of how many models exceeded the human mean are my own tally from the table values. The comparison of the paper's abstract range to its tables is my own observation. The paper is a single-author preprint, so its claims about mechanism are the author's interpretation and should not be cited as established. Only the measured rates and the suppression pattern carry weight in the skill.

The other reference files did not exist when I wrote this, so the Conflicts section relies on the project brief's descriptions of them.
