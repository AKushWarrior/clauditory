# Rules taken from the style guides

- **Sources:** The New York Times Manual of Style and Usage; the AP Stylebook; The Chicago Manual of Style; the New Oxford Style Manual (New Hart's Rules and the New Oxford Dictionary for Writers and Editors); the IBM Style Guide; the Apple Style Guide; the Google developer documentation style guide; the Microsoft Writing Style Guide; and the Red Hat Technical Writing Style Guide
- **Type:** Style manuals; the rules are this project's encodings of them
- **Where to find them:** the reference file for each guide in `chat/skills/writing-style/references/` and `chat/skills/technical-writing/references/`; the extracted rules in `lint/sources/`
- **Last reviewed:** 2026-11-05
- **Used by:** clauditory-lint 2.0.0 (`lint/styles/House`, `Oxford`, `NYT`, `AP`, `Chicago`, and `Apple`, plus `lint/terms.gen.json`)

## What the sources cover

Most of each guide is advice that needs judgment, such as when the passive voice is right or how to attribute a claim. A smaller part fixes a form exactly: how to spell, hyphenate, or capitalize a word or name; which word to use instead of a banned one; how to write a time, a date, a number, an abbreviation, or a unit; and where a mark of punctuation goes. The linter encodes only that smaller part. What each guide covers in full is in its reference file under the skills.

## What the linter applies

The rules come in two layers.

**Hand-written rules.** About 60 rules in `lint/styles` encode the writing-style skill's own choices (the House style) and the best-known mechanical rules of Oxford, NYT, AP, Chicago, and Apple. They cover the serial comma, hyphens after adverbs, apostrophes in plurals, the article before an abbreviation, `percent`, times such as `4 p.m.`, the month before the day, small numbers, American spellings, Apple product names, and more. Each rule links to the reference file for its guide.

**Extracted rules.** The rest were taken from the guides page by page. Each guide was split into 100-page chunks (the PDFs of NYT, AP, Chicago, Oxford, IBM, and Apple) or line ranges (the web guides of Microsoft and Google and the Red Hat guide's source), and a model-run extraction read every page of each chunk and recorded every rule a pattern can check: the wrong form, the right form where there is one, the kind of rule, whether the rule holds in American English, and how safe an automatic fix is. That pass produced about 12,300 candidate rules, kept in `lint/sources/`. A generator then:

- drops British-only rules, since the linter writes American English, and treats Oxford's word-choice rules as suggestions without automatic fixes, because the dictionary often prefers British idiom;
- drops rules that would flag a form the skills prescribe (email, web page, toward, contractions in conversation, and others) or that would replace a word with one the skills flag;
- settles disagreements between guides over the same word, keeping one guide's rule and dropping the rest, and drops exact duplicates;
- turns finite patterns, such as `cyber attack` or `cyber-attack`, into plain word lists, which the linter matches with a hash lookup, and keeps about 1,000 true patterns, each checked with Vale's own regular-expression engine;
- drops every entry that fires at least three times across at least two files of this repository's own prose, which is written to the skills and reviewed, so a rule that fires on it repeatedly contradicts the skills or is too broad.

About 13,600 word-list entries and 1,000 patterns survive. Word-list entries the guides call errors, and that have one replacement, report as warnings with an automatic fix; the rest report as suggestions.

## What the linter omits

Rules that need judgment stay in the skills, as do rules tied to a publication's context, such as NYT's honorifics, AP's datelines, Chicago's bibliography formats, and Hart's rules for typesetting foreign languages. Some forms a guide settles were left out because a pattern would misfire too often, such as AP's possessive of a singular name ending in s. The pruning pass also removes some correct rules that this repository's prose breaks on purpose, such as forms its reference files quote while discussing them.

## Conflicts with other references

The guides disagree often, on the serial comma, the percent sign, abbreviated months, contractions, and many single spellings. The linter never reports a rule that contradicts the writing-style or technical-writing skills, and where guides disagree with each other on a word, it reports one guide's form, so a fix can't undo another fix. The reference file for each guide records its main conflicts with the skills.

## Review notes

The extraction ran over every page of each guide's chunks; the agents reported a few pages as blank scans and some page numbers as approximate. AP's back-of-book index and several dictionary sections list headwords without rule text, and the extraction inferred some wrong forms from them, such as `air fare` from the headword `airfare`, so those rules carry suggestion severity unless the guide states the rule. The extraction was model-run and its output wasn't checked entry by entry against the guides; the generator's filters and the pruning pass are the checks on it. Three scanned Chicago chunks (pages 201 to 500) were still being read when this version was built.
