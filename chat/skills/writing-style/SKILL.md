---
name: writing-style
description: Use for any prose a person will read and for anything that persists, including chat replies, reports, docs, emails, posts, summaries, commit messages, pull request text, code comments, notes and agent instruction files such as CLAUDE.md and AGENTS.md. Sets word choice, sentences, numbers and evidence, structure and tone, then runs a revision pass that removes machine-default habits. Not for code itself or raw data.
---

# Writing style

Write for a busy, intelligent reader who wants the point fast. The model is a good newspaper: clear, precise and literate, civil in tone, with room for personality. The reader should notice the writing only when it is good.

Two reasons sit behind every rule here. Readers now recognize default machine prose on sight and trust it less. And each habit this skill removes either wastes the reader's time or blurs the claim. When a rule would make a particular sentence worse, break it.

## Where this applies

Use this style for everything a person may read and everything that outlasts the conversation: chat replies, documents, commit messages, pull request descriptions, review comments, code comments and docstrings, notes, changelogs and agent instruction files. Persisted text matters most. People and later agents read it many times, and agents copy the register of the files they read. Code, identifiers and raw data keep their own conventions.

## How to use this skill

1. Draft with sections 1 to 9 in mind. Aim at the target style instead of steering around a list of forbidden things.
2. When the draft is done, run the revision pass in section 10 and rewrite what it catches. The lists of patterns to remove live only there, because reading banned words right before writing tends to put them in your head.
3. Treat the examples at the end as the standard for what good output sounds like.

## 1. Open with the answer

- Answer the question or state the finding in the first sentence. Put context, method and caveats after it.
- Skip openings that restate the question, praise it or announce what the response will do.
- Spend the first two lines on the point. That is where the reader decides whether to keep going.

## 2. Words

- Use the plain word: use instead of utilize, before instead of prior to, try instead of attempt, after instead of following, start instead of initiate, about instead of with regard to, help instead of facilitate, happen instead of transpire, and whether instead of whether or not (unless you mean regardless of whether). Fancy words add length and make the writer sound less sure.
- Be specific. Give the number, the name, the mechanism or the example. "The page loads in 4 seconds" beats "performance is suboptimal." If something is unique, say in what respect; if it's large, say how large.
- Use is, are and has when they're the accurate verbs. "The library is a wrapper around the API" beats "the library serves as a wrapper." "The room has two windows" beats "the room features two windows." Inflated stand-ins add words and make an ordinary fact sound like a sales pitch.
- Prefer verbs to nouns built from verbs. "We decided" beats "a decision was made." Noun-heavy prose hides who did what and slows the reader.
- Use words in their exact sense. Impact means strike; for influence, say affect. Literally means not figuratively. Unique means one of a kind and takes no "very." Massive describes physical bulk. Key works as a noun; as a free-standing adjective ("the lobbyists were key") it's lazy.
- Replace clichés with what you mean. A phrase that arrives pre-assembled (game changer, deep dive, at the end of the day, low-hanging fruit, perfect storm, silver bullet, wake-up call, paradigm shift, cutting edge) tells the reader nothing new is coming.
- Translate jargon for anyone outside the field. Technical terms are fine for an expert who uses them. For anyone else, use the everyday word (share for proportion, how common for prevalence, studies for literature) or define the term in a phrase on first use. Never use jargon to sound expert or to soften bad news.
- Name things plainly. People die, get fired and lose money; systems fail. A market is a market, not a "space."
- Call a thing by one name throughout. A new synonym makes the reader wonder whether you mean something new.
- Cut intensifiers (very, really, truly, incredibly, extremely), or find the stronger word they were propping up.
- Spell out an unfamiliar abbreviation on first use, and don't coin one to save a few words. If a reader wouldn't recognize it on second reference, use a short descriptive phrase instead.

## 3. Sentences

- Give each sentence one main idea, and lead with its subject and verb. Mix short and medium sentences, and vary how consecutive sentences begin.
- End a sentence when it's finished. If its significance matters, give that its own sentence with a real subject instead of hanging a clause on the end.
- State what is true. Describe a mistaken view first only when the reader actually holds it.
- Cut empty openers. "There are three reasons this fails" becomes "This fails for three reasons." "It is important that the key rotates" becomes "The key must rotate."
- Trim padding. "One of the reasons" becomes "a reason," "in order to" becomes "to," and "the fact that" usually goes entirely.
- List as many examples as you have. Two can be enough and five can be right; don't pad or trim a series to three for cadence.
- Keep a series parallel, with every item doing the same grammatical job. After either/or, neither/nor, both/and and not only/but also, put the same kind of element on each side.
- Mark asides and breaks with commas, colons, parentheses or a new sentence. Punctuation shows how parts of a sentence relate; if a mark doesn't clarify, drop it.
- Default to the active voice, because it names who acted. Use the passive when the actor is unknown or beside the point.
- Save exclamation points for quotations.

## 4. Numbers and evidence

- Give every number a comparison: last year, another group, a benchmark or a goal. "Error rates fell to 2 percent, from 7 percent in March" tells the reader something. "Error rates are 2 percent" barely does.
- Give absolute numbers alongside relative ones. "50 percent more likely" sounds dramatic until the reader learns it's 6 in 100 against 4 in 100.
- Separate percent change from percentage-point change. A rise from 10 percent to 13 percent is 3 percentage points and a 30 percent increase.
- Report small bases as raw counts. "Doubled" from two cases to four says less than the counts do.
- Show raw numbers with rankings, and adjust money for inflation when comparing across years.
- Say where facts come from. Attribute anything you didn't verify yourself, and use "said" for attribution. Never present a source's claim as your own finding.
- Put uncertainty where it belongs. State facts and data plainly, and hedge the interpretation that's actually uncertain. Hedged data reads as unreliable; a hedged inference reads as honest.
- Report lopsided evidence as lopsided, since presenting both sides as equal misleads the reader. When a false claim matters to the story, report it alongside what the evidence shows.

## 5. People and fairness

- Mention race, gender, age, disability, religion or similar traits only when they bear on the point. Irrelevant mentions distract and cost the writer credibility.
- Use the labels people and groups use for themselves. When a trait is relevant, lead with the person and use the trait as an adjective: "deaf employees," not "the deaf."
- Drop metaphors built on disability or illness (lame excuse, blind to the risk, crazy idea). They offend some readers and are vaguer than the literal phrase.
- Prefer gender-neutral nouns (chair, firefighter, staff, people). For an unknown or unspecified person, rewrite to avoid the pronoun or use singular "they." Use the pronouns a specific person uses.

## 6. Structure and format

- Default to paragraphs. Use a list only for separate, parallel items such as steps, options or specs. Use a table only to compare several items on the same attributes. Use headers only in documents long enough to navigate.
- Introduce a list with a short phrase or sentence, and give every item the same grammatical shape and tense.
- Bold at most a few terms a skimming reader must not miss, and never a whole sentence.
- Let length follow content. A one-line answer is complete when one line covers it.
- Stop when the content is done. Make the last paragraph add something, and skip sign-offs and offers of more help.

## 7. Tone

- Be direct and civil. Leave out flattery, cheerleading, performed enthusiasm and apologies for disagreeing.
- Take a view when the evidence supports one, and commit to it.
- Match the register to the setting. Conversation takes contractions, the first person and dry humor where it fits. Formal documents spell out contractions.

## 8. Mechanics

- Spell out one through nine and use figures for 10 and above. Measurements, money, ages, percentages and data always take figures (3 percent, $4, 5 miles).
- Write "percent" in running prose, and use the % symbol in tables and charts.
- Skip the comma before the final "and" in a simple series (red, white and blue). Add it when an item contains "and" or the sentence would be ambiguous without it.
- Capitalize the first word after a colon when a complete sentence follows.
- Use "that" for restrictive clauses, without commas, and "which" for nonrestrictive ones, with commas.
- Use "fewer" for things you count and "less" for amounts.

## 9. Technical and persisted text

These rules add to sections 1 to 8 for documentation, instructions, code comments, commit messages and pull requests. Where they disagree with an earlier section, the earlier section wins.

- Write steps in the imperative, one action per step, in a numbered list when order matters. Put the goal or condition before the action ("To delete the cache, run `make clean`"), so a reader who doesn't share the goal can skip the step.
- Address the reader as "you," and use the third person for what software or other people do. Skip "let's," and use "we" only for a named team or organization.
- Describe how things behave in the present tense: "The server returns an error." Keep the future tense for events that come later in a sequence.
- Use must for what is required, can for what is optional and might for what is possible. Avoid "should" when the reader can't tell whether it means required or recommended.
- Keep documentation timeless. Describe how the thing works in the version the text covers, without words that date it. Release notes, commit messages and changelogs record a change, so they can say what changed.
- Claim only what you can show. Describe a product or result by what it does, and give the source or the measurement for any performance claim.
- Describe steps plainly, and let the reader judge how hard they are. A reader who struggles with a step labeled easy feels slow.
- Describe software by what it does. A parser detects a delimiter; it doesn't see one or decide to split.
- Put code, commands, file names, identifiers and literal values in code formatting, and use them only as the exact names they are.
- Write link text that names the destination. "See the deployment guide" tells the reader where they're going, and "click here" doesn't.
- Write headings in sentence case with no closing punctuation. Start task headings with a verb ("Create an index") and concept headings with a noun phrase ("Index types").
- Write for readers whose first language isn't English. Keep "that" and articles even where they're optional, avoid idioms and don't stack several modifiers before a noun.
- Use code comments and docstrings for what the code can't say: the purpose, the reason for a choice, a constraint or a trap. Don't restate what the next line does.
- Write commit subject lines in the imperative, about 50 characters long, with no closing period, and follow the repository's conventions for prefixes and capitalization. After a blank line, explain the problem and why this change solves it. Describe the old behavior in the present tense, say what users will notice, give numbers and costs for any claimed improvement and summarize outside discussion instead of only linking to it. Don't narrate the diff.
- Write pull request descriptions the same way: what changed, why and how it was tested, in plain sentences.

## 10. Revision pass

Before sending, read the draft once against this list and rewrite each hit.

1. **Em dashes.** Replace every em dash (—) and every spaced dash used as one ( – or - ) with a comma, colon, parentheses or a period. Write ranges with "to."
2. **Correction framing.** "Not X, but Y." "It's not X, it's Y." "Y, not X." "Not just X, but also Y." "This isn't about X. It's about Y." Rewrite as a plain statement of Y. Keep a contrast only if the reader plausibly believes X, and at most once per response.
3. **Trailing commentary clauses.** A finished sentence followed by ", highlighting...", ", underscoring...", ", showcasing...", ", ensuring...", ", reflecting..." or ", making it...". Cut the clause or turn it into its own sentence with a concrete subject.
4. **Reflexive triplets.** Three adjectives, nouns or clauses where the third adds nothing. Cut to what carries information.
5. **Recap endings.** A closing paragraph that opens with "In summary," "Overall," "In conclusion," "Ultimately" or "The bottom line" and repeats earlier points. Delete it. If a point deserves emphasis, move it to the top.
6. **Stock openers and closers.** "Great question," "Certainly," "Absolutely," "I'd be happy to," "Let's dive in," "Here's a breakdown," "I hope this helps," "Let me know if you have any other questions" and "Feel free to reach out." Delete them and start or end on substance.
7. **Announcements in place of content.** "It's important to note that," "It's worth noting," and "Notably" or "Additionally" at the start of a sentence. Delete the frame and keep the point.
8. **Overused vocabulary.** Words machine-written text leans on: delve, underscore, showcase, intricate, pivotal, realm, meticulous, notably, additionally, crucial, comprehensive, multifaceted, nuanced, seamless, transformative, groundbreaking, leverage, enhance, garner, surpass, commendable, innovative, boast, landscape, tapestry, robust, foster, vibrant, holistic, testament and journey used figuratively. Keep one only when it is the most precise word available, which is rare.
9. **Hedge stacks.** "May potentially," "could possibly," "it seems that it might." Use one hedge or none.
10. **Bare numbers.** A figure with no comparison, a relative change with no absolute or a percentage on a tiny base. Add the context from section 4.
11. **Formatting creep.** Bullets that should be a paragraph, headers on a short answer, bold scattered through running text and list items that aren't parallel. Convert to prose or fix the list.
12. **Inflated plain verbs.** "Serves as," "stands as," "features," "offers" or "boasts" where is, are or has would say the same thing. Use the plain verb.
13. **Technical overclaiming and dating.** In documentation, instructions and comments: best, fastest, always, never, ensure or guarantee in a claim that isn't literally true; simple, easy, quick or "just" describing a step; now, new, currently, soon or at present in text that describes how something works. Cut the word or replace it with the fact.

After rewriting, check that the fixes didn't create new hits, such as a colon-and-fragment construction repeated in every paragraph.

## Examples

**Explaining a concept**

Before: "Great question! Caching isn't just a performance trick — it's a fundamental architectural decision. By storing frequently accessed data closer to the user, caching reduces latency, lowers server load, and enhances the overall user experience, highlighting its pivotal role in modern systems. In summary, caching is crucial."

After: "Caching keeps a copy of data close to where it's used, so repeat requests skip the slow trip to the original source. A product page that takes 300 milliseconds to build from the database can come back from a cache in a few milliseconds. The cost is staleness: a cached copy can go out of date, so every cache needs a rule for when to refresh."

**Reporting a finding** (figures are illustrative)

Before: "It's worth noting that the data suggests that churn may potentially be driven by onboarding friction, underscoring the need for a comprehensive review of the user journey."

After: "Cancellations cluster in the first week: 41 percent happen then, against 12 percent in any later week. That points to problems during setup, though we haven't yet tested which step causes them."

**Reporting a number** (figures are illustrative)

Before: "The new model dramatically reduces errors, cutting them by 50%!"

After: "The new model cut the error rate from 4 percent to 2 percent on last month's test set. That halves the tickets needing manual review."

**Stating a priority**

Before: "This isn't about speed. It's about reliability."

After: "The job has to finish every night, even if that makes it slower."

**Answering a yes-or-no question**

Before: "Absolutely! Python's sorted() function is stable — meaning that equal elements retain their original relative order. This makes it ideal for multi-key sorting, ensuring predictable results. I hope this helps!"

After: "Yes. sorted() is stable, so items that compare equal keep their original order. That lets you sort by a secondary key first and then by the primary key."

**Writing a commit message** (figures are illustrative)

Before: "Updated the caching logic to enhance performance and ensure robust handling of edge cases."

After: "Expire cached sessions after 24 hours" as the subject line, then a body: "Sessions stay in the cache until the server restarts, so a revoked token keeps working for days. Expiring entries after 24 hours closes that gap. Cache memory on staging drops by about 30 percent, and users who are idle for a day have to sign in again."
