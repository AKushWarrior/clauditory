package lint

import (
	"os"
	"strings"
	"testing"
)

func load(t *testing.T) *Ruleset {
	t.Helper()
	rs, err := Load(Options{})
	if err != nil {
		t.Fatal(err)
	}
	return rs
}

// The hierarchy turns off the worse-ranked side of each conflict.
func TestOpposition(t *testing.T) {
	rs := load(t)
	cases := map[string]string{
		"AP.NoSerialComma": "Oxford.SerialComma",
		"Chicago.Numbers":  "NYT.SmallNumbers",
		"Apple.Webpage":    "Chicago.WebPage",
		"Microsoft.AMPM":   "NYT.Time",
	}
	for loser, winner := range cases {
		if rs.Rules[winner].Off != "" {
			t.Errorf("%s should be on", winner)
		}
		if !strings.Contains(rs.Rules[loser].Off, winner) {
			t.Errorf("%s should be off in favor of %s, got %q", loser, winner, rs.Rules[loser].Off)
		}
	}
}

func TestRanks(t *testing.T) {
	rs := load(t)
	want := map[string]int{
		"House.EmDash":        0,
		"Oxford.SerialComma":  1,
		"NYT.Percent":         1,
		"RedHat.TermsErrors":  1,
		"AP.WhoForPeople":     2,
		"Chicago.Usage":       2,
		"Apple.ProductNames":  2,
		"Chicago.PronounCase": 3,
		"AP.Words":            3,
		"Microsoft.Suspended": 5,
		"RedHat.PassiveVoice": unranked,
	}
	for name, rank := range want {
		if got := rs.Rules[name].Rank; got != rank {
			t.Errorf("%s: rank %d, want %d", name, got, rank)
		}
	}
}

func TestFixSafe(t *testing.T) {
	l, err := NewLinter(load(t))
	if err != nil {
		t.Fatal(err)
	}
	in, _ := os.ReadFile("testdata/fix-input.md")
	want, _ := os.ReadFile("testdata/fix-safe.md")
	text := string(in)
	for i := 0; i < 5; i++ {
		alerts, err := l.LintText(text, ".md")
		if err != nil {
			t.Fatal(err)
		}
		fixed, n := Fix(text, alerts, "safe")
		if n == 0 {
			break
		}
		text = fixed
	}
	if text != string(want) {
		t.Errorf("fixed text:\n%s\nwant:\n%s", text, want)
	}
}

func newTestLinter(t *testing.T) *Linter {
	t.Helper()
	l, err := NewLinter(load(t))
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func checks(t *testing.T, l *Linter, text string) map[string]int {
	t.Helper()
	alerts, err := l.LintText(text, ".md")
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]int{}
	for _, a := range alerts {
		got[a.Check]++
	}
	return got
}

// Every dash, and every run of hyphens standing in for one, is an error;
// compounds, suspended hyphens, code, and horizontal rules are not.
func TestDashes(t *testing.T) {
	l := newTestLinter(t)
	for _, s := range []string{"One -- two.", "One--two.", "One---two.", "One - two.", "One -two.",
		"One- two.", "One — two.", "One–two.", "One ― two.", "Run lint --fix now."} {
		if checks(t, l, s+"\n")["House.EmDash"] != 1 {
			t.Errorf("%q: want one House.EmDash alert", s)
		}
	}
	for _, s := range []string{"A well-known bug.", "Pre- and post-war data.", "The `a -- b` code.", "Text.\n\n---\n\nMore text."} {
		if n := checks(t, l, s+"\n")["House.EmDash"]; n != 0 {
			t.Errorf("%q: want no House.EmDash alert, got %d", s, n)
		}
	}
}

// The extracted word lists match whole words and phrases only, and skip code.
func TestTerms(t *testing.T) {
	l := newTestLinter(t)
	ix := l.terms
	if len(ix.folded)+len(ix.exact) < 1000 || len(ix.patterns) < 100 {
		t.Fatalf("word lists look empty: %d folded, %d exact, %d patterns", len(ix.folded), len(ix.exact), len(ix.patterns))
	}
	m := mask("Use `irregardless` in code, see https://example.com/irregardless.\n```\nirregardless\n```\n", ".md")
	if strings.Contains(m, "irregardless") {
		t.Errorf("mask left code or a URL unmasked: %q", m)
	}
}

// Words inside double quotes are masked, so a skill that quotes a banned
// phrase doesn't trip the rule it describes.
func TestQuotes(t *testing.T) {
	l := newTestLinter(t)
	if n := checks(t, l, "Delete \"I hope this helps\" at the end.\n")["House.StockPhrases"]; n != 0 {
		t.Errorf("quoted phrase: want no House.StockPhrases alert, got %d", n)
	}
	if n := checks(t, l, "I hope this helps.\n")["House.StockPhrases"]; n == 0 {
		t.Errorf("unquoted phrase: want a House.StockPhrases alert")
	}
	m := maskQuotes("say \"a b a\" now")
	if m != "say \"b c d\" now" {
		t.Errorf("maskQuotes: got %q", m)
	}
}

// A paragraph whose sentences all run about the same length is flagged; one
// with a short sentence in it isn't.
func TestRhythm(t *testing.T) {
	l := newTestLinter(t)
	s := "The parser reads each line of the input file and splits it on commas into fields. " +
		"The validator checks every field against the schema and records each failure it finds. " +
		"The writer formats the clean records as rows and appends them to the output table. " +
		"The reporter counts the failures by type and prints a summary when the run is done.\n"
	if checks(t, l, s)["House.Rhythm"] != 1 {
		t.Errorf("flat paragraph: want one House.Rhythm alert")
	}
	v := strings.Replace(s, "The writer formats", "Nothing else runs. The writer formats", 1)
	if n := checks(t, l, v)["House.Rhythm"]; n != 0 {
		t.Errorf("varied paragraph: want no House.Rhythm alert, got %d", n)
	}
}
