# Karr and colleagues on why detection fails for academic integrity

- **Source:** Jonathan A. Karr, Grigorii Khvatskii, Ting Hua, and Nitesh V. Chawla, "Why AI Detection Fails for Academic Integrity," arXiv preprint arXiv:2608.11256 [`cs.LG`]. First submitted August 6, 2026. The abstract page lists a later version of October 13, 2026 (v3); the HTML read here is labeled v2 and dated August 20, 2026.
- **Type:** preprint (not peer-reviewed)
- **Where to find it:** https://arxiv.org/abs/2608.11256 (HTML at https://arxiv.org/html/2608.11256)
- **First published:** 2026-08-06 (on or after the 2026-03-08 cutoff for sources that make claims about AI systems)
- **Last reviewed:** 2026-11-08
- **Used by:** `chat/skills/writing-style/SKILL.md`

## What the source covers

### Question

Universities use commercial detectors to enforce academic integrity. The authors point out that a detector gives one score for a text, and that score cannot tell light machine editing from full machine drafting. Many policies allow the first and forbid the second. The paper asks how detectors treat lightly edited text, how they treat honest human text, and what a humanizing tool does to the result.

### Design

The corpus is 642 English abstracts from OpenAlex in four fields: chemistry, computer science, political science, and theology. Half come from 2013 to 2015, before large language models, and half from 2023 to 2025. Abstracts, not student essays, were used because they are public and carry dates. Each paper yielded the original abstract and three Gemini 3 Flash versions. One was a light edit of the abstract alone, one a light edit with the full article as context, and one a fresh abstract written from the article. All versions were then run through a commercial humanizer (Undetectable AI, version 11). Two detectors, Pangram 3.2 and GPTZero, scored everything at a fixed threshold of 0.50. The authors cannot verify that 2023 to 2025 originals were written without help. They therefore call the flags on those texts flag rates and not false positives.

### Findings

On the pre-2016 originals, neither detector raised a false alarm. On 2023 to 2025 originals, Pangram flagged 14.9 percent and GPTZero 8.9 percent. Lightly edited abstracts, a stand-in for assistance that most guidelines would accept, were flagged between roughly 38 and 80 percent of the time, depending on the detector, the period, and the editing variant. The reported low end is GPTZero on 2013 to 2015 texts (37.6 percent), and the high end is Pangram on 2023 to 2025 texts (80.1 percent).

Flag rates for human originals differed sharply by field. Pangram flagged 2.1 percent of chemistry abstracts, 9.6 percent of computer science, 24.1 percent of political science, and 26.6 percent of theology, and the other detector ranked the fields the same way. Scores rose with the share of long words and the density of academic vocabulary, with rank correlations of about 0.30 to 0.35, and fell with the share of numbers and non-alphabetic tokens. The authors take this as evidence that detectors respond to stylistic cues and not only to authorship.

After humanization, fewer than 4 percent of the machine-labeled rewrites were still flagged, and the miss rate exceeded 96 percent for both detectors. Humanization also left a mark on the text itself: long-word share, academic vocabulary, and lexical variety fell, and each abstract grew by about 49 words on average. Agreement between the two detectors was moderate before humanization and weak afterward. The authors conclude that an honest user who edits lightly runs more risk of sanction than one who uses a humanizer, and that detector scores should not serve as standalone evidence of misconduct.

### Recommendations and limits

They recommend treating a flag as weak evidence, pairing it with drafting history, and saving sanctions for cases a person confirms. They also expect more false flags in non-STEM and multilingual writing, and they ask for clear norms on permitted AI use and for privacy checks before student work is uploaded. They state these limits: the texts are published abstracts, the labels are proxies, only one generator, one humanizer, and two detectors were tested, the corpus is small because of cost, and the feature links are correlational.

## What the skill applies

**Section 7, bullet against inserting typos, odd characters, or synonym swaps.** The planned bullet rests on two findings from this paper. The humanizer that beat the detectors also changed the text, and the changes were not improvements: it cut vocabulary variety and added filler length. And the detectors' scores shifted with ordinary academic style, so a writer trying to game a score would be chasing a measure that does not track authorship. The skill's advice is to write for the reader and ignore detectors. Where this paper bears on the claim that humanizer tools trade quality for evasion, the evidence is the measured loss of lexical variety and the added length, not a reader judgment of quality, which the paper did not collect.

**Section 7, bullet on revising the user's text.** The finding that light editing alone triggers a high flag rate supports caution about heavy rewriting of a person's work, because the rewrite is what drives the score.

## What the skill omits

The skill omits the detector comparison, the field breakdown, the policy recommendations for universities, and the privacy advice. It omits the humanizer's name and behavior beyond the two measured effects, since describing the tool in detail would read as a how-to. It makes no claim about which detector is better. It also does not adopt the paper's threshold analysis. The paper's argument for institutions (do not punish on a score alone) is correct but belongs to policy, not to a style skill.

## Conflicts with other references

The paper agrees with `perrone-arb-benchmark.md` that lightly rewritten human text is a blind spot, though in opposite directions. Perrone finds that some zero-shot detectors miss human text an LLM rewrote, while Karr finds that commercial detectors flag lightly edited text at high rates. The difference is in the detector family (zero-shot likelihood detectors on open-weight rewrites against commercial classifiers on edits by one proprietary model), the corpus, and the threshold. The skill treats the results as evidence that detector behavior on edited text is unstable, which supports ignoring it.

The vocabulary finding links to `kobak-excess-vocabulary.md`. Karr's detectors scored academic vocabulary and long words higher, and Kobak's measured excess words in recent abstracts are largely formal verbs and adjectives. That suggests the detectors partly key on the vocabulary shift Kobak documents, but neither paper tests that link, and I treat it as an inference.

It also connects with `smirnov-amplifying-not-learning.md`, which explains over-flagging of formal human writing, and `xu-base-models-look-human.md`, which shows evasion by paraphrase. `oppenheimer-erudite-vernacular.md` shows that needless long words cost a writer credibility with readers; Karr shows the same words cost with detectors. The two reasons point the same way, toward plain words, but the skill grounds the plain-word rule in the reader. `wikipedia-signs-of-ai-writing.md` warns against relying on detectors alone, which Karr supports.

No conflict with the style manuals.

## Review notes

I read the arXiv abstract page and the HTML full text through a fetch tool that returns a summary. The two pages disagree about versions: the abstract page showed version 3 of October 13, 2026 as current, and the HTML page was labeled version 2 of August 20, 2026. I used the August 6 submission date for first publication. The summary covered the first 100,000 of about 101,000 characters, so the last pages (probably references) were not read. I covered the design, the flag rates, field differences, the linguistic correlates, the humanization results, the agreement figures, the recommendations, and the limits. I did not open the PDF. Because I read one version and a later one exists, numbers may differ in version 3. The 38 to 80 percent range is confirmed in the text as the span of light-edit flag rates; I assembled the endpoints from the detector figures.
