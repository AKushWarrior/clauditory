# Don't Say It!: Constraints, Compliance, and Communication when Language Models Play Taboo

- **Source:** Sara Candussio, Francesca Padovani, Daniel Scalena and Malvina Nissim, "Don't Say It!: Constraints, Compliance, and Communication when Language Models Play Taboo," arXiv preprint arXiv:2607.00601. Version 1 was submitted 1 July 2026. The paper's header names CLiC-it 2026, the Twelfth Italian Conference on Computational Linguistics, held 14 to 16 September 2026 in Palermo. The first three authors are listed as equal contributors, from the University of Trieste, the University of Groningen and the University of Milano-Bicocca.
- **Type:** conference paper, posted as a preprint
- **Where to find it:** https://arxiv.org/abs/2607.00601 (code and data at https://github.com/DanielSc4/LMtaboo/)
- **Last reviewed:** 2026-11-05
- **Used by:** chat/skills/writing-style/SKILL.md

## What the source covers

### Purpose and framing

The paper uses the party game Taboo to study how language models handle a lexical ban that competes with a communicative goal. In Taboo a player describes a target word so that others can guess it, without using any of a short list of forbidden words, which are usually the most obvious clues. The task therefore asks for two things at once: avoid particular words, and still say enough to identify the target. The authors use it to compare ways of imposing the ban at different depths of the generation process, and they check the results against human players in both directions.

The paper is short (about eight pages before references and appendices). The game is played in Italian.

### Data and models

The data are 194 Italian Taboo cards transcribed by hand from a physical copy of the game. Each card has a target word and five forbidden words.

Two open-weight models are tested. Gemma 3 4B IT is a small instruction-tuned model, chosen partly because sparse autoencoders trained on it are publicly available. gpt-oss-20b is a mixture-of-experts reasoning model, tested with its reasoning trace switched on and switched off. Every description was generated with greedy decoding and a 30-word limit stated in the prompt. In the reasoning mode, gpt-oss had a budget of 2,048 tokens and a mild repetition penalty.

### The conditions

The study compares a baseline with four ways of enforcing the ban.

In the baseline, the model is asked to describe the target word and is told nothing about forbidden words. This measures how often each model would reach for the forbidden words on its own.

In the prompting condition, the prompt lists the forbidden words and tells the model to avoid them and any word sharing their root. This is the condition that matters for the skill: the forbidden words are named in the instruction the model reads just before it writes. The prompt does not tell the model to avoid the target word itself, so the authors can measure how often it gives the answer away.

The a priori constrained condition masks, before generation, every vocabulary token whose surface form shares a prefix with a stem or lemma of a forbidden word, so those tokens cannot be produced. The appendix shows this can be blunt. One card's forbidden phrase contained two common function words, and the mask grew to more than 2,000 tokens, including ordinary prepositions and articles. Across the dataset the mask averaged about 470 tokens per card.

The online constrained condition lets the model write freely and checks each completed word against the forbidden stems. When a word matches, generation backs up to the start of that word, blocks the offending token at that position and continues. This applies the ban only where a violation actually occurs.

The steering condition, run on Gemma only, does not mention the forbidden words in the prompt at all. For each forbidden word it finds the sparse-autoencoder feature most active when the model reads that word, then pushes the model's internal state away from those features while it writes.

### Measures

Compliance is measured by automatic detection of exact forbidden words, of words sharing a forbidden stem or lemma, of leaked target words and, for the reasoning mode, of runs that used up the token budget without an answer. Accuracy is the share of descriptions with none of these errors.

Informativeness is measured two ways. Each model acts as guesser on descriptions from both models; the authors read the guesser's next-token probabilities and record whether the first token of the target word falls in its top 1, 2, 3, 5 or 10 guesses (reported as pass at k). Eight human annotators also took part. In a guessing phase on 80 cards, they guessed targets from model descriptions, with each card seen under two conditions by two annotators. In a describing phase on 100 other cards, they wrote spontaneous and deliberate descriptions, which other annotators and the models then tried to guess.

### Results on compliance

Without any ban, the models used forbidden words often. Exact forbidden words appeared in 52 percent of Gemma's baseline descriptions, 55 percent of gpt-oss's without reasoning and 46 percent of gpt-oss's with reasoning. Strict accuracy was 34 percent, 25 percent and 34 percent.

Naming the forbidden words in the prompt cut their use sharply. Exact violations fell to 2.6 percent for Gemma (from 52.1 percent), 16.0 percent for gpt-oss without reasoning (from 55.2 percent) and 1.0 percent for gpt-oss with reasoning (from 45.9 percent). Strict accuracy rose to 91 percent, 72 percent and 97 percent. Reasoning made a large difference for gpt-oss, which the authors attribute to the reasoning trace actively policing the ban. With reasoning, almost all the remaining errors were related word forms.

Masking tokens at decoding time eliminated exact violations by construction. Gemma reached 98.5 percent accuracy with the a priori mask and 96.4 percent online. For gpt-oss, accuracy under the masks was 82 to 92 percent, and the remaining failures were mostly leaked target words, which rose to 17.5 percent without reasoning. Blocking the obvious clues seems to push the model toward naming the answer.

Steering without naming the words reached about 56 percent accuracy for Gemma, above the 34 percent baseline and far below prompting.

### Results on informativeness

The two model guessers differed a great deal. gpt-oss almost never placed the target in its top 10 (never above 10 percent of cards), while Gemma reached about 25 percent. Both agreed on the broad order of conditions.

With Gemma as clue-giver and Gemma as guesser, the target was in the top 10 guesses for 24.7 percent of baseline descriptions, 17.0 percent of online-constrained ones, 17.5 percent under the a priori mask, 21.1 percent under steering and 11.3 percent under prompting. With gpt-oss as guesser, the order was similar at lower levels (9.3 percent baseline, 8.8 percent online, 3.6 percent prompting). The authors propose that a model told in its instructions to avoid words becomes cautious and writes sparse descriptions, while a model blocked token by token is pushed toward less obvious but more precise paraphrases. Steering kept descriptions informative but at a cost in compliance, and it often made Gemma switch languages (in 45 percent of the descriptions humans saw).

Human guessers showed the same order. Annotators identified the target 38.1 percent of the time from constrained-generation descriptions, 31.0 percent from steered ones and 23.4 percent from prompted ones, about 31 percent overall. By my arithmetic these shares correspond to roughly 24 of 63, 9 of 29 and 15 of 64 descriptions, so the bases are small.

Humans were far better guessers than either model. Human descriptions were guessed correctly by other humans 85 to 95 percent of the time, while Gemma needed 5 to 10 guesses to approach what a human did in one. Spontaneous human descriptions were easier for humans to guess than deliberate ones (95 against 85 percent) and harder for models (about 5 percent against 14 percent at top 10, averaged over both guessers).

### Strengths

The design separates three levels of intervention cleanly and includes a no-ban baseline, so the effect of naming the words can be measured against the models' natural use of them. The compliance measures are automatic and well defined, with exact and related-form errors kept apart. The human study runs in both directions and grounds the model-guesser results. The authors publish code and data and report full tables per model and condition.

### Weaknesses and limits

The study is small and narrow. It uses two models, of 4 and 20 billion parameters, in Italian only, on 30-word descriptions, with greedy decoding and one prompt wording. Steering was tested on one model. The guesser models are weak, so most informativeness numbers sit in the single digits or low teens and small differences may be noise. The paper reports no confidence intervals or significance tests.

The informativeness comparison between baseline and banned conditions is not like for like, because baseline descriptions are free to use the most obvious clues. The fairer comparison for the skill's purpose is between prompting and the decoding-time conditions, which all obey the ban.

Several of the paper's own numbers disagree. Steering accuracy is given as 55.7 percent in one place and 56.7 percent in another. In the comparison on 10 shared cards, the text gives top-10 rates for two proprietary clue-givers that do not match the figure, and the figure's row for Claude Sonnet 4.6 contains values that cannot arise from 10 cards and appear to repeat another table's row. The text says Gemma reached 30 percent on human descriptions where the figure shows 40 percent. One table reference in the prompting section is unresolved in the copy read. These problems sit in the secondary analyses and do not touch the main compliance tables.

### Principles the paper implies for writers

Naming a forbidden word in a model's instructions makes the model avoid it reliably, especially when the model can reason before answering. The cost shows up elsewhere: an instructed model plays safe and says less. Enforcing the ban on output, after or during writing, keeps more of the content.

## What the skill applies

The paper supports step 2 of "How to use this skill" in SKILL.md in two ways.

First, it argues against the rationale the skill used to give for that step. An earlier version of the skill said that reading a list of banned words before drafting tends to put those words in the writer's head. If that were true of these models, naming the forbidden words should have raised their use or at least failed to lower it. It did the opposite: exact use fell from about half of all descriptions to between 1 and 16 percent, depending on the model. The skill no longer makes the priming claim, and this paper is one of the reasons. The skill's current wording for step 2 says nothing about banned words coming back.

Second, the paper supports the reason the skill now gives. Descriptions written under a named ban were the least informative of the banned conditions, for both model guessers and for human guessers, while the online condition (write freely, then fix each violation as it appears) kept more of the content. That matches the claim in step 2 that constraints applied while drafting tend to make a model plan a thinner answer, and that drafting freely before revising keeps the content. The online condition is a decoding procedure in which the writer never revises anything, so the match to section 10 is an analogy. It is the same order of work, though: generate first, then remove what breaks the rule.

The compliance results also support keeping explicit lists in section 10 at all. A named check works on the named items, so the revision pass names the words, phrases and marks it removes instead of describing them in general terms.

## What the skill omits

The skill leaves out the decoding-time masking and the steering method. A writer cannot mask tokens or edit internal features. They matter here only as comparisons.

It leaves out the game-specific results: leaked targets, the human and model guessing gap and the contrast between spontaneous and deliberate human descriptions. None of them bears on prose style.

It does not cite the paper's rates. Two small models writing 30-word Italian clues are a long way from a frontier model writing a report, and the skill's rules are meant to hold across models and lengths.

It does not adopt the finding that reasoning improves compliance as a rule. Whether the writer reasons before answering is set by the system running it, and a style guide cannot change that.

## Conflicts with other references

one-token-away-from-collapse.md is the main support for step 2, and this paper agrees with it from a different setting. That paper found that small surface bans placed in the prompt cut long English answers to about half their content while the bans themselves were obeyed, and that rewriting a free draft under the ban recovered most of it. This paper finds high compliance and sparser output under a named ban in short Italian descriptions. Neither paper found that naming a banned item makes the model produce it.

The reference that the skill previously relied on for the priming claim aged out under the eight-month rule in AGENTS.md and was removed. Its finding, from a single model on one-word answers, was that naming a forbidden word can raise the chance of producing it when the model's natural pull toward that word is strong. This paper's prompting results do not show that effect for whole descriptions, though its design does not test the one-word case directly. The skill no longer depends on either reading.

freeburg-em-dashes.md found that an explicit em dash ban lowered the rate in every model tested, without always reaching zero. That agrees with this paper's compliance results, and the two together are the folder's evidence that naming a target in a check works on the target.

The style manuals (nyt-manual-of-style.md, ap-stylebook.md, chicago-manual-of-style.md) and the technical guides put their lists of words to avoid in front of the writer. With the priming concern gone, the remaining reason the skill does not copy that layout is the loss of content at drafting time described here and in the One Token paper. The manuals are written for human editors who look entries up while revising, which is closer to the skill's revision pass than to its drafting stage, so there is no real conflict.

anthropic-prompting-docs.md favors telling a model what to do in place of what to avoid. This paper's prompting condition is a pure prohibition, and it produced high compliance with thinner output. That is consistent with the docs' advice that negative framing is a weak way to steer generation, if the measure is the quality of the answer and not the avoidance of the word.

## Review notes

On 2026-11-05 I read the full text of arXiv 2607.00601 from a local text copy, covering the abstract, all six sections, the acknowledgments and AI-use declaration, the reference list and appendices A through E, including the prompt templates with their English translations, generation settings, the comparison of the two masking methods, the steering details, the failure examples and the three per-model result tables. Figures were available only as extracted labels and captions, so values that appear only in figures (such as the 10-card comparison) were read from the extracted numbers, which may have lost their layout. I did not open the code repository or rerun anything.

The version 1 date of 1 July 2026 comes from the review brief, and the identifier's July 2026 prefix agrees with it. I do not know which version the local copy is. The paper's header names CLiC-it 2026, but I could not confirm that the arXiv text is the version accepted there, so I list the paper as a conference paper posted as a preprint without claiming it passed review in this form.

The raw counts behind the human guessing shares are my own arithmetic: they are the smallest whole-number fractions that give the reported percentages, and they sum to 156 against the 160 guessed entries the paper reports, so the split is my reconstruction. The averages for spontaneous and deliberate descriptions guessed by models are my averages over the two guessers in the paper's table. The argument that these results count against the priming rationale is my reading. The paper does not discuss priming or style guides, and its no-ban baseline measures the models' natural use of the forbidden words, not their use after merely seeing them mentioned without a ban. A design that names the words without forbidding them would test priming more directly.
