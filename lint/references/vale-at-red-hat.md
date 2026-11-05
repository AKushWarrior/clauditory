# Vale at Red Hat

- **Source:** Red Hat Documentation, "vale-at-red-hat," the official Vale style for Red Hat documentation, MIT License, commit 7cfb4ab (September 29, 2026)
- **Type:** Open source rule set for the Vale prose linter
- **Where to find it:** https://github.com/redhat-documentation/vale-at-red-hat
- **Last reviewed:** 2026-11-05
- **Used by:** clauditory-lint 1.0.0 (`lint/styles/RedHat`)

## What the source covers

Red Hat's documentation team maintains this style to check documentation against the Red Hat supplementary style guide and the IBM Style guide that Red Hat builds on. It holds 37 Vale rules plus a spelling dictionary and one Tengo script. The rules fall into four groups.

The largest group is terminology. Four substitution rules (TermsErrors, TermsWarnings, TermsSuggestions, and CaseSensitiveTerms) map several hundred words and phrases to Red Hat's preferred forms, at three levels of certainty. Errors cover forms Red Hat never uses, such as "utilize" or "check box"; warnings cover forms to avoid unless existing text uses them; and suggestions depend on context. A separate rule (DoNotUseTerms) names terms to drop entirely, with a message for each, and another (ConsciousLanguage) replaces terms such as master and slave. A hyphenation list (Hyphens) settles open, closed, and hyphenated forms of technical compounds.

The second group covers technical text: undefined acronyms, PascalCase or camelCase words that belong in code formatting, user-replaced values, session IDs in URLs, merge-conflict markers, and links to repositories that need approval. The third covers documentation structure: sentence-case headings, no end punctuation in headings, no gerunds in topic titles, and self-referential text such as "this topic describes."

The fourth group covers general prose: passive voice, contractions, sentences that start with a conjunction, simple words, slashes, ellipses, em dashes, the serial comma, repeated words, spacing, symbols, sentence length, and a readability grade.

Most rules report problems only. Substitution rules carry Vale's replace action, and one rule (Using) carries an edit action that rewrites a noun followed by "using" as "by using." Vale itself never applies these actions; its editor integration offers them as quick fixes.

## What the linter applies

The linter vendors the style unchanged and treats it as the primary source for technical communication. The terminology rules, the technical-text rules, and the structure rules apply at that rank. The linter applies the fixes from TermsErrors, CaseSensitiveTerms, Hyphens, SmartQuotes, and RepeatedWords as safe fixes, because each match has one replacement. ConsciousLanguage, SimpleWords, and Using apply only with `--fix=all`, because their replacements sometimes change a sentence's meaning or grammar.

The general-prose rules (passive voice, simple words, slashes, ellipses, spacing, symbols, and repeated words) stay on, but the linter ranks them by their real domain. A Red Hat rule about grammar or word choice ranks below Oxford, NYT, AP, and Chicago in that domain, so it applies only where none of those disagrees.

## What the linter omits

- **Spelling.** The spelling rule flags model names, author names, and other proper nouns outside Red Hat's dictionary, which in this repository's text means most flagged words are correct. It is off by default and turns on with `--spelling`.
- **Release notes and repository links.** ReleaseNotes applies only to Red Hat release notes, and GitLinks enforces Red Hat's internal policy on which repositories may be linked.
- **Rules the skills or a better-ranked source settle.** Contractions is off because the writing-style skill allows contractions in conversation. Conjunctions is off because NYT allows a sentence to start with "And" or "But." Abbreviations is off because NYT writes some abbreviations with periods. EmDash, OxfordComma, and SentenceLength duplicate House and Oxford rules; EmDash also reads inside code spans. ReadabilityGrade is off because the skill sets no grade-level target.
- **Unusable fixes.** DoNotUseTerms puts its explanation where the replacement would go, so its matches are never fixed automatically, and TermsWarnings maps some words to replacements that change the meaning (it suggests "you" for "he"), so it reports without fixing.

## Conflicts with other references

Red Hat and Apple disagree on several terms, and Red Hat, as the primary technical source, wins wherever both have a rule. Red Hat's rules on word choice and grammar rank below the news and book manuals, so NYT's abbreviations and contractions in conversation prevail over Red Hat's rules on them. Chicago and Red Hat agree on the serial comma and on closing up prefixes, and the linter reports such agreements once, under the better-ranked rule.

## Review notes

I cloned the repository at the commit above, read every rule file, and ran the style against the repository's own prose and against a test file. The rules' guidance comes from Red Hat's supplementary style guide, which this repository's technical-writing references cover in more depth (`chat/skills/technical-writing/references/red-hat-technical-writing-style-guide.md`). The style declares that it needs Vale 2.20 or later, and the linter runs it on Vale 3.26.0 without errors.
