# Maier, Zaiss, and Bayer on beating the style detector

- **Source:** Andreas Maier, Moritz Zaiss, and Siming Bayer, "Beating the Style Detector: Three Hours of Agentic Research on the AI-Text Arms Race," arXiv preprint arXiv:2605.02620 [`cs.CL, cs.LG`]. Version 1 was submitted May 4, 2026.
- **Type:** preprint (not peer-reviewed)
- **Where to find it:** https://arxiv.org/abs/2605.02620 (HTML at https://arxiv.org/html/2605.02620v1)
- **First published:** 2026-05-04 (on or after the 2026-03-08 cutoff for sources that make claims about AI systems)
- **Last reviewed:** 2026-11-08
- **Used by:** `chat/skills/writing-style/SKILL.md`

## What the source covers

### Purpose and setup

The paper does three things in one. It replicates an ACL 2026 study by Baumler and colleagues on whether people can post-edit a model draft until it sounds like themselves. It extends that study by asking how well current frontier models imitate a person's style from a single writing sample. It then turns the same data into a detection problem and tries to beat the detector. The authors describe the work as done by an agentic research harness with a human reviewing the results, in about three hours. They release the code, the 648 imitation drafts, the trained detectors, and the adversarial runs.

### The original study and the replication

In the original study, 81 participants each did four tasks in which they edited an o4-mini draft to sound like themselves and two control tasks in which they wrote without help. That gave 486 task records, and the original authors released the logs but not the analysis code. The replication re-implements the analysis from the logs. All seven preregistered hypotheses hold in the same direction after correction for multiple tests, though several effect sizes differ in size from the originals, which the authors attribute mostly to an aggregation choice the original paper left unstated. The correlation between how similar people felt a text was to their own writing and how similar an embedding model measured it to be is also reproduced (r near 0.24, 648 observations).

### Style imitation

Style similarity is the cosine similarity between author embeddings from LUAR-MUD, a frozen model trained on Reddit text. The ceiling is the similarity between two unaided texts by the same author, about 0.70. The o4-mini draft scores about 0.50, and the human post-edit about 0.55, which closes about 24 percent of the gap to the ceiling. Two frontier models, Claude Opus 4.7 and GPT-5.5, each wrote 100 to 200 word drafts from a task description and one earlier text by the participant. They scored about 0.64 and 0.65 and closed about 71 and 74 percent of the gap. They beat the human post-edit on about 77 and 79 percent of tasks. The held-out design matters: the sample shown to the model was never the text it was scored against, which removed a leak that had inflated an earlier run. The comparison is between workflows, since people edited an existing draft while the models wrote from scratch.

### Detection

The authors then trained a linear classifier on the embeddings to separate unaided human texts from machine drafts, validating by participant so that no author appeared in both training and test sets. Detection stayed strong: AUC was 0.999 for the original o4-mini draft, 0.971 for the human post-edit, 0.952 for Opus, and 0.931 for GPT-5.5. Better imitation lowered detectability only a little.

Diagnostics show what the detector used. A classifier that saw only text length reached 0.880 for GPT-5.5 drafts but 0.517 for Opus drafts. The authors conclude that most of the GPT-5.5 detection is a length confound (its drafts ran longer than the human controls), while Opus drafts carry a real stylistic signature. A classifier trained on one model's drafts also detected the other's at 0.89 to 0.91, which suggests a shared signature across frontier models. Labels shuffled at random gave chance results, as expected.

### Adversarial test

Five of the most confidently flagged Opus drafts were each handed to an Opus agent that could see the detector's score and had 20 rounds to lower it while keeping content and length. The mean margin fell from clearly machine to just past the boundary, and two of five drafts crossed over. None had stopped improving after 20 rounds. The authors read this as showing that a capable model can lower its own detectability with modest effort against a fixed detector.

### Limits the authors state

The detector and the generators are all models, and the detector measures author style, not quality or whether the writer would endorse the text. Only one sample per writer was shown. Drafts were short. The detection figures describe spring 2026 and are a lower bound on detectability, and the adversarial result is an upper bound on what a 20-round agent can remove. I noticed several small inconsistencies in the companion tables (a claim of exact recovery for one hypothesis that its own table does not support), which do not touch the numbers used here.

## What the skill applies

**Section 7, new bullet on writing as the user.** The imitation result shows that a model given the writer's own sample moves most of the way toward that writer's style, far more than a person editing a generic draft. That supports the planned instruction to match the user's voice from samples. The skill takes the practical lesson, not the benchmark: samples help, and a draft that ignores them lands in the shared machine register.

**Section 6, cutting to the needed length.** The length diagnostic is the evidence. Length alone separated GPT-5.5 drafts from human text almost as well as the full style classifier. A reader or tool can tell the draft by size before reading a word, so the skill asks the writer to cut to what the content needs.

**Section 7, bullet against detector games.** The paper shows that detection survives strong imitation and that evasion needed an automated loop against a known detector. The skill does not use that loop. It does take the reminder that style matching is a writing aim and not a trick.

## What the skill omits

The skill omits the embedding model, the classifier, and every AUC value, since an agent writing a reply has no access to them and no reason to optimize a score. It omits the adversarial loop entirely: the skill's goal is text that reads well to a person, not text that moves a detector. It omits the replication detail and the human post-edit comparison. It does not adopt the claim that models can lower their own detectability as advice. The paper's drafts were short and written from a task prompt, so the percentages of gap closed are not a promise for every task.

## Conflicts with other references

The paper pairs naturally with `van-nuenen-voice-under-revision.md`. Maier finds that a sample of the writer's work closes most of the style gap, and van Nuenen finds that a rewrite without a sample drifts toward one register. Together they support the voice-matching bullet. They do differ in method: Maier measures style with an author embedding, and van Nuenen uses counts of word and punctuation features.

On length, `gude-grammar-lexicon-two-generations.md` finds longer sentences in newer models, and Maier finds longer drafts. The two measure different things (sentence length against whole-draft length), and the skill uses them for separate rules.

Maier's detector results do not conflict with `smirnov-amplifying-not-learning.md`, `karr-ai-detection-academic-integrity.md`, or `perrone-arb-benchmark.md`, but the setting differs. Maier's classifier compares unaided short texts with drafts from the same participants, so authorship and topic are controlled. The other three study deployed or general-purpose detectors, which fail on formal human writing and on rewritten text. The skill reads them as one message: detection is unreliable in both directions, so the skill never relies on detector behavior.

The adversarial section sits uneasily with the skill's ban on evasion tactics. Maier reports a method as a finding, and the skill declines to use it. `freeburg-em-dashes.md` and `kobak-excess-vocabulary.md` concern individual marks and words, and the length finding adds a global signal those files do not cover. I read those two in full and the others only in part.

## Review notes

I read the arXiv abstract page and the HTML full text of version 1 through a fetch tool that returns a summary. I covered the replication tables, the style-gap table, the held-out protocol, the detection table and diagnostics, the adversarial setup, and the limitations. I did not read the original Baumler paper, the released code or data, or the supplementary material. Version 1 is the only version listed. I could not confirm the exact AUC ranges the abstract reports as 0.93 to 1.00 beyond the four values in the table. The 71 to 75 percent range matches the abstract, while my table reading gives 71 and 74; I use the abstract's range.
