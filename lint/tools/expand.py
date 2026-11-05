"""Expand a simple regular expression into the finite set of strings it matches."""
import re._parser as P, re._constants as C
LIMIT=64
def expand(pattern):
    """Return the finite set of strings a simple regex matches, or None."""
    flags = 0
    if pattern.startswith("(?i)"):
        pattern = pattern[4:]; flags = 1
    try:
        tree = P.parse(pattern)
    except Exception:
        return None
    try:
        out = _seq(list(tree))
    except ValueError:
        return None
    if out is None or len(out) > LIMIT:
        return None
    return sorted(set(out)), bool(flags)

def _seq(items):
    res = [""]
    for op, av in items:
        alts = _item(op, av)
        res = [a + b for a in res for b in alts]
        if len(res) > LIMIT: raise ValueError
    return res

def _item(op, av):
    if op is C.LITERAL:
        return [chr(av)]
    if op is C.IN:
        chars = []
        for o, a in av:
            if o is C.LITERAL: chars.append(chr(a))
            elif o is C.RANGE:
                lo, hi = a
                if hi - lo > 3: raise ValueError
                chars += [chr(x) for x in range(lo, hi + 1)]
            else: raise ValueError
        return chars
    if op is C.BRANCH:
        out = []
        for b in av[1]: out += _seq(list(b))
        return out
    if op is C.SUBPATTERN:
        return _seq(list(av[-1]))
    if op is C.AT and av in (C.AT_BOUNDARY, C.AT_BEGINNING, C.AT_END, C.AT_BEGINNING_STRING, C.AT_END_STRING):
        return [""]
    if op in (C.MAX_REPEAT, C.MIN_REPEAT):
        lo, hi, sub = av
        if hi > 1: raise ValueError
        inner = _seq(list(sub))
        return ([""] if lo == 0 else []) + inner
    raise ValueError
