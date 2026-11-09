# Vale Microsoft style

- **Source:** Vale project, "Microsoft," a Vale implementation of the Microsoft Writing Style Guide, MIT License, commit ea5fa73 (September 28, 2026)
- **Type:** Community-maintained rule set for the Vale prose linter
- **Where to find it:** https://github.com/vale-cli/Microsoft
- **Last reviewed:** 2026-11-08
- **Used by:** clauditory-lint 2.1.0 (`lint/styles/Microsoft`)

## What the source covers

The style holds 47 rules that check text against the Microsoft Writing Style Guide. Microsoft doesn't publish it; the Vale project maintains it from the guide's public pages. It covers word choice and technical formatting. Word choice includes a term list, wordiness, jargon, militaristic and gender-biased terms, and words that define people by disability. Technical formatting covers URLs, units, "auto-" compounds, and suspended hyphens. Its date and time rules ask for "July 31, 2016," no ordinal dates, and "AM" and "PM" in capitals. The rest cover punctuation (dashes without spaces, quotation marks, semicolons, exclamation points, and the serial comma), headings, voice (passive voice, the first person, and "we"), and sentence length.

## What the linter applies

The linter vendors the style unchanged and uses it for technical communication where Red Hat and Apple have no rule. The rules that stay on are those no other source contradicts or already covers: the term list, wordiness, jargon, the bias-free and accessibility lists, militaristic terms, "auto-" compounds, ordinals such as `firstly`, suspended hyphens, units, URL phrasing, heading colons and acronyms, the general list of words to avoid, and the date and range formats, which agree with NYT. None of its fixes runs automatically, because several of its replacements predate current usage, and its term list reports as suggestions.

## What the linter omits

- **Rules another source settles.** `4 PM` loses to NYT's `4 p.m.`, and the rule that turns click and tap into select loses to Apple, which uses click and tap. The rule that prefers "address" over "URL" loses to Red Hat, which uses URL.
- **Rules that contradict the skills.** The skills allow the first person and "we" in conversation, allow contractions in conversation and spell them out in formal documents, ban em dashes rather than closing them up, and don't discourage semicolons or questions.
- **Duplicates.** The serial comma, passive voice, headings, heading punctuation, hyphens after adverbs ending in ly, "(s)" plurals, quotation-mark punctuation, spacing, ellipses, Latin abbreviations, acronym definitions, adverbs, exclamation points, and sentence length each have a rule elsewhere.
- **Noise.** The vocabulary rule flags words for a manual check without stating a rule, the uppercase rule flags every acronym, and the rule on negative numbers asks for an en dash that plain text can't rely on.
- **One term.** The term list maps "agent" to "personal digital assistant," which doesn't fit software agents, so the linter ignores that match.

## Conflicts with other references

Microsoft writes "AM" and "PM" in capitals, where NYT writes "a.m." and "p.m."; the skill follows NYT. Microsoft and Apple disagree on input verbs; the linter follows Apple, whose verbs match the input device. Microsoft and the writing-style skill disagree on contractions and the first person; the skill is the owner's and overrides every source.

## Review notes

I cloned the repository at the commit above, read every rule file, and checked each rule that stays on against the repository's own prose. The Microsoft Writing Style Guide itself is covered in `chat/skills/technical-writing/references/microsoft-writing-style-guide.md`. I didn't compare the style's term list entry by entry with the current guide.
