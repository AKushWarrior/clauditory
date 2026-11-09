# One Token Away from Collapse: The Fragility of Instruction-Tuned Helpfulness

- **Source:** Erfan Baghaei Potraghloo, Seyedarmin Azizi, Souvik Kundu, and Massoud Pedram, "One Token Away from Collapse: The Fragility of Instruction-Tuned Helpfulness," arXiv preprint arXiv:2604.13006 `[cs.CL]`. Version 1 was submitted April 14, 2026; the copy read for this review is version 2, dated April 27, 2026. The first two authors are listed as equal contributors. Three authors are at the University of Southern California and one is at Intel AI.
- **Type:** preprint (not peer-reviewed)
- **Where to find it:** https://arxiv.org/abs/2604.13006
- **Last reviewed:** 2026-11-05
- **Used by:** `chat/skills/writing-style/SKILL.md`

## What the source covers

### Purpose and framing

The paper asks how well the helpfulness of an instruction-tuned model holds up when the prompt adds a small rule about surface form, such as a ban on one punctuation mark or one common word. The authors start from the usual assumption that instruction tuning teaches a general skill at giving thorough answers, so a model should explain a topic about as well with or without commas. They report the opposite. When a trivial lexical rule is added to the same question, the model does not write the same answer with the banned token removed. It switches to a much shorter, flatter answer. The authors call this constraint-induced response collapse. Most of the paper tries to show three things about it. The lost material is real content and not padding. The model could have written the full answer under the rule. And instruction tuning is what makes the model fragile in this way.

The paper describes itself as diagnostic. It measures and explains the effect and leaves fixes to later work.

### Prompts, constraints, and models

The main set has 40 prompts in four groups of 10: explanation, how-to advice, analysis or comparison, and technical detail. They are ordinary informational questions of the kind that draw a multi-paragraph answer, such as how vaccines work or how a compiler translates code.

The authors define eight lexical constraints, each added to the prompt as one extra sentence. Three ban a punctuation mark (commas, colons, or semicolons). One bans lists and dashes used as list markers and asks for prose paragraphs. Two ban words: one bans the article the, and the other bans a set of five discourse markers that includes however, therefore, and additionally. Two combine rules (commas with colons, and commas with lists). The answer format stays free prose throughout; nothing asks for structured output.

Three open-weight instruction-tuned models get the detailed analysis: Llama 3.1 8B Instruct, Qwen 2.5 7B Instruct, and Mistral 7B Instruct v0.3. Four more get the behavioral tests only. They are Llama 3.3 70B Instruct, Qwen3 30B-A3B Instruct (a mixture-of-experts model with about 3 billion active parameters), OLMo 3 7B Instruct, and the closed-weight GPT-4o-mini. The base versions of the three main open models serve as controls.

Each prompt and constraint pair got three samples at temperature 0.7 with a cap of 1,024 tokens. The study used three judging methods. In independent scoring, a judge model rates one answer alone on four dimensions. In pairwise comparison, a judge sees the unconstrained and constrained answers side by side, in random order, and rates each for `comprehensiveness` and usefulness. The third method is a blinded study with 10 human raters with graduate STEM training. The judge models were GPT-4o-mini, GPT-4o, and Claude Sonnet 4.6. The paper reports the percent change in the `comprehensiveness` score and the share of pairs the unconstrained answer won.

The constraints were followed. Automated checks found satisfaction above 90 percent for most model and constraint pairs, with an average near 94 percent for Llama and 96 percent for Qwen. The one weak case was Llama on colons, at about 66 percent, because the model kept writing step headers that end in a colon. The authors read that miss as a sign of how tightly the model is tied to its usual answer templates. The collapse therefore appears in answers that obey the rule.

### Main result

All three main models lost `comprehensiveness` under every type of constraint. With GPT-4o-mini as judge, the average change was a 40 percent loss for Qwen, 23 percent for Llama, and 14 percent for Mistral, and the unconstrained answer won 98, 98, and 77 percent of pairs. With GPT-4o as judge, across all seven instruction-tuned models, losses ran from 17 percent for Mistral to 48 percent for Qwen 2.5, and the unconstrained answer won between 78 and nearly 100 percent of pairs. GPT-4o-mini lost 31 percent. The 70B Llama lost more than the 8B Llama (35 percent against 26 percent), so scale did not protect the model in this sample. A third judge from another model family, Claude Sonnet 4.6, gave lower absolute scores but found losses of the same size and in the same order.

The damage did not depend on the banned token being a formatting mark. Banning the article cost Llama about 23 percent and Qwen about 42 percent, close to the cost of banning commas (27 and 61 percent). Combining two bans did little more harm than one, which the authors read as a switch between strategies rather than a gradual loss. The list ban was the mildest, at 12 to 13 percent for Llama and Qwen, and for Mistral it made no consistent difference: one judge scored it slightly better than baseline and another slightly worse.

The paper's opening figure shows the comma ban cutting one Llama answer from 685 tokens to 297. The supplementary discussion gives the same two numbers as Llama's average length under the comma ban. It adds a 41 percent drop in distinct content words and the loss of most formatting (bold, headers, code blocks, and lists). It is therefore unclear whether the figure shows a typical case or the average.

The effect replicated on MT-Bench, an outside set of 80 questions in eight categories, with losses of 17 percent for Llama and 41 percent for Qwen under three constraints. Llama's math answers were the one exception, which the authors explain by their short, formulaic style.

### Lost content or lost padding

The authors ran two checks against the reading that constrained answers are only shorter. In the first, GPT-4o pulled 11 to 20 atomic factual claims from each unconstrained answer and then checked, with loose paraphrase matching, whether each claim appeared in the constrained answer. Over 192 pairs, constrained answers kept about half the claims (49.8 percent) and about half the length (49.0 percent). Claims and words fell together, which is what a real loss of content predicts. For the list ban, answers kept 89 percent of their length but only 58 percent of their claims. This analysis used eight prompts and is judged end to end by a model, and the authors present it as supporting evidence.

In the second, the human raters scored each answer on four information criteria (coverage, `comprehensiveness`, correctness, and helpfulness) and two surface criteria (concision and readability), across 960 blinded pairs. Information scores fell 1.5 to 2.3 times as much as surface scores. For Mistral, information fell 16 percent and surface 7 percent; for Qwen, 44 percent and 29 percent. Human ratings of `comprehensiveness` tracked the model judges closely, and humans tended to see slightly larger losses. The authors argue that if raters were only rewarding length or formatting, the two kinds of score would fall by similar amounts.

### The planning failure behind the collapse

The planning analysis is the part of the paper the skill relies on. The authors test whether the model cannot write a full answer under the rule or simply does not plan to.

The test is a two-pass protocol on 10 prompts and two constraints (commas and the article). For each prompt the model produced three answers. One was unconstrained. One was a single-pass answer with the rule in the prompt. The third was a two-pass answer, made by rewriting the unconstrained answer under the rule with an explicit instruction to keep it as thorough as before. Single-pass answers kept about half the baseline length for Llama and Mistral (49 and 52 percent) and only 13 percent for Qwen. Two-pass answers kept 96 percent for Llama, 91 percent for Mistral, and 59 percent for Qwen. Qwen recovered well for the article ban (from 16 percent to 79 percent) and poorly for the comma ban (from 9 percent to 40 percent). The authors treat that as a second, smaller limit on capability for the most frequent tokens in heavily tuned models. The paper's worked example shows a 402-word answer collapsing to a 55-word paragraph under single-pass generation and coming back as a 323-word rewrite with the same seven-step structure.

Two further analyses support the planning account. Linear probes trained on the hidden state at the last prompt token predicted the length of the coming answer. The best layer was near the middle of the network, and the explained share of variance ran between about 0.51 and 0.94 across five instruction-tuned models. The more a model collapsed, the better the probe predicted its length, with a correlation of 0.92 across those five models. On base models the same probes did worse than guessing the mean. A token-level comparison found that constrained and unconstrained runs diverge within the first one to three generated tokens. In one example the unconstrained model opened with a bold title and the constrained one opened with a plain sentence. A perplexity check under the base model found that the comma-free or article-free two-pass rewrites were only slightly less likely than ordinary text for Llama and Mistral. The full constrained answer is therefore not an improbable kind of text. For Qwen, the collapsed single-pass answer was in fact less likely under the base model than the full rewrite.

The authors' reading is that the decision to write a short answer is made while the model reads the prompt, before it writes anything, and that the capacity for the full answer is still there.

### Instruction tuning as the cause

Base models under the same rules showed small and inconsistent effects. Llama base lost 8 percent and Mistral base lost 11 percent, with win rates near chance, and Qwen base gained 7 percent. The paper attributes the gap to instruction tuning, which in the authors' account teaches a narrow set of high-quality answer templates. When a rule blocks a token those templates depend on, the model has no other full-length strategy and falls back to a minimal one.

### Rules closer to real deployments

The authors also tested four instructions of the kind found in production system prompts, on Llama and Qwen. They were a formal, professional tone; a ban on conversational openers, with the instruction to start on the first substantive sentence; a request to hedge claims instead of stating them flatly; and a request for plain language with no jargon. The average losses were 23 percent for Llama and 34 percent for Qwen, close to the lexical bans. The ban on openers cost Qwen 40 percent of `comprehensiveness` and about three quarters of its words (from 448 words to 113), even though it restricts only the first few tokens. The authors propose that the trained opener works as a trigger that starts the full answer template. The hedging request cost both models about 26 percent. The plain-language request did the most damage, but the authors note that some loss there is intended, since it asks the model to drop technical material.

### Evaluation blind spot

Independent scoring missed most of the loss. On Llama with GPT-4o-mini as judge, independent scores fell 3.5 percent on average where pairwise comparison found 23 percent, and for the list ban independent scoring found no loss at all against 13 percent pairwise. The authors explain that a judge with no full answer to compare against rates a short, locally coherent answer as good.

### Strengths

The paper tests seven instruction-tuned models from five families, uses three judges from two families, adds a blinded human study, and checks the result against an outside benchmark. The base-model controls separate the effect of tuning from the effect of the rule itself. The two-pass test speaks directly to whether the content is recoverable, and the claim-coverage and human-criteria checks address the obvious objection that judges prefer long answers. The authors report satisfaction rates, per-constraint tables, and their judge prompts.

### Weaknesses and limits

The models are small or mid-sized and mostly from 2024 and 2025. The largest open model is 70B and the only closed model is GPT-4o-mini, so the paper says nothing direct about current frontier models. The internal analyses cover only three models of 7 to 8 billion parameters.

The two-pass test is small: 10 prompts and two constraints. Its rewrite step also carried an explicit instruction to keep the answer thorough, which the single-pass condition lacked, so part of the recovery may come from that instruction and not only from the order of drafting and constraining. Recovery is measured only by word count; the judges and the claim check were not applied to the rewrites.

The prompts are 40 general informational questions and the answers are capped at 1,024 tokens. The lexical bans are deliberately artificial stress tests, as the authors say, and the realistic rules were tested on only two models. The claim-coverage check uses eight prompts and is judged by a model at every step. The human raters number 10.

Some numbers in the paper do not match each other. The main text gives the realistic-rule losses as 22 to 34 percent, while a later discussion section gives 24 to 32 percent, which matches an earlier three-rule pilot the paper also reports. The abstract gives probe fits up to 0.94, while the figure discussion describes Qwen 2.5 at 0.93 and the table lists 0.925, with the 0.94 belonging to Qwen3. The cross-judge appendix heading has no text under it in the copy read, although the table appears elsewhere. The mismatches are small and do not change the findings.

The causal language is stronger than the design. Converging evidence from probes, rewrites, and base models fits a planning account well, but none of these is an intervention on the planning step itself.

### Principles the paper implies for writers

A rule about surface form, given to a model at the moment it plans an answer, can cost far more content than the rule itself requires, and the model will still obey the rule. The cost is largest for frequent tokens and for rules that block the opening the model would normally use. The same model can usually meet the rule without the loss if it first writes freely and then rewrites. A reader who sees only the constrained answer may not notice what is missing.

## What the skill applies

The paper is the main support for step 2 of "How to use this skill" in SKILL.md. That step tells the writer to finish the draft first and then run the section 10 revision pass. It gives the reason that constraints applied while drafting tend to make a model plan a thinner answer, while drafting freely before revising keeps the content. That reason is a direct reading of the paper's single-pass and two-pass results. The rule in the prompt cut answers to about half their length for two of three models. Applying the same rule as a rewrite of a free draft kept most of the length. The word "tends" in the skill reflects that the effect varies by model and by rule.

The paper also supports keeping the lists of patterns to remove only in section 10 and writing sections 1 to 8 as descriptions of the target style. Several of the paper's constraints closely match items in the revision pass. The discourse-marker ban includes additionally, which item 7 and item 8 name. The opener ban targets the stock openers in item 6. The list ban resembles the formatting rules in section 6 and item 11. The paper found that each of these, given as a drafting-time rule, cost content. Section 10 applies them to a finished draft, which is the two-pass arrangement.

The hedging result is consistent with section 4's rule to state facts plainly and hedge only the interpretation that is uncertain. A blanket instruction to hedge every claim cost both tested models about a quarter of their content. The skill never asks for blanket hedging, and item 9 of the revision pass removes stacked hedges.

## What the skill omits

The skill does not cite the paper's percentages. They describe 2024 and 2025 models of 7 to 70 billion parameters on short informational answers, and attaching them to the skill's rules would give false precision for current models and for long-form writing.

It leaves out the probing, token-divergence, and perplexity analyses. They explain where the decision forms inside a model, and a writer cannot act on them.

It leaves out the evaluation findings. The gap between independent and pairwise scoring matters to anyone testing the skill, and a later evaluation of the skill should compare drafts side by side. That is advice for the people maintaining this repository and does not belong in the skill file.

It does not adopt the paper's call to train models to tolerate constraints, which is advice for model builders.

The skill does not copy the two-pass protocol's instruction word for word, but section 10 asks for the same thing. It tells the writer to keep every point the original sentence carried while rewriting each hit, and to check that fixes did not create new hits. The paper's recovery was measured with its own instruction present, so the skill's version of the second pass is close to the one the paper tested but not identical.

## Conflicts with other references

taboo-negative-constraints.md is the other reference behind step 2, and the two agree. The Taboo study found that naming forbidden words in the prompt made models avoid them reliably and made their descriptions less informative. That is the same trade this paper reports: high compliance at drafting time and a thinner answer. The Taboo study works with 30-word descriptions in Italian and this paper with long English answers, so they are independent checks on the same pattern in different settings.

The skill previously justified step 2 with a different mechanism, the claim that a list of banned words read before drafting primes the writer to use those words. That claim came from a reference that has since aged out under the eight-month rule in AGENTS.md and was removed. The Taboo study's results point against it. This paper does not test priming at all; its banned tokens were obeyed more than 90 percent of the time. The skill's design (positive drafting sections, lists only in the revision pass) survives the change, but the reason for it is now loss of content, not reappearance of the banned words.

freeburg-em-dashes.md found that an explicit ban on the em dash lowered its rate in all four models tested. That fits this paper's finding that models obey surface bans. Freeburg did not measure content, so the two do not conflict; together they say a ban works on the banned item and, per this paper, can cost content when given at drafting time. The skill names the em dash only in section 10, which suits both.

anthropic-prompting-docs.md recommends telling a model what to do in place of what to avoid. This paper's realistic constraints include positive instructions (a formal tone, plain language) that also cost content, so the paper suggests that the timing of a rule matters as well as its framing. The skill follows both: its drafting sections are positive, and its lists come after the draft.

There is one internal tension. Section 1 of SKILL.md tells the writer, while drafting, to skip openings that restate or praise the question, and section 6 tells the writer to default to paragraphs. Those resemble the opener ban and the list ban, both of which cost content in this paper when placed in the prompt. The skill states both rules mainly as positive targets (open with the answer; use a list only for parallel items) and repeats the removal in section 10, items 6 and 11. The paper tested blunt prohibitions on small models, and a style guide read by a frontier model is a different case, so the tension is noted here and the skill is unchanged.

durik-hedges-credibility.md concerns how hedges affect readers' trust. This paper adds a separate cost of forced hedging, a loss of content in the writer's output. The two point the same way for section 4.

## Review notes

On 2026-11-05 I read the full text of arXiv 2604.13006 version 2 from a local text copy. That covered the abstract, all eight sections, and the appendix sections A through Q. The appendix material included the prompt list, constraint definitions, judge prompts, the two-pass example, satisfaction rates, MT-Bench replication, claim-coverage tables, the human-evaluation protocol, probing tables, token divergence, deployment-constraint details, and the perplexity analysis. The figures were available only as extracted labels and captions, so figure-only values are taken from captions, text, and the numerical appendix tables. I did not read version 1. The version 1 date of April 14, 2026 comes from the arXiv listing as given in the review brief, and the identifier's April 2026 prefix agrees with it. That date falls within eight months of this review, as AGENTS.md requires. Differences between versions 1 and 2 are unknown to me.

The figures for Llama and Qwen constraint satisfaction, the claim-coverage percentages, and the per-model losses come straight from the paper's tables. The description of the list ban's effect on Mistral as inconsistent is my reading of the two judges' tables. The observation that the two-pass rewrite carried an extra instruction is from the paper's own description of the protocol; the point that this confounds the comparison is mine. The match between the paper's constraints and particular revision-pass items is my mapping. The paper does not discuss style guides or revision checklists.

I did not open any code or data. The paper does not link a repository in the copy read.
