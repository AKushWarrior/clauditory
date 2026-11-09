# Smirnov on why text detectors amplify typicality

- **Source:** Alexander Smirnov, "Amplifying, Not Learning: The Price of Out-of-Distribution Generalization in AI-Text Detection," arXiv preprint arXiv:2605.21653. Version 1 was submitted May 20, 2026; the page also lists version 2 of August 26, 2026, which is the version read here.
- **Type:** preprint (not peer-reviewed)
- **Where to find it:** https://arxiv.org/abs/2605.21653 (HTML at https://arxiv.org/html/2605.21653v2)
- **First published:** 2026-05-20 (on or after the 2026-03-08 cutoff for sources that make claims about AI systems)
- **Last reviewed:** 2026-11-08
- **Used by:** `chat/skills/writing-style/SKILL.md`

## What the source covers

### Claim

Detectors are used in schools, hiring, and publishing, and they often label fluent, formal human writing as machine text. The paper argues that better calibration cannot fix this. A detector trained to tell human text from machine text mostly strengthens a direction that already exists inside the pretrained model's representation, which the author calls the typicality axis. It orders text by how predictable it is to a language model. Fluent, low-surprise text sits on the machine side whether a person or a model wrote it. The same axis lets a detector catch text from generators it never saw, and it is the reason the detector flags careful human prose. The author calls the pairing the price of generalizing to new generators.

### Evidence

The author tested a fine-tuned ELECTRA encoder, other encoders, decoder detectors from 124 million to 8 billion parameters, and zero-shot detectors that score text by likelihood (GLTR, DetectGPT, Binoculars, and Fast-DetectGPT). The study audited 56 deployed detectors in all. Training anchored on a human-versus-ChatGPT corpus (HC3). Over-flagging was measured on professional journalism and on formal student essays from native speakers. Held-out generator tests used SemEval-2024 Task 8 and EvoBench.

The headline figures come from one detector and one test. The fine-tuned ELECTRA detector gave the median formal-native student essay (600 essays, matched in length) a probability of 99.5 percent of being machine text. Text from a model sampled at a high temperature (90 texts) was cleared with a central score near 10.5 percent. Raising the sampling temperature makes model text less predictable, so the detector lets it through, while polished human prose stays predictable and is condemned. A separate deployed detector flagged 56 percent of formal essays at a false alarm rate of 1 percent.

The typicality axis is computed on a frozen encoder as the difference between the average representation of machine text and of human text. A related direction built from a reference model's predictability correlates with it at about 0.79. Eight working detectors share one component that explains about 32 percent of the variance in their scores. Chance would give 12.5 percent.

### The no-go result

The author tried many fixes: penalties for correlation with the axis, adversarial and invariance training objectives, erasure of the axis from the representation, threshold changes, and ensembles. None removed over-flagging while keeping detection of unseen generators. Across the 56 deployed detectors, fairness and cross-generator strength trade against each other (rank correlation about minus 0.85). To catch 90 percent of text from a new generator, the main detector had to flag 69 percent of formal human writing from other corpora.

A training-free operator that removes the typicality component can relocate the bias and revive a detector that had stopped working. It leaves the cross-generator score unchanged. It also lowers detection of 2025 commercial generators at a strict operating point (true positive rate at 1 percent false alarms, 0.818 down to 0.741).

A linear probe trained on roughly 25 labels per class on a frozen model matched or beat the fully fine-tuned detector on some unseen-generator tests. The author notes that the margin shrinks in a cleaner comparison. The probe is also fragile on a single base model.

### Replication and limits

The over-flag mechanism replicates in Chinese and in Python code, where idiomatic, predictable human functions are flagged about twice as often as unusual ones. The collapse of cross-generator detection under matched typicality does not replicate there. A French replication was inconclusive. The author scopes the claim: typicality dominates but does not account for everything, and a smaller authorship residual remains. The no-go result applies to the cross-corpus, new-generator setting. Magnitudes depend on the training recipe and the era of the generators.

## What the skill applies

**Section 2, choosing the exact word.** The planned rule says to prefer the concrete, exact word to the expected one, because model prose drifts to the most probable phrasing and that predictability is what statistical detectors score. The paper supplies the mechanism for the second half of the claim: predictability under a language model is the quantity the detectors track. The first half, that probable phrasing reads as generic, comes from the other sources and from the skill's existing emphasis on specifics.

**Section 7, bullet against detector games.** The paper shows why chasing a detector fails. A detector's verdict depends on how predictable the text is, not on whether a person wrote it, so the score says little about authorship. The skill therefore sets the writing target as clear, specific prose for a reader and leaves detector scores out of the loop.

## What the skill omits

The skill omits every detector, data set, and operator in the paper, along with the axis computation, the probe results, the replication in Chinese and code, and the audit of deployed products. They help people who build or buy detectors, not people who write. It does not adopt the advice to raise sampling temperature to look less machine-like, because an agent in a chat or coding session cannot set sampling, and because high temperature lowers accuracy and clarity. It also does not turn typicality into a rule that predictable phrasing is always bad. A common word is often the right word, and the skill's rule is about the exact word, not the rare one.

## Conflicts with other references

The paper is compatible with `xu-base-models-look-human.md`. Xu finds that detectors rate base-model text as human and instruction-tuned text as machine. Smirnov explains why that could happen: instruction tuning makes text more predictable, and predictability is what detectors read. The two papers use different detectors and tests, so the link is my inference, not either author's claim.

Against `kobak-excess-vocabulary.md` and `juzek-ward-lexical-overrepresentation.md`, the relation is complementary. They list specific words that rose in model text. Smirnov describes a broader property of which those words may be examples. The skill keeps the word list for the revision pass and adds the exact-word rule for the quieter drift.

The paper reinforces `karr-ai-detection-academic-integrity.md`, where detector scores rose with academic vocabulary and long words, and `perrone-arb-benchmark.md`, where detectors missed rewritten human text. All three support the skill's refusal to treat a detector as authority.

There is a mild tension with `oppenheimer-erudite-vernacular.md`, which argues for plain words. Smirnov shows that plain, predictable text scores as machine text. The skill resolves it by asking for the exact word, which may be plain or unusual, and by never choosing words to move a score.

## Review notes

I read the arXiv abstract page and the HTML full text of version 2 through a fetch tool that returns a summary. I did not read version 1, so any difference between versions is unchecked. The summary covered the method, data sets, the typicality axis, the figures for 99.5 and 10.5 percent, the no-go analysis, the operator, the probe, the Chinese and code replications, and the stated limits. I did not open the PDF, the supplementary material, or the cited detectors. The 10.5 percent figure is a location read from a figure panel in the paper, and the same figure reports a mean of 0.07 for the high-temperature text, so the two are different summary statistics. I report 10.5 percent as the paper's headline value for the central score.
