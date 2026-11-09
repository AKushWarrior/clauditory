# Gude and colleagues on the grammar and lexicon of two generations of LLMs

- **Source:** Adrián Gude, Roi Santos-Ríos, Francis Bond, Dan Flickinger, Carlos Gómez-Rodríguez, and Olga Zamaraeva, "More Aligned, Less Diverse? Analyzing the Grammar and Lexicon of Two Generations of LLMs," arXiv preprint arXiv:2605.06030 [`cs.CL`]. Version 1 was submitted May 7, 2026.
- **Type:** preprint (not peer-reviewed)
- **Where to find it:** https://arxiv.org/abs/2605.06030 (HTML at https://arxiv.org/html/2605.06030v1)
- **First published:** 2026-05-07 (on or after the 2026-03-08 cutoff for sources that make claims about AI systems)
- **Last reviewed:** 2026-11-08
- **Used by:** `chat/skills/writing-style/SKILL.md`

## What the source covers

### Purpose

The paper asks whether newer language models write news text that matches human grammar and vocabulary better than older models do. Most earlier comparisons used surface counts such as word frequency or perplexity. The authors instead parse every sentence with a formal grammar and compare the distributions of grammatical constructions and word types. Their question has a practical edge. Instruction tuning made models more coherent and more obedient, and the authors want to know what it cost in range.

### Data

The human baseline is the lead paragraphs of New York Times articles from two periods, a sample from 2023 and a sample from February to May 2025. The authors chose the 2025 sample because the models they tested were unlikely to have seen it in training. Human text is stable across the two periods, which makes it a fair fixed reference.

Two model cohorts wrote the comparison text. The 2023 cohort is base models released before October 2023: LLaMA at four sizes, Falcon 7B, and Mistral 7B. The 2025 cohort is instruction-tuned models trained on data through 2024: Qwen 2.5 at three sizes, LLaMA 3.3 70B, Mistral 7B v0.3, and GPT-4o. Base models received the headline and the first three words of the lead and continued from there. Instruction-tuned models received a system prompt that described a professional journalist and a user prompt that asked for one multi-sentence lead paragraph with no title, bullets, or preamble. Sampling used a temperature of 0.7 and a cap of 1,000 new tokens. The largest models ran at reduced numerical precision.

### Method

Each sentence is parsed with the English Resource Grammar, a hand-built grammar in the HPSG formalism that tries to accept all well-formed English and reject the rest. The grammar assigns each sentence its construction types and each word its lexical type. The authors report a parse rate near 94 percent on varied English. They measure diversity with two indices borrowed from ecology, Shannon-Wiener and Simpson, applied to the distribution of construction types and lexical types. They also record the average sentence length in tokens, the share of sentences of 15 tokens or fewer, the share of fragments, and the processing cost of each parse.

### Findings on sentence length and fragments

The finding most relevant to the skill concerns sentence length. Human sentences in both periods average about 22 tokens, and a third of them (33 percent in 2023, 32 percent in 2025) run 15 tokens or fewer. The 2023 base models write slightly shorter sentences, 18 to 19 tokens on average, with 35 to 39 percent of them short. The 2025 instruction-tuned models write sentences of about 26 to 30 tokens, and only 1 to 4 percent of their sentences are short. The text of the paper states the reduction in short sentences as a factor of 9 to 30. My own division of the table values gave a range of about 8 to 32, so the exact factor depends on rounding, and the direction and size are not in doubt.

Fragments follow the same pattern. Humans produce fragments in 12 to 13 percent of sentences, and the 2023 models in 10 to 13 percent. The 2025 models produce them in about 2 percent. The paper calls this a fivefold drop; the table suggests closer to sixfold.

Parsing results back the picture of regular, tidy text. The 2025 model sentences parse successfully more often (about 98 percent against about 94 percent for humans) and cost less to parse, and far fewer of them exceed the parser's resource limits. The authors read this as structural regularity, not greater complexity.

### Findings on diversity

Human news text is the most varied in construction types. The 2025 models form a lower band, and the 2023 base models sit above them. The authors link part of the gap to newer models avoiding specifics such as names, dates, and measures, and note that the 2025 models favor modification and coordination. Lexical diversity shows the sharper change. The 2023 models rank highest, humans fall in the middle, and the 2025 models rank lowest. The most distinctive word types in 2023 output were utterance particles and personal pronouns. In 2025 output they were plain staples such as adjectives, count nouns, and transitive verbs. Punctuation variety also shrank: the newer models, GPT-4o in particular, used few em dashes, semicolons, and quotations. The authors guess that post-training aimed at reducing false statements played a part, and they label that guess as a hypothesis.

The diversity values appear only in figures, so the text gives no exact index numbers.

### Limits the authors state

The study covers one genre, English news leads. Output depends on prompts and sampling. Quantization may have affected the largest models. The grammar parses about 94 percent of ordinary news text, so some sentences drop out. By my reading, the comparison also confounds model generation with prompt format, because base and instruction-tuned models received different prompts. The authors treat the claim that instruction tuning causes the loss of diversity as a hypothesis for later work.

## What the skill applies

**Section 3 (sentences).** The planned rule says to write short sentences often and to let sentence length range widely. The paper supplies the clearest measured gap behind it: models of 2025 write longer sentences than professional journalists and almost never write the short ones that make up a third of human news prose. The planned wording that model prose clusters at 25 to 30 words comes from the average lengths in this paper. The existing guidance to mix short and medium sentences stays, and the new rule makes the short end explicit.

**Revision pass item 15 (flat rhythm).** The paper gives the reason to check a paragraph whose sentences all run about the same length. A paragraph with no short sentence is the pattern the 2025 models produce. The bundled linter's `House.Rhythm` rule tests it, using the same 25 to 30 word clustering as its premise.

**Section 6 (length).** The paper supports the planned rule only indirectly. Longer sentences add length to a text, but the study fixed the output to one paragraph and did not measure total length.

**Section 2 (word choice).** The finding that newer models reach for basic word categories and lose lexical range is consistent with the planned rule to choose the exact, concrete word over the expected one. It does not prove the rule, because the paper measures word types, not how readers react to them.

## What the skill omits

The skill leaves out the formal-grammar method, the diversity indices, and the parsing costs, since an agent writing a reply cannot compute them. It does not copy the human averages as targets. A writer who aims at 22 tokens per sentence would produce a new flat rhythm, so the skill asks for a wide range and short sentences without a target. It omits the claim about instruction tuning as the cause, which the authors themselves leave open, and the guess about post-training and punctuation. It also leaves out the comparison between the 2023 base models and everything else, because an agent cannot choose to be a base model. The lexical diversity result is used only as supporting context for the word-choice rule, not as a rule of its own.

## Conflicts with other references

The paper agrees with `reinhart-llm-grammatical-style.md`, which found that instruction-tuned models pack text with long noun-heavy constructions and that base models sit closer to human rates. Gude and colleagues reach the same split from sentence length and diversity, using different measures and a different genre.

On the em dash, the paper's remark that GPT-4o produced few dashes fits `freeburg-em-dashes.md`, where GPT-4o scored well below GPT-4.1 and several other models. It does not undercut the skill's flat em dash ban, since Freeburg shows rates vary widely by model and release.

The paper's lexical result points the opposite way from `van-nuenen-voice-under-revision.md`, where rewrites showed a richer vocabulary than the originals. The two studies differ in task (open generation of news leads against rewriting personal stories), in genre, and in the measure used, so they do not contradict each other. The skill takes the practical point from both: model output moves away from the individual writer's range.

The paper's finding of narrowed vocabulary sits beside `kobak-excess-vocabulary.md` and `juzek-ward-lexical-overrepresentation.md`, which concern a few overused words, not overall diversity. The skill treats the two as complementary. The word list handles the marked words, and the exact-word rule handles the quieter drift toward the expected one.

I resolved no direct conflict with the style manuals in this folder. They advise varied sentence length in general terms, and this paper adds a measured gap. I read the other files only in part (headers and sections relevant to these points), so the comparisons should be rechecked against the full texts.

## Review notes

I read the arXiv abstract page and the HTML full text of version 1 through a fetch tool that returns a summary of the page. That reading covered the data, prompting, generation settings, methods, the sentence length and fragment table, the parsability results, the diversity discussion, the punctuation remark, and the limitations. I did not open the PDF, the supplementary material, or the cited sources. Version 1 is the only version listed.

Two numbers did not reproduce exactly. The paper's text gives a 9 to 30 times drop in short sentences, while dividing the table values gives about 8 to 32. The text gives a fivefold drop in fragments, and the table suggests about sixfold. I report the text's figures as the paper's and mention the table's range. The diversity indices appear only in figures that I did not inspect, so I describe only their ordering, as the text does. The task brief's figures (humans near 32 to 33 percent, 2025 models 1 to 4 percent) match the table.
