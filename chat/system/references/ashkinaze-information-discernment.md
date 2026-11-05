# Information Discernment in Large Language Models

- **Source:** Joshua Ashkinaze, Laura Kurek, Alina Faisal, Tongyuan Miao, Mariam Joseph, Ceren Budak and Eric Gilbert, "Information Discernment in Large Language Models," University of Michigan, arXiv preprint, 2026, arXiv 2607.19355 (version 1; the arXiv listing dates its submission May 22, 2026)
- **Type:** preprint (not peer-reviewed), with a pre-registered user study
- **Where to find it:** https://arxiv.org/abs/2607.19355
- **Last reviewed:** 2026-11-05
- **Used by:** `chat/system/prompt.md`

## What the source covers

The paper asks whether a model that meets outside information which conflicts with what it believes updates more for reliable sources and more for claims that bring it closer to the truth. The authors call the first property source discernment and the second truth discernment, and they find that 13 models have almost neither. Models moved about as much for an unreliable outlet as a reliable one, responded more to how popular a source was than to how reliable it was, and moved about as much toward a false claim as toward a true one. Larger and newer models improved on truth discernment but not on source discernment. Simple prompts helped a little.

### Design

Each trial has two turns. The model answers a question with a known numeric answer, which gives its prior. It is then told that a named news outlet said the answer is some other number and asked what it thinks now, which gives its posterior. The claimed number is the true answer multiplied by one third, one half, one, two or three, so the claim can be correct, mildly wrong or badly wrong, and that accuracy varies independently of the source.

Questions come from five sets with 4,248 test questions in all: numeric items from TriviaQA and NumerSense, two new sets built from the General Social Survey and World Bank indicators (chosen because models are unlikely to have memorized answers like a country's teacher-to-pupil ratio in a given year) and real search queries from the TREC question set, whose answers the authors looked up by hand. Sources are 132 news websites with reliability scores from Lin et al., who combined several existing rating schemes into one score, and popularity from page-rank data. The sources are balanced across three popularity levels and two reliability levels, 22 per cell.

The authors sampled about 11,000 question, source and claim combinations and ran them through 13 models under five prompting conditions, for about 670,000 trials. The models are Claude 3.5 Sonnet, Gemini 2.0 Flash and 2.5 Flash, GPT-3.5 Turbo, GPT-4o and GPT-4o mini, GPT-4.1 and GPT-4.1 mini, GPT-5 and GPT-5 mini, Mixtral 8x7B and Qwen2.5 7B and 14B, chosen as matched pairs to compare size and age.

### Metrics

There are three, each tied to a normative rule. Source discernment is the rank correlation between how far the model moved and the source's reliability score. Truth discernment is the rank correlation between how far the model moved and how much the claim would improve on the prior. Correct defense is how often a model with a correct prior keeps it when shown a wrong claim. Each has a pairwise companion measure that asks how often the model moved further for the more reliable source or the more helpful claim, where 50 percent is chance.

A pre-registered survey of 299 U.S. adults who use language models, matched to the population on age and sex, found that most agreed with all three rules and said a model that broke them would lose their trust and their use. This establishes that users care about the properties. It says nothing about whether the metrics measure them well.

### Findings

**Source discernment is near zero.** Across models it ranged from about 0.02 to 0.06, against zero for no discernment, and the pairwise measure sat at 51 or 52 percent against a chance rate of 50. Update size tracked popularity (a correlation of about 0.07) more than reliability (about 0.03). The authors call this twice as much; in absolute terms both are small, and the honest summary is that neither signal drove updates much.

**Truth discernment is near zero and sometimes negative.** It ranged from -0.25 for Claude 3.5 Sonnet to 0.22 for GPT-5. When models with a wrong prior updated, they moved away from the truth at least as often as toward it. Sorting all trials by outcome, 28 percent were correct priors held, 21 percent were wrong priors moved toward the truth, 20 percent were wrong priors left alone, 22 percent were wrong priors moved further from the truth and 8 percent were correct priors abandoned.

**Correct defense varies widely.** It ran from 0.53 for Claude 3.5 Sonnet to 0.92 for GPT-5.

**Scale and recency help one kind of discernment.** Within matched pairs, larger and newer models did better on truth discernment and correct defense but not on source discernment. The authors suggest truth discernment draws on general reasoning, while source discernment needs knowledge about particular outlets that training may not supply.

**Models partly know which sources are reliable and do not use what they know.** Under the prompt that asks the model to rate a source before updating, the model's ratings correlated with the expert scores at about 0.41, which is moderate. Measured against the model's own ratings, its updates showed source discernment of about 0.46, far above the 0.09 it showed against expert scores and still far from consistent use. The authors read this as two failures stacked: imperfect knowledge of sources and weak application of that knowledge.

**Updating works best where it is least needed.** On datasets where a model's priors were already accurate, it more often updated in the right direction. Models were also far more accurate on wealthy OECD countries than on others. Outside information helped least where the model knew least.

### Prompting interventions

Four prompts were compared with the bare baseline. Asking the model to rate the source's reliability before answering raised source discernment from about 0.02 to about 0.09 and correct defense from 0.74 to 0.83. Step-by-step reasoning raised the two truth measures modestly. A prompt asking the model to first judge whether its original answer was correct raised correct defense. A Bayesian prompt that told the model to update in proportion to the source's reliability raised source discernment slightly but made truth discernment negative (about -0.12) and dropped correct defense to 0.45, because the model moved too far and gave up correct answers. The authors recommend the rating and defense prompts as cheap fixes for deployed retrieval systems.

### Strength of the evidence

The design is clean. Source reliability and claim accuracy vary independently, the reliability scores come from an aggregate of several rating schemes, the new datasets reduce the chance that models answer from memory, the samples are large and the survey adds evidence that the targets matter to users. Parsing errors were handled with a trained filter that removed about 5 percent of rows and left results almost unchanged.

The weaknesses are mostly about what the setting can show. Claims arrive as a bare attribution with no article, no reasoning and no evidence, so the only thing a model can judge a source by is its name. Real research gives a model the text, the method and other sources to compare. Reliability is one score per outlet, though an outlet's reliability depends on the topic. All sources are news sites. The perturbations are synthetic. The metrics measure how far a model moves, not in which direction, and an ideal model would not reach a correlation of one with reliability alone, since claim accuracy should matter too, so the small values are hard to read against a ceiling. The prompting results are one wording each.

### Transfer to a frontier model following a system prompt

Transfer is weak on model generation. The paper was posted in 2026 but its newest models are GPT-5 and Gemini 2.5 Flash, and its only Claude model is Claude 3.5 Sonnet, several generations behind the models this prompt targets. The finding that newer models improve on truth discernment but not source discernment gives some reason to think the source result persists, but that is a trend across older models. The setting also differs. The prompts were placed in the user turn, and the task was a bare number with no tools. A current model with search can look up an unfamiliar outlet, which is the remedy the prompt asks for and which the paper did not test. The direction should carry over: name recognition and popularity likely still stand in for reliability, models likely know more about source quality than they apply, and an instruction to weigh sources by reliability can overshoot into over-updating.

## What the prompt applies

**Popularity as a weak signal (While working, informational goals, second bullet, second sentence).** The prompt lists popularity among the cues that are weak signals of reliability. This paper is the source for that word. Its models moved more with an outlet's popularity than with its reliability, though both effects were small.

**Weigh each source on its own record (same bullet, last sentence).** The paper's main failure is flat weighting: models moved about the same amount for every outlet, reliable or not. Discounting every source equally, the failure the prompt names, is the same flatness at a lower level. The prompt's instruction to weigh each source individually is the paper's first rule stated as an instruction. The gap is that the paper's best prompt for that rule only raised discernment from about 0.02 to about 0.09, so the paper gives modest grounds to expect an instruction to fix the problem.

**Provenance and what independent sources say about them (same bullet, first sentence).** The finding that models' own reliability ratings are moderately accurate and weakly applied supports asking the model to judge a source explicitly. The rating prompt, which forced that judgment, was the most effective intervention for source discernment. The prompt asks for more than the paper tested: it asks the model to find out what independent sources say about a source, which a model with search can do and the paper's models could not. That extension rests on the human fact-checking literature; this paper does not test it.

**Hold correct answers against claims (Handling claims, third and fourth bullets).** The prompt asks for a counterfactual check before giving ground and says to update only on new evidence or a better argument. The paper's correct-defense results and its 8 percent of abandoned correct answers show models giving up correct priors because a named outlet said otherwise. The defense prompt, which asked the model to judge its original answer first, helped. The prompt's counterfactual check and evidence rule aim at the same failure, though the paper tested neither.

**Treating claims as hypotheses (Handling claims, first bullet).** The truth-discernment failure, moving toward a claim regardless of whether it improves on the prior, is what this bullet aims to stop. The paper supports the problem and offers step-by-step reasoning as a partial fix. It does not test the bullet's wording.

## What the prompt omits

**Proportional updating.** The prompt does not tell the model to update in proportion to a source's reliability. The paper's Bayesian prompt did that and made models over-update and abandon correct answers. Leaving it out follows the evidence.

**An explicit reliability score.** The paper's best intervention asked the model to write down a reliability rating before answering. The prompt asks for a judgment of provenance without a numeric rating, because a written score on every source in every chat would add length and the gain in the paper was small.

**Caution where priors are weak.** The finding that models integrate outside information worst in domains where they know least could support a line about extra care with unfamiliar topics or low-resource regions. The prompt leaves it out because it already asks the model to confirm load-bearing claims and to say what it could not check, which applies with most force in exactly those domains.

**The user study.** The survey shows that users want these behaviors. It has no direct use in prompt text.

## Conflicts with other references

Most of the other reference files named below did not yet exist in `chat/system/references/` when this section was first written, so it works from the sources. I read the full texts of Pradhan and Goley and Khurana et al., the abstract and results of Zhu et al. and the abstract of Onweller et al. I did not read Wineburg and McGrew.

**Pradhan and Goley (`pradhan-trust-dont-verify.md`, arXiv 2606.05403).** The papers disagree on prompting. This paper reports that prompts improved discernment; Pradhan and Goley report that verification prompts produced general distrust instead of selectivity. Both results are consistent with a small effect that can overshoot: the gain here was from about 0.02 to about 0.09, and the Bayesian prompt here overshot as clearly as the checklist there. The papers also test different things. This one tests whether a model weighs an outlet by its reputation, with no content to inspect. Pradhan and Goley test whether a model checks the content of a source with no reputation. The prompt's single rule on provenance and method covers both cases.

**Khurana et al. (`khurana-authoritybench.md`, arXiv 2606.13104).** Khurana et al. found that the prestige of a cited journal made no difference to model errors, while the mere presence of a citation did. That fits this paper's finding that reliability barely moved models and something like familiarity moved them a little more. Neither paper finds that models weight sources by reputation in a useful way.

**Zhu et al. (`zhu-misknow-agent.md`, arXiv 2607.20891).** Zhu et al. found deep research agents more likely to adopt a false conclusion from a document attributed to a high-authority institution than a low-authority one, by about 15 percentage points. That is a reputation effect this paper did not find for news outlets. The likely difference is that Zhu et al.'s documents carried full text written to match the institution's style, while this paper gave only a name. Both results agree that the cues models respond to are not well matched to actual reliability.

**Wineburg and McGrew (`wineburg-mcgrew-lateral-reading.md`).** I did not read this study. From its reported findings, fact-checkers judged unfamiliar sources by checking what others said about them. That is the remedy for the gap this paper documents, where models had only a name to go on and moderate knowledge of what it was worth. The prompt's instruction to check what independent sources say about a source joins the two.

## Review notes

I read the full text of version 1 from a local text extraction, including the appendices on source selection, power analysis, metric details, prompts and parsing quality checks. I did not read the user study appendix in detail or the appendix tables comparing model sizes and ages, and relied on the main text for those results. Figures did not survive extraction; the per-prompt values I report come from text that accompanied Figure 2 in the extraction and may be slightly misaligned. The arXiv identifier points to July 2026 while the listing gives a May 22, 2026 submission date; either date meets the recency rule. The study's models are older than its posting date suggests, which limits how far its numbers apply to current Claude models.
