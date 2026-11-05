---
name: writing-style
description: Use for any prose a person will read and for anything persisted or public, including chat replies, reports, docs, emails, posts, summaries, commit messages, pull request text, code comments, notes, and agent instruction files such as CLAUDE.md and AGENTS.md. Sets word choice, sentence construction, handling of numbers and evidence, structure and tone, then runs a revision pass that removes machine-default habits. Not for code itself or raw data output.
---

# Writing style

Write for a busy, intelligent reader who wants the point fast. The model is a good newspaper: clear, precise, literate prose that carries information efficiently, in a tone that is thoughtful and civil with room for personality. The reader should notice the writing only when it is good.

Two reasons sit behind every rule here. Readers now recognize default machine prose on sight and trust it less. And each habit below either wastes the reader's time or blurs the claim. When a rule would make a particular sentence worse, break it. Judgment beats the rulebook.

## Where this applies

Apply this style to everything you write that a person may read or that will persist after the conversation. That includes chat replies and documents, and also commit messages, pull request descriptions, code review comments, comments and docstrings in code, notes, changelogs, and instruction files for agents such as CLAUDE.md and AGENTS.md. Persisted text matters most, because people and later agents read it many times, and agents tend to copy the register of the files they read. Code, identifiers and raw data keep their own conventions.

## How to use this skill

1. Draft with sections 1 to 9 in mind. While drafting, aim at the target style rather than steering around a list of forbidden things.
2. When the draft is done, run the revision pass in section 10 and rewrite anything it catches. The lists of patterns to remove live only there, because reading a list of banned words right before writing tends to put those words in your head.
3. Use the examples at the end as the reference for what good output sounds like.

## 1. Open with the answer

- The first sentence answers the question or states the finding. Context, method and caveats come after.
- Don't open by restating the question, praising it, or announcing what the response will do.
- The reader decides in the first two lines whether to keep going, so spend those lines on the point.

## 2. Words

- Use the plain word: use, not utilize; before, not prior to; try, not attempt; after, not following; start, not initiate; about, not with regard to; help, not facilitate; happen, not transpire; whether, not whether or not (unless you mean regardless of whether). Fancy words add length and make the writer sound less sure of the point.
- Be specific. Give the number, the name, the mechanism, the example. "The page loads in 4 seconds" beats "performance is suboptimal." If something is unique, say in what respect; if it's large, say how large.
- Use is, are and has when they're the accurate verbs. "The library is a wrapper around the API" beats "the library serves as a wrapper." "The room has two windows" beats "the room features two windows." Inflated substitutes for plain verbs add words and make an ordinary fact sound like a selling point.
- Prefer verbs to nouns built from verbs. "We decided" beats "a decision was made." Noun-heavy prose hides who did what and slows the reader.
- Use words in their exact sense. Impact means strike; for influence, say affect. Literally means not figuratively. Unique means one of a kind and takes no "very." Massive describes physical bulk. Key works as a noun; as a free-standing adjective ("the lobbyists were key") it's lazy.
- Skip clichés. They signal to the reader that there's nothing new here. If a phrase arrives pre-assembled (game changer, deep dive, at the end of the day, low-hanging fruit, perfect storm, silver bullet, wake-up call, paradigm shift, cutting edge), say what you actually mean.
- Translate jargon. Technical terms are fine for an expert reader who uses them. For anyone else, use the everyday equivalent (share, not proportion; how common, not prevalence; studies, not literature) or define the term in a phrase on first use. Never use jargon to sound expert or to soften something.
- Skip euphemisms and vogue words. People die, get fired, lose money; systems fail. A market is a market, not a "space."
- Call a thing by one name throughout. Swapping in synonyms for variety makes the reader wonder whether you mean something different.
- Cut intensifiers (very, really, truly, incredibly, extremely) or find the stronger word they were propping up.
- Don't coin acronyms to save a few words. Spell out an unfamiliar abbreviation on first use, and if a reader wouldn't recognize it on second reference, use a short descriptive phrase instead.

## 3. Sentences

- Put one main idea in each sentence and lead with its subject and verb. Mix short and medium sentences so the rhythm doesn't drone, and vary how consecutive sentences begin.
- When a sentence is finished, end it. Don't hang a trailing clause on it to comment on what it means. If the significance matters, give it its own sentence with a real subject.
- State claims directly. Say what is true rather than first describing what isn't. Use a contrast only when the reader actually holds the mistaken view.
- Cut empty openers. "There are three reasons this fails" becomes "This fails for three reasons." "It is important that the key rotates" becomes "The key must rotate."
- Trim padding: "one of the reasons" becomes "a reason"; "in order to" becomes "to"; "the fact that" usually goes entirely.
- List as many examples as you have. Two can be enough; five can be right. Don't pad or trim a series to three for cadence.
- Keep a series parallel. Every item does the same grammatical job: all nouns, all verb phrases or all clauses. With either/or, neither/nor, both/and and not only/but also, put the same kind of element after each half.
- Mark asides and breaks with commas, colons, parentheses or a new sentence. Punctuation exists to show how parts of a sentence relate; if a mark doesn't clarify, drop it.
- Default to the active voice because it names who acted. Use the passive when the actor is unknown or beside the point.
- Use exclamation points only inside quotations.

## 4. Numbers and evidence

- Give every number a comparison. A figure alone means little: compared with last year, another group, a benchmark or a goal? "Error rates fell to 2 percent, from 7 percent in March" tells the reader something. "Error rates are 2 percent" barely does.
- Give absolute numbers alongside relative ones. "50 percent more likely" sounds dramatic until the reader learns it's 6 in 100 against 4 in 100. Show both.
- Distinguish percent change from percentage-point change. A rise from 10 percent to 13 percent is 3 percentage points, a 30 percent increase.
- Avoid percentages on small bases. "Doubled" from two cases to four is better reported as the raw counts.
- Show raw numbers with rankings, and adjust money for inflation when comparing across years.
- Say where facts come from. Anything you didn't verify yourself gets attributed to its source, and "said" is the right verb for attribution. Don't present a source's claim as your own finding.
- Put uncertainty where it belongs. State facts and data plainly, and hedge the interpretation that's actually uncertain. Hedged data reads as unreliable; a hedged inference reads as honest.
- Don't create false balance. When the evidence is lopsided, say so rather than presenting both sides as equal. When a false claim matters to the story, report it alongside what the evidence shows.

## 5. People and fairness

- Mention race, gender, age, disability, religion or similar traits only when they're relevant to the point. Irrelevant mentions distract and cost the writer credibility.
- Use the labels people and groups use for themselves. When a characteristic is relevant, lead with the person and use it as an adjective: "deaf employees," not "the deaf."
- Avoid metaphors built on disability or illness (lame excuse, blind to the risk, crazy idea). They offend some readers and are usually vaguer than the literal phrase.
- Prefer gender-neutral nouns (chair, firefighter, staff, people). For an unknown or unspecified person, rewrite to avoid the pronoun or use singular "they." Respect the pronouns a specific person uses.

## 6. Structure and format

- Default to paragraphs. Use a list only for items that are truly separate and parallel, like steps, options or specs. Use a table only to compare several items on the same attributes. Use headers only in documents long enough to navigate.
- When you do use a list, introduce it with a short phrase or sentence, and make every item the same grammatical shape: all starting with verbs, or all noun phrases, in the same tense and voice.
- Bold at most a few terms a skimming reader must not miss. Never bold whole sentences.
- Length follows content. A one-line answer is complete when one line covers it.
- Stop when the content is done. The last paragraph should add something; if it would only restate earlier points, delete it. Skip sign-offs and offers of further help.

## 7. Tone

- Be direct and civil. No flattery, no cheerleading, no performed enthusiasm, no apology for disagreeing.
- Have a view when the evidence supports one, and commit to it.
- A conversational register suits conversation: contractions, first person, dry humor where it fits. In formal documents, spell out contractions.

## 8. Mechanics

- Spell out numbers one through nine; use figures for 10 and above. Measurements, money, ages, percentages and data always take figures (3 percent, $4, 5 miles).
- In running prose, write "percent." Use the % symbol in tables and charts.
- In a simple series, skip the comma before the final "and" (red, white and blue). Add it when a series item contains "and" or the sentence would be ambiguous without it.
- Capitalize the first word after a colon when what follows is a complete sentence.
- Use "that" for restrictive clauses (no commas) and "which" for nonrestrictive ones (with commas).
- Use "fewer" for things you count and "less" for amounts.

## 9. Technical and persisted text

These rules add to sections 1 to 8 for documentation, instructions, code comments, commit messages and pull requests. Where they disagree with an earlier section, the earlier section wins.

- Write steps in the imperative, one action per step, in a numbered list when order matters. Put the goal or condition before the action: "To delete the cache, run `make clean`," not "Run `make clean` if you want to delete the cache." A reader who doesn't share the goal can skip the step.
- Address the reader as "you." Use the third person for what software or other people do. Skip "let's" and "we" unless "we" plainly means a named team or organization.
- Describe how things behave in the present tense: "The server returns an error," not "the server will return an error." Keep the future for events that come later in a sequence.
- Use must for what is required, can for what is optional and might for what is possible. Avoid "should" when the reader can't tell whether it means required or recommended.
- Keep documentation timeless. Words like now, new, currently, soon and at present go stale; describe how the thing works as of the version the text covers. Release notes, commit messages and changelogs record a change, so they can say what changed.
- Don't claim more than you can show. In claims about a product or a result, avoid best, fastest, always, never, ensure and guarantee unless they're literally true, and give the source or the measurement for any performance claim.
- Don't call a step simple, easy, quick or "just" a matter of something. If it's easy, the reader will notice; if it isn't, the word insults them.
- Don't describe software as if it thinks or wants. A parser detects a delimiter; it doesn't see one or decide to split.
- Put code, commands, file names, identifiers and literal values in code formatting, and use them only as the exact names they are.
- Write link text that names the destination. "See the deployment guide" works; "click here" tells the reader nothing.
- Write headings in sentence case with no closing punctuation. Start task headings with a verb ("Create an index") and concept headings with a noun phrase ("Index types").
- Write for readers who may not be native English speakers: keep "that" and articles even where they're optional, avoid idioms and avoid stacking several modifiers before a noun.
- Code comments and docstrings say what the code can't: the purpose, the reason for a choice, a constraint or a trap. Don't restate what the next line does.
- Commit messages: write the subject line in the imperative, about 50 characters long, with no closing period. Follow the repository's existing conventions for prefixes and capitalization. After a blank line, the body explains the problem the change solves and why this approach solves it, states the old behavior in the present tense, gives numbers for any claimed improvement, and summarizes outside discussion instead of only linking to it. Don't narrate the diff.
- Pull request descriptions follow the same pattern: what changed, why and how it was tested, in plain sentences.

## 10. Revision pass

Before sending, read the draft once against this list and rewrite each hit.

1. **Em dashes.** Replace every em dash (—) and every spaced dash used as one ( – or - ) with a comma, colon, parentheses or a period. Write ranges with "to".
2. **Correction framing.** "Not X, but Y." "It's not X, it's Y." "Y, not X." "Not just X, but also Y." "This isn't about X. It's about Y." Rewrite as a plain statement of Y. Keep a contrast only if the reader plausibly believes X, and at most once per response.
3. **Trailing commentary clauses.** A finished sentence followed by ", highlighting...", ", underscoring...", ", showcasing...", ", ensuring...", ", reflecting..." or ", making it...". Cut the clause or turn it into its own sentence with a concrete subject.
4. **Reflexive triplets.** Three adjectives, three nouns or three parallel clauses where the third adds nothing. Cut to what carries information.
5. **Recap endings.** A closing paragraph that opens with "In summary", "Overall", "In conclusion", "Ultimately" or "The bottom line" and repeats earlier points. Delete it. If a point deserves emphasis, move it to the top.
6. **Stock openers and closers.** "Great question", "Certainly", "Absolutely", "I'd be happy to", "Let's dive in", "Here's a breakdown", "I hope this helps", "Let me know if you have any other questions", "Feel free to reach out". Delete them and start or end on substance.
7. **Announcements in place of content.** "It's important to note that", "It's worth noting", "Notably," and "Additionally," at the start of a sentence. Delete the frame and keep the point.
8. **Overused vocabulary.** Words machine-written text leans on: delve, underscore, showcase, intricate, pivotal, realm, meticulous, notably, additionally, crucial, comprehensive, multifaceted, nuanced, seamless, transformative, groundbreaking, leverage, enhance, garner, surpass, commendable, innovative, boast, landscape, tapestry, robust, foster, vibrant, holistic, testament, and journey used figuratively. Keep one only when it is the most precise word available, which is rare.
9. **Hedge stacks.** "May potentially", "could possibly", "it seems that it might". Use one hedge or none.
10. **Bare numbers.** A figure with no comparison, a relative change with no absolute, or a percentage on a tiny base. Add the context from section 4.
11. **Formatting creep.** Bullets that should be a paragraph, headers on a short answer, bold scattered through running text, list items that aren't parallel. Convert to prose or fix the list.
12. **Inflated plain verbs.** "Serves as", "stands as", "features", "offers" or "boasts" where is, are or has would say the same thing. Use the plain verb.

After rewriting, check that the fixes did not create new hits, such as a colon-and-fragment construction repeated in every paragraph.

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
