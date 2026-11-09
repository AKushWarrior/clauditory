# Kalashnikova and colleagues on how well people spot machine-written text

- **Source:** Natalia Kalashnikova, Nicolas De Bufala, Sophie Fayad, and Laurent Cervoni, "TURING: Evaluating Human Abilities to Identify AI-Generated Texts," Proceedings of LREC 2026, pages 4527 to 4535, published May 2026.
- **Type:** peer-reviewed conference paper
- **Where to find it:** https://lrec.elra.info/lrec2026-main-355 (ACL Anthology entry at https://aclanthology.org/2026.lrec-1.355/, PDF at https://aclanthology.org/2026.lrec-1.355.pdf)
- **First published:** 2026-05 (on or after the 2026-03-08 cutoff for sources that make claims about AI systems)
- **Last reviewed:** 2026-11-08
- **Used by:** `chat/skills/writing-style/SKILL.md`

## What the source covers

### Question and data

The paper measures how well people can tell machine-written text from human-written text, and which cues they say they rely on. The study is in French. It uses 500 texts across ten genres: newspaper articles, jokes, literary criticism, political speeches, discussions, emails, theater plays, poems, recipes, and novels. Of the 500, 268 were machine-generated and 232 human-written. The generating models were mid-2024 systems, named in a footnote as GPT-3.5, Claude 3, Gemini 1.5, Mistral Small, and Copilot. The paper does not report the generation prompts or name the sources of the human texts.

### Procedure

Two hundred fourteen participants, apparently employees of the research center's company, judged texts one at a time as machine, human, or unsure. After each answer they saw the true origin, and every ten texts they saw a performance recap. They could optionally rate the text on ten descriptor pairs (for example incoherent against coherent, monotonous against varied, or repetitive against direct) and could add a free comment. The study logged review time. Cleaning removed unsure answers, answers faster than 7 seconds, and answers implying an implausible reading speed, which left 7,943 judgments out of 9,164 collected. Accuracy was almost the same at every stage of cleaning.

### Findings

Accuracy across all judgments was 59.3 percent, with a confidence interval of about 58 to 60 percent. That is above chance and far from reliable. Participants leaned toward calling texts human: about 52 percent of all labels were human, 43.6 percent of machine texts were called human, and 37.3 percent of human texts were called machine. Genre mattered. Articles, jokes, and poems were easiest, and political speeches, theater, and recipes were hardest. Novels had the lowest recall for machine texts, and articles the highest. Accuracy also varied by the model that produced the text, from about 61 percent for Copilot text to about 51 percent for Claude text. A small Gemini sample scored lower still. The paper words these figures inconsistently, so the metric behind the model breakdown is uncertain.

The descriptors barely changed accuracy; only the coherence rating survived correction, with a modest benefit. They did, however, predict the label. When a reader marked a text as monotonous or lacking variety, about 80 percent labeled it machine-written, against about 30 percent when the reader chose the opposite descriptor. The authors suggest that some annotators used the positive descriptors as a checklist. The same pattern appeared for weak cohesion and coherence.

The authors also tested measured text features: length, sentence count, tokens per sentence, lexical diversity, and syntactic depth. They used models clustered by annotator, adjusted for genre, and applied a Holm correction. After correction, every feature's effect on the label fell within chance. Learning across the session was negligible, at about 0.1 to 0.2 percent per text. Age did not matter after adjustment, though other personal traits did. Two automatic detectors (Binoculars and Fast-DetectGPT) reached about 88 percent accuracy on the same texts, and the paper reports the pair together without separate figures.

### Limits the authors state

The generating models date from mid-2024. Generation protocols were not standardized, so poor prompting may have left cues that stronger prompts would remove. The study is limited to French and to employees of one company. By my reading, the missing prompts and human sources limit reproducibility, the speed thresholds were set by the authors, and some subgroups are tiny.

## What the skill applies

**Section 3, varied rhythm.** The planned rule says to write short sentences often and let sentence length range widely, because readers read flat rhythm as machine text. This paper supports the reader side of that claim: when a reader judges a text monotonous, the reader is far more likely to call it machine-written. It supplies no evidence that a particular rhythm causes the judgment, because the measured sentence features had no reliable effect on labels. The skill uses it for the point that readers perceive monotony and act on it, and uses `gude-grammar-lexicon-two-generations.md` for the measured length gap.

**Section 7, bullet against detector games.** The overall 59 percent accuracy and the humanization bias show that people also get this wrong. The skill concludes that a text should be written for the reader's benefit and not tuned to pass as human.

**Revision pass item 15.** The paper gives a human-reader reason to include a rhythm check. It says nothing about the check's thresholds.

## What the skill omits

The skill omits the genre breakdown, the model breakdown, the review-time analysis, and the annotator traits, since none changes how an agent writes a reply. It omits the 88 percent detector result and the idea of combining human and tool judgments, which concern detection. It does not treat the descriptor pairs as a checklist: the authors themselves note that readers may have used them that way, and a checklist for naturalness would invite a new formula. It also does not rely on the paper for any claim that specific linguistic features mark machine text, since the paper found none after correction.

## Conflicts with other references

The paper sits uneasily with `gude-grammar-lexicon-two-generations.md` in one respect. Gude and colleagues measure a large gap in sentence length between current models and human writers. Kalashnikova and colleagues find that neither sentence length nor lexical diversity reliably moved judgments. The studies differ in language (English against French), in generator era (2025 models against mid-2024 models), and in method (parsed corpora against reader labels). The skill keeps the rhythm rule on the strength of the corpus measurements and the perceived-monotony finding, and treats the reader-labeling result as a caution that rhythm alone does not make a text read as human.

The paper also relates to `bakhshi-rhetorical-miscalibration.md` and `boggia-artificial-epanorthosis.md`, which concern rhetorical patterns that readers may notice. Readers' stated reasons here (monotony, repetition, weak cohesion) overlap with the repetition those papers describe, though I did not read them closely enough to say how far.

Against `smirnov-amplifying-not-learning.md` and `karr-ai-detection-academic-integrity.md`, the point of agreement is that no signal, human or statistical, is reliable on its own. The paper's finding that tools beat people at 88 percent is the opposite of what those two report for deployed detectors on formal human text, but the studied detectors and corpora differ.

## Review notes

I read the LREC page, the ACL Anthology page, and the PDF of the proceedings paper through a fetch tool that returns a summary. The LREC site's own PDF link was blocked by its robots rules, so I used the Anthology copy. I covered the data set table, procedure, cleaning, results, descriptor analysis, feature analysis, detector comparison, and limitations. I did not inspect figures directly. I could not confirm per-genre accuracy values, because the text gives rankings only, and I could not confirm separate accuracy values for the two detectors. The model-level accuracies come from a figure caption that the paper words inconsistently. The data are French and from mid-2024 generators, so the paper is weaker evidence for current English-language models than its publication date suggests.
