package lint

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/AKushWarrior/clauditory/lint/internal/vale/regex"
)

// The extracted word lists are too large for Vale's backtracking regex
// engine, which tries every alternative at every position. They are plain
// words and phrases, so they are matched here with a hash lookup over the
// text, after the parts that aren't prose have been masked out.

//go:embed terms.gen.json
var termsJSON []byte

// term is one entry from terms.gen.json: [rule, kind, exact, wrong, text]
// for a literal, or [rule, kind, false, pattern, text, "re"] for a regular
// expression. kind is "sub" (replace the match with text) or "flag" (text
// is the message), and exact means a literal match is case-sensitive.
type term struct {
	rule  string
	sub   bool
	exact bool
	wrong string
	text  string
}

type pattern struct {
	term
	re *regex.Regexp
}

type termIndex struct {
	patterns []pattern
	exact    map[string][]term
	folded   map[string][]term
	maxWords int
	maxRunes int
}

func loadTerms() (*termIndex, error) {
	var rows [][]any
	if err := json.Unmarshal(termsJSON, &rows); err != nil {
		return nil, fmt.Errorf("terms.gen.json: %w", err)
	}
	ix := &termIndex{exact: map[string][]term{}, folded: map[string][]term{}}
	for _, row := range rows {
		if len(row) < 5 {
			continue
		}
		t := term{rule: row[0].(string), sub: row[1].(string) == "sub", exact: row[2].(bool),
			wrong: row[3].(string), text: row[4].(string)}
		if len(row) == 6 {
			re, err := regex.Compile(t.wrong)
			if err != nil {
				return nil, fmt.Errorf("terms.gen.json: %s: %w", t.rule, err)
			}
			ix.patterns = append(ix.patterns, pattern{term: t, re: re})
			continue
		}
		if t.exact {
			ix.exact[t.wrong] = append(ix.exact[t.wrong], t)
		} else {
			k := strings.ToLower(t.wrong)
			ix.folded[k] = append(ix.folded[k], t)
		}
		if n := len(strings.FieldsFunc(t.wrong, isSep)); n > ix.maxWords {
			ix.maxWords = n
		}
		if n := len([]rune(t.wrong)); n > ix.maxRunes {
			ix.maxRunes = n
		}
	}
	return ix, nil
}

func isSep(r rune) bool { return r == ' ' || r == '-' || r == '/' }

func isWord(r rune) bool { return unicode.IsLetter(r) || unicode.IsDigit(r) }

var (
	reFence    = regexp.MustCompile("(?m)^[ \t]*(```|~~~)")
	reInline   = regexp.MustCompile("`[^`\n]*`")
	reURL      = regexp.MustCompile(`https?://[^\s)>\]]+|www\.[^\s)>\]]+`)
	reLinkDest = regexp.MustCompile(`\]\([^)\n]*\)`)
	reAngle    = regexp.MustCompile(`<[^>\n]*>`)
	reComment  = regexp.MustCompile(`(?s)<!--.*?-->`)
	reValeOff  = regexp.MustCompile(`(?s)<!--\s*vale off\s*-->.*?(<!--\s*vale on\s*-->|\z)`)
	reFront    = regexp.MustCompile(`(?s)\A---\n.*?\n---\n`)
)

// mask blanks out, with spaces of the same length, everything in a Markdown
// document that isn't prose: front matter, code blocks, code spans, URLs,
// link destinations, HTML tags and comments, and vale off regions. Line
// breaks are kept, so positions still map to the same lines and columns.
func mask(text, ext string) string {
	b := []byte(text)
	blank := func(lo, hi int) {
		for i := lo; i < hi; i++ {
			if b[i] != '\n' {
				b[i] = ' '
			}
		}
	}
	apply := func(re *regexp.Regexp) {
		for _, m := range re.FindAllIndex(b, -1) {
			blank(m[0], m[1])
		}
	}
	md := ext == ".md" || ext == ".markdown" || ext == ".mdx" || ext == ""
	if md {
		apply(reValeOff)
		apply(reFront)
		// Fenced code blocks.
		lines := strings.SplitAfter(string(b), "\n")
		pos, in := 0, ""
		for _, ln := range lines {
			t := strings.TrimLeft(ln, " \t")
			fence := ""
			if strings.HasPrefix(t, "```") {
				fence = "```"
			} else if strings.HasPrefix(t, "~~~") {
				fence = "~~~"
			}
			if in != "" || fence != "" {
				blank(pos, pos+len(ln))
			}
			if fence != "" {
				if in == "" {
					in = fence
				} else if in == fence {
					in = ""
				}
			}
			pos += len(ln)
		}
		apply(reInline)
		apply(reLinkDest)
		apply(reComment)
		apply(reAngle)
	}
	apply(reURL)
	return string(b)
}

// lintTerms matches the word lists against the prose in text.
func (l *Linter) lintTerms(text, ext string) []*Alert {
	ix := l.terms
	if ix == nil {
		return nil
	}
	var alerts []*Alert
	for ln, line := range strings.Split(mask(text, ext), "\n") {
		alerts = append(alerts, l.lintPatterns(line, ln+1)...)
		rs := []rune(line)
		for i := 0; i < len(rs); i++ {
			if !isWord(rs[i]) || (i > 0 && (isWord(rs[i-1]) || rs[i-1] == '-' || rs[i-1] == '\'' || rs[i-1] == '’' || rs[i-1] == '_' || rs[i-1] == '.' && i > 1 && isWord(rs[i-2]))) {
				continue
			}
			best := -1
			var hits []term
			seps := 0
			for e := i + 1; e <= len(rs) && e-i <= ix.maxRunes; e++ {
				c := rs[e-1]
				if !(isWord(c) || isSep(c) || c == '.' || c == '\'' || c == '’' || c == '&') {
					break
				}
				if isSep(c) {
					if seps++; seps >= ix.maxWords {
						break
					}
					continue
				}
				if !isWord(c) && c != '.' {
					continue
				}
				if e < len(rs) {
					n := rs[e]
					if isWord(n) || n == '_' {
						continue
					}
					if (n == '-' || n == '\'' || n == '’') && e+1 < len(rs) && isWord(rs[e+1]) {
						continue
					}
				}
				s := string(rs[i:e])
				var found []term
				for _, t := range ix.exact[s] {
					if l.Rules.Get(t.rule).Off == "" {
						found = append(found, t)
					}
				}
				for _, t := range ix.folded[strings.ToLower(s)] {
					if l.Rules.Get(t.rule).Off == "" {
						found = append(found, t)
					}
				}
				if len(found) > 0 {
					best, hits = e, found
				}
			}
			if best < 0 {
				continue
			}
			match := string(rs[i:best])
			for _, t := range hits {
				alerts = append(alerts, l.termAlert(t, match, ln+1, i+1, best))
			}
			i = best - 1
		}
	}
	return alerts
}

// lintPatterns runs the regular-expression entries on one line. Each
// pattern runs only on a line that contains a literal it requires.
func (l *Linter) lintPatterns(line string, ln int) (alerts []*Alert) {
	if strings.TrimSpace(line) == "" {
		return nil
	}
	lowered := strings.ToLower(line)
	rs := []rune(line)
	for _, p := range l.terms.patterns {
		if !p.re.MightMatch(lowered) || l.Rules.Get(p.rule).Off != "" {
			continue
		}
		func() {
			defer func() { recover() }()
			for _, loc := range p.re.FindAllStringIndex(line, -1) {
				if loc[1] <= loc[0] || loc[1] > len(rs) {
					continue
				}
				alerts = append(alerts, l.termAlert(p.term, string(rs[loc[0]:loc[1]]), ln, loc[0]+1, loc[1]))
			}
		}()
	}
	return alerts
}

func (l *Linter) termAlert(t term, match string, line, start, end int) *Alert {
	a := &Alert{Check: t.rule, Match: match, Line: line, Span: [2]int{start, end}, Entry: t.wrong}
	r := l.Rules.Get(t.rule)
	switch {
	case t.sub && r.Fix == "none":
		a.Severity = "suggestion"
		a.Message = fmt.Sprintf("Consider '%s' instead of '%s'.", t.text, match)
	case t.sub:
		a.Severity = "warning"
		a.Message = fmt.Sprintf("Use '%s' instead of '%s'.", t.text, match)
	default:
		a.Severity = "suggestion"
		a.Message = t.text
	}
	if t.sub {
		a.Action.Name, a.Action.Params = "replace", []string{t.text}
	}
	return a
}
