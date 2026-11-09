# Linux kernel: Submitting patches

- **Source:** The Linux kernel development community. *Submitting patches: the essential guide to getting your code into the kernel.* `Documentation/process/submitting-patches.rst` in the Linux kernel source tree. Maintained collectively in the kernel repository, with no single author named; its reading list credits earlier guides by Andrew Morton, Jeff Garzik, and Greg Kroah-Hartman and a mail from Linus Torvalds on the canonical patch format.
- **Type:** official documentation
- **Where to find it:** https://www.kernel.org/doc/html/latest/process/submitting-patches.html (the rendered page showed the version label 7.4.0-rc1 on the day of review)
- **Last reviewed:** 2026-11-05
- **Used by:** `chat/skills/writing-style/SKILL.md`

## What the source covers

The document tells a contributor how to get a change accepted into the Linux kernel. It is a working manual for a specific process. A change is sent as a plain text email to a mailing list, and maintainers and reviewers argue about it in reply. If it is accepted, its description becomes a permanent entry in the `git` history of one of the largest and longest-lived codebases in existence. Most of the document is about that process. A smaller part, concentrated in two sections, is about how to write the description of a change, and that part is what the writing-style and technical-writing skills draw on.

The document says at the start that it is a terse collection of suggestions and points elsewhere for the full development process and a pre-submission checklist. It assumes the contributor uses `git`.

### Describing a change

The section on describing changes is the core of the document for the skill's purposes. It gives a sequence of instructions, and each one has an operational reason.

The first is to state the problem. Any change, from a one-line fix to a large feature, exists because something was wrong or missing, and the description has to persuade the reviewer that the problem is real and worth fixing. The document frames this in terms of the reviewer's attention: the opening has to give the reviewer a reason to keep reading past the first paragraph.

The second is to describe the impact on users. Crashes and lockups make their own case. Many bugs are quieter, and the document asks the writer to explain what a problem could do to users even when it was found by reading code and not by a failure in the field. The reason is specific to the kernel's distribution model. Most Linux systems run kernels from stable or vendor trees that pick individual patches from upstream, and the people choosing those patches decide from the description. So the description should include whatever helps that choice: the conditions that trigger the problem, log excerpts, crash descriptions, regressions in performance or latency.

The third is to quantify optimizations and trade-offs. A claim of better performance, lower memory use, a smaller stack footprint, or a smaller binary needs numbers. The writer also has to describe the costs that are not obvious. The document's reasoning is that optimizations usually trade one resource against another (processor time, memory, readability) or, for heuristics, favor one workload over another. The reviewer can weigh benefit against cost only if both are on the page.

Only after the problem is established does the document ask for the technical description of the fix, written in plain language so the reviewer can check that the code does what the author intends. It asks for one problem per patch and treats a long description as a sign that the patch should be split.

A resubmitted patch has to include its full description and justification again. Saying that this is the next version is not enough, because maintainers will not dig through earlier versions or linked pages, and some reviewers never saw the earlier versions. The document's word for the goal is self-contained.

The document asks for the imperative mood, phrased as if the writer were giving the codebase an order. It contrasts this with descriptions that start from the patch as subject or from the author in the first person. Git's own SubmittingPatches uses nearly the same wording and example, so the two projects share this convention directly.

On referring to other commits, the document asks for the commit's one-line summary along with its hash, so that a reader knows what the commit is about without looking it up. It also asks for at least 12 characters of the hash. The reason given is collisions: the kernel repository holds so many objects that a short hash which is unique today may not be unique in five years.

On outside material, the document asks for `Link:` tags pointing to earlier discussion or background, preferably through the lore.kernel.org archive, and asks the writer to check that each link works. It then adds that the explanation should be understandable without those links, so the writer should summarize the relevant points of the discussion in the description itself. For bug fixes it asks for a `Closes:` tag pointing to a public report. It also asks for a `Fixes:` tag naming the commit that introduced the bug, written on one line even when it runs past the wrap limit so that scripts can parse it.

### The canonical patch format

The second section that bears on writing describes the shape of a patch email. The subject has a bracketed prefix (the word PATCH, an optional version or RFC marker, and a zero-padded sequence number in a series), then a subsystem name, a colon, and a summary phrase. The body has an optional author line, the explanation wrapped at 75 columns, a blank line, the sign-off lines, a `---` marker, comments meant only for reviewers, and then the diff.

The guidance on the summary phrase is the most specific writing advice in the document. The subsystem prefix identifies the area of the kernel being changed. The summary must not be a file name and must not be repeated across the patches of a series. The document explains why the phrase matters. It becomes a global identifier for the patch and propagates into the `git` history. People search for it to find the discussion, and months later it may be the only text a developer sees while scanning thousands of patches in a one-line log view. For those reasons the summary must be at most 70 to 75 characters. It must say both what the patch changes and why the patch might be needed. The document admits that being brief and descriptive at once is hard.

The explanation body is aimed at a competent reader who has long forgotten the discussion behind the patch. The document asks for symptoms of the failure, such as kernel log or kernel oops messages, because people searching the history for a fix search on those symptoms. It wants enough detail that a reader weeks or years later can understand why the patch exists. For a fix to compile failures it asks for enough of the error output to make the patch findable, and no more. A subsection on backtraces makes the same point about trimming: full log output contains timestamps, module lists, and register dumps that distract, and a useful backtrace keeps only the lines that show the call chain.

The commentary section separates what is permanent from what is temporary. Anything relevant only to the current review, such as a diffstat or a list of changes between versions, goes below the `---` marker, where patch tools strip it automatically when the patch is applied. Version history placed above the marker has to be removed by hand.

### What else the document covers

Most of the document deals with matters the skill does not touch. It explains how to obtain the right source tree and how to split a change into logical patches that each build and run. It covers how to run the style checker (and when its warnings can be ignored) and how to choose maintainers and mailing lists. It says why patches go inline in plain text and not as attachments, how to answer reviewers, why replies should be interleaved and trimmed, how long to wait before resending, and how to supply base-tree information for automated testing. A long middle stretch defines the attribution and certification tags. These are the Developer's Certificate of Origin behind `Signed-off-by:` and the meanings of `Acked-by:`, `Co-developed-by:`, `Reported-by:`, `Tested-by:`, `Reviewed-by:`, and `Suggested-by:`, including the rule that most of them need the named person's permission. Every one of these tags names a person with a known identity; the document does not accept anonymous contributions.

### Strengths and limits as evidence

The document records practice. Nothing in it is measured, and its rules are not tested against alternatives. Its authority comes from the scale and age of the project it governs. The conventions were refined over decades by maintainers who read large volumes of patches, and many of them are now built into tools (`git format-patch`, the style checker, bots that act on `Closes:` tags).

Its strength is that almost every rule has a concrete reason tied to a reader's task. The reviewer needs to weigh cost against benefit, so trade-offs must be stated. The stable-tree maintainer chooses patches from descriptions, so user impact must be explicit. The developer scanning a one-line log sees only the summary, so the summary must state both the change and its reason. These reasons transfer well to any project with a long-lived history, which is why the skill can use them outside the kernel.

The limits are also clear. Some numbers are hedged or loose: the summary limit is given as a range of 70 to 75 characters rather than a single figure, and the document never says why. Some rules follow from the email workflow and lose their force elsewhere; the 75-column wrap and the bracketed subject prefix exist because patches travel as email. The document's own examples are not consistent on capitalization after the subsystem prefix: some summaries start with a lowercase verb and some with a capital, and no rule settles it. Its own prose leans on "should" for both requirements and recommendations, which makes some rules hard to rank. And it was written for experienced C developers working on one project, so its advice about evidence (logs, backtraces, benchmark figures) assumes a technical reader throughout.

## What the skill applies

The source shaped section 9 of the skill before version 2.0.0, mostly its commit-message bullet, with smaller contributions to sections 1, 2, and 4. The commit-message bullet, the bullet on claiming only what you can show, the timeless-documentation bullet, and the old revision item 13 moved to the technical-writing skill (its sections 3 and 8 and its first revision item). That skill's copy of this file maps the kernel to them clause by clause. The clauses cover the imperative subject, the subject length, prefixes, the problem first, what users will notice, numbers and costs for a claimed improvement, summaries of outside discussion, and the instruction not to narrate the diff. What remains in writing-style is the kernel's support for sections 1, 2, and 4 and for the scope statement.

**Section 1: open with the answer.** The kernel's argument that a description must convince the reviewer in its first paragraph is the same argument section 1 makes for opening with the answer. Technical-writing applies the point to a commit body, where the problem comes before the fix.

**Section 2: plain and specific words.** The kernel's request for the technical description in plain language, and its preference for concrete symptoms, logs, and conditions over general claims, match section 2's rules on plain and specific words. The skill did not need the kernel for those rules, but the kernel is the clearest statement of them for commit messages.

**Section 4: numbers and evidence.** The obligation to give numbers and costs for a claimed improvement now belongs to technical-writing, and section 4 governs how the numbers are presented: a figure needs a comparison, and a relative change needs the absolute values beside it. The kernel does not spell out either point, but a claim of improvement is a before-and-after comparison by nature, and a reviewer weighing costs needs the absolute figures. Read together, the kernel supplies the obligation and section 4 supplies the form. Section 4's rule against presenting lopsided evidence as balanced states the same honesty in general terms, since a commit message that reports only the gain gives a one-sided case.

**Self-contained text that outlasts the discussion.** The kernel's target reader for the body, a competent person who has forgotten the discussion, fits the skill's account of persisted text in its scope section: people and later agents read it many times, long after the conversation.

## What the skill omits

The skill no longer holds the kernel's rules for commit messages. Technical-writing does. Its copy of this file records what that skill leaves out. That covers the requirement that the subject state why, the commit-reference format, the advice on symptoms, logs, and backtraces, the process and format machinery of the email workflow, and the rules on splitting changes.

## Conflicts with other references

This file used to record conflicts over subject length, capitalization after the prefix, body wrap width, the commit-reference format, tense for the old behavior, and attribution trailers. The rules involved now belong to technical-writing's section 8 or to AGENTS.md. That skill's copies of this file and of `git-submitting-patches.md` cover them. AGENTS.md now says commit messages follow section 8 of the technical-writing skill.

**Fit with sections 1 to 8.** Section 9 says that writing-style wins where technical-writing disagrees with it. Nothing in the kernel's description advice conflicts with sections 1 to 8. Its most specific rules (numbers and costs for claims, user impact, the problem first, and self-contained text) reinforce them.

**Should.** The kernel's own prose leans on `should` for both requirements and recommendations. The usage is the document's style and not a rule it sets. The modal-verb rule it would trip now belongs to technical-writing, whose copy of this file notes the point.

## Review notes

I read the full local copy of `Documentation/process/submitting-patches.rst` from start to finish. I checked the rendered page at kernel.org on November 4, 2026. Its section headings matched the local copy, and it confirmed the 70 to 75 character summary limit, the 75-column wrap, the 12-character hash minimum, and the imperative-mood rule. The page showed the documentation version label 7.4.0-rc1 and no revision date, so I could not tie the local copy to a specific kernel commit. I compared the relevant rules against the local copy of git's `Documentation/SubmittingPatches`, reading its sections on the commit message, the subject and its prefix, the present tense, the imperative mood, commit references, and trailers. I did not read the companion `git-submitting-patches.md` reference, which I was writing in parallel. Any difference in how the two files describe git's rules should be checked against git's text. I did not follow the kernel document's external reading list (the guides by Morton, Garzik, and Kroah-Hartman, Torvalds's mail, and Kleen's paper). I also skipped the linked kernel documents on the development process, coding style, and stable rules. I compared the SKILL.md and AGENTS.md text first as it stood on disk on November 4, 2026. By then the commit-message rules in both had changed to a subject of about 50 characters, and the section 9 commit bullet had gained clauses on what users will notice and on the costs of a claimed improvement. That version also moved the dated and overclaiming words out of section 9 and into revision item 13.

On November 5, 2026, after the skill's reorganization into version 2.0.0, I updated this file. I replaced the section 9 mappings and the commit-message conflicts with pointers to technical-writing's copy of this file. I kept the links to sections 1, 2, and 4 and to the scope statement, and I brought the prose into the skill's style. I took the "What the source covers" section from technical-writing's copy, which already had the serial comma, and set the review date to November 5, 2026. I did not re-read the kernel document for this update.
