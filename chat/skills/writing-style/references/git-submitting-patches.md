# Git's guidelines for submitting patches

- **Source:** The Git project, *Submitting Patches* (Documentation/SubmittingPatches in the git.git source tree, master branch), maintained by the Git maintainer and contributors. Read together with the DISCUSSION section of the *git-commit* manual page (Documentation/git-commit.adoc, master branch). Living documents with no fixed edition; the copies reviewed were current on the master branch in November 2026.
- **Type:** official documentation
- **Where to find it:** https://github.com/git/git/blob/master/Documentation/SubmittingPatches and https://github.com/git/git/blob/master/Documentation/git-commit.adoc (rendered as the git-commit manual page at https://git-scm.com/docs/git-commit)
- **Last reviewed:** 2026-11-05
- **Used by:** chat/skills/writing-style/SKILL.md

## What the source covers

SubmittingPatches is the contributor guide for the Git project itself. It tells a developer how to prepare a change, describe it, certify it, and send it to the project's mailing list, where every change is reviewed in public before the maintainer merges it. About a third of the document concerns writing: how to split work into commits, what a commit message must explain, how its first line is formed, how to refer to other commits, and what goes in a cover letter. The rest covers the review cycle, branch choice, tests, sign-off and trailers, the mechanics of mailing patches, conflict handling, subsystem maintainers, the project's CI setup, and workarounds for particular mail clients.

It is a set of house rules written by practitioners, with no study or data behind it. Its authority comes from long use: the conventions it describes are the ones Git's own history follows, and many other projects copied them, partly because Git's tools display the first line of a message as the commit's title. The reasons it gives are practical and stated from the reader's side: a reviewer has to judge the change, and a future maintainer has to understand it years later.

### Why the log message matters

The section on describing changes opens by ranking the log message as equal in importance to the code. Its argument is that well-commented code can explain how it works, while only the message can say why it exists. It gives two reasons a later developer needs the why. First, the code may not do what the author meant, and a written statement of intent lets someone fix it to match; the document adds that authors often find their own bugs while writing that statement. Second, the code may deliberately cover only part of a problem, and recording why the author stopped where they did lets a successor decide whether to extend the work or leave it alone. The worked example is a change applied to directories but not to files: stating the property of directories that justified it lets a reader check whether files share that property.

The section on describing changes is the strongest part of the document, because it ties a writing rule to a concrete maintenance task. It also explains why the guide wants reasons over narration: the diff already shows what changed.

### The subject line

The first line of the message is a short description. SubmittingPatches calls 50 characters a soft limit and points to the git-commit manual for the reason. That manual page's DISCUSSION section is brief: it says a message need not follow any format, recommends a single short summary line of at most 50 characters, a blank line, and then a fuller description, and explains that Git treats the text before the first blank line as the commit's title throughout its tools. Its example is git-format-patch, which turns the title into an email subject and the rest into the email body. The two documents differ slightly in strength: the manual page frames 50 as a ceiling within an optional recommendation, while SubmittingPatches frames it as a soft limit.

The first line ends with no period. By convention it usually starts with an area prefix: a file name or short identifier for the part of the code being changed, followed by a colon. The document's examples use a documentation prefix and a file-name prefix. When a contributor is unsure which prefix to use, the guide tells them to run `git log --no-merges` on the files they are changing and copy what earlier commits did. That instruction is the general principle behind the prefix rule: the history already in the repository sets the convention.

After the prefix, the first word is lowercase unless it would be capitalized anywhere in a sentence. The guide's counterexample is a ref name that Git always writes in capitals, which keeps its capitals after the prefix. The lowercase rule applies only to the word that follows the prefix.

### What the body must say

The body explains the problem the change solves: what is wrong with the code as it stands without the change. It then justifies the solution by saying why the result is better. Where relevant, it names alternatives the author considered and dropped, and it records how design or viability concerns raised in review were settled, so that the history shows why this approach won. An earlier section adds that the explanation should be detailed enough for a reader to judge whether the change is a good idea without reading the patch, and that a good description summarizes the point in the subject and covers the motivation, the approach, and, for a revised series, how it differs from the previous version. If the description grows long, the guide reads that as a sign the commit should be split.

### Tense, mood, and self-containment

Two grammatical rules follow. The description of the status quo is in the present tense: the message says what the code does today when given some input, and avoids describing it as something it used to do. The guide adds that a word such as currently is unnecessary, because by project convention the problem statement always describes the code without the change.

The change itself is described in the imperative mood, as an order given to the codebase. The guide contrasts this with two forms it rejects: a sentence whose subject is the patch, and a first-person past-tense account of what the author did. The rule covers the whole description of the change, the subject line included.

In the same paragraph, the guide asks for messages that can be understood without outside resources. In place of a link to a mailing-list archive, the author summarizes the points of the discussion that matter. The summary puts the reasoning into the history itself, where it survives if the archive moves or disappears.

### Referring to other commits

The guide lists three reasons to cite another commit: it introduced the bug being fixed, it introduced the feature being extended, or it conflicts with the current work in a trial merge. A citation gives the abbreviated hash followed by the cited commit's subject and date in parentheses, and the guide names a `git show` pretty format that produces this form. Pairing the hash with the subject means a reader learns what the cited commit did without looking it up, which is the self-containment principle applied to cross-references.

### Cover letters and the AI policy

The cover letter, which introduces a series of patches, has its own short list of aims: show that the value of the change outweighs its regression risk and maintenance cost, make the series understandable without prior context, keep later versions readable for people who missed earlier rounds, and summarize what changed since the previous round. An experimental section invites a two-to-five-line topic summary suitable for the release notes.

A section on artificial intelligence is directly relevant to this skill. It raises two problems with AI-generated contributions: it is unclear whether a contributor can make the certification of origin the project requires for substantial generated content, and generated code, messages, and documentation are often wrong in ways that persist even after correction. The project says it will reject submissions that look machine-generated, read as overly formal or bloated, look plausible but make no sense, or that the sender cannot explain. It recommends using AI to guide one's own work, to debug, and to check a draft against the project's style, over asking for a complete solution and pasting it in. The section is a maintainer's policy statement, with no measurement behind it, but it records that at least one major project now treats inflated machine prose in commit messages as grounds for rejection.

### What else the document covers

The remaining material is about process. It describes the patch life cycle: posting to the list, replying to review on the list before sending a new version, rewriting history so each version is a clean replacement, sending, as a rough default, no more than one new version a day, and the path from the `seen` branch through `next` to `master`. It explains how to choose a base branch, asks for tests for every fix and feature, asks for documentation updates, and states the project's preference for gradually moving to US spelling. It sets out the Developer's Certificate of Origin and the `Signed-off-by:` trailer, the rules for identity in sign-offs, and a list of credit trailers with their meanings and ordering. It covers patch generation with `git format-patch` and `git send-email`, subject tags such as RFC and version numbers, inline patches over attachments, the advice against PGP-signing patches, a step-by-step procedure for resolving conflicts with other topics, subsystems with separate maintainers, GitHub CI, and hints for specific mail clients.

### Strengths and limits

The writing rules are short and concrete, and each comes with a reason grounded in how the history gets used. They were written for a large project with a public review process, so some of them, such as the area prefix and the lowercase word after it, are local conventions that the guide itself tells contributors to verify against the existing log. The 50-character figure is presented as a soft limit with no evidence for the exact number; it reflects how Git and common tools display titles. The document says nothing about wrapping the body at a fixed width.

## What the skill applies

The guide shaped the commit-message, pull-request, and code-comment bullets of the skill's old section 9, and it supported that section's timeless-documentation bullet. Those bullets moved unchanged to the technical-writing skill in version 2.0.0 (its sections 3, 7, and 8), and that skill's copy of this file maps them clause by clause. Section 9 of writing-style now directs writers of commit messages and pull requests to technical-writing. Two connections remain in writing-style.

**Where this applies, and the reason for the skill.** The skill's scope statement says persisted text matters most because people and later agents read it many times, and it names commit messages and pull request descriptions among that text. SubmittingPatches makes the same argument for log messages. Its AI section supports the skill's opening premise that readers recognize and distrust inflated machine prose: Git's maintainers now name overly formal and bloated text as a reason to reject a contribution.

**Section 1, open with the answer.** The guide's request that a commit's first line summarize the point, and that the body let a reader judge the change without reading the patch, is section 1's rule applied to a commit. The guide is not a source for section 1, which comes from the newspaper model, but it agrees with it.

## What the skill omits

The skill no longer holds any of the guide's rules for commits, pull requests, or comments. Technical-writing does, and its copy of this file records what that skill leaves out: the mailing-list workflow, the Developer's Certificate of Origin and trailers, the testing and documentation requirements, the format for citing other commits, the record of alternatives considered, and the AI policy's advice on how contributors should use AI tools.

## Conflicts with other references

The conflicts this file used to record (subject-line length, capitalization after the prefix, links against summaries, and body wrapping) concern rules that now live in technical-writing's section 8. That skill's copy of this file and of linux-kernel-submitting-patches.md cover them. AGENTS.md now says commit messages follow section 8 of the technical-writing skill, so the repository's commit rules and the skill no longer point at writing-style for that material.

**Dashes.** The AsciiDoc source of SubmittingPatches itself uses an em dash in one sentence. The dash is the document's typography, and it carries no rule; the skill's ban on em dashes (revision item 1) is the owner's choice and applies to the agent's own text, commit messages included.

Apart from linux-kernel-submitting-patches.md, no reference in this folder addresses commit messages in detail, so the guide has no conflict with the news manuals or the research papers.

## Review notes

I read the full local copy of Documentation/SubmittingPatches from the master branch, saved in the session scratchpad, from start to end. I downloaded Documentation/git-commit.adoc from the master branch on GitHub and read its DISCUSSION section in full; the section is short and gives only the 50-character recommendation and the title convention. I read the commit-message parts of the kernel's Documentation/process/submitting-patches.rst (the sections on describing changes and on the canonical patch format) to check the conflicts, but did not review that document in full, since it has its own reference file. That file did not yet exist in the folder when I wrote this one, so I could not check my account of the kernel rules against it.

I could not confirm the exact commit of the SubmittingPatches copy I read or when the AI policy section was added; the citation describes it as current on master in November 2026. An earlier version of the skill set a 72-character outer limit on the subject. I found no support for it in the git or kernel documents, and the skill dropped it before the commit rules moved to technical-writing.

On November 5, 2026, after the skill's reorganization into version 2.0.0, I updated this file. I replaced the section 9 mappings and the commit-message conflicts with pointers to technical-writing's copy of this file, kept the two connections that still apply to writing-style, and brought the prose into the skill's style. I took the "What the source covers" section from technical-writing's copy, which already carried the serial comma, and set the review date to November 5, 2026. I did not re-read SubmittingPatches for this update.
