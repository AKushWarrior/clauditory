#!/usr/bin/env python3
"""Generate the extracted rules from the rule lists mined from each guide.

Reads lint/sources/*.yaml (rules extracted from each style guide), drops
rules that don't apply to American English or that contradict the skills,
settles conflicts between guides, and writes:

- lint/terms.gen.json: the words, phrases, and patterns, matched by terms.go
- lint/rules.gen.yml: registry entries for every generated rule

Usage: python3 lint/tools/genrules.py [--recheck BINARY] [--prune FILE]

--recheck names a program that reads one regex per line and prints OK or
ERR for each, using Vale's regex engine. --prune names a file of entries to
drop, one "Style<TAB>wrong" per line, usually produced by prune.py.
"""

import argparse
import collections
import glob
import json
import os
import re
import subprocess
import sys

import yaml

HERE = os.path.dirname(os.path.abspath(__file__))
ROOT = os.path.dirname(HERE)
sys.path.insert(0, HERE)
from expand import expand  # noqa: E402

STYLE = {"nyt": "NYT", "ap": "AP", "chicago": "Chicago", "oxford": "Oxford",
         "ibm": "IBM", "apple": "Apple", "microsoft": "Microsoft",
         "google": "Google", "redhat": "RedHat"}

ORDER = {
    "grammar": ["Oxford", "AP", "Chicago", "NYT", "RedHat", "Apple", "Google", "IBM", "Microsoft"],
    "wording": ["NYT", "Chicago", "AP", "RedHat", "Apple", "Google", "IBM", "Microsoft", "Oxford"],
    "technical": ["RedHat", "Apple", "Google", "IBM", "Microsoft", "NYT", "Chicago", "AP", "Oxford"],
}

# Forms the skills prescribe; no generated rule may flag them.
PROTECTED = {
    "web page", "web pages", "website", "websites", "email", "emails", "online",
    "username", "usernames", "filename", "filenames", "database", "databases",
    "internet", "login", "log in", "logs in", "logged in", "cannot", "all right",
    "no one", "toward", "percent", "data", "setup", "set up", "backup", "back up",
    "for example", "that is", "and", "or", "use", "before", "after", "about",
    "because", "if", "whether", "with", "you", "we", "they", "their", "can",
    "must", "might", "may", "should", "need", "help", "start", "try", "said",
    "say", "says", "american", "british", "english", "is", "are", "has", "have",
}

# Contractions: the skill allows them in conversation, so no rule may ban them.
CONTRACTION = re.compile(r"(?i)\w+(?:n['’]t|['’](?:s|re|ve|ll|d|m))|can['’]t|won['’]t|let['’]s")

PREFIX = "Ext"
LINKS = {
    "NYT": "chat/skills/writing-style/references/nyt-manual-of-style.md",
    "AP": "chat/skills/writing-style/references/ap-stylebook.md",
    "Chicago": "chat/skills/writing-style/references/chicago-manual-of-style.md",
    "Oxford": "chat/skills/writing-style/references/oxford-new-harts-rules.md",
    "IBM": "chat/skills/technical-writing/references/ibm-style-guide.md",
    "Apple": "chat/skills/technical-writing/references/apple-style-guide.md",
    "Microsoft": "chat/skills/technical-writing/references/microsoft-writing-style-guide.md",
    "Google": "chat/skills/technical-writing/references/google-developer-style-guide.md",
    "RedHat": "chat/skills/technical-writing/references/red-hat-technical-writing-style-guide.md",
}


def rank(style, domain):
    order = ORDER[domain]
    return order.index(style) if style in order else len(order)


def settle(a, b):
    """Return the domain whose order decides between rules a and b."""
    if "wording" in (a["domain"], b["domain"]) or a["domain"] != b["domain"]:
        return "wording"
    return a["domain"]


def boundary(pat):
    """Add RE2-style word boundaries where the pattern begins or ends with a
    word character, so Vale's literal prefilter keeps working."""
    flags = "(?i)" if pat.startswith("(?i)") else ""
    body = pat[4:] if flags else pat
    head = body[3:] if body.startswith("(?:") else body
    lead = bool(re.match(r"[\w\[]|\\[wd]", head)) and not body.startswith(("\\b", "(?<", "^"))
    trail = bool(re.search(r"[\w\]]$|\\[wd]$|[)+*?}]$", body)) and not body.endswith(("\\b", "$", "\\.", "\\)")) \
        and not re.search(r"\(\?[=!][^()]*\)$", body)
    return flags + ("\\b" if lead else "") + body + ("\\b" if trail else "")


def load():
    rows = []
    for path in sorted(glob.glob(os.path.join(ROOT, "sources", "*.yaml"))):
        for r in yaml.safe_load(open(path)) or []:
            if isinstance(r, dict):
                rows.append(r)
    return rows


def handwritten():
    """Forms the hand-written styles prefer and flag."""
    pref, dis = set(), set()
    for path in glob.glob(os.path.join(ROOT, "styles", "*", "*.yml")):
        if os.path.basename(path).startswith(PREFIX):
            continue
        try:
            d = yaml.safe_load(open(path)) or {}
        except Exception:
            continue
        for k, v in (d.get("swap") or {}).items():
            e = expand(str(k))
            if e:
                dis.update(x.lower() for x in e[0])
            if isinstance(v, str) and len(v) < 40 and not v.startswith(("Do not", "Use ")):
                for alt in v.split("|"):
                    pref.add(alt.strip().lower())
    return pref, dis


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--recheck", default=os.environ.get("RECHECK"))
    ap.add_argument("--prune", default=os.path.join(ROOT, "sources", "prune.tsv"))
    args = ap.parse_args()

    prune = set()
    if os.path.exists(args.prune):
        for line in open(args.prune):
            if "\t" in line and not line.startswith("#"):
                s, w = line.rstrip("\n").split("\t", 1)
                prune.add((s, w))

    stats = collections.Counter()
    rules = []
    for r in load():
        style = STYLE.get(str(r.get("source", "")).lower())
        kind = r.get("kind")
        wrong = r.get("wrong")
        if not style or kind not in ("substitution", "existence", "casing") or not wrong:
            stats["dropped: malformed"] += 1
            continue
        if r.get("american") is False:
            stats["dropped: British usage"] += 1
            continue
        wrong = str(wrong)
        right = r.get("right")
        right = None if right is None else str(right)
        if kind != "existence" and not right:
            kind = "existence"
        domain = r.get("domain") if r.get("domain") in ORDER else "wording"
        fix = r.get("fix") if r.get("fix") in ("safe", "review", "none") else "none"
        context = str(r.get("context") or "general").lower()
        if context not in ("general", "technical writing", "technical", "computing", "software", "documentation"):
            fix = "none"
        if style == "Oxford" and domain == "wording":
            fix = "none"
        if right and "$" in right:
            fix = "none"
        if kind == "existence":
            fix = "none"
        note = re.sub(r"\s+", " ", str(r.get("note") or "")).strip()
        exp = expand(wrong)
        lits = None
        if exp:
            lits = [x for x in exp[0] if x.strip()]
            if not lits or any(not re.search(r"\w", x) for x in lits):
                lits = None
        rules.append({"style": style, "kind": kind, "wrong": wrong, "right": right, "domain": domain,
                      "fix": fix, "note": note, "entry": str(r.get("entry") or ""), "lits": lits})

    # Drop rules that contradict the skills or that fire on reviewed house prose.
    pref, dis = handwritten()
    protected = PROTECTED | pref
    kept = []
    for r in rules:
        lows = [x.lower() for x in r["lits"]] if r["lits"] else []
        bp = boundary(r["wrong"])
        if (r["lits"] and any((r["style"], x) in prune for x in lows)) or \
                any((r["style"], p) in prune for p in (r["wrong"], bp, "(?i)" + bp)):
            stats["dropped: fires on reviewed house prose"] += 1
            continue
        if r["lits"] and any(x in protected for x in lows):
            stats["dropped: flags a form the skills prescribe"] += 1
            continue
        if (r["lits"] and any(CONTRACTION.fullmatch(x) for x in r["lits"])) or "n['’]t" in r["wrong"] or "n't" in r["wrong"]:
            stats["dropped: bans contractions"] += 1
            continue
        if (r["right"] and re.search("[—–]|--", r["right"])) or re.search(r"(?i)\bem[- ]?dash", r["note"]) and "avoid" not in r["note"].lower():
            stats["dropped: prescribes a dash"] += 1
            continue
        if r["right"] and r["right"].lower() in dis:
            stats["dropped: replacement is a form the skills flag"] += 1
            continue
        if r["right"] and r["lits"] and r["right"].lower() in lows:
            stats["dropped: replaces a form with itself"] += 1
            continue
        kept.append(r)
    rules = kept

    # Settle conflicts between guides over the same literal forms.
    for i, r in enumerate(rules):
        r["i"] = i
    flagged = collections.defaultdict(list)
    prescribed = collections.defaultdict(list)
    for r in rules:
        for x in r["lits"] or []:
            flagged[x.lower()].append(r)
        if r["right"]:
            prescribed[r["right"].lower()].append(r)
    gone = set()
    for r in rules:
        lows = [x.lower() for x in r["lits"]] if r["lits"] else []
        why = None
        for x in lows:
            for o in prescribed.get(x, []):
                if o is not r and rank(o["style"], settle(r, o)) < rank(r["style"], settle(r, o)):
                    why = "flags a form a higher-ranked guide prescribes"
        if not why and r["right"]:
            for o in flagged.get(r["right"].lower(), []):
                if o is not r and rank(o["style"], settle(r, o)) < rank(r["style"], settle(r, o)):
                    why = "prescribes a form a higher-ranked guide flags"
        if not why:
            for x in lows:
                for o in flagged.get(x, []):
                    if o is r or o["i"] in gone:
                        continue
                    ro, rr = rank(o["style"], settle(r, o)), rank(r["style"], settle(r, o))
                    if ro < rr or (ro == rr and o["i"] < r["i"]):
                        why = "duplicates a higher-ranked rule"
                        break
                if why:
                    break
        if why:
            gone.add(r["i"])
            stats["dropped: " + why] += 1
    rules = [r for r in rules if r["i"] not in gone]

    # Check the remaining regular expressions with Vale's engine.
    regex = [r for r in rules if not r["lits"]]
    for r in regex:
        r["pattern"] = boundary(r["wrong"])
    if args.recheck and regex:
        out = subprocess.run([args.recheck], input="\n".join(r["pattern"] for r in regex) + "\n",
                             capture_output=True, text=True).stdout.split("\n")
        bad = {r["i"] for r, line in zip(regex, out) if not line.startswith("OK")}
        stats["dropped: regex Vale can't compile"] += len(bad)
        rules = [r for r in rules if r["i"] not in bad]

    def rule_name(r, literal):
        dom = {"wording": "", "grammar": "Grammar", "technical": "Tech"}[r["domain"]]
        kind = {"substitution": "Terms", "casing": "Casing", "existence": "Avoid"}[r["kind"]]
        fix = {"safe": "", "review": "Review", "none": "Flag"}[r["fix"]] if r["kind"] != "existence" else ""
        return f"{r['style']}.{PREFIX}{'' if literal else 'Pattern'}{dom}{kind}{fix}"

    def message(r):
        return r["note"] or f"Check this form against the {r['style']} entry for {r['entry']}."

    registry = {}
    terms = []
    seen = set()
    for r in rules:
        if not r["lits"]:
            continue
        name = rule_name(r, True)
        exact = r["kind"] == "casing"
        text = r["right"] if r["kind"] != "existence" else message(r)
        for x in r["lits"]:
            key = (name, x if exact else x.lower())
            if key in seen:
                continue
            seen.add(key)
            terms.append([name, "flag" if r["kind"] == "existence" else "sub", exact, x, text])
        registry[name] = {"domain": r["domain"], "fix": r["fix"], "topic": "term", "literal": True}
        stats[f"literal entries for {r['style']}"] += len(r["lits"])

    for path in glob.glob(os.path.join(ROOT, "styles", "*", PREFIX + "*.yml")):
        with open(path) as f:
            if f.readline().startswith("# Generated by lint/tools/genrules.py"):
                os.remove(path)
    npat = 0
    for r in rules:
        if r["lits"]:
            continue
        name = rule_name(r, False)
        ci = r["wrong"].startswith("(?i)")
        pat = r["pattern"] if ci or r["kind"] == "casing" else "(?i)" + r["pattern"] if r["kind"] == "existence" else r["pattern"]
        text = r["right"] if r["kind"] != "existence" else message(r)
        terms.append([name, "flag" if r["kind"] == "existence" else "sub", False, pat, text, "re"])
        registry[name] = {"domain": r["domain"], "fix": r["fix"], "topic": "term", "literal": True}
        stats[f"pattern entries for {r['style']}"] += 1
        npat += 1

    with open(os.path.join(ROOT, "terms.gen.json"), "w") as f:
        f.write("[\n" + ",\n".join(json.dumps(t, ensure_ascii=False) for t in sorted(terms, key=lambda t: (t[0], t[3]))) + "\n]\n")
    with open(os.path.join(ROOT, "rules.gen.yml"), "w") as f:
        f.write("# Generated by lint/tools/genrules.py. Don't edit by hand.\n")
        yaml.safe_dump({"rules": registry}, f, sort_keys=True, width=1000)

    for k, v in sorted(stats.items()):
        print(f"{v:6d}  {k}")
    print(f"{len(terms) - npat:6d}  literal terms; {npat} pattern rules")


if __name__ == "__main__":
    main()
