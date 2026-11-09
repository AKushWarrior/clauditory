# Xu and colleagues on why base models look human to detectors

- **Source:** Yixuan Even Xu, Ziqian Zhong, Aditi Raghunathan, Fei Fang, and J. Zico Kolter, "Base Models Look Human To AI Detectors," arXiv preprint arXiv:2605.19516 [`cs.CL, cs.AI, cs.LG`]. Version 1 was submitted May 19, 2026.
- **Type:** preprint (not peer-reviewed)
- **Where to find it:** https://arxiv.org/abs/2605.19516 (HTML at https://arxiv.org/html/2605.19516v1)
- **First published:** 2026-05-19 (on or after the 2026-03-08 cutoff for sources that make claims about AI systems)
- **Last reviewed:** 2026-11-08
- **Used by:** `chat/skills/writing-style/SKILL.md`

## What the source covers

### Starting observation

Two commercial detectors, GPTZero and Pangram, are used by schools and other institutions. The authors found that both rate text from base language models (models that only continue text, without instruction tuning) as almost entirely human, while text from the instruction-tuned versions of the same models is flagged. The paper then asks what the detectors are really detecting, and builds a tool to show the answer.

### The base against instruct test

The authors continued human-written prefixes with Llama-3 8B in both forms. GPTZero gave the base continuations an average human probability of 96.7 percent and Pangram gave 98.8 percent. For the instruction-tuned continuations the figures were 30.3 percent and 17.1 percent. The paper says similar gaps appear when the prefix is itself machine-written and on Qwen3 8B, but the summary I read gave no figures for those. Human prefixes made continuations look a little more human than machine prefixes did, which the authors count as evidence that local context matters to the detectors.

### The humanization method

The authors then built Humanization by Iterative Paraphrasing (HIP). They took about 11,757 human and machine text pairs, with the machine side produced by a small commercial model. They lightly fine-tuned a base model to turn machine-style text back into the human original, in a plain source and target format with loss only on the target. At test time they applied the resulting paraphraser repeatedly, ten rounds in the main runs. Models covered Qwen3 at five sizes from 0.6 billion to 14 billion parameters and Llama-3 at 8 billion and 70 billion. Evaluation used 256 machine-written passages across eight categories. A model rated each rewrite for how well it kept the original meaning on a scale from 0 to 10.

The comparison included four baselines: plain paraphrase by prompting, an 11 billion parameter supervised paraphraser (DIPPER), a method that swaps in look-alike characters (SilverSpeak), and a reinforcement-learning attack (StealthRL). Only DIPPER gave a non-trivial balance of meaning and evasion among the baselines, and HIP on the smaller Llama-3 8B beat all four on both detectors. In one worked example, the original passage scored 0 percent human on both detectors. After two rounds, Pangram reached 100 percent and GPTZero 20.9 percent. After ten rounds both reached 100 percent, with the meaning score at 9 of 10. Across models, detector scores rose over rounds while meaning scores fell slowly.

### The interpretation

The authors argue that these detectors respond to artifacts of instruction tuning and to local context, and not to any stable signature of machine writing. An appendix experiment supports this: a plain fine-tune of the instruction-tuned model on web-educational text recovered much of the lost human-likeness, and so did the paraphraser. A fine-tune of a commercial OpenAI model through its API behaved differently, with detector scores falling fast and the balance staying weak. The authors suspect the platform adds its own post-training but could not verify it.

### Limits and ethics

Detectors may adapt to the attack, so the results may weaken. The authors acknowledge that an evasion method can be misused, and they frame the work as a diagnosis of a weakness. Computing the study cost about 300 GPU hours, roughly 200 dollars of API use, and many thousands of dollars of detector fees at list prices.

## What the skill applies

**Section 2, choosing the exact word.** The planned rule links model prose that drifts to the probable phrasing with the quantity that detectors score. Xu supplies the link on the instruction-tuning side: what detectors flag is the habit that instruction tuning adds, and a base model without that habit passes. That fits the skill's view that the problem lies in default habits, which no list of forbidden strings captures.

**Section 7, bullet against detector games.** The paper shows evasion needs a purpose-built pipeline that is trained, iterated, and run against a detector, and that it costs meaning. A writer cannot reproduce that in a reply, and the skill does not ask for it. The finding that detector scores track a training artifact is also a reason not to treat a score as a measure of quality.

## What the skill omits

The skill omits HIP, its training recipe, the baselines, and the cost figures. Evading detectors is not a writing goal, and the method is a paraphrasing pipeline an agent cannot run mid-reply. It omits the advice, implied by the base-model result, to write like a base model or to run a second paraphrasing pass. Repeated paraphrase lowers fidelity, and the skill's rule is to change only what is broken. It omits the detector pricing and the OpenAI fine-tuning anecdote.

## Conflicts with other references

The paper fits `smirnov-amplifying-not-learning.md` closely. Smirnov argues detectors read how predictable a text is; Xu shows they read a post-training signature. Those are compatible if instruction tuning makes text more predictable, though neither paper tests that link, and I treat it as an inference.

It also fits `reinhart-llm-grammatical-style.md`, which found base models close to human rates on its grammatical features and instruction-tuned models far from them, and `juzek-ward-lexical-overrepresentation.md`, whose candidate causes include reinforcement learning from human feedback. In all three, the instruction-tuning stage carries the machine-like habits. The skill's revision pass targets the habits and not the model.

Against `one-token-away-from-collapse.md`, which concerns the fragility of instruction-tuned helpfulness, the point of contact is the same stage of training. I did not compare the two beyond that. `karr-ai-detection-academic-integrity.md` and `perrone-arb-benchmark.md` also test evasion and detector failure. The skill treats the group together: detectors fail in several directions, and none of the failures is a writing technique.

No conflict with the style manuals.

## Review notes

I read the arXiv abstract page and the HTML full text of version 1 through a fetch tool that returns a summary. The summary covered the base against instruct results, the data sets, the method, the baselines, the results, the appendix experiments, and the limits. I did not open the PDF or the appendix tables, and I did not read the detectors' documentation. Only version 1 is listed. The paper gives exact base against instruct figures for Llama-3 8B on human prefixes only; I could not confirm the figures for Qwen3 or for machine prefixes. I could not confirm the size of the HIP gain on each model, because the summary reported it as a trend. The meaning scores come from a model judge, not human raters.
