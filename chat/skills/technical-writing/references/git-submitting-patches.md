# Git's guidelines for submitting patches

- **Source:** The Git project, *Submitting Patches* (Documentation/SubmittingPatches in the git.git source tree, master branch), maintained by the Git maintainer and contributors. Read together with the DISCUSSION section of the *git-commit* manual page (Documentation/git-commit.adoc, master branch). Living documents with no fixed edition; the copies reviewed were current on the master branch in November 2026.
- **Type:** official documentation
- **Where to find it:** https://github.com/git/git/blob/master/Documentation/SubmittingPatches and https://github.com/git/git/blob/master/Documentation/git-commit.adoc (rendered as the git-commit manual page at https://git-scm.com/docs/git-commit)
- **Last reviewed:** 2026-11-05
- **Used by:** chat/skills/technical-writing/SKILL.md

## What the source covers

SubmittingPatches is the contributor guide for the Git project itself. It tells a developer how to prepare a change, describe it, certify it, and send it to the project's mailing list, where every change is reviewed in public before the maintainer merges it. About a third of the document concerns writing: how to split work into commits, what a commit message must explain, how its first line is formed, how to refer to other commits, and what goes in a cover letter. The rest covers the review cycle, branch choice, tests, sign-off and trailers, the mechanics of mailing patches, conflict handling, subsystem maintainers, the project's CI setup, and workarounds for particular mail clients.

It is a set of house rules written by practitioners, with no study or data behind it. Its authority comes from long use: the conventions it describes are the ones Git's own history follows, and many other projects copied them, partly because Git's tools display the first line of a message as the commit's title. The reasons it gives are practical and stated from the reader's side: a reviewer has to judge the change, and a future maintainer has to understand it years later.

### Why the log message matters

The section on describing changes opens by ranking the log message as equal in importance to the code. Its argument is that well-commented code can explain how it works, while only the message can say why it exists. It gives two reasons a later developer needs the why. First, the code may not do what the author meant, and a written statement of intent lets someone fix it to match; the document adds that authors often find their own bugs while writing that statement. Second, the code may deliberately cover only part of a problem, and recording why the author stopped where they did lets a successor decide whether to extend the work or leave it alone. The worked example is a change applied to directories but not to files: stating the property of directories that justified it lets a reader check whether files share that property.

This is the strongest part of the document, because it ties a writing rule to a concrete maintenance task. It also explains why the guide wants reasons over narration: the diff already shows what changed.

### The subject line

The first line of the message is a short description. SubmittingPatches calls 50 characters a soft limit and points to the git-commit manual for the reason. That manual page's DISCUSSION section is brief: it says a message need not follow any format, recommends a single short summary line of at most 50 characters, a blank line, and then a fuller description, and explains that Git treats the text before the first blank line as the commit's title throughout its tools. Its example is git-format-patch, which turns the title into an email subject and the rest into the email body. The two documents differ slightly in strength: the manual page frames 50 as a ceiling within an optional recommendation, while SubmittingPatches frames it as a soft limit.

The first line ends with no period. By convention it usually starts with an area prefix: a file name or short identifier for the part of the code being changed, followed by a colon. The document's examples use a documentation prefix and a file-name prefix. When a contributor is unsure which prefix to use, the guide tells them to run `git log --no-merges` on the files they are changing and copy what earlier commits did. That instruction is the general principle behind the prefix rule: the history already in the repository sets the convention.

After the prefix, the first word is lowercase unless it would be capitalized anywhere in a sentence. The guide's counterexample is a ref name that Git always writes in capitals, which keeps its capitals after the prefix. The lowercase rule applies only to the word that follows the prefix.

### What the body must say

The body explains the problem the change solves, meaning what is wrong with the code as it stands without the change. It then justifies the solution by saying why the result is better. Where relevant, it names alternatives the author considered and dropped, and it records how design or viability concerns raised in review were settled, so that the history shows why this approach won. An earlier section adds that the explanation should be detailed enough for a reader to judge whether the change is a good idea without reading the patch, and that a good description summarizes the point in the subject and covers the motivation, the approach, and, for a revised series, how it differs from the previous version. If the description grows long, the guide reads that as a sign the commit should be split.

### Tense, mood, and self-containment

Two grammatical rules follow. The description of the status quo is in the present tense: the message says what the code does today when given some input, and avoids describing it as something it used to do. The guide adds that a word such as currently is unnecessary, because by project convention the problem statement always describes the code without the change.

The change itself is described in the imperative mood, as an order given to the codebase. The guide contrasts this with two forms it rejects: a sentence whose subject is the patch, and a first-person past-tense account of what the author did. The rule covers the whole description of the change, the subject line included.

In the same paragraph, the guide asks for messages that can be understood without outside resources. In place of a link to a mailing-list archive, the author summarizes the points of the discussion that matter. This puts the reasoning into the history itself, where it survives if the archive moves or disappears.

### Referring to other commits

The guide lists three reasons to cite another commit: it introduced the bug being fixed, it introduced the feature being extended, or it conflicts with the current work in a trial merge. A citation gives the abbreviated hash followed by the cited commit's subject and date in parentheses, and the guide names a `git show` pretty format that produces this form. Pairing the hash with the subject means a reader learns what the cited commit did without looking it up, which is the self-containment principle applied to cross-references.

### Cover letters and the AI policy

The cover letter, which introduces a series of patches, has its own short list of aims: show that the value of the change outweighs its regression risk and maintenance cost, make the series understandable without prior context, keep later versions readable for people who missed earlier rounds, and summarize what changed since the previous round. An experimental section invites a two-to-five-line topic summary suitable for the release notes.

A section on artificial intelligence is directly relevant to this skill. It raises two problems with AI-generated contributions: it is unclear whether a contributor can make the certification of origin the project requires for substantial generated content, and generated code, messages, and documentation are often wrong in ways that persist even after correction. The project says it will reject submissions that look machine-generated, read as overly formal or bloated, look plausible but make no sense, or that the sender cannot explain. It recommends using AI to guide one's own work, to debug, and to check a draft against the project's style, over asking for a complete solution and pasting it in. This is a maintainer's policy statement, with no measurement behind it, but it records that at least one major project now treats inflated machine prose in commit messages as grounds for rejection.

### What else the document covers

The remaining material is about process. It describes the patch life cycle: posting to the list, replying to review on the list before sending a new version, rewriting history so each version is a clean replacement, sending, as a rough default, no more than one new version a day, and the path from the `seen` branch through `next` to `master`. It explains how to choose a base branch, asks for tests for every fix and feature, asks for documentation updates, and states the project's preference for gradually moving to US spelling. It sets out the Developer's Certificate of Origin and the `Signed-off-by:` trailer, the rules for identity in sign-offs, and a list of credit trailers with their meanings and ordering. It covers patch generation with `git format-patch` and `git send-email`, subject tags such as RFC and version numbers, inline patches over attachments, the advice against PGP-signing patches, a step-by-step procedure for resolving conflicts with other topics, subsystems with separate maintainers, GitHub CI, and hints for specific mail clients.

### Strengths and limits

The writing rules are short and concrete, and each comes with a reason grounded in how the history gets used. They were written for a large project with a public review process, so some of them, such as the area prefix and the lowercase word after it, are local conventions that the guide itself tells contributors to verify against the existing log. The 50-character figure is presented as a soft limit with no evidence for the exact number; it reflects how Git and common tools display titles. The document says nothing about wrapping the body at a fixed width.

## What the skill applies

**Section 8, bullet 1 (commit messages).** Most of this bullet traces to SubmittingPatches and the DISCUSSION section of the git-commit manual page. It moved unchanged from writing-style's section 9 when technical-writing became its own skill.

- The imperative subject line comes from the guide's imperative-mood rule. The skill states it for the subject; the guide applies it to the whole description of the change, and the skill's closing instruction not to narrate the diff carries the same intent into the body.
- The length of about 50 characters comes from the soft limit in SubmittingPatches and the recommendation in the git-commit manual.
- No closing period on the subject comes directly from the guide.
- Following the repository's conventions for prefixes and capitalization generalizes two Git rules, the area prefix and the lowercase word after it. The guide's advice to check the existing log for the current convention is what let the skill state the rule as deference to the repository, so it does not impose Git's prefix style on repositories that capitalize subjects and use no prefix.
- The blank line between subject and body comes from the git-commit manual, which explains that Git treats the text before the first blank line as the commit's title.
- Explaining the problem and why the change solves it condenses the guide's list of what a message must contain: what is wrong without the change, and why the result is better.
- Describing the old behavior in the present tense comes from the guide's rule for the status quo, which describes the code as it stands without the change.
- Summarizing outside discussion instead of only linking to it comes from the guide's request that a message be understandable without outside resources.
- Not narrating the diff rests on the guide's argument that the code shows how and only the message can say why.

The clauses on what users will notice and on numbers and costs for a claimed improvement come from the Linux kernel's guide (linux-kernel-submitting-patches.md). SubmittingPatches has no such rules, though it is consistent with them.

**Section 8, bullet 2 (pull requests).** Asking a pull request description to say what changed, why, and how it was tested adapts the guide's aims for a cover letter (make the series understandable without prior context, and show that its value outweighs its risk) and its requirement of tests for every fix and feature to the pull-request workflow most repositories use.

**Section 7, bullet 1 (code comments).** The rule that comments state the purpose, the reason for a choice, a constraint, or a trap draws on the guide's division of labor between comments and the log message. The guide accepts that comments can explain how code works and reserves the reason for a change for the message. Its example of a change applied to directories but not files, where the author records why the work stops short, is the model for the skill's mention of constraints and reasons.

**Section 3, bullet 1 (timeless documentation).** The bullet says that commit messages record a change and can say what changed. The guide's remark that currently adds nothing to a problem statement, because the statement always describes the code without the change, agrees with the skill and with revision item 1, which removes currently from text that describes how something works. The guide's present-tense rule shows how a message can describe a change without dating itself.

**Example.** The skill's one example is a commit message, and it follows the guide's shape. The subject is an imperative with no closing period and well under 50 characters. The body opens with the problem in the present tense (sessions stay in the cache until a restart, so a revoked token keeps working), then says why the change fixes it. The figures and the cost in the last sentence come from the kernel's rule on quantified trade-offs.

**The case for persisted text.** Writing-style's scope statement, which technical-writing inherits, says persisted text matters most because people and later agents read it many times. SubmittingPatches makes the same argument for log messages, and its AI section records that Git's maintainers reject contributions that read as overly formal or bloated, which supports both skills' premise that readers recognize and distrust inflated machine prose.

## What the skill omits

The skill leaves out the mailing-list workflow: recipients, reply etiquette, pacing of new versions, patch generation with format-patch and send-email, subject tags such as RFC and version numbers, inline patches, PGP, conflict resolution against the integration branches, and mail-client hints. They are procedures for one project's process and say nothing about prose.

It leaves out the Developer's Certificate of Origin, sign-off, and credit trailers. They are legal and attribution conventions set per repository, and AGENTS.md has its own rule for this repository: commits go in under the owner's name, with no trailers crediting an AI model or session.

It leaves out the testing and documentation requirements, except that a pull-request description reports how the change was tested, and the advice to split a commit whose description grows long. Commit granularity is an engineering decision more than a writing one.

It leaves out Git's format for citing another commit (abbreviated hash, subject, and date). The kernel uses a different format, and the skill's deference to repository conventions covers the choice.

It leaves out the request to record alternatives considered and the outcome of review debates. The body rule asks for the problem and the reasoning, and the instruction to summarize outside discussion covers the cases where a review thread exists.

It leaves out the AI policy's advice on how contributors should use AI tools, which addresses people and their workflow, not the text an agent writes.

## Conflicts with other references

Of the references in this folder, only the kernel's guide addresses commit messages. Red Hat's warning that revision-history entries drift into telegraphic style, recorded in red-hat-technical-writing-style-guide.md, agrees with Git's preference for full explanations, and nothing in the IBM, Google, Microsoft, Apple, or Oxford files bears on commits. The real tensions are with the kernel guide and with AGENTS.md.

**Subject-line length (linux-kernel-submitting-patches.md).** Git gives about 50 characters as a soft limit; the kernel sets a ceiling of 70 to 75 characters and asks the summary to say why the change is needed as well as what it does. The skill follows Git's figure, which also satisfies the kernel's ceiling, and leaves the why to the body. AGENTS.md asks for about 50 characters too.

**Capitalization after the prefix (linux-kernel-submitting-patches.md).** Git lowercases the first word after an area prefix unless it would be capitalized anyway. The kernel has no rule, and its examples go both ways. AGENTS.md uses no prefix and capitalizes the first word, as does the skill's example. The skill defers to the repository's existing convention, which is Git's own advice for prefixes.

**Links and summaries (linux-kernel-submitting-patches.md).** Git asks for a summary of a discussion in place of an archive link. The kernel asks for Link tags and also for an explanation that makes sense without them. The skill takes the common ground: summarize, and don't rely on a link alone.

**Body wrapping (linux-kernel-submitting-patches.md).** SubmittingPatches sets no wrap width. The kernel wraps at 75 columns and AGENTS.md asks for about 72. The skill states no width, so in this repository the AGENTS.md figure applies.

**AGENTS.md.** The commit section of AGENTS.md agrees with the guide apart from the prefix: an imperative subject with no closing period, a blank line, then plain sentences on the problem and the reasoning, with nothing the diff already shows. It also asks the body to name the files or behavior that changed, which fits the guide and the skill. One pointer in AGENTS.md is now out of date. It says commit messages follow section 9 of the writing-style skill, and that section now only directs the writer to the technical-writing skill, whose section 8 holds the rules. The content AGENTS.md summarizes still matches.

**House choices.** The AsciiDoc source of SubmittingPatches uses an em dash in one sentence. The dash is the document's typography and carries no rule; the owner's ban on em dashes applies to an agent's own text, commit messages included.

## Review notes

The title, metadata, and "What the source covers" section come from the copy of this file in chat/skills/writing-style/references, which was reviewed on 2026-11-04 from a local copy of Documentation/SubmittingPatches on the master branch and the DISCUSSION section of git-commit.adoc. In this copy I added the serial comma where it was missing and repaired one non-parallel sentence in the closing assessment (short, concrete, and each comes with a reason). The review date is now 2026-11-05.

For the mapping, I checked the guide against technical-writing 1.0.0, whose commit and pull-request bullets are the same as writing-style's former section 9 bullets, and against the AGENTS.md commit section as it stood on 2026-11-05. I did not re-read SubmittingPatches in full; the local copy at scratchpad/tech/git-submittingpatches.txt was available, and I relied on the writing-style copy's account of it. I read the applies, omits, and conflicts sections of the IBM, Red Hat, and Oxford files in this folder, which were being written at the same time, and found nothing in them about commit messages beyond Red Hat's remark on revision histories. I could not confirm the commit of SubmittingPatches that the original review read or when its AI section was added.
