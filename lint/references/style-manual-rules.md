# Rules written from the style manuals

- **Sources:** New Hart's Rules in the New Oxford Style Manual (Oxford University Press, 2016); The New York Times Manual of Style and Usage; the AP Stylebook; The Chicago Manual of Style; the Apple Style Guide; and the writing-style skill for the House rules
- **Type:** Style manuals; the rules are this project's own Vale encodings of them
- **Where to find them:** the reference files for each manual in `chat/skills/writing-style/references/` and `chat/skills/technical-writing/references/`
- **Last reviewed:** 2026-11-05
- **Used by:** clauditory-lint 1.0.0 (`lint/styles/House`, `Oxford`, `NYT`, `AP`, `Chicago`, and `Apple`)

## What the sources cover

No published Vale style exists for Oxford, NYT, AP, Chicago, or Apple; a search of GitHub and the Vale package list on November 5, 2026 found none. The five manuals are books or web guides written for people. Most of their rules need judgment, such as when to use the passive voice or how to attribute a claim, and a linter can't check those. A smaller set fixes a form exactly: how to write a time or a date, where a comma or a period goes, which spelling of a word to use, and which article goes before an abbreviation. The linter encodes only that smaller set. What each manual covers in full is in its reference file under the skills.

## What the linter applies

<!-- vale off -->

**House (owner).** These rules encode the writing-style skill's own choices, which outrank every source: em dashes and spaced hyphens used as dashes; correction framing; trailing participles; recap openers; stock openers and closers; announcements; the overused-vocabulary list; hedge stacks; inflated verbs; the plain-word list; intensifiers; clichés; e.g. and i.e. in running text; etc.; and/or; "(s)" plurals; sentences longer than 40 words; a bare "This" or "These" before a verb; exclamation points; cannot, all right, no one, and "consists of"; all-figure dates; and "between 10 to 20." Each maps to a rule in sections 2, 3, 8, or 10 of the skill.

**Oxford (grammar and punctuation, primary).** New Hart's Rules supplies the serial comma, no hyphen after an adverb ending in -ly, no apostrophe in plurals such as the 1990s and CDs, the article chosen by the sound of an abbreviation, and no space before punctuation. For quotation marks, Hart's gives the American practice for American text: periods and commas go inside the closing mark. The linter writes American text, so it follows that practice.

**NYT (word choice and formatting, primary).** The Times manual supplies "percent" in prose, times written as 4 p.m., noon, and midnight, the month before the day, no ordinal days, one through nine spelled out, no figure at the start of a sentence, commas in large numbers, lowercase "internet," and the spellings email, website, online, OK, toward, among, while, and nonprofit.

**AP (grammar secondary, word choice tertiary).** The Stylebook supplies "who" for people, no comma before Jr. or Sr., and the spellings afterward, backward, forward, upward, downward, and adviser. Its rule against the serial comma is encoded but turned off, because Oxford ranks higher in grammar.

**Chicago (word choice secondary, grammar tertiary).** The Manual's usage glossary supplies irregardless, could care less, orientate, preventative, supposably, alot, based off of, off of, try and, and very unique; "different from" before a noun; object pronouns after a preposition; closed-up prefixes; American spellings; "web page"; and the comma after the year in a full date. Its rule that spells out numbers through one hundred is encoded but turned off, because NYT's one-through-nine rule ranks higher.

**Apple (technical secondary).** The Apple Style Guide supplies the capitalization of Apple product names, Wi-Fi, and Bluetooth; onscreen; keyboard shortcut instead of hot key; press, not hit, for Return; and no "please" in instructions. Its "webpage" is encoded but turned off, because spelling a general word is a question of word choice, where Chicago ranks higher.
<!-- vale on -->

## What the linter omits

Rules that need judgment stay in the skills: whether a sentence is passive for a good reason, whether a list is parallel, whether "that" or "which" fits, whether a number is data, and most of NYT's and AP's entries on people, titles, and places. Rules that depend on publication context also stay out, such as NYT's honorifics and AP's datelines. Some forms the manuals settle were left out because a pattern would misfire too often, such as AP's possessive of a singular name ending in s, which a linter can't tell from a plural possessive.

## Conflicts with other references

The hierarchy settles the conflicts among these manuals: the serial comma (Oxford over AP), small numbers (NYT over Chicago), times (NYT over Microsoft), and "web page" (Chicago over Apple). Where the manuals agree with Red Hat or Microsoft, as on "4 p.m.," the serial comma, and no hyphen after -ly adverbs, the linter reports the problem once under the better-ranked rule.

## Review notes

I checked the Oxford rules against the text of New Hart's Rules in the New Oxford Style Manual. The NYT, AP, Chicago, and Apple rules come from this repository's reference files for those manuals and from general knowledge of their entries; I didn't check each word-list entry against the current edition, and I dropped two NYT spellings (health care and cellphone) that I couldn't confirm. Every rule ran against a test file and against the repository's own prose before it shipped, and patterns that flagged correct text more than rarely were narrowed or dropped.
