# Linux kernel: Submitting patches

- **Source:** The Linux kernel development community. *Submitting patches: the essential guide to getting your code into the kernel.* `Documentation/process/submitting-patches.rst` in the Linux kernel source tree. Maintained collectively in the kernel repository, with no single author named; its reading list credits earlier guides by Andrew Morton, Jeff Garzik, and Greg Kroah-Hartman and a mail from Linus Torvalds on the canonical patch format.
- **Type:** official documentation
- **Where to find it:** https://www.kernel.org/doc/html/latest/process/submitting-patches.html (the rendered page carried the version label 7.4.0-rc1 on the day of review)
- **Last reviewed:** 2026-11-05
- **Used by:** `chat/skills/technical-writing/SKILL.md`

## What the source covers

The document tells a contributor how to get a change accepted into the Linux kernel. It is a working manual for a specific process. A change travels as an email in plain text to a mailing list, and maintainers and reviewers argue about it in reply. If the change is accepted, its description becomes a permanent entry in the `git` history of one of the largest and longest-lived codebases in existence. Most of the document is about that process. A smaller part, concentrated in two sections, is about how to write the description of a change, and that part is what the writing-style and technical-writing skills draw on.

The document says at the start that it is a terse collection of suggestions and points elsewhere for the full development process and a pre-submission checklist. It assumes the contributor uses `git`.

### Describing a change

The section on describing changes is the core of the document for the skill's purposes. It gives a sequence of instructions, and each one has an operational reason.

The first is to state the problem. Any change, from a one-line fix to a large feature, exists because something was wrong or missing, and the description has to persuade the reviewer that the problem is real and worth fixing. The document frames this in terms of the reviewer's attention: the opening has to give the reviewer a reason to keep reading past the first paragraph.

The second is to describe the impact on users. Crashes and lockups make their own case, but many bugs are quieter. The document asks the writer to explain what a problem could do to users even when it was found by reading code rather than by a failure in the field. The reason is specific to the kernel's distribution model. Most Linux systems run kernels from stable or vendor trees that pick individual patches from upstream, and the people choosing those patches decide from the description. So the description should carry whatever helps that choice: the conditions that trigger the problem, log excerpts, crash descriptions, regressions in performance or latency.

The third is to quantify optimizations and tradeoffs. A claim of better performance, lower memory use, a smaller stack footprint, or a smaller binary has to come with numbers. The writer also has to describe the costs that are not obvious. The document's reasoning is that optimizations usually trade one resource against another (processor time, memory, readability) or, for heuristics, favor one workload over another, and the reviewer can only weigh benefit against cost if both are on the page.

Only after the problem is established does the document ask for the technical description of the fix, written in plain words so the reviewer can check that the code does what the author intends. It asks for one problem per patch and treats a long description as a sign that the patch should be split.

A resubmitted patch has to carry its full description and justification again. Saying that this is the next version is not enough, because maintainers will not dig through earlier versions or linked pages, and some reviewers never saw the earlier versions. The document's word for the goal is self-contained.

The document asks for the imperative mood, phrased as if the writer were giving the codebase an order, and contrasts this with descriptions that start from the patch as subject or from the author in the first person. Git's own SubmittingPatches uses nearly the same wording and example, so the two projects share this convention directly.

On referring to other commits, the document asks for the one-line summary of the commit along with its hash, so that a reader knows what the commit is about without looking it up. It asks for at least 12 characters of the hash. The reason given is collisions: the kernel repository holds so many objects that a short hash which is unique today may not be unique in five years.

On outside material, the document asks for `Link:` tags pointing to earlier discussion or background, preferably through the lore.kernel.org archive, and asks the writer to check that each link works. It then adds that the explanation should be understandable without those links: the writer should summarize the relevant points of the discussion in the description itself. For bug fixes it asks for a `Closes:` tag pointing to a public report. It also asks for a `Fixes:` tag naming the commit that introduced the bug, written on one line even when it runs past the wrap limit so that scripts can parse it.

### The canonical patch format

The second section that bears on writing describes the shape of a patch email. The subject carries a bracketed prefix (the word PATCH, an optional version or RFC marker, and a zero-padded sequence number in a series), then a subsystem name, a colon, and a summary phrase. The body holds an optional author line, the explanation wrapped at 75 columns, a blank line, the sign-off lines, a `---` marker, comments meant only for reviewers, and then the diff.

The guidance on the summary phrase is the most specific writing advice in the document. The subsystem prefix identifies the area of the kernel being changed. The summary must not be a file name and must not be repeated across the patches of a series. The document explains why the phrase matters. It becomes a global identifier for the patch and propagates into the `git` history. People search for it to find the discussion, and months later it may be the only text a developer sees while scanning thousands of patches in a one-line log view. For those reasons the summary must be at most 70 to 75 characters, and it must say both what the patch changes and why the patch might be needed. The document admits that being brief and descriptive at once is hard.

The explanation body is aimed at a competent reader who has long forgotten the discussion behind the patch. The document asks for symptoms of the failure, such as kernel log or `oops` messages, because people searching the history for a fix search on those symptoms. It wants enough detail that a reader weeks or years later can understand why the patch exists. For a fix to compile failures it asks for enough of the error output to make the patch findable, and no more. A subsection on backtraces makes the same point about trimming: full log output carries timestamps, module lists, and register dumps that distract, and a useful backtrace keeps only the lines that show the call chain.

The commentary section separates what is permanent from what is temporary. Anything relevant only to the current review, such as a diffstat or a list of changes between versions, goes below the `---` marker, where patch tools strip it automatically when the patch is applied. Version history placed above the marker has to be removed by hand.

### What else the document covers

Most of the document deals with matters the skill does not touch. It explains how to obtain the right source tree, how to split a change into logical patches that each build and run, and how to run the style checker (and when its warnings can be ignored). It also covers how to choose maintainers and mailing lists, why patches go inline in plain text rather than as attachments, how to answer reviewers, and why replies should be interleaved and trimmed. Further topics are how long to wait before resending and how to supply base-tree information for automated testing. A long middle stretch defines the attribution and certification tags: the Developer's Certificate of Origin behind `Signed-off-by:`, and the meanings of `Acked-by:`, `Co-developed-by:`, `Reported-by:`, `Tested-by:`, `Reviewed-by:`, and `Suggested-by:`, including the rule that most of them need the named person's permission. Every one of these tags names a person with a known identity; the document does not accept anonymous contributions.

### Strengths and limits as evidence

The document records practice. Nothing in it is measured, and its rules are not tested against alternatives. Its authority comes from the scale and age of the project it governs. The conventions were refined over decades by maintainers who read large volumes of patches, and many of them are now built into tools (`git format-patch`, the style checker, bots that act on `Closes:` tags).

Its strength is that almost every rule carries a concrete reason tied to a reader's task. The reviewer needs to weigh cost against benefit, so tradeoffs must be stated. The stable-tree maintainer chooses patches from descriptions, so user impact must be explicit. The developer scanning a one-line log sees only the summary, so the summary must carry both the change and its reason. These reasons transfer well to any project with a long-lived history, which is why the skill can use them outside the kernel.

The limits are also clear. Some numbers are hedged or loose: the summary limit is given as a range of 70 to 75 characters rather than a single figure, and the document never says why. Some rules follow from the email workflow and lose their force elsewhere; the 75-column wrap and the bracketed subject prefix exist because patches travel as email. The document's own examples are not consistent on capitalization after the subsystem prefix: some summaries start with a lowercase verb and some with a capital, and no rule settles it. Its own prose leans on "should" for both requirements and recommendations, which makes some rules hard to rank. And it was written for experienced C developers working on one project, so its advice about evidence (logs, backtraces, benchmark figures) assumes a technical reader throughout.

## What the skill applies

The kernel's guide shaped section 8 of the technical-writing skill, mostly its commit-message bullet, and it supports section 3's rule on claims and the first item of the revision pass. These rules moved unchanged from writing-style's section 9 when technical-writing became its own skill.

**Section 8, bullet 1: imperative subject line.** The kernel and Git both ask for the imperative mood, with the same reasoning: the description reads as an order given to the codebase. The kernel applies the mood to the whole description of the change. The skill states it for the subject line, and the body's account of the old behavior is in the present tense, a rule that comes from Git.

**Section 8, bullet 1: subject length.** The skill's figure of about 50 characters comes from Git's soft limit. The kernel's ceiling of 70 to 75 characters is well above it, so a subject that meets the skill also meets the kernel. The kernel's reasons for keeping the summary short (it is what people scan in one-line logs and search for months later) support the skill's choice of a short target. The skill does not adopt the kernel's requirement that the subject also say why the change is needed.

**Section 8, bullet 1: prefixes and capitalization.** The instruction to follow the repository's conventions covers the kernel's subsystem prefix and Git's area prefix without choosing between them. It also absorbs the inconsistency in the kernel's own examples about capitalizing the word after the prefix.

**Section 8, bullet 1: the problem first.** The body explains the problem and why the change solves it. The order matches the kernel's sequence of establishing the problem before describing the fix. The kernel's argument that the reviewer must be persuaded in the first paragraph is writing-style section 1's rule to open with the answer, applied to a commit body.

**Section 8, bullet 1: what users will notice.** The clause comes from the kernel's instruction to describe the impact on users, including for problems found by reading code. The kernel's own reason concerns its stable and vendor trees, whose maintainers pick fixes by reading descriptions. The skill takes the general point that a reader of the history needs to know what a problem does to the people who run the code, and a description of what is wrong inside the code does not say that.

**Section 8, bullet 1: numbers and costs for a claimed improvement.** Both halves come from the kernel's instruction to quantify optimizations and tradeoffs. Improvements in speed, memory, stack use, or binary size need figures, and the costs that are not obvious need describing so the reviewer can weigh one against the other. Writing-style section 4 governs how the numbers are presented, with a comparison for each figure and absolute values beside a relative change.

**Section 8, bullet 1: summarizing outside discussion.** The kernel says this twice, in its advice to make the explanation understandable without linked pages and in its rule that a resubmitted patch carry its full description. Git says the same. The kernel's reasoning (the reader may never have seen the earlier thread, and links can die) is why the skill treats a bare link as incomplete.

**Section 8, bullet 1: don't narrate the diff.** The kernel's emphasis on why a patch exists supports the closing instruction, as does its request for a plain account of intent that the reviewer can check the code against. The kernel never says outright not to repeat the diff. Git and AGENTS.md state the rule more directly.

**Section 3, bullet 2 (claim only what you can show).** The bullet asks for the source or the measurement behind any performance claim. The measurement half has its clearest statement in the kernel's quantification rule; the source half and the rule's general form come from the Google guide's excessive-claims page.

**Section 3, bullet 1 (timeless documentation).** The kernel's split between permanent changelog text and temporary review notes below the patch separator is a narrower form of the distinction this bullet draws: change records may say what changed, and documentation should not go stale.

**Section 9, item 1 (overclaiming and dating).** The item catches claims of best, fastest, or guaranteed that are not literally true. It rests mainly on the Google guide, with the kernel's demand for numbers behind optimizations as support from the commit-message side.

**Example.** The skill's commit-message example applies the kernel's rules directly. Its body names the user-visible problem (a revoked token keeps working for days) and gives a figure for the gain (cache memory on staging drops by about 30 percent). It also states the cost (users idle for a day have to sign in again). The last sentence is the kernel's tradeoff rule at work.

## What the skill omits

**The subject must state why.** The kernel asks the summary to carry both the change and its reason. The skill keeps the subject to about 50 characters and leaves the reason to the body, which is closer to Git's practice; a 50-character subject rarely has room for both.

**Commit references.** The kernel's format for citing another commit (at least 12 characters of the hash plus the one-line summary) is left to the repository's conventions, since Git uses a different format.

**Symptoms, logs, and backtraces.** The advice to include failure symptoms for later searchers and to trim backtraces to the relevant call chain suits debugging-heavy projects. Writing-style's rule to be specific covers the principle, and the skill does not prescribe log content. The kernel's advice to quote enough of a compiler error to make the patch findable is related to section 4's rule to copy output from its source, but it did not shape that rule.

**Process and format machinery.** The email workflow, the bracketed subject prefix, the 75-column wrap, the patch separator, the tag system, and the Developer's Certificate of Origin are process and format machinery. So are recipients, plain text mail, interleaved replies, resending, base-tree information, and tooling. They govern how kernel patches move and are credited, and they are not rules about prose. Most do not apply to repositories that use pull requests.

**Splitting changes.** One logical change per patch, and a series that builds at every step, concern how code changes are structured. The kernel's remark that a long description signals a patch to split is a sound check, but the skill concerns the text.

## Conflicts with other references

Of the references in this folder, only Git's SubmittingPatches (git-submitting-patches.md) addresses commit messages, so the real tensions are with it and with AGENTS.md. The IBM, Red Hat, Google, Microsoft, Apple, and Oxford files do not discuss commit messages, apart from Red Hat's warning against telegraphic revision histories, which the kernel's preference for full explanations matches.

**Subject length (git-submitting-patches.md).** The kernel sets a ceiling of 70 to 75 characters; Git sets a soft limit of 50. The skill asks for about 50 with no outer bound, and AGENTS.md repeats that figure. A subject of about 50 characters satisfies all three, and the kernel's own reasoning favors brevity, so the skill contradicts only the kernel's number.

**Capitalization after the prefix (git-submitting-patches.md).** Git lowercases the word after the prefix; the kernel has no rule. The skill defers to the repository. AGENTS.md's example has no prefix and starts with a capital.

**Body wrap width.** The kernel wraps at 75 columns, AGENTS.md asks for about 72, and Git and the skill name no width. Inside this repository the AGENTS.md figure applies.

**Commit reference format (git-submitting-patches.md).** The kernel cites at least 12 hash characters and the quoted summary; Git cites an abbreviated hash, the subject, and the date. The skill leaves the format to the repository.

**Attribution trailers.** The kernel's tags credit people with known identities and reject anonymous contributions. AGENTS.md tells agents to commit as the repository owner without trailers crediting an AI model or session. The two govern different histories and do not bear on prose.

**Should (ibm-style-guide.md, red-hat-technical-writing-style-guide.md, google-developer-style-guide.md).** The kernel's own prose uses should for both requirements and recommendations. The usage is the document's style, not a rule it sets. A writer who copied its phrasing into documentation would trip the skill's section 2 modal-verb rule and revision item 2, which IBM, Red Hat, and Google support.

**AGENTS.md pointer.** AGENTS.md says commit messages follow section 9 of the writing-style skill. The section now points to the technical-writing skill, whose section 8 holds the commit rules. The summary in AGENTS.md still matches those rules.

## Review notes

The title, metadata, and "What the source covers" section come from the copy of this file in `chat/skills/writing-style/references`. That copy was reviewed on 2026-11-04 from a local copy of `Documentation/process/submitting-patches.rst`, checked against the rendered page at kernel.org (version label 7.4.0-rc1). In this copy I added the serial comma where it was missing and changed one sentence that named only the writing-style skill as the user of the commit-message material. The review date is now 2026-11-05.

For the mapping, I checked the guide against technical-writing 1.0.0, whose commit bullet is the same as writing-style's former section 9 bullet. I also checked it against the AGENTS.md commit section as it stood on 2026-11-05. I did not re-read the kernel document in full. The local copy at `scratchpad/tech/kernel-submitting-patches.rst` was available, and I confirmed only the Link and Closes tag guidance in it. I read the applies, omits, and conflicts sections of the IBM, Red Hat, and Oxford files in this folder, which were being written alongside this one, and did not check their accounts of their sources.
