# Bakhshi on rhetorical intensity versus epistemic grounding in LLM prose

- **Source:** Asim D. Bakhshi (National University of Science and Technology, Islamabad), Saying More Than They Know: A Framework for Quantifying Epistemic-Rhetorical Miscalibration in Large Language Models. arXiv preprint 2604.19768, version 1, 2026. The arXiv record shows submission on 27 March 2026; the PDF carries a date stamp of 23 April 2026. The record's comments field says the paper is under review at an Elsevier journal on writing assessment.
- **Type:** preprint (not peer-reviewed)
- **Where to find it:** https://arxiv.org/abs/2604.19768 (PDF at https://arxiv.org/pdf/2604.19768, 19 pages with appendix)
- **Last reviewed:** 2026-11-04
- **Used by:** chat/skills/writing-style/SKILL.md

## What the source covers

### Purpose and claim

The paper argues that large language models write with more rhetorical force than their evidence supports. Bakhshi calls the gap epistemic-rhetorical miscalibration. His starting point is that most work on LLM bias looks at what models say (stereotypes, false facts, toxic content) and neglects how they position their claims: how confident, how elaborate, how hedged. He says no earlier study had measured rhetorical intensity and epistemic stance together as one quantity, and he builds a framework to do it. The paper is explicitly exploratory. It describes itself as philosophically inspired and offers the metrics as a screening tool and as features for detecting machine-written text.

The theoretical case rests on three traditions. Grice's cooperative principle says the weight and elaborateness of an assertion should match what the speaker actually knows. Relevance Theory adds a processing-cost argument: a reader spends effort on every rhetorical flourish and expects a matching gain in meaning, and a sentence that makes the reader track a long enumeration to learn something already accepted gives a poor return. Brandom's inferentialism treats an assertion as a commitment, so an overclaim hands the reader a premise the writer has not earned and will never be held to. The paper illustrates this with an invented pair of sentences about sleep deprivation (one hedged and proportionate, one categorical and escalating) and an invented LLM-style policy sentence with a six-item enumeration and three stacked adjectives that delivers a platitude. These illustrations are the author's own constructions, not data.

### The taxonomy

The framework, called the ERM taxonomy, has three levels. Level 1 covers ten rhetorical devices in three groups. Sentence-level devices are tricolon (three parallel units in a row), anaphora (repeated openings), chiasmus (AB-BA inversion), erotema (a rhetorical question that asserts) and sententia (a self-contained aphoristic statement that closes an argument with general present-tense force). Argument-level devices are correctio (a mid-argument self-correction that retracts or intensifies an earlier phrasing), enumeratio (a structured list of three or more supports) and auxesis (clauses of rising intensity). Narrative-level devices are peripeteia (a sudden reversal of the argument's direction) and an expectation-violation-resolution arc.

Level 2 sorts stance markers into genuine and performed. Genuine markers tie a claim to an identifiable evidential base: modal auxiliaries such as must or might, sentence adverbials such as probably, restrictor phrases such as according to or judging by, and evidential markers that name the type of evidence (we measured, the record suggests). Performed markers, the paper's key addition, take on the sound of uncertainty without saying what the uncertainty is about. There are two kinds. Complexity tokens are stock phrases that call a topic nuanced or resistant to easy answers without saying why or for whom. Meta-epistemic gestures foreground conditionality, such as saying it depends on how one defines a term, without stating the condition.

Level 3 describes how a whole text ends. Synthetic closure resolves competing considerations into one positive conclusion. Aporetic endpoint ends in sustained, openly acknowledged uncertainty and counts as calibrated only when the evidence really does underdetermine the answer. Premature closure asserts resolution before the argument earns it, typically with certainty tokens. Speculative depth is a long chain of nested conditionals that runs beyond the evidence.

### Metrics

Three composite scores come out of the annotations. Form-meaning divergence (FMD) multiplies the per-sentence density of rhetorical devices by the per-sentence density of performed markers and divides by one plus the density of genuine markers, so it is high only when ornament and performed hesitancy coincide with little real grounding. The genuine-to-performed ratio (GPR) divides genuine density by one plus performed density. Rhetorical device distribution entropy (RDDE) cuts each text into windows of up to 50 words and measures how evenly devices spread across them, normalized from 0 to 1. The reasoning is that skilled human writers cluster devices where the argument is under strain, while habit spreads them evenly. Level 3 markers are binary per document and compared as proportions with chi-squared tests.

### Corpus

There are 225 documents, 75 in each of three groups, about 0.6 million tokens in total (roughly 2,700 tokens per document by my arithmetic). Human experts wrote for outlets such as Aeon, the Cato Institute, Stanford Social Innovation Review, The New Yorker, The Guardian, Nautilus, Wired and Psyche, between 2011 and 2022, and averaged about 93 sentences per text. Human non-experts came from Real Life Magazine and Medium, 2016 to 2022, averaging about 97 sentences. Both human groups predate November 2022 so that generative tools could not have assisted. The LLM texts date from 2024 and 2025 and averaged about 105 sentences. They came from four models: GPT-5.1 (31 texts), DeepSeek-V3.2 (19), Claude-Sonnet-4.6 (13) and Gemini-2.5-Pro (12). Each LLM text answered an open-ended argumentative question derived from the topic of a matching expert text, with a 1,500-word target and fixed instructions to take a clear position, acknowledge a counterargument and engage honestly with uncertainty. Topics were philosophy of mind, ethics, psychology, culture, technology policy and organizational communication.

### Annotation

A sentence splitter (spaCy) divided documents into sentences, and a locally hosted LLM split them into argument chunks using a seven-way Toulmin-style typology (claim, grounds, warrant, backing, qualifier, rebuttal, non-argumentative). Annotation then ran as separate passes, each returning yes-or-no judgments per marker. The paper says a trained human or a capable mid-sized LLM could do the annotation; the study itself used an LLM, though the text I read does not name which one beyond calling it capable. The paper states plainly that no formal inter-annotator reliability against human coders was established.

### Findings

Tricolon was the strongest single separator. LLM texts averaged 7.13 per document, against 3.73 for experts and 4.87 for non-experts, an effect size of 0.95 against experts (p below .001). The abstract says nearly twice the expert rate; the raw ratio is 1.91. Because LLM texts were longer in sentences, the per-sentence ratio is smaller. By my arithmetic from the reported means it is about 1.7 times the expert rate, which is still large.

Erotema ran the other way: 2.28 per document in LLM text, against 5.55 for experts and 5.11 for non-experts (effect size about 0.73). The author reads this as occasion-dependent devices being harder for a generator to produce, and concedes that the prompt's register may play a part. Correctio was also lower in LLM text (0.17 per document against 0.40 for experts), a small effect.

Performed hesitancy density was about 0.114 per sentence in LLM text and 0.057 and 0.058 in the two human groups (effect sizes 0.72 and 0.65). Because experts and non-experts matched, the author concludes that performed hesitancy marks machine text and does not mark lack of expertise. Complexity tokens averaged 7.3 per document in LLM text against about 4.7 in human text.

FMD rose in order from experts (0.009) to non-experts (0.012) to LLMs (0.017). The gap between LLMs and experts had an effect size of 0.68 and LLMs differed from both human groups at p below .001; the two human groups did not differ reliably. Genuine marker density was highest in experts (0.283), then LLMs (0.242), then non-experts (0.182), so the LLM excess lay in performed markers, not in missing genuine ones. GPR ordered experts (0.267) above LLMs (0.217) above non-experts (0.172), with the LLM and non-expert difference just short of significance (p = 0.055). RDDE was highest and least variable in LLM text (0.753 against 0.666 for experts, effect size 0.74), which the author takes as evidence of template-driven, evenly spread ornament.

At the discourse level only the aporetic endpoint differed: 24.0 percent of LLM texts, 8.0 percent of non-expert texts and 2.7 percent of expert texts. Synthetic closure was common everywhere (66.7 percent for experts, 69.3 for non-experts, 76.0 for LLMs) without a significant difference. The author reads the LLM aporetic rate as a performed-hesitancy move at document scale, open-ended wrap-ups used reflexively. Across the four models no composite metric differed significantly, which the author treats as evidence that the pattern belongs to LLM generation in general. With only 12 to 31 texts per model, that null result has little power and is better read as an absence of detected differences.

### Strengths and weaknesses of the evidence

The main strength is the contrast design: matched topics, a pre-LLM human baseline, an expert and non-expert split that lets the author separate machine habits from lack of skill, and four vendors. The performed-versus-genuine distinction is a useful idea for writers because it names a fault that word lists cannot, hedging that carries no information.

The weaknesses are substantial. The preprint is not peer-reviewed and has a single author. The annotation was done by an LLM with no reliability check, and several labels (correctio, sententia, aporetic endpoint, premature closure) require judgment about intent and warrant that models handle unevenly. An LLM annotator may also be more sensitive to LLM-style patterns. The device counts are per document and not normalized by length, while LLMs wrote longer texts. The human and LLM groups differ in period, medium and editing: published and edited articles against single-shot generations from a fixed prompt. The prompt itself told the models to engage honestly with uncertainty and to acknowledge a counterargument, which probably inflated both genuine and performed hedging and the aporetic endings; the author acknowledges this. The sample is English argumentative essays on abstract topics only. The composite scores have tiny magnitudes (FMD under 0.02, GPR under 0.3 in every group) because of how the formulas are built, so their absolute values are hard to interpret, and the claim of miscalibration is inferred from marker density, not checked against whether any claim was actually overstated. Nothing here measures true epistemic status. The paper's own limitations section says the taxonomy measures surface markers only. A small inconsistency: the correctio effect size is given as 0.40 in the table and about 0.41 in the text. The model list also includes models the author calls frontier, and the results describe them as one pool, so no vendor-level claims are supportable.

### Implications for writers

The paper implies four things for anyone producing prose. Do not reach for a three-part series by default; humans use it too, but at a lower and more occasion-driven rate. Make hedges carry information: name the source, condition or scope that makes a claim uncertain, and where none exists, state the claim plainly. Avoid stock phrases that announce complexity without locating it. Let rhetorical devices appear where the argument is under pressure and not at an even rhythm throughout. A further implication, flagged by the Level 3 results, is that an open-ended ending is a template too when the evidence does not call for it.

## What the skill applies

The skill never cites this paper in its text, and several of its rules rest on other sources as well. The match is therefore one of support and wording, not a documented derivation. I note below where I am inferring the link.

Section 9, item 4 (reflexive triplets) lines up most directly with the tricolon finding. The rule targets three adjectives, nouns or clauses where the third adds nothing, which is the pattern Bakhshi finds at a much higher rate in LLM output and spread evenly through a text. The paper also shapes the limit of the rule: human experts use tricolon at a measurable rate, so the rule cuts the padded third item and leaves genuine triples alone. Section 3 ("List as many examples as you have... Don't pad or trim a series to three for cadence") is the drafting-time counterpart of the same finding.

Section 9, item 9 (hedge stacks) and section 4 ("Put uncertainty where it belongs. State facts and data plainly, and hedge the interpretation that's actually uncertain") fit the genuine versus performed distinction. Bakhshi's point that a hedge needs an identifiable basis supports the rule's logic that one hedge or none is the target and that hedging should sit on the inference, not on the data. The paper adds a reason beyond word count: a hedge with no stated condition tells the reader nothing, and LLMs produce about twice the human density of these. That sharpens the rationale for treating a hedge as a claim about evidence.

Section 9, item 8 (overused vocabulary) includes nuanced and multifaceted, which are the sort of complexity tokens the paper codes as performed hesitancy. Bakhshi measures phrases that invoke complexity, not a vocabulary list, so he supports the idea that these words are filler and does not supply the specific words.

Section 9, item 5 (recap endings) and section 6 ("Stop when the content is done") are loosely consistent with the Level 3 results: the aporetic ending is a reflexive wrap-up used by 24 percent of LLM texts. The skill's rule concerns recaps; the paper concerns open-question closers. The paper therefore supports extending the same suspicion to a closing line that gestures at an open question without naming it, but the skill does not say so explicitly.

Section 1 and section 7 (answer first, commit to a view when evidence supports one) are consistent with the Grice-based argument that confidence should be proportionate to evidence, in both directions. Section 4's instruction to avoid false balance has the same root.

## What the skill omits

The skill does not adopt the taxonomy, the metrics (FMD, GPR, RDDE) or any numeric threshold. They are annotation-based research instruments for corpora, and an agent cannot compute them while drafting. Their scale is also hard to interpret and their reliability is unverified.

It does not ban or require specific rhetorical devices beyond the tricolon-related rule. The erotema finding (humans use rhetorical questions more) is a descriptive difference that may owe something to the prompt, and the skill's tone section does not call for rhetorical questions. Anaphora, chiasmus, correctio, auxesis, peripeteia and the expectation-violation-resolution arc showed no significant differences or only small ones, so there is no evidence to base a rule on. Sententia, the aphoristic closer, was coded but not reported as significant in the results I read; the skill's recap-ending rule covers closers on other grounds.

It does not adopt the claim that performed hesitancy is a reliable machine signature or use it for detection. The skill is about writing well, not about evading or performing detection, and a one-corpus preprint is thin support for a detector anyway. It does not adopt the philosophical framework (Grice, Relevance Theory, Brandom) as an explicit rationale, though the reasoning overlaps with its own two stated reasons.

It does not apply the discourse-level closure categories, since the evidence there is mixed: synthetic closure was common in all groups, and only the aporetic endpoint differed.

## Conflicts with other references

No direct conflict with the style manuals; they do not address rhetorical devices or hedging in these terms. I did not open the other reference files while writing this one, so the tensions below are drawn from the brief and from my reading of this paper.

Boggia (boggia-artificial-epanorthosis.md): the skill bans correction framing such as not X but Y except once when the reader holds X, because models overuse it in oratory. Bakhshi finds the nearest device in his taxonomy, correctio, is used less in LLM text than in human text. The two are not the same measure (Bakhshi codes explicit self-retraction or intensification; the contrast frame is about denying a position nobody held), and the genres differ, but a reader could take the two findings as pointing in opposite directions. The skill resolves this by keeping its rule narrow to the contrast frame and leaving self-correction unregulated.

Kobak, Juzek and Ward, Liang and Wikipedia (vocabulary files): those sources measure word frequency, while Bakhshi measures rhetorical structure and phrase types, and the numbers do not compare. His complexity-token category overlaps with words on popular lists, but he does not report individual words. The skill keeps its vocabulary list for the revision pass and treats his result as support for the category, not for any word.

Freeburg (freeburg-em-dashes.md) and the dash question: Bakhshi's example sentences use em dashes in a few places and say nothing about dashes in the findings, so there is no conflict.

Semantic gravity wells and the Anthropic prompting docs: naming a pattern can prime it. Bakhshi's taxonomy is a list of patterns to detect, which fits the skill's choice to keep such lists inside the post-draft pass. The skill does not apply them at drafting time.

Durik on hedges and credibility: if that file finds hedging helps or hurts credibility depending on conditions, it is compatible with Bakhshi's distinction between hedges with and without a stated basis. I have not checked this against the file itself.

## Review notes

I downloaded the arXiv abstract page and the full PDF and read all 19 pages of text, including the appendix tables and the reference list. I could not view the figures (violin plots, ellipses, heat map) because I worked from extracted text, so I relied on the figure captions and the numbers in the body and Table 3. The paper says annotation prompts, corpus metadata and scripts will be released after publication; they were not available, so I could not inspect the annotation prompt, the annotator model or the individual texts.

The date is ambiguous. The arXiv record says submitted 27 March 2026, the identifier begins with 2604, and the PDF says 23 April 2026. I used the arXiv record for the submission date and flagged the other. My per-sentence tricolon ratio and the tokens-per-document estimate are my own arithmetic from reported means, and they ignore that mean sentence length may differ between groups. The statement that LLM text exceeds human text on performed density is from the paper; the interpretation that this indicates overclaiming is the author's and is not independently tested. I did not read any of the other reference files, so the conflicts section reflects the brief and this paper only. The mapping to SKILL.md is my inference, since the skill file does not name its sources.
