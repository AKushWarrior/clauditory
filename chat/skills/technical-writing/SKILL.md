---
name: technical-writing
description: Use when writing documentation, how-to instructions, READMEs, code comments, docstrings, error or log messages, commit messages, pull request descriptions, changelogs, or agent instruction files. Adds rules for steps, modal verbs, code and identifiers, links, headings, and commits on top of the writing-style skill. Not for code itself.
metadata:
  version: "1.1.1"
---

# Technical writing

These rules add to the writing-style skill for technical and persisted text. Apply writing-style first, then these. Where the two disagree, writing-style wins.

## 1. Instructions and procedures

- Write steps in the imperative, one action per step, in a numbered list when order matters. Write one action as a plain sentence instead of a one-item list.
- Put the goal or condition before the action ("To delete the cache, run `make clean`"), so a reader who doesn't share the goal can skip the step. Turn "if you want to X, you can" into "to X."
- Make sure the reader knows where an action happens before you describe it ("In the settings file, set `retries` to 3").
- Put any explanation after the instruction, in the same step. When the reader must choose between paths, give the choice in one step with the alternatives as bullets under it, and start an optional step with `Optional:`.
- Stop at the last real action. Don't add a step that only says the task is done.
- Leave "please," "thank you," and "congratulations" out of instructions and messages. State the action, or state that the task is finished.

## 2. Voice, person, and tense

- Address the reader as "you," and use the third person for what software or other people do. Skip "let's," and use "we" only for a named team or organization.
- State a recommendation as an instruction ("Back up the disk first") or name who recommends it. Skip "it is recommended" and "we suggest."
- Describe how things behave in the present tense: "The server returns an error." Keep the future tense for events that come later in a sequence.
- Use must, or the imperative, for what is required; can for what is optional or possible; and might for what could happen but is uncertain. Use may only for permission. Avoid "should" and "need to" when the reader can't tell whether they mean required or recommended. Use "don't have to" for something optional and "must not" for something forbidden; they mean different things.
- Describe software by what it does. A parser detects a delimiter; it doesn't see one or decide to split. A tool reports an error; it doesn't complain. A window opens; it doesn't appear.

## 3. Claims and time

- Keep documentation timeless. Describe how the thing works in the version the text covers, without words that date it. Release notes, commit messages, and changelogs record a change, so they can say what changed.
- Claim only what you can show. Describe a product or result by what it does, and give the source or the measurement for any performance claim.
- Describe steps plainly, and let the reader judge how hard they are. A reader who struggles with a step labeled easy feels slow.

## 4. Code, identifiers, and literals

- Put code, commands, file names, identifiers, and literal values in code formatting, and use them only as the exact names they are. Copy code, commands, and output from the source instead of retyping them.
- Name the kind of thing after an identifier at first mention: "the `retries` parameter," "the `deploy` command," "the `config.yaml` file." Don't start or end a sentence with a bare identifier, and don't use a command as a verb ("grep the logs"); write "search the logs with `grep`." In chat with engineers, a familiar command verb is fine.
- Introduce a code block with a complete sentence that says what it does, such as "This command lists the open ports." A bare `Run:` tells the reader nothing. Don't start a sentence before a code block and finish it after.
- Use ISO dates (2026-11-05) in logs, file names, changelogs, and data, and the written-month form in prose.
- When the difference matters, say whether a gigabyte means 10^9 or 2^30 bytes, or write GiB for the binary unit.
- Use the common forms of computing words: email, website, web page, online, username, filename, database, internet (lowercase), and login as a noun but log in as a verb.

## 5. Structure, links, and headings

- Write headings in sentence case with no closing punctuation. Start task headings with a verb ("Create an index") and concept headings with a noun phrase ("Index types"). Avoid bare headings such as "Overview" or "Introduction," use question headings only in FAQs and troubleshooting pages, and put at least a sentence between a heading and its first subheading.
- Write link text that names the destination. "See the deployment guide" tells the reader where they're going, and "click here" doesn't. Say why you're sending the reader elsewhere.
- Put what the reader needs to finish the task in the text, and link only to extra material. When the referenced fact is short and essential, such as a limit or a required flag, state it and give the link as backup. Keep links to a couple per paragraph.
- Refer to sections, tables, and figures by name or number ("the previous section," "the table of limits"), not by position ("above," "below"), in any document whose layout may change. In a long document read out of order, redefine an abbreviation at its first use in each section, or link to a glossary.

## 6. Writing for every reader

- Write for readers whose first language isn't English. Avoid idioms and figurative phrases, don't use a verb as a noun ("the install"), and repeat "to" before each infinitive in a series. Avoid words whose meaning differs between American and British readers, such as first floor and momentarily; name the floor or the time instead.

## 7. Comments and messages

- Use code comments and docstrings for what the code can't say: the purpose, the reason for a choice, a constraint, or a trap. Don't restate what the next line does.
- Write error messages that name the specific problem ("The config file /etc/app.yaml can't be read"), not "An error occurred." Don't blame the user. Give the likely cause and a specific fix that names the file, command, or setting, most likely fix first. Suggest contacting support or an administrator only when nothing else will work, and say what they need to do.

## 8. Commits and pull requests

- Write commit subject lines in the imperative, about 50 characters long, with no closing period, and follow the repository's conventions for prefixes and capitalization. After a blank line, explain the problem and why this change solves it. Describe the old behavior in the present tense, say what users will notice, give numbers and costs for any claimed improvement, and summarize outside discussion instead of only linking to it. Don't narrate the diff.
- Write pull request descriptions the same way: what changed, why, and how it was tested, in plain sentences.

## 9. Revision pass

After the writing-style revision pass, check the draft against these items and rewrite each hit.

1. **Overclaiming and dating.** Best, fastest, always, never, ensure, or guarantee in a claim that isn't literally true; simple, easy, quick, or "just" describing a step; now, new, currently, soon, or at present in text that describes how something works. Cut the word or replace it with the fact.
2. **Vague requirements.** "You should," "you need to," or "you may" where the reader can't tell whether the step is required, optional, or permitted. Use must, can, or might.
3. **Bare identifiers and positional references.** An identifier with no noun for its kind, a command used as a verb in documentation, "above" or "below," and link text such as "here" or "this page."

Then run the writing-style skill's linter on the draft, with `--commit` for a commit message, and fix what it reports where the rule applies.

## Example

**Writing a commit message** (figures are illustrative)

Before: "Updated the caching logic to enhance performance and ensure robust handling of edge cases."

After: "Expire cached sessions after 24 hours" as the subject line, then a body: "Sessions stay in the cache until the server restarts, so a revoked token keeps working for days. Expiring entries after 24 hours closes that gap. Cache memory on staging drops by about 30 percent, and users who are idle for a day have to sign in again."
