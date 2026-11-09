// Package lint holds the clauditory-lint rule registry, the embedded Vale
// styles, and the logic that ranks, filters, and fixes Vale's alerts.
package lint

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed rules.yml
var registryYAML []byte

//go:embed rules.gen.yml
var generatedYAML []byte

//go:embed all:styles
var stylesFS embed.FS

// Version is the linter's version. Raise it with the same semantic-versioning
// rule the skills use, and record the change in CHANGELOG.md.
const Version = "2.2.0"

// unranked is the rank of a source the domain's order doesn't list.
const unranked = 99

// RuleMeta is one entry in rules.yml.
type RuleMeta struct {
	Domain  string   `yaml:"domain"`
	Topic   string   `yaml:"topic"`
	Opposes []string `yaml:"opposes"`
	Off     string   `yaml:"off"`
	Fix     string   `yaml:"fix"`
	Fixer   string   `yaml:"fixer"`
	Ignore  []string `yaml:"ignore"`
	Skip    string   `yaml:"skip"`
	Literal bool     `yaml:"literal"`
	Level   string   `yaml:"level"`
}

// Registry is the parsed rules.yml.
type Registry struct {
	Owner  string              `yaml:"owner"`
	Order  map[string][]string `yaml:"order"`
	Styles map[string]struct {
		Domain string `yaml:"domain"`
		Level  string `yaml:"level"`
	} `yaml:"styles"`
	Rules map[string]RuleMeta `yaml:"rules"`
}

// Rule is a rule after the hierarchy has been applied.
type Rule struct {
	Name    string
	Style   string
	Domain  string
	Topic   string
	Fix     string
	Fixer   string
	Rank    int
	Off     string
	Level   string
	Ignore  map[string]bool
	Skip    *regexp.Regexp
	literal bool
	opp     []string
}

// RankLabel describes the rule's rank for the rules listing.
func (r *Rule) RankLabel() string {
	switch r.Rank {
	case 0:
		return "owner"
	case unranked:
		return "-"
	}
	return fmt.Sprint(r.Rank)
}

// Options change how the registry resolves.
type Options struct {
	Spelling bool // turn on RedHat.Spelling
	Commit   bool // the text is a commit message; turn on AITellsCommits
}

// Ruleset is the resolved set of rules.
type Ruleset struct {
	Rules map[string]*Rule
	reg   *Registry
}

// Load parses the registry and resolves every rule in the embedded styles.
func Load(opt Options) (*Ruleset, error) {
	var reg Registry
	if err := yaml.Unmarshal(registryYAML, &reg); err != nil {
		return nil, fmt.Errorf("rules.yml: %w", err)
	}
	var gen struct {
		Rules map[string]RuleMeta `yaml:"rules"`
	}
	if err := yaml.Unmarshal(generatedYAML, &gen); err != nil {
		return nil, fmt.Errorf("rules.gen.yml: %w", err)
	}
	for k, v := range gen.Rules {
		if _, dup := reg.Rules[k]; !dup {
			reg.Rules[k] = v
		}
	}
	rs := &Ruleset{Rules: map[string]*Rule{}, reg: &reg}
	styles, err := fs.ReadDir(stylesFS, "styles")
	if err != nil {
		return nil, err
	}
	for _, st := range styles {
		if !st.IsDir() {
			continue
		}
		if _, ok := reg.Styles[st.Name()]; !ok {
			return nil, fmt.Errorf("style %s has no entry in rules.yml", st.Name())
		}
		files, _ := fs.ReadDir(stylesFS, "styles/"+st.Name())
		for _, f := range files {
			if f.IsDir() || !strings.HasSuffix(f.Name(), ".yml") {
				continue
			}
			name := st.Name() + "." + strings.TrimSuffix(f.Name(), ".yml")
			rs.Rules[name] = rs.newRule(name, st.Name())
		}
	}
	for name, m := range reg.Rules {
		if _, ok := rs.Rules[name]; ok {
			continue
		}
		if !m.Literal {
			return nil, fmt.Errorf("rules.yml names %s, which has no rule file", name)
		}
		r := rs.newRule(name, strings.SplitN(name, ".", 2)[0])
		r.literal = true
		rs.Rules[name] = r
	}
	if opt.Spelling {
		rs.Rules["RedHat.Spelling"].Off = ""
	}
	for name, r := range rs.Rules {
		if r.Style == "AITellsCommits" && !opt.Commit && r.Off == "" {
			r.Off = "Applies only to commit messages; pass --commit."
		}
		_ = name
	}
	rs.applyOpposition()
	return rs, nil
}

func (rs *Ruleset) newRule(name, style string) *Rule {
	m := rs.reg.Rules[name]
	r := &Rule{Name: name, Style: style, Domain: m.Domain, Topic: m.Topic, Fix: m.Fix,
		Fixer: m.Fixer, Off: m.Off, opp: m.Opposes, Ignore: map[string]bool{}}
	if m.Skip != "" {
		r.Skip = regexp.MustCompile(m.Skip)
	}
	for _, w := range m.Ignore {
		r.Ignore[strings.ToLower(w)] = true
	}
	if r.Domain == "" {
		r.Domain = rs.reg.Styles[style].Domain
	}
	r.Level = m.Level
	if r.Level == "" {
		r.Level = rs.reg.Styles[style].Level
	}
	if r.Fix == "" {
		r.Fix = "none"
	}
	r.Rank = rs.rank(style, r.Domain)
	return r
}

// rank places a source in a domain: 0 for the owner, 1 and up for the listed
// sources in order, and unranked for a source the domain doesn't list.
func (rs *Ruleset) rank(style, domain string) int {
	if style == rs.reg.Owner {
		return 0
	}
	for i, s := range rs.reg.Order[domain] {
		if s == style {
			return i + 1
		}
	}
	return unranked
}

// applyOpposition turns off the worse-ranked side of each pair of rules that
// demand opposite things. On a tie the rule that declares the opposition
// yields.
func (rs *Ruleset) applyOpposition() {
	names := make([]string, 0, len(rs.Rules))
	for n := range rs.Rules {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		a := rs.Rules[n]
		for _, bn := range a.opp {
			b := rs.Rules[bn]
			if a.Off != "" || b.Off != "" {
				continue
			}
			loser, winner := a, b
			if b.Rank > a.Rank {
				loser, winner = b, a
			}
			loser.Off = fmt.Sprintf("Conflicts with %s, which wins in %s.", winner.Name, winner.Domain)
		}
	}
}

// Disabled lists the rules that are off, sorted by name.
func (rs *Ruleset) Disabled() []*Rule {
	var out []*Rule
	for _, r := range rs.Rules {
		if r.Off != "" {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

// Sorted returns every rule, sorted by domain, rank, and name.
func (rs *Ruleset) Sorted() []*Rule {
	var out []*Rule
	for _, r := range rs.Rules {
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.Domain != b.Domain {
			return a.Domain < b.Domain
		}
		if a.Rank != b.Rank {
			return a.Rank < b.Rank
		}
		return a.Name < b.Name
	})
	return out
}

// Get returns a rule, or a default for a name the registry doesn't know.
func (rs *Ruleset) Get(name string) *Rule {
	if r, ok := rs.Rules[name]; ok {
		return r
	}
	style := strings.SplitN(name, ".", 2)[0]
	return &Rule{Name: name, Style: style, Domain: "wording", Fix: "none", Rank: unranked, Ignore: map[string]bool{}}
}

// Order lists the styles from the owner down, for the Vale config.
func (rs *Ruleset) Order() []string {
	seen := map[string]bool{rs.reg.Owner: true}
	out := []string{rs.reg.Owner}
	for _, d := range []string{"grammar", "wording", "technical"} {
		for _, s := range rs.reg.Order[d] {
			if !seen[s] {
				seen[s] = true
				out = append(out, s)
			}
		}
	}
	var rest []string
	for s := range rs.reg.Styles {
		if !seen[s] {
			rest = append(rest, s)
		}
	}
	sort.Strings(rest)
	out = append(out, rest...)
	var have []string
	for _, s := range out {
		if _, err := fs.Stat(stylesFS, "styles/"+s); err == nil {
			have = append(have, s)
		}
	}
	return have
}

// WriteConfig extracts the styles into a cache directory and writes a Vale
// config that turns off every disabled rule. It returns the config's path.
func (rs *Ruleset) WriteConfig() (string, error) {
	h := sha256.New()
	h.Write(registryYAML)
	h.Write(generatedYAML)
	fs.WalkDir(stylesFS, "styles", func(p string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			b, _ := stylesFS.ReadFile(p)
			h.Write([]byte(p))
			h.Write(b)
		}
		return nil
	})
	base, err := os.UserCacheDir()
	if err != nil || base == "" {
		base = os.TempDir()
	}
	dir := filepath.Join(base, "clauditory-lint", hex.EncodeToString(h.Sum(nil))[:16])
	if _, err := os.Stat(filepath.Join(dir, "styles")); err != nil {
		tmp := dir + ".tmp" + fmt.Sprint(os.Getpid())
		err := fs.WalkDir(stylesFS, "styles", func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			dst := filepath.Join(tmp, p)
			if d.IsDir() {
				return os.MkdirAll(dst, 0o755)
			}
			b, err := stylesFS.ReadFile(p)
			if err != nil {
				return err
			}
			return os.WriteFile(dst, b, 0o644)
		})
		if err != nil {
			return "", err
		}
		os.MkdirAll(filepath.Dir(dir), 0o755)
		if err := os.Rename(tmp, dir); err != nil {
			os.RemoveAll(tmp)
			if _, statErr := os.Stat(filepath.Join(dir, "styles")); statErr != nil {
				return "", err
			}
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "StylesPath = %s\nMinAlertLevel = suggestion\n\n[*]\nBasedOnStyles = %s\n",
		filepath.Join(dir, "styles"), strings.Join(rs.Order(), ", "))
	for _, r := range rs.Disabled() {
		if !r.literal {
			fmt.Fprintf(&b, "%s = NO\n", r.Name)
		}
	}
	ini := filepath.Join(dir, "vale-"+shortHash(b.String())+".ini")
	if err := os.WriteFile(ini, []byte(b.String()), 0o644); err != nil {
		return "", err
	}
	return ini, nil
}

func shortHash(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])[:12]
}
