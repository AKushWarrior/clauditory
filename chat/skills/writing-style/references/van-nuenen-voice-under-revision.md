# Van Nuenen on voice under revision by LLMs

- **Source:** Tom van Nuenen, "Voice Under Revision: Large Language Models and the Normalization of Personal Narrative," arXiv preprint arXiv:2604.22142. Version 1 was submitted April 24, 2026.
- **Type:** preprint (under review at the time of reading)
- **Where to find it:** https://arxiv.org/abs/2604.22142 (HTML at https://arxiv.org/html/2604.22142v1)
- **First published:** 2026-04-24 (on or after the 2026-03-08 cutoff for sources that make claims about AI systems)
- **Last reviewed:** 2026-11-08
- **Used by:** `chat/skills/writing-style/SKILL.md`

## What the source covers

### Question and design

The paper asks what happens to a person's writing when a language model revises it. The author works in digital humanities, where word counts and punctuation habits are used as evidence of style, authorship, and the integrity of a collection. If model revision shifts those features, a revised text stops being evidence of its writer.

The corpus is a stratified sample of 300 first-person narratives drawn from a public collection of personal stories that spans oral histories, recollections, three Reddit communities, and confessions. Stories run about 244 words on average. Three recent models (GPT-5.4, Claude Sonnet 4.6, and Gemini 3.1 Pro) rewrote each story under three prompts. The first asked for a general improvement. The second asked for a rewrite with no mention of improving. The third asked for improvement while keeping the author's voice and emotional tone. The full crossing gives 2,700 rewrites, and each original acts as its own control. Sampling used a temperature of 0.7.

### Features and findings

The author tracks 13 measures: coverage of the most frequent words, the share of function words, three vocabulary range indices, mean word length, character-level entropy, comma and dash frequency, mean sentence length, emotion words, first-person pronouns, and contractions. Under the general prompt, all 13 shifted for all three models, with corrected tests, and every shift pointed the same way across models. Function words, first-person pronouns, and contractions fell (contractions by about 31 percent). Vocabulary range, word length, and comma use rose. Commas rose by about two thirds, and the share of texts containing a dash rose from 26 to 72 percent. Mean sentence length rose modestly (about 7 percent), and the models differed in how they got there: two merged clauses and shortened the sentence count, while Claude split constructions and added sentences. The author reads the change as a move toward a formal register with punctuation doing more of the connecting.

Prompt wording changed the size of the shift but not its direction. Dropping the word improve left the shift intact; the average effect was about 5 percent larger. The voice-preserving prompt cut the average effect by about a third. It did not remove it: vocabulary range still rose by about 35 percent, and frequent-word coverage still fell by about 8 percent. Emotion words and word length responded most to the voice instruction, while first-person pronouns and function words barely did. The author concludes that voice prompts suppress additions such as elaboration more easily than subtractions such as lost contractions and lost first-person presence. In one example the voice prompt produced more literary, more emotive prose, which the author reads as stylization, not preservation.

### Clustering, predictability, and narrative effects

Original stories are widely scattered in a space built from character-level features. Rewrites from all three models land in one tight region. They are hard to match back to their sources: a nearest-neighbor match recovered the source in under 15 percent of cases for the best model and about 1 to 2 percent for the others. Variance across features shrank for the rewrites, and the rewrites were more predictable to a small reference model than the originals were. Meaning was largely kept. Cosine similarity was 0.84 to 0.90, against 0.22 for random pairs.

The paper also reports shifts above the word level. Narration moved from inside the events toward looking back on them, and explicit causal links gave way to summary statements of what the narrator had come to understand. The voice-preserving prompt nearly erased the first shift but not the second.

### Limits the authors state

The design cannot say whether pretraining, instruction tuning, or default system behavior causes the drift. It uses one English corpus. The stories are short, so author-attribution methods built for longer texts are used only as aggregate patterns. The markers stand in for voice and do not measure it, and no reader study shows which shifts people notice.

## What the skill applies

**Section 7, new bullet on revising the user's text.** The planned bullet says to keep the user's own phrasing where it works and to change only what is broken. The paper supplies the reason: an open instruction to improve, or even to rewrite, moves text in the same direction every time, toward a single polished register with fewer contractions and first-person forms. The skill's response is to treat the user's wording as the default and to edit narrowly.

**Section 7, matching the user's voice.** The finding that a voice instruction helps but leaves a residue is the reason the skill adds samples and a narrow edit scope to the instruction itself. A line telling the writer to preserve voice is not enough on its own.

**Section 8 and the register rules.** The loss of contractions and first-person pronouns supports the existing rule that conversation takes contractions and the first person, and that a piece keeps one approach throughout.

## What the skill omits

The skill leaves out the 13-feature battery, the effect sizes, the clustering analysis, and the source-recovery tests. Those serve digital humanities questions about corpus integrity and are not writing guidance. It omits the claims about narrative stance and explicit causation, which rest on automated measures of clause types that an agent cannot check as it writes. The skill does keep the plain advice not to compress a person's reasoning into summary. It sets no numeric targets for contractions or commas. It also does not call the drift bad in every case, since a formal document may want the register the models produce.

## Conflicts with other references

The most direct link is `freeburg-em-dashes.md`. Dash frequency in van Nuenen's rewrites rose more than fourfold, and the share of rewritten texts with a dash nearly tripled. That supports the skill's flat em dash replacement rule, and it adds a case in which the rewriting task, not free generation, produced the dashes.

`gude-grammar-lexicon-two-generations.md` finds that newer models write news with lower lexical diversity than humans, while van Nuenen finds richer vocabulary in rewrites of personal stories. The tasks, genres, and metrics differ, so the studies do not contradict each other. Each finds that the output departs from the human range, in different directions.

`reinhart-llm-grammatical-style.md` reports that instruction-tuned models use dense, noun-heavy grammar, and van Nuenen's abstraction and register findings are consistent with it. `maier-beating-style-detector.md` supplies the contrast: a model given the writer's own sample imitates that writer far better. I read other files in the folder only in part.

The paper does not conflict with the style manuals in this folder. I did not check what each says about contractions.

## Review notes

I read the arXiv abstract page and the HTML full text of version 1 through a fetch tool that returns a summary. The summary covered the first 100,000 of about 109,000 characters, so the final part (probably the last discussion, conclusion, and references) was not read. I covered the corpus, models, prompts, the marker table, the prompt-condition results, clustering, the narrative analyses, the consistency checks, and the limitations. I did not open the PDF or the supplementary material. Only version 1 is listed. All percentages and effect sizes are taken from the summary, and I rounded them. The claim that most markers move in the same direction for every model matches the abstract.
