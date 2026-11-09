package lint

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/AKushWarrior/clauditory/lint/internal/vale/core"
	"github.com/AKushWarrior/clauditory/lint/internal/vale/lint"
)

// Alert is one Vale alert, with the rule's place in the hierarchy attached.
type Alert struct {
	Action struct {
		Name   string   `json:"Name"`
		Params []string `json:"Params"`
	} `json:"Action"`
	Span     [2]int `json:"Span"`
	Check    string `json:"Check"`
	Link     string `json:"Link"`
	Message  string `json:"Message"`
	Severity string `json:"Severity"`
	Match    string `json:"Match"`
	Line     int    `json:"Line"`

	Rule  *Rule  `json:"-"`
	Entry string `json:"-"` // the word-list entry that matched, for generated rules
}

var severityRank = map[string]int{"suggestion": 0, "warning": 1, "error": 2}

// SeverityAtLeast reports whether sev is at or above min.
func SeverityAtLeast(sev, min string) bool { return severityRank[sev] >= severityRank[min] }

// Linter runs the embedded Vale engine with the resolved rules.
type Linter struct {
	Rules *Ruleset
	cfg   *core.Config
	terms *termIndex
}

// NewLinter prepares the Vale config. Vale is compiled into this program,
// so no separate Vale install is needed.
func NewLinter(rs *Ruleset) (*Linter, error) {
	ini, err := rs.WriteConfig()
	if err != nil {
		return nil, err
	}
	cfg, err := core.ReadPipeline(&core.CLIFlags{Path: ini, IgnoreGlobal: true, Output: "JSON", NoExit: true}, false)
	if err != nil {
		return nil, fmt.Errorf("vale config: %w", err)
	}
	terms, err := loadTerms()
	if err != nil {
		return nil, err
	}
	return &Linter{Rules: rs, cfg: cfg, terms: terms}, nil
}

// LintText lints text as if it were a file with the given extension.
func (l *Linter) LintText(text, ext string) (alerts []*Alert, err error) {
	defer func() {
		if r := recover(); r != nil {
			alerts, err = nil, fmt.Errorf("vale: %v", r)
		}
	}()
	if ext == "" {
		ext = ".md"
	}
	code := false
	if cs := codeComments[strings.ToLower(ext)]; cs != nil {
		text, ext, code = commentView(text, cs), ".txt", true
	}
	defer func() {
		// The comment view keeps code's indentation, which spacing rules
		// misread as extra spaces between words.
		if code && err == nil {
			kept := alerts[:0]
			for _, a := range alerts {
				if !strings.HasSuffix(a.Check, ".Spacing") {
					kept = append(kept, a)
				}
			}
			alerts = kept
		}
	}()
	vl, err := lint.NewLinter(l.cfg)
	if err != nil {
		return nil, err
	}
	text = maskQuotes(text)
	files, err := vl.LintStringAs(text, "input"+ext)
	if err != nil {
		return nil, fmt.Errorf("vale: %w", err)
	}
	for _, f := range files {
		for _, va := range f.Alerts {
			if va.Hide {
				continue
			}
			a := &Alert{Check: va.Check, Link: va.Link, Message: va.Message, Severity: va.Severity,
				Match: va.Match, Line: va.Line}
			a.Action.Name, a.Action.Params = va.Action.Name, va.Action.Params
			if len(va.Span) == 2 {
				a.Span = [2]int{va.Span[0], va.Span[1]}
			}
			alerts = append(alerts, a)
		}
	}
	alerts = dropNumberSplits(alerts, text)
	alerts = append(alerts, l.lintTerms(text, ext)...)
	return l.resolve(alerts), nil
}

// resolve attaches rules and keeps, for each overlapping group of alerts on
// one topic, only the best-ranked alert. Two fixes that would write different
// text over the same span count as one topic too.
func (l *Linter) resolve(alerts []*Alert) []*Alert {
	var kept []*Alert
	for _, a := range alerts {
		a.Rule = l.Rules.Get(a.Check)
		if a.Rule.Level != "" {
			a.Severity = a.Rule.Level
		}
		if a.Rule.Off == "" && !a.Rule.Ignore[strings.ToLower(strings.TrimSpace(a.Match))] &&
			(a.Rule.Skip == nil || !a.Rule.Skip.MatchString(a.Match)) {
			kept = append(kept, a)
		}
	}
	sort.SliceStable(kept, func(i, j int) bool {
		a, b := kept[i], kept[j]
		if a.Line != b.Line {
			return a.Line < b.Line
		}
		if a.Rule.Rank != b.Rule.Rank {
			return a.Rule.Rank < b.Rule.Rank
		}
		if severityRank[a.Severity] != severityRank[b.Severity] {
			return severityRank[a.Severity] > severityRank[b.Severity]
		}
		return a.Span[0] < b.Span[0]
	})
	var out []*Alert
	for _, a := range kept {
		dup := false
		for _, b := range out {
			if b.Line != a.Line || a.Span[0] > b.Span[1] || b.Span[0] > a.Span[1] {
				continue
			}
			if a.Check == b.Check || sameWords(a.Match, b.Match) ||
				(a.Rule.Topic != "" && a.Rule.Topic == b.Rule.Topic) ||
				(a.Rule.Domain == b.Rule.Domain && a.Action.Name == "replace" && b.Action.Name == "replace") {
				dup = true
				break
			}
		}
		if !dup {
			out = append(out, a)
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Line != out[j].Line {
			return out[i].Line < out[j].Line
		}
		return out[i].Span[0] < out[j].Span[0]
	})
	return out
}

// sameWords reports whether two matches cover the same words, ignoring case,
// spacing, and punctuation, so that two sources flagging one phrase count once.
func sameWords(a, b string) bool {
	norm := func(s string) string {
		return strings.ToLower(strings.TrimFunc(s, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) }))
	}
	na, nb := norm(a), norm(b)
	if na == "" && nb == "" {
		return strings.TrimSpace(a) != "" && strings.TrimSpace(a) == strings.TrimSpace(b)
	}
	return na != "" && na == nb
}

// Fixable reports whether the alert's fix runs at the given level
// ("safe" or "all").
func (a *Alert) Fixable(level string) bool {
	switch a.Rule.Fix {
	case "safe":
	case "review":
		if level != "all" {
			return false
		}
	default:
		return false
	}
	_, ok := a.Replacement()
	return ok
}

var digitWords = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

var timeRe = regexp.MustCompile(`^(\d{1,2})(?::(\d{2}))? ?([AaPp])\.?[Mm]\.?$`)

// Replacement returns the text the fix writes over the match.
func (a *Alert) Replacement() (string, bool) {
	m := a.Match
	switch a.Rule.Fixer {
	case "time":
		sm := timeRe.FindStringSubmatch(strings.TrimSpace(m))
		if sm == nil {
			return "", false
		}
		ap := strings.ToLower(sm[3]) + ".m."
		if sm[1] == "12" && (sm[2] == "" || sm[2] == "00") {
			if ap == "p.m." {
				return "noon", true
			}
			return "midnight", true
		}
		if sm[2] == "" || sm[2] == "00" {
			return sm[1] + " " + ap, true
		}
		return sm[1] + ":" + sm[2] + " " + ap, true
	case "spell-number":
		if len(m) > 0 && m[0] >= '1' && m[0] <= '9' {
			return digitWords[m[0]-'0'] + m[1:], true
		}
		return "", false
	}
	switch a.Action.Name {
	case "replace":
		if len(a.Action.Params) == 1 {
			return a.Action.Params[0], true
		}
	case "remove":
		return "", true
	case "edit":
		p := a.Action.Params
		if len(p) == 0 {
			return "", false
		}
		switch p[0] {
		case "regex":
			if len(p) < 3 {
				return "", false
			}
			re, err := regexp.Compile(p[1])
			if err != nil {
				return "", false
			}
			out := re.ReplaceAllString(m, p[2])
			return out, out != m
		case "trim":
			if len(p) < 2 {
				return "", false
			}
			return strings.Trim(m, p[1]), true
		case "trim_left":
			if len(p) < 2 {
				return "", false
			}
			return strings.TrimLeft(m, p[1]), true
		case "trim_right":
			if len(p) < 2 {
				return "", false
			}
			return strings.TrimRight(m, p[1]), true
		}
	}
	return "", false
}

// Fix applies the fixable alerts to text and reports how many it applied.
// Alerts that overlap an alert already applied on the same line are skipped;
// a later pass picks them up.
func Fix(text string, alerts []*Alert, level string) (string, int) {
	lines := strings.SplitAfter(text, "\n")
	byLine := map[int][]*Alert{}
	for _, a := range alerts {
		if a.Fixable(level) {
			byLine[a.Line] = append(byLine[a.Line], a)
		}
	}
	applied := 0
	for ln, as := range byLine {
		if ln < 1 || ln > len(lines) {
			continue
		}
		sort.Slice(as, func(i, j int) bool { return as[i].Span[0] > as[j].Span[0] })
		line := []rune(lines[ln-1])
		lastStart := len(line) + 1
		for _, a := range as {
			start, end := locate(line, a)
			if start < 0 || end >= lastStart {
				continue
			}
			rep, _ := a.Replacement()
			rep = matchCase(line, start, string(line[start:end]), rep, a.Rule)
			nl := append([]rune{}, line[:start]...)
			nl = append(nl, []rune(rep)...)
			line = append(nl, line[end:]...)
			lastStart = start
			applied++
		}
		lines[ln-1] = string(line)
	}
	return strings.Join(lines, ""), applied
}

// locate finds the alert's match in the line, as rune offsets [start, end).
func locate(line []rune, a *Alert) (int, int) {
	m := []rune(a.Match)
	s := a.Span[0] - 1
	if s >= 0 && s+len(m) <= len(line) && string(line[s:s+len(m)]) == a.Match {
		return s, s + len(m)
	}
	str := string(line)
	best, bestDist := -1, 1<<30
	for i := 0; ; {
		j := strings.Index(str[i:], a.Match)
		if j < 0 {
			break
		}
		pos := len([]rune(str[:i+j]))
		d := pos - s
		if d < 0 {
			d = -d
		}
		if d < bestDist {
			best, bestDist = pos, d
		}
		i += j + 1
		if i >= len(str) {
			break
		}
	}
	if best < 0 {
		return -1, -1
	}
	return best, best + len(m)
}

// matchCase capitalizes a lowercase replacement when it replaces a
// capitalized word at the start of a sentence.
func matchCase(line []rune, start int, old, rep string, r *Rule) string {
	if rep == "" || old == "" || r.Topic == "term" || r.Topic == "internet" {
		return rep
	}
	o, n := []rune(old), []rune(rep)
	if !unicode.IsUpper(o[0]) || !unicode.IsLower(n[0]) {
		return rep
	}
	i := start - 1
	for i >= 0 && unicode.IsSpace(line[i]) {
		i--
	}
	if i >= 0 && !strings.ContainsRune(".!?:#*->\"(", line[i]) && !unicode.IsDigit(line[i]) {
		return rep
	}
	n[0] = unicode.ToUpper(n[0])
	return string(n)
}

// dropNumberSplits removes sentence-start alerts that Vale raises inside a
// section number such as 4.3.5, which its sentence splitter breaks at each
// period.
var reCiteAbbr = regexp.MustCompile(`(?i)\b(pp?|ch|chap|sec|secs|fig|figs|no|nos|vol|vols|para|art|eq)\.\s*$`)

func dropNumberSplits(alerts []*Alert, text string) []*Alert {
	lines := strings.Split(text, "\n")
	out := alerts[:0]
	for _, a := range alerts {
		if a.Check == "NYT.SentenceStartFigure" && a.Line >= 1 && a.Line <= len(lines) {
			rs := []rune(lines[a.Line-1])
			i := a.Span[0] - 2
			if i >= 0 && i < len(rs) && (rs[i] == '.' || unicode.IsDigit(rs[i])) {
				continue
			}
			// A page, section, or chapter citation such as "p. 29" is
			// not a sentence that starts with a figure.
			if reCiteAbbr.MatchString(string(rs[:max(0, a.Span[0]-1)])) {
				continue
			}
			// A numbered heading such as "## 1. Instructions" is an
			// enumerator, not a sentence that starts with a figure.
			if strings.HasPrefix(strings.TrimLeft(string(rs), " "), "#") {
				continue
			}
		}
		out = append(out, a)
	}
	return out
}
