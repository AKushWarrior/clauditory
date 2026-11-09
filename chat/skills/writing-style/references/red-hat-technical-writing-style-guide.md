# Red Hat Technical Writing Style Guide

- **Source:** Red Hat, Inc., Red Hat Technical Writing Style Guide: Style Conventions for Writers and Editors, edition 7.3. The book's metadata (Book_Info.xml) gives edition 7.3. Its release notes chapter (New.xml) describes releases through 7.2, a maintenance release dated March 2026. Its revision-history appendix (Revision_History.xml) stops at revision 3-0, dated July 2021, and has not been kept up since. The copy read was the repository's state at a commit dated October 27, 2026.
- **Type:** style manual (corporate technical style guide with a word usage dictionary)
- **Where to find it:** https://github.com/StyleGuides/WritingStyleGuide (DocBook XML sources under `en-US/`, built with Publican)
- **Last reviewed:** 2026-11-05
- **Used by:** `chat/skills/writing-style/SKILL.md`

## What the source covers

### Purpose, audience, and status

The guide is Red Hat's house manual for technical writing. Its audience page and the repository README both say that it is the official guide for Red Hat training and certification content and for all other technical documentation except where stated otherwise. Both carve out one large exception: Red Hat's product documentation, written by the Customer Content Services group, follows a separate document, the Red Hat supplementary style guide, in place of this one. Rules in this guide therefore describe Red Hat training courses, exercises, and general technical writing more than the product manuals most readers would associate with the company. Its resources chapter sets the scope as software manuals, training courses, user guides, and white papers, and puts marketing and corporate content outside it. Release 7.0 (November 2024) removed the marketing-specific material and the references to the AP Stylebook.

The guide presents itself as a supplement that does not replace a complete manual. It sends readers to IBM Style as the first general reference (access is restricted to IBM and Red Hat employees, and the guide marks those links with a padlock), then to The Chicago Manual of Style and Merriam-Webster. The edition of Chicago it names is inconsistent: the grammar, punctuation, and numbers sections cite the 17th edition, while the word usage references page names the 18th. The resources page also says that company guides should be consulted before general ones and that published style references contradict each other because audiences and genres differ, which is a fair description of the conflicts this file records.

The guide is open source and accepts community contributions through GitHub issues and pull requests. That shows in its texture. Entries were written by many hands over more than a decade. Some carry editorial notes such as a remark that an entry is under review or that an example would help, and a few rules exist only in commented-out XML that the built guide does not show. The most important case is a grammar section on putting a condition before an action, which survives only as a comment; the published guide still models the pattern in examples elsewhere, as described below. This file treats the XML as published (comments excluded) as the guide's content.

### Objectives and readability targets

A short opening chapter states what Red Hat documentation should be: legally compliant, accurate and consistent, readable, comprehensible, and focused on the user without condescending to the reader or assuming their skills. It gives numeric targets: a Flesch reading-ease score of 60 to 70 and a Gunning fog index of 9 to 12, and it sets the minimum reading level of the audience at about eighth grade (ages 14 to 15). The guide does not explain how these targets were chosen or how writers should measure them, and no other chapter refers back to them.

### Grammar

The grammar chapter covers a handful of topics and defers everything else to Chicago.

**Voice.** It asks for the active voice whenever possible, because it reads more directly and cuts word count, but it says the passive is not forbidden. It names release notes and technical notes as places where the passive suits. It allows the passive when the important item should be the subject of the sentence or when a keyword should come first in a title, heading, or opening sentence. Writers are asked to weigh that front-loading against readability case by case.

**Agreement and pronouns.** A section on pronoun-antecedent agreement covers singular and plural antecedents and collective nouns. A section on countable nouns gives fewer and number for things that can be counted and less and amount for things that cannot. A section on possessives forbids them with product names and abbreviations, which are written as modifiers instead, while allowing the possessive of the company name when it means the company itself. The relative-pronoun section gives who and whom for people and that and which for things, with that for restrictive clauses, and a substitution test for choosing who or whom.

**Sentence structure.** The guide asks writers not to pack too much into one sentence and sets a ceiling of 40 words for a sentence in technical documentation. Its example breaks one long sentence into three. It allows a preposition at the end of a sentence when the alternative sounds stilted. It defines run-on sentences and comma splices and offers four repairs (a period, a semicolon, a coordinating conjunction, or a subordinating conjunction). It asks writers to repair sentence fragments, with a read-aloud test.

**Telegraphic style.** A short section warns that extreme word economy drops articles, prepositions, and verbs, and leaves sentences that are hard to parse and to translate. Its examples restore a missing article and a missing form of the verb be (a note that work remains becomes a note that work is required). It singles out revision histories, where terse entries are common, as a place where wording must stay clear for translators and customers.

**Causative verbs.** The guide asks writers to avoid constructions in which something is made to happen through have, and to replace have with a verb that names the action (apply the changes, not have the changes take effect).

**Either...or, following, and if...then.** Either...or is for a choice between exactly two options; for three or more, the guide rewrites with such as. The word following should be followed by a noun (the following steps, not the following alone), and in quiz questions it is often superfluous. An if clause is generally followed by then.

**That, this, and these.** The guide asks writers to keep the conjunction that in clauses unless the result is stilted, because it helps translation and readers whose first language is not English. A separate section asks that this, that, these, and those always be followed by a noun. The reason given is clarity for the same readers and accuracy in languages where the demonstrative changes with the gender of its noun. Its examples turn a bare this at the start of a sentence into this action, or rewrite with consequently.

**Verbosity and word order.** A verbosity section asks writers to cut unnecessary words, with examples that turn a nominalized verb back into a verb (`perform the installation of the product` becomes `install the product`). Another example cuts a sentence about a service that provides services for administration into one that says what the service does. A word-order section asks for the arrangement that creates the least ambiguity, usually subject, verb, and object with modifiers next to what they modify.

**Contractions and abbreviations.** Red Hat documents use no contractions, because they mark informal writing and cause problems in translation. The Latin abbreviations `e.g.` and `i.e.` are written out as for example and that is; the word usage entries repeat the rule.

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

**Slashes.** The guide forbids a slash to mean one of two options, so enable or disable replaces the slashed form, and A or B, or both replaces `and/or`.

**Quotation marks.** Curly quotation marks are for direct quotations and straight ones for code, syntax, hexadecimal values, and links. Punctuation goes inside the closing quotation mark only when it belongs to the quotation or to a programming element that uses quotation marks, and otherwise outside. The rule follows the logical style common in technical writing, and it differs from the American convention of putting periods and commas inside. Release 5.1 (January 2022) records an update to this rule, and the revision history holds a commented-out entry that once stated it more broadly.

**Apostrophes and exclamation points.** Apostrophes never form plurals. Possessives follow standard rules, with the guide's variants for singular names ending in s. Exclamation points do not end sentences; the mark appears only as a character in commands.

**Dashes.** The punctuation chapter has no dash section, but the dash entry in the word usage dictionary says to use a dash only for ranges (pages, letters, dates) and otherwise not at all in technical content. Commas, parentheses, or a colon take its place. The em dash entry was removed in release 7.0. The through entry separately forbids using a hyphen or dash to mean through.

**Lists and tables.** List punctuation is covered in the writing chapter (below). Table cells end with periods when every cell in the column is a complete sentence and take none when every cell is a fragment; mixed columns should be made parallel where possible and otherwise punctuated cell by cell.

**Naming characters and keys.** The chapter ends with conventions for naming punctuation marks and special characters in text (the name, a describing noun where the name includes one, and the symbol in parentheses, never the bare symbol). It adds a table of names and a table of keyboard key labels. It distinguishes pipe (for command pipelines) from vertical bar (for a logical or in expressions). These are documentation conventions with no bearing on general prose.

### Design

The design chapter collects conventions for technical publications, some of which apply beyond Red Hat.

**Headings.** The guide uses title case for headings, titles, and references to other publications' titles, with sentence case for figure captions, legends, diagram labels, and table column headers. Headings carry little punctuation and no closing period. For larger units such as chapters, the gerund form of a specific action verb is the usual choice; tasks the reader performs can take an imperative verb; purely informational sections take a noun phrase. Vague verbs such as understanding, introducing, or exploring are replaced by a specific verb or a noun phrase, and one-word titles are avoided. Headings should not contain file, command, or markup names; those belong in the text.

**User interface.** UI elements are named with the capitalization and spelling of the interface (with spelling errors corrected), without their punctuation (an ellipsis on a button label is dropped) and usually without the element type (the label alone, without the word button). Go to covers any sequence of selections in a UI. Figures need a caption or callouts only when the surrounding text does not describe them, and never both.

**Commands and examples.** The guide explains command syntax with worked examples, warns against showing force and assume-yes options without reason, asks for line continuation characters (and no manual indentation) in long commands, and sets a notation for omitted output. It asks writers not to name an editor when telling readers to view or edit a file, and not to use line numbers as reference points, because they change. Example usernames, hostnames, and IP addresses are lowercase, use reserved example domains, never use real public addresses or real people, and draw on a diverse set of realistic names when a case study needs them. A long section on escalated privileges sets conventions for sudo in training exercises.

**Object names.** Release 7.2 expanded the guidance on referring to named objects such as files, commands, variables, and resources. A sentence should be complete without the object name. The object name should work as an adjective followed by a noun that names its class (`the ls-blue router`, `the /mnt/save directory`). A sentence should not start or end with an object name. Where the noun would duplicate the name, or the object can be named in plain words, the guide prefers plain language. A related section on naming object classes asks writers who create examples to avoid names that end with their own class, so text does not end up with a doubled noun. File-system names keep their creators' capitalization, even at the start of a title.

**Abbreviations and names.** Most acronyms and initialisms are spelled out on first use, but well-known ones such as HTML are not. The expansion is lowercase unless it is a proper noun. The article follows pronunciation (an RTS, a RAM upgrade), plurals add a lowercase s with no apostrophe, and abbreviations take no possessive. Product names follow company rules. They use the full name on first use, take no articles before names that include their own noun, take no possessives, and take no line breaks inside a name. Non-breaking spaces go in defined places, and the major version number is used unless a minor release is meant.

**Admonitions and recommendations.** Note, Important, and Warning admonitions should be brief and rare, so they keep their force. Recommendations name the recommender (Red Hat recommends) in place of the impersonal it is recommended.

**Citations and links.** External sources should be authoritative, credible, and publicly accessible, and link text should describe the destination in place of a bare here.

### Overall publication design

This short chapter covers book structure. Titles combine product, version, and publication name, and a one-phrase subtitle follows. A descriptive abstract says what the publication covers, how it conveys it, and why it exists (and for whom). Introductions use the product name instead of the book name. Two rules apply to any document. Two headings should never appear in a row without text between them; if no useful text fits, one heading is probably unnecessary. And a title should not be Overview or About on its own, nor should a chapter or section have a separate introduction section; introductory material goes in the body text under the title.

### Choosing appropriate language

This chapter, the longest prose chapter, covers word choice for a global and translated audience, inclusive language, and ambiguity.

**Misleading language.** The guide tolerates best practices as a common search term but prefers recommended practices when in doubt, and notes the concern about superlatives without supporting data. Idioms such as first come, first served are to be avoided in content that will be translated.

**Slang and jargon.** An alphabetical list of about 75 terms gives replacements or says not to use them. The entries fall into a few groups. Business jargon: `best-of-breed`, `core competency`, `going-forward basis`, `net-net`, `synergy`, `value-added`, `paradigm`, `mission-critical` where nothing is critical to a mission, `solutions-based`, `silo`. Idioms and metaphors: `apples to apples`, `at the end of the day`, `boil the ocean`, `drop the ball`, `low-hanging fruit`, `on the same page`, `think outside the box`, `double-edged sword`, `kettle of fish`. Invented or informal words: `administrivia`, `automagic`, `bucketize`, `productize`, `performant`, `upskill`, `ask` as a noun. Padded phrases: `at this point in time` for now, `in concert with` for with, `right before` for immediately before. Vague claims: `improve and enhance`, `meet your needs`, `exceed your expectations`, `harness the power`, `frictionless`. Several entries apply directly to machine-written prose. The entries for `dig deeper` and `delve deeper` flag them as likely to translate awkwardly, `leverage` is to be avoided in favor of use, and `is designed to` is replaced by a statement of what the product does. The entry for `if you want` or `if you wish` turns the condition into a purpose clause that starts with to.

**Neologisms and phrasal verbs.** The guide discourages new words formed by turning verbs into nouns and nouns into verbs, and asks for one-word verbs where they exist (`click` for `click on`, `complete` for `fill in` or `fill out`, `omit` for `leave out`, `print` for `print out`).

**-ing forms.** Release 7.2 expanded a section on words ending in -ing. A gerund suits a title that names an activity in general, a sentence whose subject is the activity and a list of activities that is parallel with its lead-in. When a filler verb phrase such as is used for pushes the real action into a gerund, the sentence should be rewritten around the active verb. A gerund should not replace an instruction in a procedure step. The section's sharpest rule concerns present participles. The guide asks writers to avoid a noun followed by a word ending in -ing, with or without a comma between them, because the construction is ambiguous and hard to translate. Its examples are sentences that end in a participial phrase tacked onto a finished clause (`a table that shows dropped packets, enabling you to identify flows`; `an add-on that extends support, providing stability`). Each is rewritten with which, so that, or a to-infinitive, which makes the relation explicit. Noun phrases that name a single thing, such as error logging, are exempt.

**Anthropomorphism.** Software and hardware do not think, talk, refuse, or complain. Release 7.2 added examples, and each rewrite names the observable behavior (the tool produces errors, the proxy communicates with a server).

**Allow.** Release 7.1 added a rule that allow is for permission only (a license allows reuse, sudo allows a user to run a command). For anything else, the sentence should focus on what the user or system does (you can use the command to filter messages) or use another verb such as enable.

**Inclusive language.** The guide drops white and black as metaphors for good and bad (`allowlist`, `blocklist`, or `denylist`). It drops `master` when paired with `slave` (`main` for a default branch; `primary`, `controller`, `leader`, `worker`, `replica`, or `secondary` elsewhere). It also drops gendered terms such as `man-hours`, `he or she` in favor of they, terms with neurodiversity bias such as `sanity check`, `disabled` as a description of a person, and superlative job titles such as `guru`, `ninja`, and `rockstar`. It says the guidance will change over time.

**Redundant words.** Filler words such as actually, really, simply, and very are to be cut, along with pairs that repeat a meaning (revert back) and words that add nothing (located on for on). The word usage dictionary adds basically, essentially, totally, and just to the list.

**Ambiguity.** A section with 13 subsections catalogs common sources of ambiguity, each with an example and a rewrite. May and should are avoided because each has two senses (possibility or permission; recommendation or obligation), and the rewrites use can, might, must, or is expected to. Statements that a task is easy or a product is user-friendly are cut in favor of describing how to do the task. Proper nouns are capitalized so translators can tell them from concepts. A word with several senses is made explicit with a noun. Two infinitives in one sentence are separated so it is clear what each one belongs to. Noun-noun phrases that hide whether the first noun is plural are rewritten. Only is placed directly before the word it modifies, which is often not the verb. Pronouns with two possible antecedents are replaced by the noun. One concept keeps one term throughout, because a second term suggests a second thing. There is, there are, and there might be are rewritten so the real subject leads. `using` is replaced by `by using` or `that uses`, depending on which is meant. An opening verb phrase must modify the subject that follows it (no dangling modifiers).

**Dates, times, and numbers.** All-numeric dates are forbidden in running text because the order of day and month varies by country. The guide writes the month as a word in day-month-year order. It allows ISO 8601 (year-month-day) where space is short, with the format stated. Times use uppercase `AM` and `PM` without periods, and noon and midnight replace `12 PM` and `12 AM`. Numbers zero through nine are spelled out, as are numbers that begin a sentence, a number directly before another number, approximations, and very large values. Numerals are used for 10 and above, and for smaller numbers in the same paragraph as larger ones, as well as for negatives, fractions, percentages, decimals, measurements, section references, registers, code, and versions. Four-digit numbers take no comma; five digits and more do. Numbers for telephones use spaces and an international country code.

### Writing clearly and succinctly (translation)

This chapter (Translation.xml) covers structure for readability and translation.

**Lists.** A single sentence should not be formatted as a bulleted list. Bulleted lists are for unordered items and usually have at least three. Numbered lists are for sequences or for items referred to elsewhere. Variable (term and definition) lists suit options and arguments and are often better than tables, which are costly to translate and suit only data that needs a grid. Procedures always have a title, and nested lists go no deeper than two levels. A list must not interrupt a sentence that continues after it. A lead-in should be a complete sentence, because translation tools treat the lead-in and each item as separate segments. Items must be grammatically parallel. Items that are complete sentences end with periods; items that are fragments (including ones with a verb but no subject) take no end punctuation; and a list whose fragments are followed by sentences ends every entry with a period. The chapter also discusses normal versus compact spacing.

**Noun stacks.** Strings of modifiers and nouns are hard to translate when several readings are possible, and the guide asks for no more than three consecutive nouns.

**Grammatical gender.** Pronouns such as it and they, and abbreviations that can name two things (RPM the package or RPM the package manager), are hard to translate into languages with gendered nouns; the fix is to repeat the noun.

**Markup, code blocks, and entities.** Untranslatable terms should be marked up and should not do grammatical work in a sentence (a path or command name should be followed by a noun such as directory or command, and should not stand in for a verb). Commentary should not be placed inside a code block, where it can be mistaken for output and left untranslated. Entities should be limited to those a build requires.

### Cross-references

This chapter argues that a writer who sends readers from link to link is serving their own convenience at the reader's cost. A link should point to additional information, never to information the reader needs to finish the task in hand; that information belongs in the text, with a link at the end. Paragraphs should hold no more than a couple of links, and links do not belong in titles, headings, or captions. Vital information of up to a couple of paragraphs (or half a page, or five table rows) should be repeated where it is needed, not cross-referenced, because repetition saves the reader time and matters for safety.

### The word usage dictionary

The usage dictionary runs from 0-9 through XYZ and holds roughly 500 entries. Most settle the spelling, spacing, hyphenation, or capitalization of technical terms (`back end` as a noun and `back-end` as an adjective, `backup` and `back up`, `file system`, `plug-in`, `stand-alone`, `setup` and `set up`, and `login` and `log in`) or give Red Hat product and trademark rules. Releases 7.1 and 7.2 added entries for AI terms. AI need not be spelled out, and `gen AI` is acceptable. `large language model` and `machine learning` are not hyphenated even as modifiers, `fine-tune` is hyphenated, and `inference` can be a verb in an AI context.

A smaller group of entries sets general usage rules that bear on prose:

- **Conjunctions in one sense.** `because` is for reasons; `since` and `as` are for time only. `once` means one time and is not a conjunction for after or when. `while` is for simultaneous events, with `although` for contrast. `like` is not used for `such as`. `whether` replaces `if` for alternatives. `as long as` is for length or time, with `provided that` otherwise.
- **Modal verbs.** `can` is for possibility and `may` only for permission, and `can` is preferred over `could` and `may` because it is less tentative. `must` is for requirements and `should` never states something the user must do. The entry for `need to`, `needs`, and `needs to be` asks writers to avoid them where possible and to use `must`, `required`, or similar words. The entry for `you may` suggests deleting it.
- **Plain words.** `use` replaces `utilize`, and `open` replaces `bring up`. `examine` or `inspect` replaces `look at`, and `many` replaces `lots of` and `bunches of`. `emit` or `issue` replaces `send out`, `start` replaces `start up`, and `now` replaces `right now` and `at this point in time`. `want` replaces `desire`, `wish`, and `would like`. The entry asks writers to recast the sentence so that even `want` disappears (`to use the debugger` in place of `if you want to use the debugger`). `fail` is preferred over `crash`, `invalid` over `illegal`, and `insecure` over `nonsecure`.
- **Words to drop.** `please` is not used, because politeness terms convey the wrong tone in technical information and carry different weight across cultures. `basically`, `essentially`, `simply`, `totally`, and `very` are cut, and `just` is cut from instructions. `interesting` is avoided because it substitutes for showing why something matters. `we suggest` is replaced by a direct instruction or by `recommend`.
- **Claims.** Secure is never used as a bare claim: the text should say how a feature improves security and give verifiable data, and any comparison should name what it compares against, with a cited source for a competitor. Backward compatibility and forward compatibility are used with caution because they can read as promises. Hardened needs a qualifier saying what it was hardened for. Frictionless is called inaccurate and replaced by measured improvements. Push-button and turn-key automation are called metaphors that misstate how much human work remains. Digital transformation is called vague, and the entry asks writers to describe instead of label.
- **Referring within a document.** Above and below are avoided because online formats break page position; the text uses a cross-reference or a phrase such as previously mentioned. Refer to is used for references, and see only for graphics and output. Named is used for the name of a file or object and called for code that is invoked.
- **Reader and person.** The reader is you, not the user; the guide forbids I, we, he, and she in documentation.
- **Will.** The future tense is avoided unless necessary, so a sentence about what the next section will describe takes the present tense.
- **Spelling.** The US plurals `appendixes`, `indexes`, and `matrixes` are preferred, and `spelled` replaces `spelt`.

### Strengths and limits as evidence

The guide is a corporate house style built from editorial practice, and it cites no research. Its numeric targets (40 words, Flesch 60 to 70, fog 9 to 12, three nouns) are stated without a source. Its strengths are specific, testable rewrites. Nearly every rule includes a before-and-after pair from real Red Hat documentation. Many of those pairs target constructions that machine-generated prose produces often, such as trailing participial phrases, bare this, ambiguous `using`, `allows you to`, `is designed to`, and vague verbs of improvement. Much of its reasoning comes from translation, which gives a concrete argument for rules other guides state as taste. A writer whose text will not be translated still gains from the rules, because an ambiguity that misleads a translation tool usually slows a human reader too.

Its limits come from its scope and history. It is written for training material and documentation, so its rules about tone (contractions, `please`, and the first person are all excluded) assume formal technical text and would make a chat reply stiff. It depends on IBM Style, which most readers cannot open, for many questions it leaves unsettled. Sections were written at different times, and some disagree (dashes for ranges against the through entry; a release note that once favored straight quotation marks against the current curly-quotes rule). Its punctuation-outside-quotes and day-month-year date rules follow technical and international practice over American convention.


## What the skill applies

The guide is a secondary source for writing-style 2.0.0. Its biggest contribution is to the rules on ambiguity in sections 2 and 3 and to revision items 3, 7, 12, and 13. Most of these rules arrived in version 2.0.0. Several also rest on IBM or Google, and this section names those cases.

**Section 2, be specific (last sentence).** The rule against calling something secure, improved, faster, or easy to use without saying how comes from the secure entry, the entry on `improve` and `enhance`, and the section on not stating that something is easy. The secure entry supplies the test of naming the mechanism and the comparison. The skill adds "by how much" and faster, and it moves the rule from technical text into the core skill, so it now covers chat replies and reports.

**Section 2, put the person who acts in the subject.** This bullet comes from the section on allow (release 7.1). The guide limits allow to permission and recommends rewriting around what the user does. The skill goes further. The guide offers enable as an acceptable replacement, but the skill saves allow, enable, let, and permit for actual permission. That stricter version comes from IBM's grant-verb entries (ibm-style-guide.md).

**Section 2, prefer verbs to nouns built from verbs.** The new example, install the package in place of perform the installation of the package, adapts the guide's verbosity section.

**Section 2, break up noun stacks.** The three-noun limit is the guide's (noun stacking, in the Translation chapter). The guide gives translation as the reason. The skill gives the reader's difficulty instead, and its example is its own.

**Section 2, give connecting words one job each.** The guide's word usage entries supply the pairs: `because` for a reason in place of `since` and `as`, `after` or `when` in place of `once`, `although` in place of `while`, `such as` in place of `like`, and `whether` in place of `if`. Google states the `since` and `once` rules too (`google-developer-style-guide.md`). The skill's explanation (each of these words has a second sense that makes the reader choose) paraphrases the guide's reason for the since entry.

**Section 2, avoid words readers take two ways.** The guide's entry on biannual, bimonthly, and biweekly is one source for the biweekly and bimonthly examples. The other examples (momentarily and table) come from other references.

**Section 2, clichés.** `harness the power`, `best-of-breed`, and `on the same page` come from the guide's slang list.

**Section 2 (intensifiers and fillers).** Actually, basically, and essentially come from the guide's section on redundant words and its basically and essentially entries.

**Section 2, abbreviations.** The rule to choose a or an by pronunciation matches the guide's articles section under abbreviations. The guide also supports spelling out on first use. Neither part of the bullet depends on this guide alone.

**Section 2 (`e.g.`, `i.e.`, and `etc.`).** The guide supports writing for example and that is. The skill is looser than the guide here, since it keeps the abbreviations in parentheses and tables. Its `etc.` rule comes from other references.

**Section 2 (slashes, `and/or`, and `file(s)`).** The guide's slash and parentheses sections and its `and/or` entry supply the three rewrites. The skill adds exceptions the guide doesn't state: fixed terms such as client/server, and ampersands in names. The guide's ampersand entry allows the mark only in quoted third-party names and tight design space, so the second exception is close to it.

**Section 3, sentence length.** The ceiling of 40 words is the guide's. The target of most sentences under about 25 words comes from IBM (ibm-style-guide.md). The guide gives no evidence for its number.

**Section 3 (demonstratives and pronouns).** The bullet comes from two grammar sections (on this, that, these, and those with a noun, and on ambiguous pronouns) and from the Translation section on grammatical gender. The skill's reason, that the reader shouldn't have to search the previous sentence for the subject, replaces the guide's reasons, which are translation and second-language readers.

**Section 3, opening phrases, using, and only.** These three rules come from the guide's ambiguity catalog (verb phrases, use of using, and misplaced only). Google and Microsoft also state the using and only rules. The examples are the skill's own.

**Section 3, empty openers.** The new example, which rewrites there might be, follows the guide's section on there is and there are, which names the hedged variant.

**Section 3, trim padding.** The rule against words that repeat a meaning and the example revert back come from the section on redundant words. The other examples (past history, PIN number, the CLI interface) come from other references.

**Section 3, keep articles, that, and be.** This bullet combines the guide's sections on telegraphic style and on keeping that in clauses. The skill's example of a clipped status message is its own.

**Section 3 (`either/or` for two choices).** This comes from the section on `either/or`.

**Section 3, comma splices.** The guide's sections on run-on sentences and conjunctive adverbs support the rule against joining sentences with a comma before however or therefore. The bullet rests on other references as well.

**Section 6, lists.** Several list rules match the guide's rules on lists. The lead-in should be a complete sentence. A sentence should not continue across the list. Items should be all sentences with periods or all fragments with no end punctuation. The guide's reason is translation tools, and IBM and Microsoft state the same rules. The guide doesn't say whether fragments take a capital, so that part is the owner's decision.

**Section 7, peer register and contractions.** The guide's simply entry supports the ban on simply as a way of talking down. The guide's translation argument against contractions supports the skill's clause about spelling out contractions in text likely to be translated.

**Section 8, mechanics.** Several section 8 rules agree with the guide, but none depend on it alone:

- The serial comma.
- The ban on a colon between a verb and its object.
- A capital after a colon when a full sentence follows.
- No hyphen after an -ly adverb.
- Phrasal verbs as two words, with the noun as one (backup, setup, login).
- Hyphens before a capitalized word.
- No apostrophes in plurals.
- The that and which rule.
- The fewer and less rule.
- Cannot as one word.

The date bullet shares the guide's ban on all-numeric dates and its preference for noon and midnight. The quotation bullet adopts the resolution from the phase 1 proposals: American placement in prose, and code formatting for literals so that punctuation never attaches to them.

**Section 10, item 3 (trailing participles).** This is the guide's largest single contribution. The present-participle section (release 7.2) bans any noun followed by an -ing word and rewrites each example with which, so that, or a to-infinitive. The skill turned item 3 from a list of trigger words into a rule about the construction. It added `enabling you to`, `allowing`, `providing`, and `resulting in`, and it states the guide's reason: the reader has to guess the relation. The skill's fixes are which, so that, and a new sentence. It leaves out the to-infinitive fix.

**Section 10, item 7.** The phrase `it's interesting to note` comes from the `interesting` entry.

**Section 10, item 12.** Allows you to and enables you to come from the section on allow, and is designed to comes from the guide's slang list. IBM lists the same phrases. Lets you is not in the guide.

**Section 10, item 13.** These slips enforce rules described above that the guide supports. They are a bare This, a slash for or, `e.g.` and `i.e.` in running text, an apostrophe in a plural, a hyphen after an -ly adverb, an all-figure date, and a missing serial comma.

**Section 10, item 14.** The guide's section on easily confused words supplies affect and effect and the ensure and insure pair.

**Examples, describing a tool.** The example's before text uses `serves as`, `allowing`, `e.g.`, and `etc.`, and its after text uses `you can use`. It illustrates the allow rule and the abbreviation rule. The guide supplied both.

## What the skill omits

**Rules that went to the technical-writing skill.** The proposals on object names, modal verbs, recommendations, the if-you-want rewrite, links, headings, `please`, and anthropomorphism examples went to the technical-writing skill. Its own reference file, `chat/skills/technical-writing/references/red-hat-technical-writing-style-guide.md`, records them.

**The need, needed, and necessary question.** The guide's entry for need to asks writers to use must or required in place of need. The owner chose a general rule that word choice stays idiomatic over word-level rules, so section 2 has one bullet on idiom and says nothing about need. The guide doesn't support that bullet. Its only treatment of idiom is about avoiding figurative phrases, and it has no rule on whether a phrase sounds natural. The need-to entry survives only in technical-writing, where it has a narrower purpose.

**Part of the cliché list.** The skill took three of the guide's phrases and left out `dig deeper`, `frictionless`, `mission-critical`, and `synergy`. The word `delve` stays in item 8 on the strength of the AI-vocabulary studies. The guide's translation concern about `dig deeper` and `delve deeper` supports keeping it, but the skill doesn't depend on that concern.

**Part of the filler list.** Simply, totally, and just are not in section 2's list of fillers. Simply appears in section 7 as a sign of talking down, and just appears in the technical-writing revision pass as a word that makes a step sound easy. Totally is gone.

**One-word verbs for phrasal verbs.** The guide's preference for `omit` over `leave out` and `complete` over `fill out` didn't make it in. Section 8 has a rule about phrasal verbs, but it only covers whether to write them as one word or two. The proposal was a word-level rule, the kind the owner decided against.

**The to-infinitive fix in item 3.** The skill lists two connectors and a new sentence. It leaves out the guide's third rewrite, a to-infinitive for a purpose. So that covers most of the same cases.

**Mixed lists.** For a list where fragments are followed by sentences, the guide ends every item with a period. The skill asks the writer to reword so all items are the same kind. The skill's rule is stricter and simpler to check.

**Readability scores and a ban on contractions.** The guide gives Flesch and fog targets with no method or source, and a score target would push toward short words over exact ones. Its blanket ban on contractions fits training manuals. Section 7 keeps contractions for conversation.

**Product, UI, and publishing rules.** The skill leaves out trademarks, product names, non-breaking spaces, keyboard keys, privilege escalation, example usernames, figure captions, abstracts, and most single-term hyphenation entries. They are Red Hat-specific or concern publication production.

## Conflicts with other references

**Heading case (google-developer-style-guide.md, ibm-style-guide.md, microsoft-writing-style-guide.md).** The guide puts headings in title case and prefers gerund titles for chapters. Google, IBM, and Microsoft use sentence case. Google discourages an -ing word at the start of a heading, while IBM uses gerunds for task headings. The skill writes headers in sentence case with no closing punctuation (section 6), following Google, IBM, and Microsoft, which agree with one another on case.

**Punctuation with quotation marks (nyt-manual-of-style.md, ap-stylebook.md, chicago-manual-of-style.md, oxford-new-harts-rules.md, ibm-style-guide.md).** In technical documentation the guide puts periods and commas outside a closing quotation mark unless they belong to the quotation or to code. The American manuals put them inside. Oxford's Hart's Rules (14.5.2) and IBM both make an exception for quoted code. The skill puts periods and commas inside, following the American convention that section 8 uses throughout. It puts literals in code formatting, which settles the case that motivates the guide's rule.

**Dates (nyt-manual-of-style.md, ap-stylebook.md, ibm-style-guide.md, oxford-new-harts-rules.md).** The guide and IBM write day, month, and year (`1 April 2021`), which is also British order in Hart's Rules. The news manuals put the month first. The skill uses the American order. It shares the guide's ban on all-figure dates, and it keeps ISO dates for logs and data in technical-writing.

**Times (ap-stylebook.md, ibm-style-guide.md, oxford-new-harts-rules.md).** The guide writes AM and PM in capitals without periods. AP, IBM, and Hart's Rules use lowercase a.m. and p.m. with periods. The skill writes 4 p.m. All of these sources agree on noon and midnight.

**Commas in four-digit numbers (nyt-manual-of-style.md, ibm-style-guide.md).** The guide omits the comma in four-digit numbers, as IBM does except where consistency calls for one. The skill uses commas from four digits up, except in years, versions, and code, following the news manuals.

**Figures in a mixed set of numbers (nyt-manual-of-style.md, google-developer-style-guide.md).** The guide uses figures for every number in a paragraph that contains one of 10 or more. Google applies the rule within a sentence. By its reference file, the NYT spells out small numbers even in a mixed series. The skill keeps the NYT rule: one through nine spelled out and figures for 10 and above, even in a mixed series. An earlier version used figures for comparable numbers in the same series; version 2.0.0 dropped that rule because it contradicted the NYT. That is narrower than the guide's paragraph-wide rule.

**Hyphens after a verb (oxford-new-harts-rules.md, oxford-dictionary-for-writers.md).** The guide hyphenates up-to-date and out-of-date even after the noun. Oxford leaves compound modifiers open after a verb, and so does the skill (section 8).

**Doubled vowels after a prefix.** The guide hyphenates a prefix that ends in the same vowel the word begins with (`pre-emptive`, `semi-independent`), with listed exceptions. The skill closes up most prefixes and gives preexisting as an example. The skill's closed form applies the general closed-prefix rule it takes from IBM and the push toward fewer hyphens in Hart's Rules. I did not find a reference file in this folder that states the doubled-vowel rule either way.

**Filename (oxford-dictionary-for-writers.md).** The guide writes file name as two words. The Oxford dictionary gives filename as one, and technical-writing uses filename. This conflict affects technical-writing more than this skill.

**Allow, enable, and lets you (google-developer-style-guide.md, ibm-style-guide.md).** The guide offers enable as an acceptable substitute for allow. Google's word list replaces allows you to with lets you. IBM treats all of these grant verbs as stand-ins for a sentence about what the user does. The skill follows IBM and lists all three in item 12.

**Contractions (google-developer-style-guide.md, microsoft-writing-style-guide.md, ibm-style-guide.md).** The guide and IBM ban contractions in technical text, and Google and Microsoft encourage them. The skill uses contractions in conversation and spells them out in formal documents and in text likely to be translated (section 7).

**Serial comma (nyt-manual-of-style.md, ap-stylebook.md).** The guide requires the serial comma, and the news manuals omit it. The skill uses it as the owner's choice. The guide is one of several sources that agree.

**Chicago edition (chicago-manual-of-style.md).** The guide cites the 17th edition of Chicago in its grammar, punctuation, and numbers sections and the 18th in its word usage references. The skill's Chicago reference covers the 18th edition. I found no rule the skill took from this guide that depends on the difference.

## Review notes

I read every chapter in the built book: the preface (with audience, release notes, and feedback), objectives, grammar (with easily confused words and punctuation), design (with overall publication design), choosing appropriate language, writing clearly and succinctly, cross-references, and resources. I read all 24 word usage files, from 0-9 through XYZ, plus README.md, RELEASENOTES.md, Book_Info.xml, Revision_History.xml, and two files the build leaves out (WUG_Intro.xml and WUG_References.xml). I worked from text extracted from the XML with comments removed. I checked commented-out passages against the XML, chief among them the condition-and-action section in Grammar.xml, which the built guide doesn't show.

The repository was a shallow clone with one commit (October 27, 2026), so I could date rules only through the release notes chapter. Three version signals disagree: the book metadata says edition 7.3, the release notes stop at 7.2 (March 2026), and the revision history stops at July 2021. I could not consult IBM Style, which the guide treats as its main reference and which is restricted to IBM and Red Hat staff. I didn't confirm AP's position on since used for because. I raised it in phase 1, and it doesn't affect any conflict recorded here.

For phase 2 I read writing-style SKILL.md 2.0.0, technical-writing SKILL.md 1.0.0, and AGENTS.md. In this folder I checked the reference files for NYT, AP, Chicago, Hart's Rules, the Oxford dictionary, IBM, Google, and Microsoft for the conflicts above. Where I attribute a rule to IBM or Google, I relied on their reference files, not the guides themselves. The mapping in "What the skill applies" is my reading of which wording came from which source. Several rules have more than one source, and where IBM, Google, or another reference proposed the same rule, the credit here may be shared.
