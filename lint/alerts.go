package lint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"
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

	Rule *Rule `json:"-"`
}

var severityRank = map[string]int{"suggestion": 0, "warning": 1, "error": 2}

// SeverityAtLeast reports whether sev is at or above min.
func SeverityAtLeast(sev, min string) bool { return severityRank[sev] >= severityRank[min] }

// FindVale locates the Vale binary.
func FindVale() (string, error) {
	if p := os.Getenv("CLAUDITORY_VALE"); p != "" {
		return p, nil
	}
	if exe, err := os.Executable(); err == nil {
		p := filepath.Join(filepath.Dir(exe), "vale")
		if st, err := os.Stat(p); err == nil && !st.IsDir() {
			return p, nil
		}
	}
	if p, err := exec.LookPath("vale"); err == nil {
		return p, nil
	}
	if home, err := os.UserHomeDir(); err == nil {
		p := filepath.Join(home, "go", "bin", "vale")
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}
	return "", fmt.Errorf("can't find Vale; install it with `go install github.com/vale-cli/vale/v3/cmd/vale@v3.26.0` or set CLAUDITORY_VALE")
}

// Linter runs Vale with the resolved rules.
type Linter struct {
	Rules *Ruleset
	Vale  string
	ini   string
}

// NewLinter prepares the Vale config.
func NewLinter(rs *Ruleset, vale string) (*Linter, error) {
	ini, err := rs.WriteConfig()
	if err != nil {
		return nil, err
	}
	return &Linter{Rules: rs, Vale: vale, ini: ini}, nil
}

// LintText lints text as if it were a file with the given extension.
func (l *Linter) LintText(text, ext string) ([]*Alert, error) {
	if ext == "" {
		ext = ".md"
	}
	cmd := exec.Command(l.Vale, "--no-global", "--config="+l.ini, "--output=JSON", "--ext="+ext)
	cmd.Stdin = strings.NewReader(text)
	var out, errb bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &errb
	runErr := cmd.Run()
	var res map[string][]*Alert
	if err := json.Unmarshal(out.Bytes(), &res); err != nil {
		msg := strings.TrimSpace(errb.String() + out.String())
		if runErr != nil {
			return nil, fmt.Errorf("vale: %v: %s", runErr, msg)
		}
		return nil, fmt.Errorf("vale returned output that isn't JSON: %s", msg)
	}
	var alerts []*Alert
	for _, as := range res {
		alerts = append(alerts, as...)
	}
	return l.resolve(alerts), nil
}

// resolve attaches rules and keeps, for each overlapping group of alerts on
// one topic, only the best-ranked alert. Two fixes that would write different
// text over the same span count as one topic too.
func (l *Linter) resolve(alerts []*Alert) []*Alert {
	var kept []*Alert
	for _, a := range alerts {
		a.Rule = l.Rules.Get(a.Check)
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
