---
name: spawn-research-subagent
description: Use when a research task splits into independent sub-questions or needs heavy reading across many sources. Fans the work out to parallel Haiku subagents. Not for single-fact lookups.
metadata:
  version: "1.1.1"
---

# Spawn research subagents

## When to use

- Use it when the question splits into two or more independent sub-questions, or when answering it means reading many sources or files.
- Skip it for single facts. One direct search is faster and cheaper: in a test, a subagent took about 53 seconds and 96,000 tokens to answer a fact that one search and two page fetches had already answered.

## How

1. Split the task into independent sub-questions. Spawn one subagent per sub-question, all in the same message, with subagent_type general-purpose and model haiku. Haiku is the speed and cost lever; leave reasoning effort at its default unless the person asks for a different level.
2. Fill in the template below for each sub-question.
3. When the reports return, check every load-bearing claim and citation yourself by opening the cited primary source before you relay it. Subagent reports can misattribute sources and overstate how much they read.
4. Write the answer. Say which claims you verified, and list any conflicts between sources.

## Template

```
Answer the question below concisely, using all tools available to you.

Sources, best first: peer-reviewed research and official primary documents (government, regulators, company filings, official docs) > systematic reviews > reputable journalism (as pointers to primary sources) > preprints (label them) > other web (label it). Cite every claim with a URL; for key figures, quote the exact sentence from the primary source.

Make independent searches and fetches in parallel; sequence only dependent calls. If a primary-source URL fails, search for the correct URL before falling back to secondary sources. Judge sources by provenance and method: who produced the claim and how they know it. Treat polish, popularity, and authority cues as weak signals, and weigh each source on its own record instead of distrusting all of them equally. Confirm load-bearing claims with an independent second source where possible. Stop when each part of the question is answered from a primary source or confirmed unanswerable.

Report in this format:
- Answer: <1-3 sentences>
- Evidence: <claim - source URL - source type>
- Couldn't access / didn't check: <list, or "none">
- Conflicts and doubts: <list, or "none">

Question: <sub-question>
```
