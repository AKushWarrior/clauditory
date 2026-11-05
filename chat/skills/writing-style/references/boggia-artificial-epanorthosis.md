# Artificial Epanorthosis: Why Large Language Models Overuse a Classical Rhetorical Figure

- **Source:** Federico Boggia, "Artificial Epanorthosis: Why large language models overuse a classical rhetorical figure, and how to mitigate it," arXiv preprint 2607.21498, July 2026 (version 2, stamped 28 July 2026).
- **Type:** preprint (not peer-reviewed)
- **Where to find it:** https://arxiv.org/abs/2607.21498 (PDF at https://arxiv.org/pdf/2607.21498)
- **Last reviewed:** 2026-11-04
- **Used by:** chat/skills/writing-style/SKILL.md

## What the source covers

### Purpose and scope

The paper takes a figure of speech that readers now associate with machine prose, the self-correcting contrast, and asks three questions about it. Where does the habit come from, how much do models overuse it compared with people, and what can be done about it. It is a short essay with a small empirical core. The author is a teacher and trainer with a humanities-computing degree from the University of Pisa, and he states that the detector, measurements, experiments and paper are his own work, with an Italian reference corpus built jointly with a fellow student in a university course project. He combines a literature-style argument from classical rhetoric with a pilot measurement in English, a validation exercise and a mitigation demonstration in Italian, and a survey of intervention methods.

The argument ends on a stance that matters for anyone writing rules about this figure. The goal is to bring the figure's frequency in line with what human writers do in each genre, and it is a mistake to aim at zero.

### The rhetorical figure and how it relates to the familiar patterns

Epanorthosis is the figure in which a speaker returns to a claim just made in order to strengthen it, soften it or retract it. The Latin name is correctio. The paper separates three variants by function: an emphatic one that upgrades the claim, an attenuative one that downgrades it, and a corrective one that replaces the content. Quintilian discusses related self-correction devices and notes that simulated reconsideration can make a speaker seem spontaneous and less rehearsed. The paper's clearest classical example is a passage in Cicero's first speech against Catiline, where a flat statement is corrected upward by a single particle. It also distinguishes true epanorthosis from accumulation, such as a run of four verbs in Cicero's second Catiline speech that pile up without retracting one another. Dante's close of the Ugolino canto appears as a vernacular case where the final correction reverses the meaning of the whole scene.

Fontanier, writing in 1827, classed correction as a figure of style and epanorthosis proper (which he also called retroaction) as a figure of thought. Boggia leans on this split. If epanorthosis is a figure of thought, a general discursive operation, it can be studied in model output without asking what the model intended. The familiar surface patterns then become detectable proxies for that operation. The paper's focus is the emphatic variant, and its listed surface forms include the plain not-X-but-Y frame, not-only-X-but-Y, X-or-rather-Y, it-is-not-X-it-is-Y, a sentence-split form where a negated sentence is followed by an affirmed one, and a question answered by a flat no followed by the corrected claim. So the paper treats the constructions that readers call not-X-but-Y and not-just-X-it's-Y as the emphatic family of one old figure, and counts only a slice of that family: the frames built on a negation followed by a replacement.

### Proposed origins of the overuse

The paper describes the overuse as a trained disposition and ranks the candidate causes. It names preference tuning as the probable main driver: raters tend to prefer confident, committed answers, and an upward correction sounds confident. Web training data comes second, because copywriting and motivational text use emphatic correction to hold attention, though the paper says it has not measured this. Left-to-right generation is a minor and indirect factor. A model cannot retract tokens it has already emitted, so a learned disposition appears in the output unedited, but the paper argues the architecture does not by itself cause the habit. Likelihood-seeking decoding adds a small push, because the negate-then-replace frame is a high-probability path. The author is open that the preference-tuning claim is an inference. It borrows from earlier lexical work, which found that preference tuning is consistent with inflating certain words, and extends it to a grammatical construction without a direct test.

The paper opens its case with a published finding that models diverge from human writers in grammatical and rhetorical style, that the divergence is larger for instruction-tuned models than for base models, and that it persists as models scale. It also offers a field heuristic: three or more such frames in a short passage strongly suggest generated text but never prove it. Finally, it says it tries to practice its own advice, keeping its prose free of gratuitous correction, with the corrective turns it does use being the content-bearing kind.

### Method: how instances were found

The English detector counts the emphatic family per 10,000 words. It looks for the negation-then-replacement frames described above, including the sentence-split form. It was not validated separately. Significance comes from a label-permutation test with 5,000 iterations.

The Italian detector is a high-recall lexical pass over the not-but frame and its variants. The paper proposes a construction-level classifier to remove false positives but does not implement it, and it does not implement marker-based triggers for corrections such as the Italian equivalents of that is and let us say. Validation used 206 windows of about 150 words, 180 from the human corpus and 26 from model output. Two annotators labeled them and disagreements were adjudicated. Agreement between the annotators was weak (Cohen's kappa 0.35). Observed agreement was 0.66. Against the adjudicated labels, overall micro precision was 0.45 and recall 0.52, from only 23 true positives, and every genre except academic prose rested on seven or fewer true instances. Precision split sharply by source: 0.82 on model text and only 0.17 on human text. Recall in spoken Italian was 0.14, and five of the six missed corrections there were marked by other expressions of self-correction, so the loss reflects the channel's narrow coverage. Read plainly, the lexical pattern is a good indicator of the figure in model output and a poor one in human writing, where the same words often serve a plain contrast rather than a true correction.

### The Epanorthosis Index

The index for a genre is the model's density divided by the human baseline density for that genre. A value of 1 means human-like frequency, above 1 means overshoot and below 1 means undershoot. A corpus-level calibration error is the mean absolute log of the index across genres, which is zero for perfect calibration and penalizes over- and undershoot symmetrically. The design is simple and sensible: it compares each genre only with its own human norm, so it does not punish a genre where people use the figure heavily.

### Data and models

The English measurement covers eight prompted genres, seven with human baselines. The human texts are Wikipedia for encyclopedic prose, Wikinews for journalism, arXiv abstracts for academic writing, Conan Doyle for narrative fiction, high-voted pre-2022 Stack Exchange answers for informal Q&A, Mill's On Liberty for argument, and inaugural and commemorative addresses by Lincoln, Jefferson and Theodore Roosevelt for oratory. Promotional copy has no public-domain baseline and is measured on the models alone. The models are Claude Haiku, Sonnet and Opus, three sizes of one instruction-tuned family with a shared alignment pipeline. Unit counts are windows of text, with 17 to 28 model windows in most genres and 19 to 128 human windows.

Italian materials serve the validation, the human reference rates and the mitigation demonstrations. They include a spoken-conversation corpus of about 320 conversations, 248 open-access philosophy articles, 11 novels from 1827 to 1923, about 24,500 news articles and 260 Reddit posts. For adapter training the author used the open Qwen2.5-7B-Instruct model rather than the Claude models.

### English results

Emphatic density per 10,000 words, human against pooled model output, with the index in parentheses:

- Encyclopedic: 1.2 against 1.4 (1.2 times, essentially parity).
- Journalism: 2.4 against 2.0 (0.9 times).
- Academic abstracts: 3.6 against 7.8 (2.2 times, but not significant, p = 0.28).
- Narrative fiction: 7.5 against 4.1 (0.5 times, p = 0.38).
- Informal Q&A: 8.2 against 1.3 (0.2 times, p = 0.05).
- Argumentative essay: 17.2 against 22.0 (1.3 times, p = 0.37).
- Oratory: 14.9 against 33.5 (2.2 times, p = 0.03).
- Promotional: models alone at 26.4.

The headline is therefore two-sided. Models overshoot in oratory by about a factor of two and undershoot sharply in informal answers by about a factor of five, while encyclopedic, journalistic and argumentative prose sit near human levels. Within oratory the overshoot grows with model size: about 15 for Haiku, near the human rate, against 37 for Sonnet and 46 for Opus. Neutral connectives were less frequent in model oratory than in human oratory (1.5 against 7.5), which rules out the simple explanation that models are just heavier users of discourse markers. The promotional figure of 26.4 is notable because it is the one genre with no human anchor, so it supports no claim of overuse, only a level of use that is high.

The Italian oratory baseline was 39.6 against a human reference near 14, close to threefold, and the paper reports model rates two to five times above human Italian rates, concentrated in the larger tiers.

### Mitigation work

The paper surveys data-side fixes (counterfactual pairs of emphatic and plain text, pruning promotional instruction data), preference optimization, decoding-time control, activation steering, and prompting or constitutional clauses, with results only for a few.

On the adapter, preference optimization alone on 570 pairs failed: density on held-out prompts rose from 60.3 to 68.9. Supervised fine-tuning at rank 32 on 797 filtered pairs cut density from 75.4 to 1.7, about 98 percent. When the adapter's strength was scaled from 0 to 1, overall density fell from 61.4 to 4.7, and argument density passed near the human reference around the 0.75 setting. The scaling was not monotone everywhere: promotional density rose to 96.0 at the 0.25 setting before falling. The author reads this as a reason to tune the dial per genre. These tests had six generations per prompt and no significance tests, and content fidelity was checked only by informal inspection.

A one-line prompt asking for affirmative writing, applied to 72 Italian texts across three genres, three model sizes and four topics, reduced density by 72 percent in oratory (39.6 to 10.9, p = 0.03), 70 percent in argument (56.9 to 17.2, p = 0.004) and 48 percent in promotional text (32.3 to 16.8, not significant). The larger models went to zero while Haiku about halved. Among training-free options, best-of-six selection cut overall density from 59.9 to 20.9, and a rewrite pass for affirmative phrasing reached 0.0, which the author counts as overcorrection because it falls below the human rate. These baselines rest on two prompts per genre and the paper warns they are not comparable across its tables. Its recommended stack combines curated data, a detachable adapter trained by supervised fine-tuning, a dial, and decoding-time or steering controls where per-request adjustment is needed. It calls prompting the cheapest first move and the least durable, since training-time dispositions resist surface instructions.

### When self-correction does real work

A section of the paper argues the figure is sometimes a virtue. Chain-of-thought text with phrases that reverse an earlier step does shift later token probabilities toward alternative reasoning. Work cited by the paper finds, however, that models often cannot fix their reasoning reliably without an external signal such as a tool result, verifier or retrieved fact, and that blind revision can make answers worse. The author also notes an inversion: classical rhetoric taught speakers to simulate self-correction to seem credible, while engineering teaches machines to correct themselves to be accurate. A further section weighs whether a model's correction is a real figure or a statistical artifact, sets a Chinese Room reading against a functionalist one, and concludes that both are defensible. The practical claim it keeps is narrower: the excess is produced, it is measurable and gratuitous instances can be cut.

### Calibration, not elimination

The closing argument is that the right target is each genre's human rate, preferably exposed as an adjustable setting, and that pushing density below that rate is a failure of its own. Evaluation should combine genre-relative density, content fidelity and human judgments of appropriateness. The paper ends with a warning that the larger risk is people beginning to write the way the machines do.

### Strengths and weaknesses of the evidence

The index is a clean idea, the genre-by-genre comparison is more informative than a single overall rate, and the paper is candid about its limits. Its stated limits are real and large. Cells are small. The seven per-genre p-values are uncorrected, and the paper itself says oratory (0.03) and informal Q&A (0.05) would not survive strict correction. Human baselines for oratory, argument and narrative are from the 19th and early 20th centuries, so period style is mixed with register, and a gap in oratory could partly reflect the era. The detector catches one slice of the figure, and its validated precision on human text was low. The three measured models share one alignment pipeline, so the tier trend and the preference-tuning story cannot be separated from that pipeline, and no base-versus-aligned comparison was run. Adapter results come from a different model than the measured ones, in a different language, and the Italian human reference rates rest on one argumentative author and a few orations. The paper says its findings should not be generalized to instruction-tuned models broadly.

Because of these limits, the direction of the findings is more trustworthy than the sizes. The oratory overshoot and the Q&A undershoot are suggestive, the near-parity in encyclopedic, journalistic and argument prose is weakly supported by small samples, and the mitigation numbers show what is possible in a pilot rather than what will transfer.

## What the skill applies

The skill's revision pass, item 2 in section 9, names correction framing and lists the forms "Not X, but Y," "It's not X, it's Y," "Y, not X," "Not just X, but also Y" and "This isn't about X. It's about Y." Boggia's catalog of surface variants, including the sentence-split form and the not-only form, is the evidence that these are variants of one family, which is why the item lists several shapes and not one. The Stating a priority example at the end of SKILL.md shows the sentence-split form being rewritten as a direct statement.

Section 3 carries the principle behind the rule in positive form: state claims directly, say what is true rather than first describing what isn't, and use a contrast only when the reader actually holds the mistaken view. That last clause reflects the paper's own point that a correction does real work when something specific is being corrected.

The three-or-more heuristic shaped the limit in item 2 that a contrast is allowed at most once per response. The paper treats repetition inside a short passage as the signal that reads as machine-made, so the cap targets density, which is what the index measures.

The reflexive-triplet item (4) and the correction item share a purpose: both remove patterned rhythm that adds nothing. Boggia does not discuss triplets, so that link is the skill's own, not the paper's.

The paper's finding that models overuse the figure most in the larger tiers and in oratory and promotional registers also supports the skill's placing of this check in the revision pass for all prose, since the risk rises when a model writes in a persuasive voice.

## What the skill omits

The skill does not adopt the calibration goal. It does not ask the writer to match a genre-specific human rate, because an agent cannot measure its own density mid-draft and the user has asked for the pattern to be banned. It keeps the single permitted use for a true correction of a belief the reader holds, which is the nearest the skill comes to calibration.

It omits the Epanorthosis Index and calibration error, which are evaluation instruments for researchers. It omits the detector design, the Italian validation figures, the adapter and dial experiments, the prompting and best-of-n demonstrations, and the survey of decoding and steering methods. None of these is something a writer can apply while drafting prose in a chat or document, and most rest on pilot samples, one non-Claude model and Italian text.

The attenuative and corrective variants of the figure are left out because the skill targets the emphatic family that the evidence concerns. The discussion of reasoning traces and of the Chinese Room question has no bearing on prose style rules.

The skill also leaves out the undershoot finding for informal Q&A as grounds for encouraging the pattern in conversational replies. The finding sits at the edge of significance, rests on a Stack Exchange human baseline, and the user's explicit choice to ban the construction outweighs it.

Nothing in the skill relies on the paper's causal story about preference tuning, since it is untested and the rules work the same whatever the cause.

## Conflicts with other references

The main tension is in goal, and it is with the paper's own recommendation, not with a style manual. Boggia argues for matching each genre's human rate, and finds the models below it in informal Q&A, which is close to the register of a chat reply. The skill bans the construction outright with a one-use exception. The skill resolves this in favor of the ban for three reasons: the user asked for it, an agent cannot calibrate a rate it cannot count, and the undershoot result is the paper's weakest (p = 0.05 before correction, and not significant after it). The practical difference is small, because a ban that allows one real correction produces rates well within the human range for most genres.

A second tension is with Freeburg (freeburg-em-dashes.md) in method only: that study shows human writing also uses a feature that models overuse, which supports banning on the grounds of reader perception rather than rarity in human text. Boggia's human baselines point the same way for this figure, since human argument and oratory use it heavily. I did not re-read that file for this review, so the comparison rests on its filename and the brief.

Boggia cites the lexical work by Juzek and Ward and the skill's other reference notes on vocabulary and syntax, but those use different measures (word frequency against construction density), so their numbers cannot be compared with the index. Wikipedia's signs-of-AI-writing page, as I understand it from the brief, lists contrast constructions among popular tells. Boggia supplies a measured counterweight: the figure is common in human writing in some genres, and the lexical detector flags human text far more falsely than model text. The skill's response is to restrict the rule to the reader's belief test and the per-response cap, not to treat every contrast as a defect.

The style manuals (NYT, AP, Chicago) do not address this construction, so no precedence question arises.

## Review notes

I read the full text of the paper (version 2, 18 pages including the limitations, references and both appendices) from the text file the coordinator supplied, and checked every number in this file against it, correcting my first draft where it differed. An earlier draft had been built from page-summarizing fetches; the full text matched them on all figures but added detail on author background, observed agreement and true-positive counts, and showed that supplementary materials (scripts, adapter weights, logs) are offered for download alongside the paper. I did not download or inspect those, the English corpus (available from the author on request), or any of the paper's cited works. The classical citations (Cicero, Quintilian, Fontanier, Dante) are reported as the paper presents them, and I did not check them against the originals.

Two doubts about my own reading. First, the claim that the paper's contrast frames cover the not-just-X-it's-Y shape rests on the paper's variant list, which includes not-only and it-is-not-it-is forms, and not on a table of every counted instance. Second, the conflicts section compares this source with other reference files that did not yet exist in the references folder when I wrote it, so those comparisons draw on the brief's description of them.
