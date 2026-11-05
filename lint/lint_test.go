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
	want := map[string]string{
		"House.EmDash":        "owner",
		"Oxford.SerialComma":  "primary",
		"NYT.Percent":         "primary",
		"RedHat.TermsErrors":  "primary",
		"AP.WhoForPeople":     "secondary",
		"Chicago.Usage":       "secondary",
		"Apple.ProductNames":  "secondary",
		"Chicago.PronounCase": "tertiary",
		"AP.Words":            "tertiary",
		"Microsoft.Suspended": "tertiary",
		"RedHat.PassiveVoice": "unranked",
	}
	for name, tier := range want {
		if got := rs.Rules[name].Tier(); got != tier {
			t.Errorf("%s: tier %s, want %s", name, got, tier)
		}
	}
}

func TestFixSafe(t *testing.T) {
	vale, err := FindVale()
	if err != nil {
		t.Skip(err)
	}
	l, err := NewLinter(load(t), vale)
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
