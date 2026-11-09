# vale-ai-tells

- **Source:** Tony Burns, "vale-ai-tells," a Vale style for patterns common in AI-generated prose, MIT License, commit 51f646e (October 11, 2026)
- **Type:** Open source rule set for the Vale prose linter
- **Where to find it:** https://github.com/tbhb/vale-ai-tells
- **Last reviewed:** 2026-11-05
- **Used by:** clauditory-lint 2.1.0 (`lint/styles/AITells` and `lint/styles/AITellsCommits`)

## What the source covers

The package holds 137 rules aimed at technical documentation, plus a separate style of 16 rules for commit messages. Its author maintains it as a living rule set and tunes each rule against corpora of human-written technical prose, which the rule files' comments describe. The rules fall into several families:

- **Vocabulary and phrases** that models overuse: an overused-vocabulary list, adjective-noun pairs such as `seamless integration`, clichéd compounds, promotional puffery, and strategy buzzwords.
- **Figurative verbs and metaphors**, about 40 rules: a file that "lives in" a folder, a check that "sits" at a step, a change that "lands," growth, journey, motion, and money metaphors, each with a literal alternative.
- **Framing and structure**: strawman contrasts, contrastive formulas, mic-drop sentences, pseudo-clefts, staccato parallel sentences, stacked anaphora, and summative appositives.
- **Hedging and evasion**: stacked and defensive hedges, false balance, vague attributions, and universal claims with no stated scope.
- **Openers, closers, and headings**: opening clichés, sycophancy, closing pleasantries, conclusion markers, and announcement, marketing, and wrap-up headings.
- **Commit messages**: attribution lines, emoji, buzzwords, changelog-style lists, and file listings.

## What the linter applies

The linter vendors both styles unchanged. Most rules report as suggestions, because many of them describe habits that are fine in some sentences. About 35 high-precision rules, such as sycophancy markers, closing pleasantries, strawman contrasts, puffery, and announcement headings, report as warnings. The commit-message style runs only with `--commit`.

## What the linter omits

- **Rules that contradict the skills.** The colon rule lowercases a sentence after a colon, where the writing-style skill capitalizes it. The semicolon rule discourages semicolons, which the skill uses between related clauses. The formal-transitions rule flags "for example" and "therefore," which the skill prescribes. The consequence-clause rule flags a clause joined with "so," which the skill's own examples use. The empty-padding rules flag "given name," which the skill prescribes.
- **Rules that fire on required structure.** The explainer-headings rule flags the section headings every reference file in this repository must use.
- **Rules that fire on ordinary series.** The verb-tricolon rules flag any three parallel verbs, while the skill asks for as many items as the content has and objects only to a third item that adds nothing.
- **Duplicates.** The em-dash and double-hyphen rules duplicate the House dash rule, and the noun-string rule duplicates the skill's noun-stack guidance with more false alarms.
- **The experimental package.** Its structural rules need document-wide scripts and are opt-in even in the source.

## Conflicts with other references

The package targets technical documentation and discourages contractions in some rules' advice; the writing-style skill allows contractions in conversation, and no rule here enforces either way. Its overused-vocabulary list overlaps the skill's revision-pass vocabulary list; where both flag the same word, the linter reports it once.

## Review notes

I cloned the repository at the commit above, read the rule list and the files of every rule I turned off or raised to a warning, and ran the style over the repository's own prose. That run produced about 3,900 alerts; I sampled the most frequent rules to decide which contradict the skills. The package describes its basis as 2024 and 2025 research into vocabulary fingerprints and structural tells. It is living documentation, so it counts from the version read, but the research behind individual rules might be older than eight months, and I did not check each rule against a current study.
