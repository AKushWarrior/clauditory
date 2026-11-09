# Red Hat Technical Writing Style Guide

- **Source:** Red Hat, Inc., Red Hat Technical Writing Style Guide: Style Conventions for Writers and Editors, edition 7.3. The book's metadata (Book_Info.xml) gives edition 7.3; its release notes chapter (New.xml) describes releases through 7.2, a maintenance release dated March 2026. Its revision-history appendix (Revision_History.xml) stops at revision 3-0, dated July 2021, and has not been kept up since. The copy read was the repository's state at a commit dated October 27, 2026.
- **Type:** style manual (corporate technical style guide with a word usage dictionary)
- **Where to find it:** https://github.com/StyleGuides/WritingStyleGuide (DocBook XML sources under `en-US/`, built with Publican)
- **Last reviewed:** 2026-11-05
- **Used by:** `chat/skills/technical-writing/SKILL.md`

## What the source covers

### Purpose, audience, and status

The guide is Red Hat's house manual for technical writing. Its audience page and the repository README both say that it is the official guide for Red Hat training and certification content and for all other technical documentation except where stated otherwise, and both carve out one large exception. Red Hat's product documentation, written by the Customer Content Services group, follows a separate document, the Red Hat supplementary style guide, in place of this one. Rules in this guide therefore describe Red Hat training courses, exercises, and general technical writing more than the product manuals most readers would associate with the company. Its resources chapter sets the scope as software manuals, training courses, user guides, and white papers, and puts marketing and corporate content outside it. Release 7.0 (November 2024) removed the marketing-specific material and the references to the AP Stylebook.

The guide calls itself a supplement to other references. It sends readers to IBM Style as the first general reference (access is restricted to IBM and Red Hat employees, and the guide marks those links with a padlock), then to The Chicago Manual of Style and Merriam-Webster. The edition of Chicago it names is inconsistent: the grammar, punctuation, and numbers sections cite the 17th edition, while the word usage references page names the 18th. The resources page also says that company guides should be consulted before general ones and that published style references contradict each other because audiences and genres differ, which is a fair description of the conflicts this file records.

The guide is open source and accepts community contributions through GitHub issues and pull requests. That shows in its texture. Entries were written by many hands over more than a decade. Some carry editorial notes such as a remark that an entry is under review or that an example would help, and a few rules sit in commented-out XML that the built guide does not show. The most important case is a grammar section on putting a condition before an action, which survives only as a comment; the published guide still models the pattern in examples elsewhere, as described below. This file treats the XML as published (comments excluded) as the guide's content.

### Objectives and readability targets

A short opening chapter states what Red Hat documentation should be: legally compliant, accurate and consistent, readable, comprehensible, and focused on the user without condescending to the reader or assuming their skills. It gives numeric targets: a Flesch reading-ease score of 60 to 70 and a Gunning fog index of 9 to 12, and it sets the minimum reading level of the audience at about eighth grade (ages 14 to 15). The guide does not explain how these targets were chosen or how writers should measure them, and no other chapter refers back to them.

### Grammar

The grammar chapter covers a handful of topics and defers everything else to Chicago.

**Voice.** It asks for the active voice whenever possible, because it reads more directly and cuts word count, but it says the passive is not forbidden. It names release notes and technical notes as places where the passive suits. It also allows the passive when the important item should be the subject of the sentence or when a keyword should come first in a title, heading, or opening sentence. Writers are asked to weigh that front-loading against readability case by case.

**Agreement and pronouns.** A section on pronoun-antecedent agreement covers singular and plural antecedents and collective nouns. A section on countable nouns gives fewer and number for things that can be counted and less and amount for things that cannot. A section on possessives forbids them with product names and abbreviations, which are written as modifiers instead, while allowing the possessive of the company name when it means the company itself. The relative-pronoun section gives who and whom for people and that and which for things, with that for restrictive clauses, and a substitution test for choosing who or whom.

**Sentence structure.** The guide asks writers not to pack too much into one sentence and sets a ceiling of 40 words for a sentence in technical documentation. Its example breaks one long sentence into three. It allows a preposition at the end of a sentence when the alternative sounds stilted. It defines run-on sentences and comma splices and offers four repairs (a period, a semicolon, a coordinating conjunction, or a subordinating conjunction). It asks writers to repair sentence fragments, with a read-aloud test.

**Telegraphic style.** A short section warns that extreme word economy drops articles, prepositions, and verbs, and leaves sentences that are hard to parse and to translate. Its examples restore a missing article and a missing form of the verb be (a note that work remains becomes a note that work is required). It singles out revision histories, where terse entries are common, as a place where wording must stay clear for translators and customers.

**Causative verbs.** The guide asks writers to avoid constructions in which something is made to happen through have, and to replace have with a verb that names the action (apply the changes, not have the changes take effect).

**Either...or, following, and if...then.** Either...or is for a choice between exactly two options; for three or more, the guide rewrites with such as. The word following should be followed by a noun (the following steps, not the following alone), and in quiz questions it is often superfluous. An if clause is generally followed by then.

**That, this, and these.** The guide asks writers to keep the conjunction that in clauses unless the result is stilted, because it helps translation and readers whose first language is not English. A separate section asks that this, that, these, and those always be followed by a noun. The reason given is clarity for the same readers and accuracy in languages where the demonstrative changes with the gender of its noun. Its examples turn a bare this at the start of a sentence into this action, or rewrite with consequently.

**Verbosity and word order.** A verbosity section asks writers to cut unnecessary words. Its examples turn a nominalized verb back into a verb (`perform the installation of the product` becomes `install the product`) and cut a sentence about a service that provides services for administration into one that says what the service does. A word-order section asks for the arrangement that creates the least ambiguity, usually subject, verb, and object with modifiers next to what they modify.

**Contractions and abbreviations.** Red Hat documents use no contractions, because they mark informal writing and cause problems in translation. The Latin abbreviations `e.g.` and `i.e.` are written out as `for example` and `that is`; the word usage entries repeat the rule.

**Hyphenation.** The guide says there are no fixed rules and asks for hyphens only where clarity needs them. Prefixes and suffixes are usually closed up. Compound modifiers are hyphenated, except open source, which is never hyphenated. A compound that starts with an adverb ending in -ly is never hyphenated, citing Chicago. Hyphens separate doubled vowels after a prefix (with stated exceptions such as cooperate and reestablish), prefixes before a capitalized word and double prefixes.

**Gender and person.** Documentation uses no gender-specific pronouns except for a specific named person, and singular they is accepted. Instructions use the imperative or you; general explanations can use users. The guide forbids one in place of you, it for a person, and the first person (I, we, our). When the company is the actor, the text names Red Hat with a singular verb, so a recommendation reads as Red Hat recommending something, not as we.

**Tense.** The guide asks writers to avoid the future tense and will wherever possible and to use the simple present, because readers usually consult documentation while they use the system, and because the present is the easiest tense to translate.

### Easily confused words

A short section distinguishes affect from effect (each as noun and verb, with a note that effect as a verb is rare and usually has a clearer alternative) and assure, ensure, and insure. Other confusable pairs are scattered through the word usage dictionary and are described below.

### Punctuation

**Colons and semicolons.** Either mark can join two related clauses where the second explains or amplifies the first. The word after a colon or semicolon is lowercase unless it is a proper noun, and after a colon a complete sentence starts with a capital letter. The guide asks writers to limit both marks and to use a period where they can. A colon introduces a series, but not when the series is the object or complement of a verb in the sentence. Semicolons separate series items that contain commas, and a semicolon precedes a conjunctive adverb that joins two independent clauses.

**Colons before lists and code.** A sentence that immediately introduces a list, command, code block, or procedure ends with a colon. If other sentences come between the introduction and the item, every one of them ends with a period. A sentence never ends with the following; a noun comes after it.

**Commas.** A comma joins the clauses of a compound sentence unless they are short and closely related. Conjunctive adverbs are set off with commas. A dependent clause that comes before the main clause takes a comma; one that follows takes a comma only when it is nonrestrictive. Nonrestrictive adjectival clauses take commas and restrictive ones do not. Coordinate adjectives are separated by commas. In a series of three or more, a comma goes before the conjunction: the guide requires the serial comma.

**Parentheses.** Closely related material goes between commas and incidental material between parentheses. Phrases such as for example and that is go in parentheses when they break the sentence sharply and between commas when the break is slight. The plural marker in parentheses is forbidden: if something can be singular or plural, write the plural.

**Slashes.** The guide forbids a slash to mean one of two options, so `enable or disable` replaces the slashed form, and `A or B, or both` replaces `and/or`.

**Quotation marks.** Curly quotation marks are for direct quotations and straight ones for code, syntax, hexadecimal values, and links. Punctuation goes inside the closing quotation mark only when it belongs to the quotation or to a programming element that uses quotation marks, and otherwise outside. The rule follows the logical style common in technical writing, and it differs from the American convention of putting periods and commas inside. Release 5.1 (January 2022) records an update to this rule, and the revision history holds a commented-out entry that once stated it more broadly.

**Apostrophes and exclamation points.** Apostrophes never form plurals. Possessives follow standard rules, with the guide's variants for singular names ending in s. Exclamation points do not end sentences; the mark appears only as a character in commands.

**Dashes.** The punctuation chapter has no dash section, but the dash entry in the word usage dictionary says to use a dash only for ranges (pages, letters, dates) and otherwise not at all in technical content. Commas, parentheses, or a colon take its place. The em dash entry was removed in release 7.0. The through entry separately forbids using a hyphen or dash to mean through.

**Lists and tables.** List punctuation is covered in the writing chapter (below). Table cells end with periods when every cell in the column is a complete sentence and take none when every cell is a fragment; mixed columns should be made parallel where possible and otherwise punctuated cell by cell.

**Naming characters and keys.** The chapter ends with conventions for naming punctuation marks and special characters in text (the name, a describing noun where the name includes one, and the symbol in parentheses, never the bare symbol). It includes a table of names and a table of keyboard key labels. It distinguishes pipe (for command pipelines) from vertical bar (for a logical or in expressions). These are documentation conventions with no bearing on general prose.

### Design

The design chapter collects conventions for technical publications, some of which apply beyond Red Hat.

**Headings.** The guide uses title case for headings, titles, and references to other publications' titles, with sentence case for figure captions, legends, diagram labels, and table column headers. Headings carry little punctuation and no closing period. For larger units such as chapters, the gerund form of a specific action verb is the usual choice; tasks the reader performs can take an imperative verb; purely informational sections take a noun phrase. Vague verbs such as understanding, introducing, or exploring are replaced by a specific verb or a noun phrase, and one-word titles are avoided. Headings should not contain file, command, or markup names; those belong in the text.

**User interface.** UI elements are named with the capitalization and spelling of the interface (with spelling errors corrected), without their punctuation (an ellipsis on a button label is dropped) and usually without the element type (the label alone, without the word button). Go to covers any sequence of selections in a UI. Figures need a caption or callouts only when the surrounding text does not describe them, and never both.

**Commands and examples.** The guide explains command syntax with worked examples, warns against showing force and assume-yes options without reason, asks for line continuation characters (and no manual indentation) in long commands, and sets a notation for omitted output. It asks writers not to name an editor when telling readers to view or edit a file, and not to use line numbers as reference points, because they change. Example usernames, hostnames, and IP addresses are lowercase, use reserved example domains, never use real public addresses or real people, and draw on a diverse set of realistic names when a case study needs them. A long section on escalated privileges sets conventions for sudo in training exercises.

**Object names.** Release 7.2 expanded the guidance on referring to named objects such as files, commands, variables, and resources. A sentence should be complete without the object name. The object name should work as an adjective followed by a noun that names its class (the `ls-blue` router, the `/mnt/save` directory). A sentence should not start or end with an object name. Where the noun would duplicate the name, or the object can be named in plain words, the guide prefers plain language. A related section on naming object classes asks writers who create examples to avoid names that end with their own class, so text does not end up with a doubled noun. File-system names keep their creators' capitalization, even at the start of a title.

**Abbreviations and names.** Most acronyms and initialisms are spelled out on first use, but well-known ones such as HTML are not. The expansion is lowercase unless it is a proper noun. The article follows pronunciation (an RTS, a RAM upgrade). Plurals add a lowercase s with no apostrophe, and abbreviations take no possessive. Product names follow company rules. They use the full name on first use, take no articles before names that include their own noun, and take no possessives or line breaks inside a name. Non-breaking spaces appear in defined places, and the major version number is used unless a minor release is meant.

**Admonitions and recommendations.** Note, Important, and Warning admonitions should be brief and rare, so they keep their force. Recommendations name the recommender (Red Hat recommends) in place of the impersonal it is recommended.

**Citations and links.** External sources should be authoritative, credible, and publicly accessible, and link text should describe the destination in place of a bare here.

### Design of the whole publication

This short chapter covers book structure. Titles combine product, version, and publication name, and a one-phrase subtitle follows. A descriptive abstract says what the publication covers, how it conveys it, and why it exists (and for whom). Introductions carry the product's name instead of the book's. Two rules apply to any document. Two headings should never appear in a row without text between them; if no useful text fits, one heading is probably unnecessary. And a title should not be `Overview` or `About` on its own. Nor should a chapter or section have a separate introduction section; introductory material goes in the body text under the title.

### Choosing appropriate language

This chapter, the longest prose chapter, covers word choice for a global and translated audience, inclusive language, and ambiguity.

**Misleading language.** The guide tolerates best practices as a common search term but prefers recommended practices when in doubt, and notes the concern about superlatives without supporting data. Idioms such as first come, first served are to be avoided in content that will be translated.

**Slang and jargon.** An alphabetical list of about 75 terms gives replacements or says not to use them. The entries fall into a few groups: business jargon (`best-of-breed`, `core competency`, `going-forward basis`, `net-net`, `synergy`, `value-added`, `paradigm`, `mission-critical` where nothing is critical to a mission, `solutions-based`, `silo`); idioms and metaphors (`apples to apples`, `at the end of the day`, `boil the ocean`, `drop the ball`, `low-hanging fruit`, `on the same page`, `think outside the box`, `double-edged sword`, `kettle of fish`); invented or informal words (`administrivia`, `automagic`, `bucketize`, `productize`, `performant`, `upskill`, `ask` as a noun); padded phrases (`for now` for `at this point in time`, `with` for `in concert with`, `immediately before` for `right before`); and vague claims (`improve and enhance`, `meet your needs`, `exceed your expectations`, `harness the power`, `frictionless`). Several entries apply directly to machine-written prose. The guide flags `dig deeper` and `delve deeper` as likely to translate awkwardly, asks writers to avoid `leverage` in favor of `use`, and replaces `is designed to` with a statement of what the product does. The entry for `if you want` or `if you wish` turns the condition into a purpose clause that starts with `to`.

**Neologisms and phrasal verbs.** The guide discourages new words formed by turning verbs into nouns and nouns into verbs, and asks for one-word verbs where they exist (`click` for `click on`, `complete` for `fill in` or `fill out`, `omit` for `leave out`, `print` for `print out`).

**-ing forms.** Release 7.2 expanded a section on words ending in -ing. A gerund suits a title that names an activity in general, a sentence whose subject is the activity and a list of activities that is parallel with its lead-in. When a filler verb phrase such as is used for pushes the real action into a gerund, the sentence should be rewritten around the active verb. A gerund should not replace an instruction in a procedure step. The section's sharpest rule concerns present participles: the guide asks writers to avoid a noun followed by a word ending in -ing, with or without a comma between them, because the construction is ambiguous and hard to translate. Its examples are sentences that end in a participial phrase tacked onto a finished clause ("a table that shows dropped packets, enabling you to identify flows"; "an add-on that extends support, providing stability"). Each is rewritten with `which`, `so that`, or a to-infinitive, which makes the relation explicit. Noun phrases that name a single thing, such as error logging, are exempt.

**Anthropomorphism.** Software and hardware do not think, talk, refuse, or complain. Release 7.2 added examples, and each rewrite names the observable behavior (the tool produces errors, the proxy communicates with a server).

**Allow.** Release 7.1 added a rule that allow is for permission only (a license allows reuse, sudo allows a user to run a command). For anything else, the sentence should focus on what the user or system does (you can use the command to filter messages) or use another verb such as enable.

**Inclusive language.** The guide drops `white` and `black` as metaphors for good and bad (`allowlist`, `blocklist`, or `denylist` instead) and `master` when paired with `slave` (`main` for a default branch; `primary`, `controller`, `leader`, `worker`, `replica`, or `secondary` elsewhere). It also drops gendered terms such as `man-hours`, `he or she` in favor of `they`, terms that carry neurodiversity bias such as `sanity check`, `disabled` as a description of a person, and superlative job titles such as `guru`, `ninja`, and `rockstar`. It says the guidance will change over time.

**Redundant words.** Filler words such as actually, really, simply, and very are to be cut, along with pairs that repeat a meaning (revert back) and words that add nothing (located on for on). The word usage dictionary adds basically, essentially, totally, and just to the list.

**Ambiguity.** A section with 13 subsections catalogs common sources of ambiguity, each with an example and a rewrite. `May` and `should` are avoided because each has two senses (possibility or permission; recommendation or obligation), and the rewrites use `can`, `might`, `must`, or `is expected to`. Statements that a task is easy or a product is user-friendly are cut in favor of describing how to do the task. Proper nouns are capitalized so translators can tell them from concepts. A word with several senses is made explicit with a noun. Two infinitives in one sentence are separated so it is clear what each one belongs to. Noun-noun phrases that hide whether the first noun is plural are rewritten. `Only` is placed directly before the word it modifies, which is often not the verb. Pronouns with two possible antecedents are replaced by the noun. One concept keeps one term throughout, because a second term suggests a second thing. `There is`, `there are`, and `there might be` are rewritten so the real subject leads. `Using` is replaced by `by using` or `that uses`, depending on which is meant. An opening verb phrase must modify the subject that follows it (no dangling modifiers).

**Dates, times, and numbers.** All-numeric dates are forbidden in running text because the order of day and month varies by country. The guide writes the month as a word in day-month-year order and allows ISO 8601 (year-month-day) where space is short, with the format stated. Times use uppercase `AM` and `PM` without periods, and noon and midnight replace `12 PM` and `12 AM`. Numbers zero through nine are spelled out, as are numbers that begin a sentence, a number directly before another number, approximations, and very large values. Numerals are used for 10 and above, and for smaller numbers in the same paragraph as larger ones, as well as for negatives, fractions, percentages, decimals, measurements, section references, registers, code, and versions. Four-digit numbers take no comma; five digits and more do. Numbers to be dialed use spaces and international country codes.

### Writing clearly and succinctly (translation)

This chapter (Translation.xml) covers structure for readability and translation.

**Lists.** A single sentence should not be formatted as a bulleted list. Bulleted lists are for unordered items and usually have at least three. Numbered lists are for sequences or for items referred to elsewhere. Variable (term and definition) lists suit options and arguments and are often better than tables, which are costly to translate and suit only data that needs a grid. Procedures always have a title, and nested lists go no deeper than two levels. A list must not interrupt a sentence that continues after it. A lead-in should be a complete sentence, because translation tools treat the lead-in and each item as separate segments. Items must be grammatically parallel. Items that are complete sentences end with periods; items that are fragments (including ones with a verb but no subject) take no end punctuation; and a list whose fragments are followed by sentences ends every entry with a period. The chapter also discusses normal versus compact spacing.

**Noun stacks.** Strings of modifiers and nouns are hard to translate when several readings are possible, and the guide asks for no more than three consecutive nouns.

**Grammatical gender.** Pronouns such as it and they, and abbreviations that can name two things (RPM the package or RPM the package manager), are hard to translate into languages with gendered nouns; the fix is to repeat the noun.

**Markup, code blocks, and entities.** Untranslatable terms should be marked up and should not do grammatical work in a sentence (a path or command name should be followed by a noun such as directory or command, and should not stand in for a verb). Commentary should not be placed inside a code block, where it can be mistaken for output and left untranslated. Entities should be limited to those a build requires.

### Cross-references

This chapter argues that a writer who sends readers from link to link is serving their own convenience at the reader's cost. A link should point to additional information, never to information the reader needs to finish the task in hand; that information belongs in the text, with a link at the end. Paragraphs should hold no more than a couple of links, and links do not belong in titles, headings, or captions. Vital information of up to a couple of paragraphs (or half a page, or five table rows) should be repeated where it is needed, not cross-referenced, because repetition saves the reader time and matters for safety.

### The word usage dictionary

The usage dictionary runs from 0-9 through XYZ and holds roughly 500 entries. Most settle the spelling, spacing, hyphenation, or capitalization of technical terms (`back end` as a noun and `back-end` as an adjective, `backup` and `back up`, `file system`, `plug-in`, `stand-alone`, `setup` and `set up`, and `login` and `log in`) or give Red Hat product and trademark rules. Release 7.1 and 7.2 added entries for AI terms. AI need not be spelled out, and `gen AI` is acceptable. `Large language model` and `machine learning` are not hyphenated even as modifiers, `fine-tune` is hyphenated, and `inference` can be a verb in an AI context.

A smaller group of entries sets general usage rules that bear on prose:

- **Conjunctions in one sense.** `Because` is for reasons; `since` and `as` are for time only. `Once` means one time and is not a conjunction for `after` or `when`. `While` is for simultaneous events, with `although` for contrast. `Like` is not used for `such as`. `Whether` replaces `if` for alternatives. `As long as` is for length or time, with `provided that` otherwise.
- **Modal verbs.** Can is for possibility and may only for permission, and can is preferred over could and may because it is less tentative. Must is for requirements and should never states something the user must do. The entry for need to, needs, and needs to be asks writers to avoid them where possible and to use must, required, or similar words. The entry for you may suggests deleting it.
- **Plain words.** `Use` replaces `utilize`, `open` replaces `bring up`, `examine` or `inspect` replaces `look at`, and `many` replaces `lots of` and `bunches of`. `Emit` or `issue` replaces `send out`, `start` replaces `start up`, and `now` replaces `right now` and `at this point in time`. `Want` replaces `desire`, `wish`, and `would like`, and the entry asks writers to reword so that even `want` disappears (`to use the debugger` in place of `if you want to use the debugger`). `Fail` is preferred over `crash`, `invalid` over `illegal`, and `insecure` over `nonsecure`.
- **Words to drop.** Please is not used, because politeness terms convey the wrong tone in technical information and carry different weight across cultures. Basically, essentially, simply, totally, and very are cut, and just is cut from instructions. Interesting is avoided because it substitutes for showing why something matters. We suggest is replaced by a direct instruction or by recommend.
- **Claims.** Secure is never used as a bare claim: the text should say how a feature improves security and give verifiable data, and any comparison should name what it compares against, with a cited source for a competitor. Backward compatibility and forward compatibility are used with caution because they can read as promises. Hardened needs a qualifier saying what it was hardened for. Frictionless is called inaccurate and replaced by measured improvements. Push-button and turn-key automation are called metaphors that misstate how much human work remains. Digital transformation is called vague, and the entry asks writers to describe instead of label.
- **Referring within a document.** Above and below are avoided because online formats break page position; the text uses a cross-reference or a phrase such as previously mentioned. Refer to is used for references, and see only for graphics and output. Named is used for the name of a file or object and called for code that is invoked.
- **Reader and person.** The reader is you, not the user; the guide forbids I, we, he, and she in documentation.
- **Will.** The future tense is avoided unless necessary, so a sentence about what the next section will describe takes the present tense.
- **Spelling.** US plurals `appendixes`, `indexes`, and `matrixes` are preferred; `spelled` replaces `spelt`.

### Strengths and limits as evidence

The guide is a corporate house style built from editorial practice, and it cites no research. Its numeric targets (40 words, Flesch 60 to 70, fog 9 to 12, three nouns) are stated without a source. Its strengths are specific, testable rewrites. Nearly every rule includes a before-and-after pair from real Red Hat documentation. Many of those pairs target constructions that machine-generated prose produces often, such as trailing participial phrases, bare `this`, ambiguous `using`, `allows you to`, `is designed to`, and vague verbs of improvement. Much of its reasoning comes from translation, which gives a concrete argument for rules other guides state as taste. A writer whose text will not be translated still gains from the rules, because an ambiguity that misleads a translation tool usually slows a human reader too.

Its limits come from its scope and history. It is written for training material and documentation, so its rules about tone (it forbids contractions, `please`, and the first person) assume formal technical text and would make a chat reply stiff. It depends on IBM Style, which most readers cannot open, for many questions it leaves unsettled. Sections were written at different times, and some disagree (dashes for ranges against the `through` entry; a release note that once favored straight quotation marks against the current curly-quotes rule). Its punctuation-outside-quotes and day-month-year date rules follow technical and international practice over American convention.


## What the skill applies

The guide is a secondary source for technical-writing 1.0.0. It is the main source for the identifier rule in section 4, the link rule in section 5, and items 2 and 3 of the revision pass. It also supports several bullets that came from earlier sections of writing-style.

**Section 1, goal before action.** The skill turns an `if you want` sentence into a purpose clause that starts with `to`. That rewrite comes from the guide's entries on `if you want` and `if you wish` and on `want`. The guide states the general rule of putting the condition before the action only in a commented-out block of its grammar chapter, so the published guide teaches it through examples. The bullet's main source remains the Google guide, as google-developer-style-guide.md records.

**Section 1, `please`.** The rule against `please` in instructions comes from the guide's `please` entry. The guide gives two reasons: technical text needs an authoritative tone, and politeness terms carry different weight across cultures. The skill adds `thank you` and `congratulations` and applies the rule to messages as well. Both additions come from other references.

**Section 2, `you` and `we`.** The guide's section on pronouns, which forbids the first person and asks for `you`, supports this bullet. The guide forbids `we` entirely. The skill allows `we` for a named team or organization, which matches the guide's practice of naming Red Hat as the actor.

**Section 2, recommendations.** The rule to state a recommendation as an instruction or name who recommends it, and to skip it is recommended and we suggest, comes from the guide's section on making recommendations and its we suggest entry. The guide's version names the company. The skill generalizes it to any named source.

**Section 2, present tense.** The guide's tense section and its will entry support this bullet. The guide's reason is that readers use documentation while they use the system.

**Section 2, modal verbs.** Using may only for permission comes from the guide's section on avoiding may and should and from its can, may, and you may entries. Avoiding should and need to when the reader can't tell whether something is required comes from the should, must, and need to entries. The skill narrows the need to rule. The guide says to avoid need to wherever possible and use must or required. The skill objects only when the force is unclear. The pair don't have to and must not comes from other references.

**Section 2, software by what it does.** The example of a tool that reports an error instead of complaining adapts the anthropomorphism section, which release 7.2 expanded. The parser and window examples come from other references.

**Section 3, describing steps plainly and claiming only what you can show.** The guide's section on not stating that something is easy, its simply and just entries, and its secure entry support both bullets. Revision item 1 removes simple, easy, and just. The timeless-documentation rule and most of the overclaiming list come from the Google guide.

**Section 4, naming the kind of thing after an identifier.** This bullet comes from the guide's section on object names, expanded in release 7.2, and its section on using markup. The guide asks for an object name to work as an adjective followed by a noun for its class, for no sentence to start or end with an object name, and for no name to do a verb's job. IBM's rule on file names (ibm-style-guide.md) agrees. The skill limits the noun to the first mention and allows a familiar command verb in chat with engineers. Both exceptions go beyond the guide, which applies its rule throughout documentation.

**Section 4, introducing code blocks.** Introducing a code block with a complete sentence, and not splitting a sentence around a block, follows the guide's colon rule for lead-ins to commands and code and its list rule against continuing a sentence after a list. The skill's example is its own.

**Section 4, ISO dates.** The guide allows year-month-day dates where space is short, with the format defined. The skill limits ISO dates to logs, file names, changelogs, and data. That limit is narrower than the guide's.

**Section 4, common forms of computing words.** The guide's entries agree with username as one word, web page as two words, lowercase internet, and login as a noun beside log in as a verb. Its file name entry conflicts with the skill's filename (see below).

**Section 5, headings.** The rule against bare headings such as Overview or Introduction, and the rule to put text between a heading and its first subheading, come from the guide's chapter on overall publication design. The guide also bans About as a heading. The skill doesn't name it, but its rule against bare headings covers it.

**Section 5, links.** Three practices come from the guide's cross-reference chapter (the additional-information test, the information-to-link ratio, and repetition). They are linking only to extra material, stating short essential facts in the text with the link as backup, and keeping to a couple of links per paragraph. The guide's citation section supports link text that names the destination.

**Section 5, positional references.** The guide's above and below entries support referring to sections by name, not by position. Their reason is that online formats move content. Redefining abbreviations in each section of a long document comes from other references.

**Section 6, readers whose first language isn't English.** The guide's language chapter supports avoiding idioms. Its section on invented words supports not using a verb as a noun. Its section on infinitives supports repeating to before each infinitive in a series.

**Section 9, revision item 2 (vague requirements).** This item enforces the modal-verb rule and draws on the same entries: should, need to, and you may.

**Section 9, revision item 3 (bare identifiers and positional references).** This item enforces the object-name rule, the above and below entries, and the citation section's rule against here as link text.

## What the skill omits

**Rules from the guide's design chapter that are specific to Red Hat.** The skill leaves out conventions for UI element names, figure captions, privilege escalation, example usernames and hostnames, line continuation, the notation for omitted output, editor names, product names, trademarks, non-breaking spaces, version numbers, and admonitions. Most concern Red Hat training material or publishing.

**Naming object classes.** The guide asks writers who create example names to avoid ending a name with its own class. That rule is about designing examples, not writing about them, and the skill doesn't take it up.

**Links in headings.** The guide bans links in titles, headings, and captions. The skill doesn't state this.

**Need to as a general rule.** The guide's entry treats need to as something to avoid in all requirements. The skill keeps it only where its force is unclear. That fits the owner's decision to avoid word-level rules in favor of a general rule on idiom (writing-style, section 2).

**Title case and gerund headings.** These conflict with the skill, which follows Google, IBM, and Microsoft on heading case (see below).

**Release notes and the passive.** The guide allows the passive in release notes. The skill's commit and changelog rules don't mention voice, and writing-style's passive exception (the actor is unknown or beside the point) already covers these cases.

**Telegraphic revision histories.** The guide warns that revision-history entries drift into telegraphic style. The skill's commit-message rules ask for imperative subject lines and full sentences in the body, which serves the same purpose. The skill doesn't cite the guide for them.

## Conflicts with other references

These conflicts concern the technical-writing skill. Each cited file is in `chat/skills/writing-style/references`, because technical-writing grew out of that skill's section 9 and shares its sources.

**Heading case (google-developer-style-guide.md, ibm-style-guide.md, microsoft-writing-style-guide.md).** The guide uses title case and suggests gerunds for chapter and section titles, with imperatives for tasks. Google, Microsoft, and IBM use sentence case. Google discourages an -ing word at the start of a heading, while IBM uses gerunds for tasks. The skill uses sentence case, starts task headings with a verb, and starts concept headings with a noun phrase. It follows Google, which agrees with IBM and Microsoft on sentence case and discourages an -ing opening because it translates inconsistently.

**File name or filename (oxford-dictionary-for-writers.md).** The guide writes file name as two words, and the Oxford dictionary gives filename. The skill writes filename, following the Oxford dictionary and the newer technical guides, and the closed form matches the one-word username the skill also uses.

**Punctuation with quoted literals (oxford-new-harts-rules.md, ibm-style-guide.md, nyt-manual-of-style.md).** The guide puts punctuation outside quotation marks around anything that isn't part of the quotation. Hart's Rules and IBM do the same only for quoted code. The skill puts literals in code formatting (section 4), and writing-style section 8 keeps the American placement for prose, so the guide's rule never applies.

**Should, may, and need to (google-developer-style-guide.md, microsoft-writing-style-guide.md).** The guide avoids may and should outright. Google and Microsoft keep should for recommendations and ask writers to avoid it only when its force is unclear. The skill follows Google and Microsoft on should and the guide on may.

**ISO dates (ibm-style-guide.md, ap-stylebook.md).** The guide allows ISO dates wherever space is short. IBM and AP write dates in words in prose. The skill keeps ISO dates to logs, file names, changelogs, and data.

**First person (microsoft-writing-style-guide.md).** The guide forbids we. Microsoft allows we where the company must be clear that it is speaking. The skill allows we for a named team or organization, which is close to Microsoft's rule.

## Review notes

The reading behind this file is the same as for the writing-style reference: every chapter of the built guide, all 24 word usage files, and the repository's metadata files, with commented-out XML checked separately. For phase 2 I read technical-writing SKILL.md 1.0.0, writing-style SKILL.md 2.0.0, and AGENTS.md, and I checked the writing-style reference files cited above for the conflicts. The conflicts cite files in the writing-style references folder, which holds the sources the two skills share.

I could not consult IBM Style, which the guide treats as its primary reference. The guide's condition-before-action rule exists only as a comment in Grammar.xml, so I cite the guide for the if-you-want rewrite and not for the general rule. The mapping above is my reading of which wording came from which source. Where the Google, IBM, or Microsoft reference files record the same rule, credit is shared.
