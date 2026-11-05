# Anthropic Prompt Engineering and Skill Authoring Documentation

- **Source:** Anthropic, Claude Platform documentation: Prompt engineering overview; Prompting best practices (including its sections on being clear and direct, adding context and motivation, using examples, roles, and output formatting); the Sonnet 5.5 prompting guide; and Skill authoring best practices. Anthropic, current web edition, read November 2026.
- **Type:** official documentation
- **Where to find it:** https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/overview ; https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/claude-prompting-best-practices ; https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/prompting-claude-sonnet-5-5 ; https://platform.claude.com/docs/en/agents-and-tools/agent-skills/best-practices
- **Last reviewed:** 2026-11-04
- **Used by:** chat/skills/writing-style/SKILL.md

## What the source covers

### Purpose and scope

These pages are Anthropic's own guidance to developers on how to write instructions for Claude models. They are living documents, revised with each model release, so statements in them are tied to specific model versions. The overview page is short. It sets preconditions (you need a definition of success, a way to test against it, and a first draft to improve), points out that some failures are better fixed by switching models than by rewording a prompt, and sends readers to the best practices page as the main reference. It lists the techniques that page covers: clarity and examples, XML structuring, role prompting, thinking, and prompt chaining. It also mentions two tutorials, one of which was written for older Claude 3 models, and tells readers to follow current guidance where the two differ.

The best practices page is long and mixes timeless advice with model-specific tuning notes. Large parts of it concern tool use, thinking budgets, multi-window agent workflows, subagents, vision and frontend design, and migration between model versions. Those parts are irrelevant to prose style. The parts that bear on writing are the general principles near the top and the section on output and formatting.

### Being clear and direct

The central teaching is to treat the model as a capable new colleague who has none of your context. The more precisely you say what you want, the better the result. The page offers a simple check: show the prompt to a colleague with little background, and if that person would be confused, the model probably will be too. It recommends stating output format and constraints explicitly and using numbered steps or bullets when order or completeness matters. It also says to ask outright for extra effort when you want it. A request for an analytics dashboard gets a basic result, while the same request with an instruction to include as many relevant features as possible and go beyond the basics gets a fuller one. The implication is that recent models follow what is written literally and do not reliably infer unstated ambitions.

### Context and motivation

The page argues that a rule with a reason attached generalizes better than a bare rule. Its example replaces a flat prohibition on ellipses with a sentence explaining that the output will be read aloud by a text-to-speech engine that cannot pronounce them. With the reason in hand, the model can extend the rule to related cases (other symbols a speech engine would stumble on) that the author never listed. The page states this generalization claim directly but gives no measured effect size, so it rests on Anthropic's experience rather than on a published experiment.

### Examples

Examples are described as one of the most dependable ways to steer format, tone and structure. The page gives three requirements: they should be relevant (close to the real use case), diverse (covering edge cases and varied enough that the model does not copy an accidental pattern), and structured (each wrapped in an example tag, with several wrapped together in an examples tag, so the model can tell them from instructions). It recommends three to five examples. It adds that you can ask Claude to review your examples for relevance and diversity or to generate more from them. A separate note says that worked examples also shape how the model reasons in its thinking blocks, and recommends presenting each as a problem, a method and an answer. The fetch for the older standalone multishot page returned the same best practices content, which suggests the material was folded into one page. A separate use-examples URL returned no readable content, so I could not confirm whether it exists as its own page.

### Roles

A single sentence in the system prompt naming a role focuses behavior and tone. The page's illustration is a one-line description of a coding assistant that specializes in Python. It does not say how much role prompting helps on any task, and it offers no comparison of roles.

### Structure with tags

The page recommends XML-style tags to separate instructions, context and input, with consistent names and nesting for hierarchical material, and says tags reduce misreading when a prompt mixes several kinds of content. For long inputs of twenty thousand tokens or more, it advises putting documents first and the question last, and claims this can improve quality by as much as 30 percent on complex multi-document inputs. That figure is Anthropic's own and the page does not describe the test behind it.

### Output formatting and markdown

This is the section most relevant to a writing skill. It gives four levers for controlling formatting:

1. Tell the model what to do and skip telling it what to avoid. A request for smoothly flowing prose paragraphs works better than a ban on markdown.
2. Use tags as format indicators, such as asking for the prose sections of a reply inside a named prose tag.
3. Match the prompt's own style to the output you want. The page says that removing markdown from the prompt itself tends to reduce markdown in the reply.
4. Supply a detailed formatting block. The sample block asks for clear flowing prose with complete paragraphs for reports, documents and explanations. It limits markdown to inline code, code blocks and simple headings, rules out bold and italics, reserves lists for truly discrete items or explicit requests, and asks for items woven into sentences rather than strings of short bullets, with the aim of readable text that guides the reader through ideas instead of fragmenting them.

The page also warns that one newer model already formats less by default, so the sample block can strip out structure the content needs. The advice there is to remove the block or replace it with the shorter rule in that model's own guide. A parallel sample handles LaTeX, which models tend to use for math by default, by asking for plain text characters for operators.

### Communication style and verbosity

The page says recent models are more concise, direct and conversational than earlier ones. They report facts rather than celebrate progress, and they may skip a summary after a run of tool calls unless one is requested. One model line runs longer by default, so conciseness has to be requested there. The Sonnet 5.5 guide adds that progress notes between tool calls are short by default and that an instruction to think less does not reliably cut thinking, since effort level is the control for that. That guide has no section on markdown, chat formatting or prose style, so it does not add writing guidance beyond what the best practices page gives. It does include scope-control wording, such as telling the model to stop and report when the requested work is done and to mention possible extras rather than do them. That is relevant by analogy to padded endings, though the guide frames it for coding tasks.

### Over-forceful emphasis

A recurring theme is that newer models respond strongly to the system prompt. The page recommends replacing capitalized, absolute phrasing such as critical and must with plain statements of when to do something, because forceful language now causes overtriggering. The same applies to anti-laziness prompts and blanket defaults. The pattern across the page is to use conditional, reasoned instructions in place of shouting.

### Skill authoring guidance

The skill authoring page covers how to write an Agent Skill, which is the delivery format for this writing skill. Its main points:

- **Conciseness.** The context window is shared among the system prompt, conversation, other skills' metadata and the user's request, so a skill should add only what the model does not already know. The page's test is whether the model really needs each explanation, and its example contrasts a short code snippet with a longer version that explains what PDFs are.
- **Description.** The description is what the model uses to choose among potentially a hundred or more skills. It should say what the skill does and when to use it, include trigger terms, and be written in the third person. It is capped at 1,024 characters, and names are capped at 64 characters.
- **Length.** Keep the body of the main skill file under 500 lines and split into separate files as it approaches that size.
- **Progressive disclosure.** Only metadata loads at startup. The main file loads when the skill is relevant, and supporting files load only if needed. References should sit one level below the main file, because nested references risk partial reads. Reference files over 100 lines should open with a table of contents.
- **Degrees of freedom.** Match specificity to fragility. Open-ended tasks suit general guidance. Fragile, sequence-critical operations suit exact scripts. Writing style is the open-ended kind.
- **Examples.** For style-dependent output, input and output pairs convey format and level of detail more clearly than descriptions.
- **Terminology.** Pick one term for each concept and keep it.
- **Workflows and feedback loops.** Complex tasks benefit from sequential steps and a validate-then-fix loop before finalizing.
- **Evaluation.** Build tests before writing extensive documentation: observe failures without the skill, write about three scenarios that target them, measure a baseline, write the minimum that closes the gaps, and iterate. Test across smaller and larger models, since a skill tuned on one may need more detail for a smaller one.

### Strength and generality of the evidence

This is official guidance from the model's maker, which makes it authoritative about how Claude models behave, and several claims are consistent with each other across pages. It is also practitioner advice. It cites few measurements, gives no sample sizes, and tunes much of its content to specific model versions that will be superseded. The claims most useful here (positive framing, reasons, matching prompt style to output style, few diverse examples) are general and stable across the versions discussed, but their size is unquantified. Nothing in these pages studies reader reactions to AI prose, so they speak to how to instruct a model and not to which prose habits readers dislike. That second question belongs to the research references in this folder.

## What the skill applies

**Positive framing and the split between drafting and revision (How to use this skill, steps 1 and 2).** The docs recommend telling the model what to do in place of what to avoid. The skill's drafting sections describe the target style, such as open with the answer, use the plain word, state claims directly, and write in paragraphs. Items to remove appear only in section 10. This keeps the ban lists out of the drafting stage, which fits both the docs' advice and the priming concern raised in the semantic gravity wells reference.

**Reasons attached to rules (the opening of SKILL.md and many bullets).** The docs claim a reason lets the model generalize. The skill opens by naming two reasons behind every rule (readers recognize and distrust default machine prose, and each habit wastes time or blurs a claim) and tells the writer to break a rule when it would worsen a sentence. Individual bullets also carry rationale, for example the note that fancy words add length and make the writer sound less sure, or that bare numbers mean little without a comparison. This is the docs' context-and-motivation principle applied across the whole file in place of capitalized commands.

**Matching prompt style to output style (section 6, Structure and format).** The docs say prompt formatting leaks into output formatting. The skill instructs paragraphs by default and limits lists, tables, headers and bold to cases where they earn their place, which echoes the docs' sample block closely: lists for truly discrete items, no scattered bold, no strings of short bullets. The revision pass item on formatting creep is the same idea in a checklist.

**Examples (the Examples section).** The docs recommend three to five relevant, diverse examples with clear boundaries. The skill gives five before and after pairs covering different situations (a concept explanation, a finding, a number, a priority, a yes-or-no answer), so the examples are diverse in form and not five copies of one pattern. The skill separates them with bold labels and before and after markers; it does not use tags, which is a gap against the docs' structuring advice but a minor one in a markdown file.

**Description and length (frontmatter, overall size).** The skill description is in the third person, says what the skill does and when to use it, names the kinds of prose it covers, and states what it excludes, which matches the docs' description guidance. The whole file is about 140 lines, well under the 500-line guideline, and the skill relies on one main file rather than nesting references, in line with the one-level-deep advice.

**Conciseness (throughout).** The skill spends its words on judgment calls and concrete replacements rather than on explaining what good writing is in general. That follows the docs' test of adding only what the model does not already know, though some content (the plain-word list, the number rules) is borrowed from style manuals because the model's defaults differ from them.

**Open-ended degree of freedom.** Writing style is an open-field task by the docs' own classification, so the skill gives direction and a closing instruction to use judgment instead of scripting exact output.

**Terminology consistency.** The skill's own advice to call a thing by one name mirrors the docs' skill-authoring rule, and the skill file itself keeps its terms steady (revision pass, hit).

## What the skill omits

The skill leaves out everything about tools, thinking, effort, prefilling, subagents, long-running agent state, vision, frontend design and migration, since none of it concerns prose. It does not adopt XML tags for its own structure, partly because the skill is read as a document and partly because its examples are short. It does not use a role line, as the docs' role advice is about focusing a model on a domain and the skill's opening paragraph already sets a reader-centered target. It omits the docs' sample prose-only formatting block wholesale. That block forbids bold and bans anything beyond simple headings, whereas the skill allows a few bold terms and headers in long documents, so copying it would have contradicted the skill's own rules. The skill also does not follow the docs' evaluation advice inside SKILL.md. Building test scenarios and baselines is a development process for the repository, not content for the skill file. Finally, the skill does not reproduce the docs' warnings about model-specific behaviors, since those change by version and would date the file.

## Conflicts with other references

The main tension is with the revision pass. The docs and the semantic gravity wells reference both favor positive instructions, yet section 10 of the skill contains an explicit list of phrases, words and constructions to remove. The skill resolves this by confining the lists to a post-draft step, so the drafting guidance stays positive. The docs do not say a negative check is harmful after generation, only that negative framing is a weaker way to steer generation, so the compromise is consistent with them. Still, the docs offer no evidence on whether a post-draft checklist helps, and the skill depends on the research references (Kobak, Juzek and Ward, Liang, the Wikipedia guide) for what belongs on the list.

A second tension is internal. The docs say prompt style leaks into output style, but SKILL.md is itself written largely in bullets, while telling the writer to prefer paragraphs. The skill follows the docs on content and the user's explicit choices on the numbered revision list, but the bulleted form of its rules may nudge output toward lists. The Reinhart and Boggia references concern the model's output habits, not prompt format, so they do not settle this. No change was made, because the rules are parallel, separate items where the skill itself allows lists.

A third point involves emphasis. The docs advise against forceful, absolute phrasing, while the user explicitly asked for hard bans on em dashes and on correction framing. The skill states those bans plainly with a stated procedure and a narrow exception rather than in capitals or with threats, which fits the docs' preference for calm, specific wording and the user's instruction at once.

No real conflict exists between this source and the style manuals (NYT, AP, Chicago), since these docs say nothing on punctuation, numbers or word choice. Their advice to be clear, specific and direct about requirements agrees with the plain-word and open-with-the-answer rules the manuals support.

## Review notes

Read in full or in substantial part through the fetch tool, which returns model-written summaries of each page and not the raw page text: the prompt engineering overview, the best practices page (fetched three times with different questions), the Sonnet 5.5 prompting guide, and the skill authoring best practices page. The fetch of the be-clear-and-direct URL returned the content of the best practices page, so I treated that material as covered by the combined page and could not verify whether a separate page exists at that address. The multishot URL likewise returned the best practices content. The use-examples URL returned only generic page metadata, so nothing from it is used. Because every page reached me as a summary, exact wording, tag names, token counts and the 30 percent long-context figure should be rechecked against the live pages before anyone relies on them. The pages are updated often, and model names and version-specific advice here reflect what the fetch returned on the review date. I did not read the model migration guide, the Anthropic blog post on prompt engineering, or the Claude Fable 5.1 guide, which the docs say contains a shorter formatting rule. I did not test any claim about prompt behavior myself.
