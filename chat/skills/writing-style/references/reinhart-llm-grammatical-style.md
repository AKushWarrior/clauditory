# Do LLMs Write Like Humans? Variation in Grammatical and Rhetorical Styles

- **Source:** Reinhart, A., Markey, B., Laudenbach, M., Pantusen, K., Yurko, R., Weinberg, G., and Brown, D. W. "Do LLMs write like humans? Variation in grammatical and rhetorical styles." Proceedings of the National Academy of Sciences 122 (2025), e2422455122. Preprint: arXiv 2410.16107.
- **Type:** peer-reviewed journal article
- **Where to find it:** DOI 10.1073/pnas.2422455122; preprint at https://arxiv.org/html/2410.16107v2
- **Last reviewed:** 2026-11-05
- **Used by:** chat/skills/writing-style/SKILL.md

## What the source covers

### Purpose and framing

The paper asks how large language models differ from people in the grammar and rhetoric of the text they produce, rather than in vocabulary alone. Earlier work on machine-written text mostly looked at word choice, punctuation, sentence length, or detector accuracy. The authors instead apply the feature framework of the linguist Douglas Biber, which was built to describe register and genre variation in human writing and speech. Their claim is that this lens shows what an LLM does to a text's style, and that those differences are systematic across models, persist as models grow, and are larger after instruction tuning. The significance statement adds that instruction-tuned models keep the dense style even when asked to match informal speech and writing. They also report that instruction tuning limits a model's ability to shift style to suit a genre, which they call genre misalignment.

### Design: parallel corpora

The central design choice is a parallel corpus. Each human text was cut at a sentence boundary into two chunks of roughly 500 words. A model received the first chunk as context and a style sample, and was asked to write about 500 more words in the same style, tone, and diction. The model's continuation is then compared with the real second chunk of the same document. This design gives every human passage a matched passage from each model, written from the same starting point, which controls for topic and genre far better than comparing unrelated human and machine corpora.

Two corpora were used. HAP-E (Human-AI Parallel English) was built from 12,000 human texts across six categories: academic writing, news, fiction, spoken word such as podcast transcripts, blogs, and TV and movie scripts. After removing refusals and unusable short outputs, it holds 8,290 texts, or 66,320 chunks counting the human second chunk and six model continuations per text. HAP-E was the main analysis corpus. CAP (COCA AI Parallel) was built from the Corpus of Contemporary American English, which spans eight registers (spoken, fiction, magazines, newspapers, academic, blogs, web pages, and TV or movie subtitles). After the same filtering it holds 9,615 texts and 76,920 chunks. CAP serves as a check that the HAP-E findings generalize to a differently built sample. A third, smaller check used part of the M4 corpus, where arXiv abstracts are paired with abstracts that GPT-3.5 generated from their titles. That check tests a different model and a single academic genre.

### Models

Six models were used. From OpenAI there were GPT-4o and GPT-4o Mini, accessed through the API, with no base versions available. From Meta there were four Llama 3 variants: the 8B and 70B base models, and the 8B and 70B instruction-tuned models. The Llama set is what makes the study informative about causes, because the same pretrained network appears with and without instruction tuning. A base model continues text; an instruct model has been further trained to follow instructions and answer questions. The study therefore has a clean base versus instruct contrast only within Llama 3, and a size contrast (8B versus 70B, GPT-4o versus Mini) within each family.

### Features and statistics

The linguistic measurements come from Biber's tagset of 66 features, spanning lexical complexity indices and grammatical categories. Examples include mean word length, nominalizations, agentless passives, hedges, various subordinate and coordinated clause types, and downtoners. The paper names the tagging tool (pseudobibeR) only in its discussion. Figure 3 plots the 15 features the random forest ranked most important, which besides the four above include mean word length, adverbs, other nouns, clausal coordination, downtoners, attributive adjectives, demonstratives, agentless passives, postnominal past participles, infinitives, and prepositional phrases. Rates there are shown relative to the human rate on a log scale, with significance marked. Differences were tested with paired Wilcoxon signed-rank tests with Bonferroni correction, and effect sizes are paired Cohen's d values. The pairing follows directly from the parallel design: each model continuation is compared with its own human counterpart.

### Main findings: grammatical style

The strongest and most consistent result is that the four instruction-tuned models use a dense, noun-heavy grammar far more than people do. The features that separate them most from human text are present participial clauses (the "-ing" modifier tacked onto a main clause), "that" clauses serving as the subject of a sentence, nominalizations (nouns derived from verbs or adjectives), and phrasal coordination (joining words or phrases with "and" or "or" rather than joining clauses).

The paper gives GPT-4o's rates against the human paired text. Present participial clauses appear about 5.3 times as often (d about 1.38). "That" clauses as subjects appear about 2.6 times as often (d about 0.77). Nominalizations appear about 2.1 times as often (d about 1.23). Phrasal coordination appears about 1.9 times as often (d about 0.81). Across all instruction-tuned models, the introduction summarizes the range as two to five times the human rate for participial clauses and about 1.5 to 2 times for nominalizations. The agentless passive, a passive with no named actor, moves the other way: GPT-4o uses it at roughly half the human rate. This difference matters for any reader who assumes machine prose is passive. In these data it tends toward participial and nominal constructions and away from agentless passives.

Models also differ from each other. Both GPT-4o models avoid clausal coordination (joining full clauses), while every Llama 3 variant uses it more than humans do. Both GPT-4o models use downtoners such as words meaning barely or nearly more than humans, while every Llama 3 variant uses them less. So no single "AI style" covers all models on every feature, though the instruction-tuned models share the dense-information profile.

Size does not repair the gap. Larger models show the same direction of difference as their smaller siblings, and the authors observe that scaling up does not appear to make text more human-like on these measures, even though the larger models do better on other benchmarks. Their reading of the same-family confusion in the classifier is that size alone does not produce a dramatic style difference. The Llama base models sit much closer to human rates on most features, while the instruct models and GPT-4o spread widely across the feature set.

### Main findings: vocabulary

The paper also compares word rates, using lemmatized words that occur more than once per million words in the human text. The word-rate comparison (Figure 4) plots usage per 1,000 words for each model against the human second chunk, with bands marking ten times more and ten times less than human use. Instruction-tuned models overuse a recognizable set of words. For GPT-4o the largest ratios include camaraderie (about 162 times the human rate), tapestry (about 155), intricate (about 119), underscore (about 107), unspoken, amidst, palpable, solace, fleeting, and unravel. GPT-4o Mini overlaps heavily, adding grapple, ignite, vibrant, and cacophony. The Llama instruct models favor a different set (unease, palpable, continuation, pang, prioritize, reminder, grapple). Tapestry appears in about 23 percent of GPT-4o outputs and amidst in about 27 percent. Certain obscenities appear more than 100 times less often than in human text.

The authors point out artifacts. The word continuation is inflated in the Llama instruct models because many outputs open with a line announcing the continuation. A first name that tops one Llama base list comes largely from a single document that repeats it hundreds of times. The Llama base models' overrepresented words are mostly proper names and topical oddities, which reads as noise rather than a style signal. That contrast is informative: word choice is most human-like in the base models, so the biased vocabulary appears to come from post-training and not from the pretraining text. The authors add that many of the overused words are ordinary in fiction but conspicuous in academic, news, or scripted writing, and that in GPT-4o several of them express abstract relations among things, which the authors connect to grand but empty summarizing sentences.

### Main findings: classification

A random forest trained on the 66 features to identify the source among seven (human plus six models) reached 66 percent accuracy on a held-out quarter of HAP-E (14,535 texts), against about 14 percent by chance. Most errors were between models of the same family at different sizes, such as 8B versus 70B or GPT-4o versus Mini. Confusion between humans and models was low: 4.2 percent of model texts were labeled human and 9.8 percent of human texts were labeled as model-written. Pairwise classifiers (one model against humans) reached 93 to 98 percent accuracy, including when trained on one corpus and tested on the other. Lasso logistic regression did about as well for most models, but only reached about 75 percent for the Llama base models, which the authors read as a sign that those models differ from humans through interactions among features, not additive shifts in single ones. The summary pattern is that instruction tuning makes text easier to tell apart from human writing.

Generalization is weaker. Classifiers trained on HAP-E and applied to the M4 abstracts lost accuracy. Those built for instruct models still beat chance on GPT-3.5 text, while those built for Llama base models fell to chance (50 percent). Instruct models share detectable habits, but the habits do not transfer cleanly across models and registers.

### Interpretation: why instruction tuning shifts style

The base-versus-instruct contrast in Llama 3 is the paper's main causal evidence. Because the base models stay near human rates and the instruct versions depart from them, the authors conclude that the stylistic drift comes from the instruction-tuning stage and not from what the pretraining text contained. They offer candidate mechanisms: preferences expressed by human raters during tuning, or tuning tasks such as summarization that reward compressed, information-packed prose. They say plainly that the tuning recipes are not public, so the mechanism cannot be pinned down. The interpretation of why participial clauses and nominalizations would be rewarded is therefore a hypothesis, not a result.

### Rhetoric, genre, and implications for writers

The title promises rhetorical analysis, but the evidence is grammatical. The paper's rhetorical argument is that the dense features it measures are associated with informational prose and with human judgments of mastery, and that models apply that register everywhere. The genre misalignment claim is stated in the introduction and discussion: models do not shift style with context the way people do, so their output fits poorly with the conventions of academic writing, spoken interaction, or news. The authors warn that heavy reliance could yield awkward, confusing, or unpersuasive text in creative, instructional, and argumentative writing, and that novices may skip the work of fitting a text to its audience. They explicitly say they are not proposing a detector or policing student work, and describe the differences as teachable moments for revision.

### Strengths and weaknesses of the evidence

The parallel design is a real strength, as are the large samples, paired tests, two corpora, and a held-out generalization check. The inclusion of base and instruct versions of one model family gives an unusual chance to separate causes.

There are limits. The main text pools results across genres. It reports no feature rates or classifier results by category, so the genre misalignment claim rests on interpretation and on the pooled pattern, not on a tested genre-by-model comparison. The study does not collect reader judgments, so the statements about awkward or unpersuasive prose are untested. Prompts were fixed to one continuation task, and the full prompt wording and sampling settings are in the supplement, which I could not read. A different task, such as answering a question or writing to a brief, might yield different rates. A continuation prompt that tells the model to imitate a given passage is also not how most people use these tools, and it may understate or overstate default style. Two models are closed (GPT-4o and Mini), and only one open family has base versions. The models are 2024 vintages, and later releases may have shifted. Participial and nominal rates are descriptive: the paper does not show that readers dislike them. Finally, a few headline vocabulary ratios are inflated by one-off artifacts, so the word list is best read for its shape (abstract, literary, emotionally charged words) rather than for specific multipliers.

### What the findings imply for writers

The discussion contrasts experts, who can spot and revise misaligned output (the authors' examples are a therapist collating notes and someone writing thank-you notes), with learners who hand over the work of shaping a text for an audience, and it says the models are not bad technically or morally. The discussion also lists passive voice among the distinguishing features and says the findings reproduce earlier work showing more nominalizations and fewer human subjects and epistemic stance markers in ChatGPT-4 text.

For someone drafting with or editing model output, the paper identifies concrete grammatical habits to look for: sentences that end with a trailing participial phrase, abstract noun phrases standing where a verb would do, "that" clauses as subjects, paired nouns or adjectives joined by "and," and a literary vocabulary out of place in plain reporting. It also implies something about what to avoid over-correcting. Model text is not passive-heavy, so a rule that simply bans the passive aims at the wrong target.

## What the skill applies

The source shapes four parts of SKILL.md.

Section 3 (Sentences), the rule about ending a finished sentence instead of hanging a trailing clause on it, and the matching item 3 of the revision pass in section 10 (trailing participles). The participial-clause finding is the strongest and most replicated result in the paper, at several times the human rate. The skill's wording names the specific forms (a clause beginning with an "-ing" verb after a comma) and tells the writer to cut it, name the relation it leaves implicit, or give it its own sentence with a concrete subject. The paper supplies the evidence that this habit is both frequent in model prose and a marker of it. Since version 2.0.0 the item covers every trailing participle, which matches the grammatical form the paper measures. The paper itself says nothing about commentary function (highlighting or underscoring significance); the item's examples of that kind come from the other references, and this source only establishes the grammatical form and its frequency.

Section 2 (Words), the rule to prefer verbs to nouns built from verbs. The nominalization result, about double the human rate for GPT-4o, supports the skill's instruction to write "we decided" rather than "a decision was made." The skill's stated cost, that noun-heavy prose hides who did what, is the skill's own reasoning, but the frequency data justifies listing it among the defaults to correct rather than as a stylistic preference.

Section 10, item 8 (overused vocabulary). The paper's lists of overrepresented words, including tapestry, intricate, and underscore, are among the supporting evidence for those entries. The skill does not use the paper's multipliers, and treats this paper as one of several vocabulary sources.

Section 3, the rule to keep a series parallel, and section 10, item 4 (reflexive triplets). The phrasal coordination finding, about double the human rate, supports giving attention to pairs and series that join nouns or adjectives by "and." The paper does not measure triplets or the usefulness of a third item, so the triplet rule rests on other sources; this one only shows that coordinated phrases are a model habit worth watching.

Section 1 (Open with the answer) and section 4 (numbers and evidence) do not draw on this paper.

## What the skill omits

The skill leaves out all of the classifier material. Detector accuracy, feature interactions in the lasso comparison, and cross-corpus transfer describe how separable model text is, not how a writer should write. Including them would also invite the skill to be read as a detection tool, which the authors themselves disclaim.

It omits the full Biber feature set. Most of the 66 features, such as downtoners, clausal coordination, and hedges, differ in direction between GPT-4o and Llama, so a single rule about them would be correct for one family and wrong for the other. The skill targets only the features that are consistently elevated across the instruction-tuned models.

It omits the vocabulary multipliers and model-by-model word lists. They are partly artifacts (continuation, a repeated first name), they reflect 2024 model versions, and the skill's own list in section 10 is deliberately short and kept only in the revision pass.

It omits the genre misalignment argument as a rule, because it is an interpretation without a per-genre test in the main text. The skill does state in its principles that a writer should match register to reader, but that comes from the style manuals, not from this paper.

It omits the causal story about instruction tuning. That story explains why models write this way but does not change what the skill tells an agent to do. It is also a hypothesis the authors could not verify.

It omits the agentless-passive finding as a basis for any rule. The skill defaults to active voice for accountability reasons from other sources, not because this paper says models overuse the passive. They do not.

## Conflicts with other references

The one real tension is with the voice guidance in rhodes-active-passive-voice.md and with SKILL.md section 3, which defaults to active voice. A reader might assume that advice targets a machine habit. This paper shows the reverse for GPT-4o: it uses agentless passives at about half the human rate. The skill therefore cannot justify its active-voice default as a correction of typical model output. It rests the default on accountability (naming who acted) and on the allowance for passives when the actor is unknown or unimportant, and Rhodes finds no comprehension difference by voice. The skill keeps the default and does not cite this paper for it.

A second, milder tension is with the vocabulary papers (kobak-excess-vocabulary.md, juzek-ward-lexical-overrepresentation.md, liang-ai-modified-peer-reviews.md). This paper measures rate ratios against paired human text from a continuation task in 2024 models. Those papers use different corpora and different measures, so their word lists and magnitudes are not comparable. The overlap is qualitative: words such as intricate and underscore appear in several lists. The skill treats overlap across measures as supporting evidence for including a word and does not compare magnitudes.

Against the style manuals, there is no direct conflict. The manuals do not address participial clauses or nominalization in measured terms, though NYT and AP's general advice to prefer plain verbs agrees with the direction of the skill's rule.

## Review notes

I read the full text of the arXiv v2 manuscript (compiled August 25, 2025, arXiv date August 21, 2025) on 2026-11-04, and rechecked every number in this file against it. The file I was given is a draft-marked manuscript with a placeholder PNAS DOI and volume, so I could not confirm the published citation details (volume 122, article e2422455122) or whether the PNAS version differs from this text; the DOI and article number come from the task brief. The supplementary appendix (Tables S1 to S12, Figure S1, exact prompts, sampling settings, corpus composition by genre) is not part of the text I read, so claims that depend on it are marked as unavailable. The main text pools results across genres and reports no per-genre figures, which I confirmed in the full text. Early passes of this file relied on a model-written digest; the corrections after the full read were the complete author list, the significance-statement point about informal styles, the Figure 3 feature list and Figure 4 details, softer wording on size effects, and the discussion details on experts and learners. All headline figures matched the digest. The rate ratios and effect sizes quoted are GPT-4o's only; other models' values are shown in the paper only graphically.
