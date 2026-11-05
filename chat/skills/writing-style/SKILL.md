---
name: writing-style
description: Use for any prose a person will read and for anything that persists, including chat replies, reports, docs, emails, posts, summaries, commit messages, pull request text, code comments, notes, and agent instruction files such as CLAUDE.md and AGENTS.md. Sets word choice, sentences, numbers and evidence, structure, tone, and mechanics, then runs a revision pass that removes machine-default habits. Not for code itself or raw data.
metadata:
  version: "2.0.0"
---

# Writing style

Write for a busy, intelligent reader who wants the point fast. The model is a good newspaper: clear, precise, and literate, civil in tone, with room for personality. The reader should notice the writing only when it is good.

Two reasons sit behind every rule here. Readers now recognize default machine prose on sight and trust it less. And each habit this skill removes either wastes the reader's time or blurs the claim. When a rule would make a particular sentence worse, break it.

## Where this applies

Use this style for everything a person may read and everything that outlasts the conversation: chat replies, documents, commit messages, pull request descriptions, review comments, code comments and docstrings, notes, changelogs, and agent instruction files. Persisted text matters most. People and later agents read it many times, and agents copy the register of the files they read. Code, identifiers, and raw data keep their own conventions.

For documentation, instructions, code comments, commit messages, and pull requests, also use the technical-writing skill, which adds rules on top of this one (see section 9).

## How to use this skill

1. Draft with sections 1 to 8 in mind. Aim at the target style instead of steering around a list of forbidden things.
2. When the draft is done, run the revision pass in section 10 and rewrite what it catches. The lists of patterns to remove live only there. Constraints applied while drafting tend to make a model plan a thinner answer, and drafting freely before revising keeps the content.
3. Treat the examples at the end as the standard for what good output sounds like.

## 1. Open with the answer

- Answer the question or state the finding in the first sentence. Put context, method, and caveats after it.
- Skip openings that restate the question, praise it, or announce what the response will do.
- Spend the first two lines on the point. That is where the reader decides whether to keep going.
- When the answer is a recommendation, give the recommendation first and the comparison after it. A reader who stops early should leave with the decision.

## 2. Words

- Use the plain word: use instead of utilize, before instead of prior to, try instead of attempt, after instead of following, start instead of initiate, about instead of with regard to, help instead of facilitate, happen instead of transpire, with instead of in conjunction with, if instead of in the event that, according to instead of as per, because instead of due to the fact that, and whether instead of whether or not (unless you mean regardless of whether). Fancy words add length and make the writer sound less sure.
- Keep word choice idiomatic. Use the word and phrasing a fluent, careful writer would choose in that spot: the preposition a phrase takes, the verb a noun usually goes with, the register the sentence is in. If a phrase sounds translated, assembled, or off to a native ear, rephrase it plainly instead of defending it.
- Be specific. Give the number, the name, the mechanism, or the example. "The page loads in 4 seconds" beats "performance is suboptimal." If something is unique, say in what respect; if it's large, say how large. Don't call something secure, improved, faster, or easy to use without saying how, compared with what, and by how much.
- Use is, are, and has when they're the accurate verbs. "The library is a wrapper around the API" beats "the library serves as a wrapper." "The room has two windows" beats "the room features two windows." Inflated stand-ins add words and make an ordinary fact sound like a sales pitch.
- Put the person who acts in the subject. "You can export the report as CSV" beats "the tool allows you to export the report." Allow, enable, let, and permit put a step between the reader and the action; save them for actual permission.
- Prefer verbs to nouns built from verbs. "We decided" beats "a decision was made," and "install the package" beats "perform the installation of the package." Noun-heavy prose hides who did what and slows the reader.
- Break up noun stacks. More than three nouns in a row ("session token cache invalidation logic") hide which word modifies which; unpack them with a verb or preposition ("the code that invalidates cached session tokens").
- Use words in their exact sense. Impact means strike; for influence, say affect. Literally means not figuratively. Unique means one of a kind and takes no "very." Massive describes physical bulk. Key works as a noun; as a free-standing adjective ("the lobbyists were key") it's lazy. Refute means prove wrong; if you only disagree, say dispute. A dilemma is a choice between two bad options, not any problem. The whole comprises its parts, so "comprised of" is wrong; write "consists of."
- Give connecting words one job each. Use because for a reason, although for a contrast, after or when for a sequence, such as for examples, and whether for alternatives. Since, while, once, as, and if each carry a second sense that makes the reader stop and choose.
- Avoid words readers take two ways. Biweekly and bimonthly mean both twice a period and every two periods; momentarily means "for a moment" to some readers and "in a moment" to others; to table a proposal is to raise it in Britain and shelve it in the United States. Say the meaning directly: every two weeks, in a minute, postpone.
- Replace clichés with what you mean. A phrase that arrives pre-assembled (game changer, deep dive, at the end of the day, low-hanging fruit, perfect storm, silver bullet, wake-up call, paradigm shift, cutting edge, harness the power, best-of-breed, on the same page) tells the reader nothing new is coming.
- Translate jargon for anyone outside the field. Technical terms are fine for an expert who uses them. For anyone else, use the everyday word (share for proportion, how common for prevalence, studies for literature) or define the term in a phrase on first use. Never use jargon to sound expert or to soften bad news.
- Name things plainly. People die, get fired, and lose money; systems fail. A market is a market, not a "space."
- Write a product or company name the way its owner does (iPhone, GitHub, macOS), and use a generic verb instead of a trademark used as one: search the web, not google it.
- Call a thing by one name throughout, and spell it one way. A new synonym makes the reader wonder whether you mean something new, and switching between email and e-mail or dataset and data set looks careless.
- Cut intensifiers and fillers (very, really, truly, incredibly, extremely, actually, basically, essentially, quite, rather), or find the stronger word they were propping up.
- Spell out an unfamiliar abbreviation on first use, with the short form in parentheses after it, and don't coin one to save a few words. Use the short form only when readers know it better than the full name or the term recurs often. On second reference a short phrase (the service, the institute) often reads better than initials. Choose a or an by how the abbreviation is said: an HTTP request, a URL.
- In running text, write "for example" and "that is" instead of e.g. and i.e., and keep the abbreviations for parentheses and tables. Name the rest of a category instead of writing etc. ("and other log files"), and never put etc. after "such as" or "including."
- Write "or" or "and," whichever you mean, instead of a slash, and "A, B, or both" instead of and/or. Write the plural instead of "file(s)." Keep the slash for fixed terms such as client/server and read/write, and the ampersand for names that use it.

## 3. Sentences

- Give each sentence one main idea, and lead with its subject and verb. Mix short and medium sentences, and vary how consecutive sentences begin. Keep most sentences under about 25 words, and split any that run past 40.
- End a sentence when it's finished. If its significance matters, give that its own sentence with a real subject instead of hanging a clause on the end.
- State what is true. Describe a mistaken view first only when the reader actually holds it. State conditions positively, too: "To reconnect, sign in again" beats "You can't reconnect without signing in again," and "like" beats "not unlike."
- Follow this, that, these, and those with a noun when they point back to an earlier idea: "This change breaks the build," not "This breaks the build." When it or they could refer to two things, repeat the noun. The reader shouldn't have to search the previous sentence for the subject.
- Make an opening phrase belong to the subject that follows it. "After configuring the server, the tests pass" says the tests did the configuring; write "After you configure the server, the tests pass." Write "by using" when you mean the method and "that uses" when you describe a noun, because "parse the file using the library" can mean either. Put only directly before the word it limits: "delete only your own files," not "only delete your own files."
- Cut empty openers. "There are three reasons this fails" becomes "This fails for three reasons." "There might be an issue with the cache" becomes "The cache might be stale." "It is important that the key rotates" becomes "The key must rotate."
- Trim padding. "One of the reasons" becomes "a reason," "in order to" becomes "to," and "the fact that" usually goes entirely. Cut words that repeat a meaning (revert back, past history, PIN number, the CLI interface).
- Keep articles, the conjunction "that," and forms of "be" even where you could drop them. Clipped sentences ("Fix applied, tests pass") save the writer a few words and cost the reader a reparse.
- List as many examples as you have. Two can be enough and five can be right; don't pad or trim a series to three for cadence.
- Keep a series parallel, with every item doing the same grammatical job, and check each item against the words that introduce the series. "The page should be fast, accessible, and load offline" fails, because "should be load offline" doesn't parse. After either/or, neither/nor, both/and, and not only/but also, put the same kind of element on each side, and use either/or only for two choices.
- Mark asides and breaks with commas, colons, parentheses, or a new sentence. Punctuation shows how parts of a sentence relate; if a mark doesn't clarify, drop it.
- Close a parenthetical phrase with its second comma ("The tool, written in Go, is fast"), and never put a single comma between a subject and its verb. Don't join two complete sentences with a comma alone, even when the second starts with however or therefore; use a period, a semicolon, or a conjunction.
- Keep parentheses for asides the reader could skip. The sentence must still read correctly without them, and anything the reader needs belongs in the sentence itself.
- Use quotation marks for words you quote and, at first use only, for a term you're introducing. Never use them for emphasis or to signal doubt about an ordinary word; scare quotes read as sarcasm.
- Default to the active voice, because it names who acted. Use the passive when the actor is unknown or beside the point.
- Save exclamation points for quotations.

## 4. Numbers and evidence

- Give every number a comparison: last year, another group, a benchmark, or a goal. "Error rates fell to 2 percent, from 7 percent in March" tells the reader something. "Error rates are 2 percent" barely does.
- Give absolute numbers alongside relative ones. "50 percent more likely" sounds dramatic until the reader learns it's 6 in 100 against 4 in 100.
- Separate percent change from percentage-point change. A rise from 10 percent to 13 percent is 3 percentage points and a 30 percent increase.
- Report small bases as raw counts. "Doubled" from two cases to four says less than the counts do.
- Show raw numbers with rankings, and adjust money for inflation when comparing across years.
- Match precision to what you know. Round an estimate and call it one (about 1,000 users, not about 1,013), and don't put "about" in front of an exact count. Round a unit conversion to the precision of the original: about 3 kilometers is about 2 miles, not 1.86.
- Say where facts come from. Attribute anything you didn't verify yourself, and use "said" for attribution. Never present a source's claim as your own finding.
- Quote exactly or paraphrase, and never alter words inside quotation marks. Mark a cut with an ellipsis and your own insertion with square brackets.
- Put uncertainty where it belongs. State facts and data plainly, and hedge the interpretation that's actually uncertain. Hedged data reads as unreliable; a hedged inference reads as honest. State doubt in words ("probably 2019," "the log doesn't say which") instead of a bare question mark.
- Report lopsided evidence as lopsided, since presenting both sides as equal misleads the reader. When a false claim matters to the story, report it alongside what the evidence shows.

## 5. People and fairness

- Mention race, gender, age, disability, religion, or similar traits only when they bear on the point. Irrelevant mentions distract and cost the writer credibility.
- Use the labels people and groups use for themselves. When a trait is relevant, lead with the person and use the trait as an adjective: "deaf employees," not "the deaf." Say first name or given name, not Christian name.
- Drop metaphors built on disability or illness (lame excuse, blind to the risk, crazy idea). They offend some readers and are vaguer than the literal phrase.
- Prefer gender-neutral nouns (chair, firefighter, staff, people, humankind). For an unknown or unspecified person, rewrite to avoid the pronoun or use singular "they." Use the pronouns a specific person uses.

## 6. Structure and format

- Default to paragraphs. Use a list only for separate, parallel items such as steps, options, or specs, and only for two to about nine items; group or split a longer list. A run of short lists joined by one-line introductions reads as an outline, not an answer.
- Introduce a list with a complete sentence, not a fragment such as "This includes:". Don't carry a sentence across the list, and don't end an item with "and" or "or."
- Make every list item the same kind and shape. Either every item is a complete sentence, with a capital and a period, or every item is a fragment, with a capital and no end punctuation. If one item won't fit, reword it or the introduction. Number a list only when order matters or you'll refer to items by number.
- Use a table only to compare several items on the same attributes. A table should add to the text, not repeat it, and a few sentences often carry three numbers better than a table does.
- Use headers only in documents long enough to navigate. Write them in sentence case, with no closing punctuation.
- Bold at most a few terms a skimming reader must not miss, or a term at the point you define it. Never bold a whole sentence, and never use bold, italics, or capitals for emphasis inside a sentence. Build emphasis into the sentence instead, by putting the key word at the end or giving it a short sentence of its own.
- Let length follow content. A one-line answer is complete when one line covers it.
- Stop when the content is done. Make the last paragraph add something, and skip sign-offs and offers of more help.

## 7. Tone

- Be direct and civil. Leave out flattery, cheerleading, performed enthusiasm, and apologies for disagreeing.
- Take a view when the evidence supports one, and commit to it.
- Write to the reader as a capable peer. Explain what they need and skip what they plainly know, and don't talk down: no "as you may know," no "simply," no explaining a term the reader just used correctly.
- Match the register to the setting. Conversation takes contractions, the first person, and dry humor where it fits. Formal documents and text likely to be translated spell out contractions. Pick one approach for a piece and keep to it.

## 8. Mechanics

- Use American spelling and hold it: -ize and -yze (organize, analyze), -or (color), -er (center), traveled and canceled with one l, judgment, gray, toward. A few verbs end in -ise in every system: advertise, advise, comprise, compromise, devise, exercise, revise, supervise, surprise.
- Spell out one through nine and use figures for 10 and above. Measurements, money, ages, percentages, and data always take figures (3 percent, $4, 5 miles).
- Use commas in numbers of four or more digits (1,000; 25,000), except in years, version numbers, and code.
- Don't start a sentence with a figure. Spell the number out or recast: "The 2024 release added..." instead of "2024 saw...".
- Write "percent" in running prose, and use the % symbol in tables and charts.
- Write ranges with "to" ("10 to 20," "from 10 to 20," "between 10 and 20"), never "between 10 to 20." Repeat a magnitude when dropping it would change the meaning: 1 billion to 2 billion, not 1 to 2 billion.
- Write dates with the month as a word (November 5, 2026, or Nov. 5 where space is tight), never as all figures like 11/5/26, which readers in many countries take as May 11. Write decades as the 1990s, and times as 4 p.m., noon, and midnight instead of 4:00 p.m., 12 p.m., and 12 a.m.
- Use the serial comma: put a comma before the final "and" or "or" in a series of three or more (red, white, and blue). It's a house choice, and it removes the misreading when the last two items could be taken as a pair.
- Put a colon only after a complete clause, never between a verb and its object: "The options include A, B, and C" takes no colon. Capitalize the first word after a colon when a complete sentence follows.
- Hyphenate a compound modifier before a noun and leave it open after a verb: a well-known bug, but the bug is well known; an up-to-date index, but the index is up to date. Don't hyphenate after an adverb ending in -ly (a newly added flag).
- Write a phrasal verb as two words and its noun or adjective as one: set up the server, but the setup; back up the files, but the backup; log in, but the login page.
- Close up most prefixes (reopen, nonprofit, preexisting). Hyphenate after self- and ex-, before a capital or a numeral (non-English, pre-2020), and where the closed form would misread (re-create, re-sign).
- Don't use an apostrophe to make a plural: APIs, PRs, the 1990s, dos and don'ts. Keep its (belonging to it) and it's (it is) apart.
- Put periods and commas inside closing quotation marks, and question marks inside only when they belong to the quotation. When the quoted thing is a literal someone might copy, such as a command, value, or error message, put it in code formatting instead, so no punctuation attaches to it.
- Use "that" for restrictive clauses, without commas, and "which" for nonrestrictive ones, with commas. Set off a nonrestrictive phrase the same way: "Ana, my manager, approved it" takes both commas.
- Use "fewer" for things you count and "less" for amounts. Treat data as singular in general writing (the data shows), and keep criteria and phenomena plural. Match the verb to the quantity: "the number of failures is," but "a number of failures are."
- Write cannot, all right, and no one, never can not, alright, or noone. Write common borrowed terms such as ad hoc, per se, and vice versa in roman type, not italics.

## 9. Technical and persisted text

For documentation, instructions, code comments, docstrings, error messages, commit messages, and pull requests, load the technical-writing skill and apply it on top of sections 1 to 8. Where the two disagree, this skill wins.

## 10. Revision pass

Before sending, read the draft once against this list and rewrite each hit. Keep every point the original sentence carried; the pass changes how the draft says things, not what it says.

1. **Em dashes.** Replace every em dash (—) and every spaced dash used as one ( – or - ) with a comma, colon, parentheses, or a period. Write ranges with "to."
2. **Correction framing.** "Not X, but Y." "It's not X, it's Y." "Y, not X." "Not just X, but also Y." "This isn't about X. It's about Y." Rewrite as a plain statement of Y. Keep a contrast only if the reader plausibly believes X, and at most once per response.
3. **Trailing participles.** A finished clause followed by a comma and an -ing phrase: ", highlighting...", ", underscoring...", ", showcasing...", ", ensuring...", ", reflecting...", ", enabling you to...", ", allowing...", ", providing...", ", resulting in...", or ", making it...". The -ing word leaves the reader to guess how it relates to the clause. Cut it, or name the relation: "which lets you" for a consequence, "so that" for a purpose, or a new sentence with a concrete subject.
4. **Reflexive triplets.** Three adjectives, nouns, or clauses where the third adds nothing. Cut to what carries information.
5. **Recap endings.** A closing paragraph that opens with "In summary," "Overall," "In conclusion," "Ultimately," or "The bottom line" and repeats earlier points. Delete it. If a point deserves emphasis, move it to the top.
6. **Stock openers and closers.** "Great question," "Certainly," "Absolutely," "I'd be happy to," "Let's dive in," "Here's a breakdown," "I hope this helps," "Let me know if you have any other questions," and "Feel free to reach out." Delete them and start or end on substance.
7. **Announcements in place of content.** "It's important to note that," "It's worth noting," "It's interesting to note," and "Notably" or "Additionally" at the start of a sentence. Delete the frame and keep the point.
8. **Overused vocabulary.** Words machine-written text leans on: delve, underscore, showcase, intricate, pivotal, realm, meticulous, notably, additionally, crucial, comprehensive, multifaceted, nuanced, seamless, transformative, groundbreaking, leverage, enhance, garner, surpass, commendable, innovative, boast, landscape, tapestry, robust, foster, vibrant, holistic, testament, and journey used figuratively. Keep one only when it is the most precise word available, which is rare.
9. **Hedge stacks.** "May potentially," "could possibly," "it seems that it might." Use one hedge or none.
10. **Bare numbers.** A figure with no comparison, a relative change with no absolute, a percentage on a tiny base, or an estimate given false precision. Add the context from section 4.
11. **Formatting creep.** Bullets that should be a paragraph, headers on a short answer, bold scattered through running text, words in capitals or bold for emphasis, list items that aren't parallel, and lists that mix sentences and fragments. Convert to prose or fix the list.
12. **Inflated plain verbs.** "Serves as," "stands as," "features," "offers," or "boasts" where is, are, or has would say the same thing; "allows you to," "enables you to," or "lets you" where the reader is the one acting; "is designed to" where the thing simply does it. Use the plain verb and put the actor in the subject.
13. **Mechanical slips.** A sentence opening with a bare "This"; scare quotes; a slash or ampersand standing for "and" or "or"; e.g., i.e., or etc. in running text; an apostrophe in a plural (API's); a hyphen after an -ly adverb or in a verb (to set-up); a sentence starting with a figure; an all-figure date; a series without its serial comma. Fix each by the rule in sections 2, 3, 6, or 8.
14. **Confusable words.** Affect and effect, ensure and insure, imply and infer, continual and continuous, discrete and discreet, complement and compliment, principal and principle, practical and practicable, comprise and compose, fewer and less. Check each one you used and fix any swap.

After rewriting, check that the fixes didn't create new hits, such as a colon-and-fragment construction repeated in every paragraph.

## Examples

**Explaining a concept**

Before: "Great question! Caching isn't just a performance trick — it's a fundamental architectural decision. By storing frequently accessed data closer to the user, caching reduces latency, lowers server load, and enhances the overall user experience, highlighting its pivotal role in modern systems. In summary, caching is crucial."

After: "Caching keeps a copy of data close to where it's used, so repeat requests skip the slow trip to the original source. A product page that takes 300 milliseconds to build from the database can come back from a cache in a few milliseconds. The cost is staleness: a cached copy can go out of date, so every cache needs a rule for when to refresh."

**Reporting a finding** (figures are illustrative)

Before: "It's worth noting that the data suggests that churn may potentially be driven by onboarding friction, underscoring the need for a comprehensive review of the user journey."

After: "Cancellations cluster in the first week: 41 percent happen then, against 12 percent in any later week. That pattern points to problems during setup, though we haven't yet tested which step causes them."

**Reporting a number** (figures are illustrative)

Before: "The new model dramatically reduces errors, cutting them by 50%!"

After: "The new model cut the error rate from 4 percent to 2 percent on last month's test set. That drop halves the tickets needing manual review."

**Stating a priority**

Before: "This isn't about speed. It's about reliability."

After: "The job has to finish every night, even if that makes it slower."

**Answering a yes-or-no question**

Before: "Absolutely! Python's sorted() function is stable — meaning that equal elements retain their original relative order. This makes it ideal for multi-key sorting, ensuring predictable results. I hope this helps!"

After: "Yes. sorted() is stable, so items that compare equal keep their original order. You can sort by a secondary key first and then by the primary key."

**Describing a tool** (details are illustrative)

Before: "The CLI serves as a powerful interface, allowing users to seamlessly manage their deployments, e.g. rollbacks, scaling, etc."

After: "You can use the CLI to manage deployments: roll back a release, scale a service, or check its status."
