# Linux kernel: Submitting patches

- **Source:** The Linux kernel development community. *Submitting patches: the essential guide to getting your code into the kernel.* `Documentation/process/submitting-patches.rst` in the Linux kernel source tree. Maintained collectively in the kernel repository, with no single author named; its reading list credits earlier guides by Andrew Morton, Jeff Garzik and Greg Kroah-Hartman and a mail from Linus Torvalds on the canonical patch format.
- **Type:** official documentation
- **Where to find it:** https://www.kernel.org/doc/html/latest/process/submitting-patches.html (the rendered page carried the version label 7.4.0-rc1 on the day of review)
- **Last reviewed:** 2026-11-04
- **Used by:** chat/skills/writing-style/SKILL.md

## What the source covers

The document tells a contributor how to get a change accepted into the Linux kernel. It is a working manual for a specific process: a change travels as a plain-text email to a mailing list, maintainers and reviewers argue about it in reply, and if it is accepted its description becomes a permanent entry in the `git` history of one of the largest and longest-lived codebases in existence. Most of the document is about that process. A smaller part, concentrated in two sections, is about how to write the description of a change, and that part is what the writing-style skill draws on.

The document says at the start that it is a terse collection of suggestions and points elsewhere for the full development process and a pre-submission checklist. It assumes the contributor uses `git`.

### Describing a change

The section on describing changes is the core of the document for the skill's purposes. It gives a sequence of instructions, and each one comes with an operational reason.

The first is to state the problem. Any change, from a one-line fix to a large feature, exists because something was wrong or missing, and the description has to persuade the reviewer that the problem is real and worth fixing. The document frames this in terms of the reviewer's attention: the opening has to give the reviewer a reason to keep reading past the first paragraph.

The second is to describe the impact on users. Crashes and lockups make their own case, but many bugs are quieter, and the document asks the writer to explain what a problem could do to users even when it was found by reading code rather than by a failure in the field. The reason is specific to the kernel's distribution model. Most Linux systems run kernels from stable or vendor trees that pick individual patches from upstream, and the people choosing those patches decide from the description. So the description should carry whatever helps that choice: the conditions that trigger the problem, log excerpts, crash descriptions, regressions in performance or latency.

The third is to quantify optimizations and trade-offs. A claim of better performance, lower memory use, a smaller stack footprint or a smaller binary has to come with numbers. The writer also has to describe the costs that are not obvious. The document's reasoning is that optimizations usually trade one resource against another (processor time, memory, readability) or, for heuristics, favor one workload over another, and the reviewer can only weigh benefit against cost if both are on the page.

Only after the problem is established does the document ask for the technical description of the fix, written in plain English so the reviewer can check that the code does what the author intends. It asks for one problem per patch and treats a long description as a sign that the patch should be split.

A resubmitted patch has to carry its full description and justification again. Saying that this is the next version is not enough, because maintainers will not dig through earlier versions or linked pages, and some reviewers never saw the earlier versions. The document's word for the goal is self-contained.

The document asks for the imperative mood, phrased as if the writer were giving the codebase an order, and contrasts this with descriptions that start from the patch as subject or from the author in the first person. Git's own SubmittingPatches uses nearly the same wording and example, so the two projects share this convention directly.

On referring to other commits, the document asks for the one-line summary of the commit along with its hash, so that a reader knows what the commit is about without looking it up, and for at least 12 characters of the hash. The reason given is collisions: the kernel repository holds so many objects that a short hash which is unique today may not be unique in five years.

On outside material, the document asks for `Link:` tags pointing to earlier discussion or background, preferably through the lore.kernel.org archive, and asks the writer to check that each link works. It then adds that the explanation should be understandable without those links: the writer should summarize the relevant points of the discussion in the description itself. For bug fixes it asks for a `Closes:` tag pointing to a public report, and for a `Fixes:` tag naming the commit that introduced the bug, written on one line even when it runs past the wrap limit so that scripts can parse it.

### The canonical patch format

The second section that bears on writing describes the shape of a patch email. The subject carries a bracketed prefix (the word PATCH, an optional version or RFC marker and a zero-padded sequence number in a series), then a subsystem name, a colon and a summary phrase. The body holds an optional author line, the explanation wrapped at 75 columns, a blank line, the sign-off lines, a `---` marker, comments meant only for reviewers, and then the diff.

The guidance on the summary phrase is the most specific writing advice in the document. The subsystem prefix identifies the area of the kernel being changed. The summary must not be a file name and must not be repeated across the patches of a series. The document explains why the phrase matters: it becomes a global identifier for the patch, it propagates into the `git` history, people search for it to find the discussion, and months later it may be the only text a developer sees while scanning thousands of patches in a one-line log view. For those reasons the summary must be at most 70 to 75 characters, and it must say both what the patch changes and why the patch might be needed. The document admits that being brief and descriptive at once is hard.

The explanation body is aimed at a competent reader who has long forgotten the discussion behind the patch. The document asks for symptoms of the failure, such as kernel log or oops messages, because people searching the history for a fix search on those symptoms. It wants enough detail that a reader weeks or years later can understand why the patch exists. For a fix to compile failures it asks for enough of the error output to make the patch findable, and no more. A subsection on backtraces makes the same point about trimming: full log output carries timestamps, module lists and register dumps that distract, and a useful backtrace keeps only the lines that show the call chain.

The commentary section separates what is permanent from what is temporary. Anything relevant only to the current review, such as a diffstat or a list of changes between versions, goes below the `---` marker, where patch tools strip it automatically when the patch is applied. Version history placed above the marker has to be removed by hand.

### What else the document covers

Most of the document deals with matters the skill does not touch. It explains how to obtain the right source tree, how to split a change into logical patches that each build and run, how to run the style checker (and when its warnings can be ignored), how to choose maintainers and mailing lists, why patches go inline in plain text rather than as attachments, how to answer reviewers, why replies should be interleaved and trimmed, how long to wait before resending, and how to supply base-tree information for automated testing. A long middle stretch defines the attribution and certification tags: the Developer's Certificate of Origin behind `Signed-off-by:`, and the meanings of `Acked-by:`, `Co-developed-by:`, `Reported-by:`, `Tested-by:`, `Reviewed-by:` and `Suggested-by:`, including the rule that most of them need the named person's permission. Every one of these tags names a person with a known identity; the document does not accept anonymous contributions.

### Strengths and limits as evidence

The document records practice. Nothing in it is measured, and its rules are not tested against alternatives. Its authority comes from the scale and age of the project it governs: the conventions were refined over decades by maintainers who read large volumes of patches, and many of them are now built into tools (`git format-patch`, the style checker, bots that act on `Closes:` tags).

Its strength is that almost every rule carries a concrete reason tied to a reader's task. The reviewer needs to weigh cost against benefit, so trade-offs must be stated. The stable-tree maintainer chooses patches from descriptions, so user impact must be explicit. The developer scanning a one-line log sees only the summary, so the summary must carry both the change and its reason. These reasons transfer well to any project with a long-lived history, which is why the skill can use them outside the kernel.

The limits are also clear. Some numbers are hedged or loose: the summary limit is given as a range of 70 to 75 characters rather than a single figure, and the document never says why. Some rules follow from the email workflow and lose their force elsewhere; the 75-column wrap and the bracketed subject prefix exist because patches travel as email. The document's own examples are not consistent on capitalization after the subsystem prefix: some summaries start with a lowercase verb and some with a capital, and no rule settles it. Its own prose leans on "should" for both requirements and recommendations, which makes some rules hard to rank. And it was written for experienced C developers working on one project, so its advice about evidence (logs, backtraces, benchmark figures) assumes a technical reader throughout.

## What the skill applies

The source shaped section 9 of SKILL.md, mostly its commit-message bullet, with smaller contributions to section 4 and section 1.

**Imperative subject line.** The commit-message bullet in section 9 asks for a subject line in the imperative. The kernel document and git's SubmittingPatches both ask for it, with the same reasoning: the description reads as an instruction to the codebase. The kernel applies the mood to the whole description of the change, subject and body alike. The skill states it only for the subject line and does not repeat it for the body, where the body's account of the old behavior is in the present tense (a rule that comes from git, which the kernel does not state).

**Subject length.** The skill's figure of about 50 characters comes from git's soft limit. The kernel's ceiling of 70 to 75 characters is well above it, so a subject that meets the skill also meets the kernel. The kernel's reasons for keeping the summary short (it is what people scan and search for months later) support the skill's choice of a short target. The skill does not adopt the kernel's further requirement that the subject state why the change is needed; it leaves the why to the body.

**Prefixes and capitalization.** The bullet's instruction to follow the repository's existing conventions for prefixes and capitalization covers the kernel's subsystem prefix and git's area prefix without choosing between them. This wording also absorbs the inconsistency in the kernel's own examples about capitalizing the word after the prefix. The kernel's ban on file names as the summary and on repeating the same summary across a series are not in the skill.

**Stating the problem first.** The bullet says the body explains the problem the change solves and why this approach solves it. That order matches the kernel's sequence: establish the problem, then describe the fix. The kernel's argument that the reviewer must be convinced in the first paragraph is the same argument section 1 makes for opening with the answer, applied to a commit body.

**User-visible impact.** The bullet says the body states what users will notice. This clause comes from the kernel's instruction to describe user-visible impact, including for problems found by reading code. The kernel's own reason is specific to its downstream trees, whose maintainers pick fixes by reading descriptions. The skill keeps the general point that applies to any project: a reader of the history needs to know what the problem does to the people who run the code, and a description of what is wrong inside the code does not tell them that. The clause also matches section 2's call for specifics, since the kernel's examples of impact are concrete (crashes, lockups, latency spikes, the conditions that trigger them).

**Numbers for claimed improvements, with their costs.** The bullet asks for numbers for any claimed improvement along with its costs. Both halves come from the kernel's instruction to quantify optimizations and trade-offs: improvements in speed, memory, stack use or binary size need figures, and the non-obvious costs of the change need describing so the reviewer can weigh one against the other. The cost clause brings the skill in line with the kernel's view that an optimization usually trades one resource or workload against another. Section 9's bullet on claiming only what you can show, which asks for the source or measurement behind any performance claim, and revision item 13, which catches claims of best, fastest or guaranteed that are not literally true, rest on the same instruction together with the corporate style guides. Section 4 governs how the numbers are presented: a figure needs a comparison, and a relative change needs the absolute values beside it. The kernel does not spell out either point, but a claim of improvement is a before-and-after comparison by nature, and a reviewer weighing costs needs the absolute figures. Read together, the kernel supplies the obligation to give numbers and costs, and section 4 supplies the form. Section 4's rule against false balance is the general version of the same honesty: a commit message that reports only the gain presents a one-sided case.

**Summarizing outside discussion.** The bullet asks the body to summarize outside discussion instead of only linking to it. The kernel says this twice: in its advice to make the explanation understandable without the linked pages, and in its rule that a resubmitted patch must carry its full description. Git says the same. The kernel's reasoning (the reader may never have seen the earlier thread, and links can die) is the reason the skill treats a bare link as incomplete.

**Self-contained text that outlasts the discussion.** The kernel's target reader for the body, a competent person who has forgotten the discussion, fits the skill's general account of persisted text in its scope section: people and later agents read it many times, long after the conversation. The kernel's split between permanent changelog text and temporary review notes below the `---` marker is a narrower version of the distinction section 9 draws in its timeless-documentation bullet, where commit messages may record what changed but documentation should not go stale.

**Plain English and specifics.** The kernel's request for the technical description in plain English, and its preference for concrete symptoms, logs and conditions over general claims, match section 2's rules on plain and specific words. The skill did not need the kernel for those rules, but the kernel is the clearest statement of them for commit messages.

**Don't narrate the diff.** The kernel's emphasis on why a patch exists, and its request that the description give the reviewer a plain account of intent to check the code against, support the bullet's closing instruction, though the kernel never says outright not to restate the diff. The rule itself is stated more directly in AGENTS.md and the general technical guides.

## What the skill omits

**The subject must state why.** The kernel asks the summary to carry both the change and its reason. The skill keeps the subject to about 50 characters and leaves the reason to the body, which is closer to git's practice. A 50-character subject rarely has room for both.

**Commit references.** The kernel's format for citing another commit (at least 12 characters of the hash plus the commit's one-line summary) is not in the skill. Git uses a different format, and the right choice depends on the repository, so the skill leaves it to the repository's conventions.

**Symptoms, logs and backtraces.** The advice to include failure symptoms for searchers and to trim backtraces to the relevant call chain is specific to debugging-heavy projects. The skill's general rule to be specific covers the principle, and it does not prescribe log content.

**Process and format machinery.** The skill ignores the email workflow, the bracketed subject prefix, the 75-column wrap, the `---` marker, the tag system and the Developer's Certificate of Origin, the rules on choosing recipients, plain-text mail, interleaved replies, resending, base-tree information and tooling. These govern how kernel patches move and are credited. They are not rules about prose, and most do not apply to repositories that use pull requests. The skill keeps only the general instruction to follow a repository's own conventions.

**Splitting changes.** The rules about one logical change per patch, and a series that builds at every step, are about structuring code changes. The skill is about the text that describes them, so it does not include them, though the kernel's observation that a long description signals a patch to split is a sound check for any writer.

## Conflicts with other references

The real tensions are with git's SubmittingPatches (documented in `git-submitting-patches.md`) and with the commit-message section of AGENTS.md. None of the other references in this folder addresses commit messages, so there is no conflict with them.

**Subject length.** The kernel sets a ceiling of 70 to 75 characters. Git sets a soft limit of 50 characters. SKILL.md section 9 asks for about 50 characters with no outer bound, and AGENTS.md repeats that figure and defers to section 9. The figures are compatible in one direction: a subject of about 50 characters satisfies all three. The skill followed git's tighter target. The kernel's own reasoning (the summary is read in one-line log views and searched for) favors brevity, so the skill's choice does not contradict the kernel's purpose, only its stated number.

**Capitalization after the prefix.** Git says the word after the area prefix is lowercase unless it would be capitalized mid-sentence anyway. The kernel has no rule, and its examples go both ways. The skill does not choose; it tells the writer to follow the repository's existing conventions. AGENTS.md's example has no prefix, and it starts with a capitalized verb.

**Body wrap width.** The kernel wraps the explanation at 75 columns. AGENTS.md asks for about 72. Git's SubmittingPatches, in the text reviewed, gives no wrap width. SKILL.md is silent. The difference is small and comes from tooling, and inside this repository the AGENTS.md figure applies.

**Commit reference format.** The kernel cites a commit by at least 12 hash characters followed by its quoted one-line summary in parentheses. Git cites a commit by an abbreviated hash followed by its subject and date in parentheses. The skill takes no side and leaves the format to the repository.

**Present tense for the old behavior.** The skill's clause about describing the old behavior in the present tense comes from git, which states it explicitly. The kernel does not address tense for the problem statement. The kernel is silent on the point, so the two documents do not disagree.

**Attribution trailers.** The kernel's tags credit people with known identities, and it rejects anonymous contributions. AGENTS.md tells agents to commit as the repository owner and not to add co-author trailers that credit an AI model. The two do not conflict: the kernel's rules govern who can certify and be credited for a kernel patch, and the AGENTS.md rule governs this repository's history. Neither rule bears on prose style.

**Precedence.** SKILL.md ranks the technical guides below the news manuals and says that where section 9 disagrees with sections 1 to 8, the earlier sections win. Nothing in the kernel's description advice conflicts with sections 1 to 8. Its most specific rules (numbers and costs for claims, user impact, problem first, self-contained text) reinforce them.

## Review notes

I read the full local copy of `Documentation/process/submitting-patches.rst` from start to finish. I checked the rendered page at kernel.org on 2026-11-04: its section headings matched the local copy, and it confirmed the 70 to 75 character summary limit, the 75-column wrap, the 12-character hash minimum and the imperative-mood rule. The page showed the documentation version label 7.4.0-rc1 and no revision date, so I could not tie the local copy to a specific kernel commit. I compared the relevant rules against the local copy of git's `Documentation/SubmittingPatches`, reading its sections on the commit message, the subject and its prefix, the present tense, the imperative mood, commit references and trailers. I did not read the companion `git-submitting-patches.md` reference, which was being written at the same time, so any difference in how the two files describe git's rules should be checked against git's text. I did not follow the kernel document's external reading list (the guides by Morton, Garzik and Kroah-Hartman, Torvalds's mail and Kleen's paper) or the linked kernel documents on the development process, coding style and stable rules. The SKILL.md and AGENTS.md text compared here is the version on disk on 2026-11-04, after the commit-message rules in both were changed to a subject of about 50 characters, and after the section 9 commit bullet gained its clauses on what users will notice and on the costs of a claimed improvement. That later SKILL.md version also moved the dated and overclaiming words out of section 9 into revision item 13.
