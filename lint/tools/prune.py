#!/usr/bin/env python3
"""Find extracted rules that fire on reviewed house prose.

Lints the repository's own prose, which follows the skills and has been
reviewed, with the extracted rules, and lists every extracted entry that
fires at least MIN times across at least FILES files. An entry that fires
that often on text written to the house style contradicts the skills or is
too broad, so genrules.py drops it.

Usage: python3 lint/tools/prune.py LINTER [--min 3] [--files 2]
Writes lint/sources/prune.tsv, keeping entries already listed there.
"""

import argparse
import collections
import glob
import json
import os
import subprocess

HERE = os.path.dirname(os.path.abspath(__file__))
LINT = os.path.dirname(HERE)
REPO = os.path.dirname(LINT)


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("linter")
    ap.add_argument("--min", type=int, default=3)
    ap.add_argument("--files", type=int, default=2)
    args = ap.parse_args()

    corpus = [os.path.join(REPO, p) for p in ("README.md", "AGENTS.md", "chat/system/prompt.md", "lint/README.md")]
    corpus += glob.glob(os.path.join(REPO, "chat", "**", "*.md"), recursive=True)
    out = subprocess.run([args.linter, "--format=json", "--fail=none", *sorted(set(corpus))],
                         capture_output=True, text=True).stdout
    alerts = [a for a in json.loads(out) if ".Ext" in a["rule"]]

    hits = collections.Counter()
    files = collections.defaultdict(set)
    for a in alerts:
        entry = a.get("entry")
        if not entry:
            continue
        style = a["rule"].split(".")[0]
        is_pattern = ".ExtPattern" in a["rule"]
        key = (style, entry if is_pattern else entry.lower())
        hits[key] += 1
        files[key].add(a["path"])

    path = os.path.join(LINT, "sources", "prune.tsv")
    keep = set()
    if os.path.exists(path):
        for line in open(path):
            if "\t" in line and not line.startswith("#"):
                keep.add(tuple(line.rstrip("\n").split("\t", 1)))
    new = {k for k, n in hits.items() if n >= args.min and len(files[k]) >= args.files}
    with open(path, "w") as f:
        f.write("# Extracted entries that fire on the repository's reviewed prose; genrules.py drops them.\n")
        for s, w in sorted(keep | new):
            f.write(f"{s}\t{w}\n")
    print(f"{len(new - keep)} new entries pruned, {len(keep | new)} in total")


if __name__ == "__main__":
    main()
