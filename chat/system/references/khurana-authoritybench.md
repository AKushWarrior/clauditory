# Authority, Truth, and Citation Bias: A Large-Scale Multi-Domain Benchmark for Studying Epistemic Susceptibility in Large Language Models

- **Source:** Aryan Khurana, Aravind Ramana RN and Dhruv Kumar, "Authority, Truth, and Citation Bias: A Large-Scale Multi-Domain Benchmark for Studying Epistemic Susceptibility in Large Language Models," BITS Pilani, arXiv preprint, June 11, 2026, arXiv 2606.13104 (version 1); the arXiv listing says it was accepted to the AI4GOOD and EIML workshops at ICML 2026
- **Type:** workshop paper (accepted to two ICML 2026 workshops; the arXiv version is a preprint, and workshop review is lighter than main-conference review)
- **Where to find it:** https://arxiv.org/abs/2606.13104
- **Last reviewed:** 2026-11-05
- **Used by:** `chat/system/prompt.md`

## What the source covers

The paper builds a benchmark, AuthorityBench, to measure how a citation attached to a claim changes a model's judgment of that claim, apart from whether the claim is true. Its main finding is that adding a citation, real or invented, made every one of seven models more error-prone on average than giving the bare claim. The worst case was an invented citation attached to a true claim, which led models to reject correct facts more often. The prestige of the cited venue and the apparent nationality of the author made no measurable difference.

### Design

The benchmark crosses two factors: whether the claim is true or false, and whether the citation is real or invented. Each claim also appears with no citation as a baseline. The 110,282 base claims come from four existing datasets: FEVER for general knowledge (claims checked against Wikipedia), SciQ for science, CaseHOLD for law and MedMCQA for medicine. Multiple-choice items were turned into statements by pairing a question with its correct answer or a wrong option. That gives 220,564 prompts in the four cited conditions.

Invented citations draw an author surname from a dataset of common names by country, a venue from 480 real journals and law reviews ranked into four prestige tiers by standard ranking sources, and a year from 1980 onward. Real citations come from existing scientific, medical and legal datasets. FEVER has no citation records, so its "real" citations were borrowed from the other three domains, which means that in general knowledge a real citation is a genuine paper unrelated to the claim. Citations were placed in 40 templates of seven kinds, from a leading "according to" phrase to a footnote, and half were drawn from the same domain as the claim and half from a different one.

Seven models were tested. Gemma 3 4B, Llama 3.1 8B and Phi-4 Mini ran on the full set. Gemma 4 31B, Claude Haiku 4.5, GPT-5.4 mini and DeepSeek V3.2 ran on a 15,000-prompt stratified subset for cost reasons; the authors checked that two of the small models gave similar results on the subset and the full set. A small judge model, Qwen3-8B, given the ground-truth label, marked each answer as a hallucination or not. Two annotators checked 1,500 outputs, with a kappa of 0.83; the paper does not say clearly whether that figure measures agreement between the two people or between the people and the judge.

### Findings

**Citations raise error rates.** Without a citation, error rates ran from about 12 percent (Gemma 4 31B) to about 31 percent (Gemma 3 4B); Claude Haiku 4.5 was at about 25 percent and GPT-5.4 mini at about 19 percent. Averaged across true and false claims, every model made more errors once any citation was added.

**An invented citation on a true claim is the worst case.** In every model this condition had the highest error rate. The increase over the model's baseline on true claims ran from about 3 percentage points (DeepSeek V3.2, Gemma 3 4B, Phi-4 Mini) to 22 points (Llama 3.1 8B), with Claude Haiku 4.5 up 15 points to 47 percent and GPT-5.4 mini up 20 points to 37 percent. Real citations also raised errors on true claims in six of seven models, by up to about 16 points. The larger and closed models were not more resistant.

**On false claims, model families split.** For Llama 3.1 8B, Claude Haiku 4.5 and GPT-5.4 mini, an invented citation on a false claim reduced errors, by about 7 to 14 points, as though the citation prompted closer scrutiny. For the Gemma models and Phi-4 Mini it increased them slightly. Read together with the true-claim result, the first group's pattern looks like a citation setting off general suspicion: errors fall on false claims and rise on true ones. The authors frame it as the citation acting as a credibility check in some families and an endorsement in others.

**Domain matters; prestige and names do not.** General knowledge was the weakest domain, with error rates on true claims with invented citations near 77 percent for Llama 3.1 8B and Claude Haiku 4.5 and 69 percent for GPT-5.4 mini. Law was the most resistant; the authors suggest legal citation formats are recognizable enough to raise the bar. Venue prestige produced spreads of under 4 points with no consistent ordering, so an invented low-tier journal did as much as an invented top journal. Author surname region made no real difference. Older citation years were slightly more disruptive, by up to about 2 points.

**Format matters a little and differently by model.** A leading "according to" phrase was usually the riskiest placement and a bare author-and-year mention the safest, with spreads of 4 to 17 points, but no format was safest for every model.

### Strength of the evidence

The benchmark is large and the two-by-two design is a real improvement on earlier work, which tested invented citations only on false claims. The finding that citations disrupt judgment of true claims is new and replicated across seven models.

Several weaknesses limit how far the numbers can be pushed. The paper never shows the instruction the models received, so it is unclear whether they were asked to judge the claim, answer a question or comment freely, and what counted as a hallucination on a true claim has to be inferred. The judge is an 8-billion-parameter model that cannot check whether a citation exists. The general-knowledge domain, where the effects are largest, is the one where real citations were borrowed from unrelated fields and invented citations name academic journals for Wikipedia facts, so the mismatch itself may be what the models reacted to. Some reported lifts do not match the baseline and condition rates in the same tables, apparently because full-set and subset baselines are mixed. Main results lack confidence intervals in the text. Four of the seven models are small open models, and the frontier models ran on a subset.

### Transfer to a frontier model following a system prompt

Transfer is limited. Only one Claude model was tested, Claude Haiku 4.5, a small model, and no current large frontier model was included. The models had no tools, so they could not look up a citation. A model following this prompt with search can open the cited source, which removes the condition the benchmark depends on: judging a citation by how it looks. The finding most likely to carry over is the shape of the effect, while its size may differ: a citation's presence changes how a model treats a claim regardless of whether the citation is real or relevant, and in some models it triggers suspicion that hurts true claims as much as it helps on false ones.

## What the prompt applies

**Open a cited source before relying on it (While working, informational goals, third bullet, second sentence).** The prompt tells the model to open any cited source it relies on, including one from a subagent, and check that it says what the citation claims. This paper is the clearest support for why the model cannot judge a citation by looking at it: real and invented citations moved models in similar ways, and venue prestige, which a model might use as a shortcut, carried no information about the claim and no weight in the models' behavior. The authors' own conclusion, that models should treat a citation as evidence to check and not as authority, matches the rule. The gap is that the paper never tested checking. It shows the problem the rule addresses and leaves open whether the rule fixes it.

**Authority cues as weak signals (same section, second bullet, second sentence).** The prompt treats authority cues as weak signals. The paper shows a citation's presence acting as a strong signal in models that should have treated it as weak. It also shows prestige being ignored, so the models' error was reacting to citations as a category, with no extra trust for prestigious ones. The prompt's wording covers that, since a citation is itself an authority cue.

**Weigh each source on its own record (same bullet, last sentence).** In Claude Haiku 4.5, GPT-5.4 mini and Llama 3.1 8B, a citation seemed to raise suspicion across the board, cutting errors on false claims and raising them on true ones. That is uniform discounting, the failure this sentence names. The paper did not test any instruction against it.

**Citations from subagents (same section, third bullet, second sentence).** The prompt extends the citation check to subagent reports. The paper has nothing on subagents; its relevance is that a cited claim arriving from any source should be checked against the source.

## What the prompt omits

**Format and domain guidance.** The paper found leading "according to" phrasing the riskiest and law the most resistant domain. The prompt does not mention either. The format effects differ by model and the paper says safe-format guidance has to be model-specific, and a rule to open the source makes the format irrelevant.

**Author identity and recency.** The null results on surname region and the small effect of citation year give the prompt nothing to act on.

**Citation-free reasoning.** One response to the paper would be to tell the model to judge a claim first without its citation. The prompt does not, because it asks the model to check the citation itself, which is stronger when tools are available, and the paper did not test the alternative.

## Conflicts with other references

Most of the other reference files named below did not exist in `chat/system/references/` when this file was first written, so the notes come from the sources: the full texts of Pradhan and Goley and Ashkinaze et al., the abstract and results of Zhu et al. and the abstract of Onweller et al. All four files now exist, and the points below agree with them.

**Pradhan and Goley (`pradhan-trust-dont-verify.md`, arXiv 2606.05403).** The papers agree closely. Pradhan and Goley found that a seniority title had no effect on how models weighed a source, matching this paper's null for venue prestige. Both find that presentation does the work: analytical register there, citation presence here. Both find suspicion that does not discriminate: in Pradhan and Goley it came from verification prompts, and here it seems to come from the citation itself.

**Ashkinaze et al. (`ashkinaze-information-discernment.md`, arXiv 2607.19355).** Ashkinaze et al. found that models barely weighted news outlets by reliability and slightly more by popularity. That matches this paper's finding that models do not use the reputation of a venue in a useful way. The papers measure different cues, outlet popularity and citation presence, and neither tests whether looking the source up helps.

**Zhu et al. (`zhu-misknow-agent.md`, arXiv 2607.20891).** Zhu et al. found that deep research agents adopted false conclusions more often from documents attributed to high-authority institutions, by about 15 percentage points over low-authority ones, and most often from documents written in the style of a research paper. That conflicts with this paper's prestige null. The likely difference is that Zhu et al.'s institutions came with full documents written to match, while here prestige was a journal name in a citation string. The prompt treats authority as a weak signal in either case, so the conflict does not change the rule, but it does mean the prompt's grouping of authority cues rests on mixed evidence about how much models respond to them.

**Onweller et al. (`onweller-cited-not-verified.md`, arXiv 2605.06635).** Onweller et al. found that frontier deep research reports cited sources that supported the claim only 39 to 77 percent of the time, though the links worked. That is the production side of this paper's problem: models write citations that look sound and are wrong, and this paper shows models trust citations that look sound. Together they make the case for the prompt's rule to open the source from both directions.

## Review notes

I read the full text of version 1 from a local text extraction. The paper is 10 pages with no appendix, and figures did not survive extraction, so per-model results shown only in Figures 3 to 5 come from the text's description. The paper does not print the task instruction given to models or the judge prompt, so my account of what counted as a hallucination on a true claim is inferred. I could not reconcile some lift figures with the rates in the same tables; I report the paper's lift figures. The kappa of 0.83 is ambiguous between human-human and human-judge agreement. The arXiv listing gives June 11, 2026 for version 1, which meets the recency rule. I did not check the workshop acceptance beyond the arXiv comment field.
