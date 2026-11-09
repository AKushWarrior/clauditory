# Perrone and Romano on the ARB authorship-rewriting benchmark

- **Source:** Gaetano Perrone and Simon Pietro Romano, "ARB: A Matched Authorship-Rewriting Benchmark Data set for AI-Text Detector Evaluation," arXiv preprint arXiv:2607.29539. Version 1 was submitted July 31, 2026.
- **Type:** preprint and data set release (not peer-reviewed)
- **Where to find it:** https://arxiv.org/abs/2607.29539 (HTML at https://arxiv.org/html/2607.29539v1)
- **First published:** 2026-07-31 (on or after the 2026-03-08 cutoff for sources that make claims about AI systems)
- **Last reviewed:** 2026-11-08
- **Used by:** `chat/skills/writing-style/SKILL.md`

## What the source covers

### Purpose

Standard detector benchmarks compare human text with text a model generated from scratch. Real use is messier. People hand a model their own draft and ask it to rewrite. The authors ask whether a detector's score on the standard benchmark predicts how it treats human text that a model has rewritten, and they build a data set to test it.

### Data set

ARB starts from 1,800 human texts of 150 to 500 words, 600 each from news summaries (XSum), creative writing prompts (WritingPrompts), and web text (OpenWebText). Four open-weight models produce the variants: Llama-3.2-3B, Qwen2.5-7B, Mistral-7B, and Gemma-2-9B. For every source text and model, there are four matched versions. The human original. A model generation written from a topic derived from the source (Free-LLM). A rewrite of the human text by the model (H2L). A rewrite by the same model of its own generation (LLM2L). That gives 7,200 matched sets and 23,400 texts. The shared system prompt asks for output only, with no preface or markdown, and avoids any wording about humanizing or evading detectors. Sampling used a temperature of 0.7. About 0.7 percent of outputs were flagged for refusals, non-English text, duplication, or length and were handled in metadata.

### Detectors and metric

Five detectors were tested without retuning: two supervised encoders (BERT-Defense and RoBERTa-Defense), RADAR, Binoculars, and Fast-DetectGPT. The main metric is recall at a false-positive rate of 1 percent on the human texts, which fixes the cost of wrongly accusing a person and asks how much machine text still gets caught. AUROC is a secondary measure. Confidence intervals come from a bootstrap over blocks of data set and generator.

### Findings

Fast-DetectGPT caught 91.2 percent of Free-LLM text and Binoculars 93.5 percent. On human text that a model had rewritten, recall fell to 30.8 percent for Fast-DetectGPT and 15.1 percent for Binoculars, drops of about 60 and 78 points. RADAR fell from 66.8 to 12.2 percent. On LLM text rewritten by the same model, recall stayed at 78.3 percent for Fast-DetectGPT and 83.0 percent for Binoculars, a loss of 13 and 10 points. In AUROC terms, Binoculars went from 0.983 to 0.666 for rewritten human text and 0.952 for same-model rewrites. The two supervised encoders sat near zero recall in every condition at this strict operating point, so their small gains on rewritten human text do not show resilience.

The paper reads these results two ways. Detectors that work on directly generated text do not transfer to human text revised by a model, and the same detectors cope fairly well with a second model pass over machine text. A model leaves its fingerprint when it writes, and a rewrite of human text leaves much less of it.

### Limits the authors state

The design compares benchmark conditions and is not a clean causal test of origin, because Free-LLM text is generated from a topic and not paraphrased from the human source. Only four open-weight generators were tested, so results should not be assumed for proprietary systems. The three domains are English. Text quality was not filtered. The two-pass LLM2L texts run shortest. The study covers five detectors chosen to represent families and excludes watermarking.

## What the skill applies

**Section 7, bullet against detector games.** Rewritten human text slips past some detectors, and detection of machine text is unreliable across conditions. That undercuts any reason to tamper with text to match a score. Together with Karr, the paper shows that scores can fall for the wrong reasons and rise for the wrong reasons. The skill takes the rule from this: write for the reader, never insert artifacts to fool a detector.

**Section 7, bullet on revising the user's text.** The finding that a model rewrite changes how a text is scored supports the view that a rewrite alters the text in ways beyond what the user asked for. It is a modest support, since the paper measures detector behavior and not the text's content or voice, and `van-nuenen-voice-under-revision.md` does that work.

## What the skill omits

The skill omits the data set, the detector families, the metrics, and the individual recall tables. It omits the question of which detector is best, and it omits the observation that same-model rewriting is easier to catch than rewriting of human text. It draws no advice from the finding that human text rewritten by a model avoids detection, because that would invite using a rewrite as a way to launder text, and the skill does not treat evasion as a goal. The paper's data are on open-weight models of 3 to 9 billion parameters, so the skill makes no claim that the percentages carry over to current frontier models.

## Conflicts with other references

The paper pairs with `karr-ai-detection-academic-integrity.md` as noted there: Perrone finds that likelihood-based detectors miss rewritten human text, Karr finds that commercial detectors flag lightly edited text often. The detector families, corpora, and edit types differ, so these are not contradictory. Both leave the writer with the same conclusion, that scores on edited text are unstable.

It extends `smirnov-amplifying-not-learning.md`, which argues that detectors track how predictable text is. A model rewrite of human text might pass as human because the rewrite keeps the human's source structure; that is my speculation, and Perrone does not test it. It also fits `xu-base-models-look-human.md`, where detectors respond to instruction-tuning artifacts, and `maier-beating-style-detector.md`, where detection still works on short drafts when the classifier is trained for the task. The settings differ, since Maier's detector is trained on the participants' own writing and ARB's detectors are used as released.

`freeburg-em-dashes.md` and `kobak-excess-vocabulary.md` concern marks and words that readers can see. This paper concerns scores that readers cannot see. The skill keeps the revision-pass patterns because readers notice them, and it does not claim the patterns control a detector score. No conflict with the style manuals.

## Review notes

I read the arXiv abstract page and the HTML full text of version 1 through a fetch tool that returns a summary. I covered the data set construction, prompts, decoding settings, detectors, the recall and AUROC table, and the limitations. The summary confirmed the four headline values: 30.8 and 15.1 percent on rewritten human text, and 78.3 and 83.0 percent on same-model rewrites. I did not open the PDF, the released data, or the supplementary material. Only version 1 is listed. The abstract gives the drops as roughly 60 to 78 points for the first pair and about 10 to 13 points for the second, which matches the table. I did not check the detectors' own papers.
