// Command clauditory-lint checks prose against the Clauditory writing-style
// rules, ranks the sources behind each rule by the owner's hierarchy, and
// fixes the problems that have one mechanical fix.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	lint "github.com/AKushWarrior/clauditory/lint"
)

const usage = `Usage:
  clauditory-lint [flags] [file ...]   lint files, or standard input with no file or "-"
  clauditory-lint rules                list every rule with its domain, fix setting, and status
  clauditory-lint hook                 run as a Claude Code PostToolUse hook
  clauditory-lint version              print the version

Flags:
`

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "rules":
			os.Exit(cmdRules(os.Args[2:]))
		case "hook":
			os.Exit(cmdHook())
		case "version", "--version", "-version":
			fmt.Println("clauditory-lint", lint.Version)
			os.Exit(0)
		}
	}
	os.Exit(cmdLint(os.Args[1:]))
}

type lintFlags struct {
	fix      string
	format   string
	min      string
	fail     string
	ext      string
	spelling bool
	commit   bool
}

func cmdLint(args []string) int {
	fs := flag.NewFlagSet("clauditory-lint", flag.ContinueOnError)
	var f lintFlags
	fs.StringVar(&f.fix, "fix", "", `apply fixes: "safe", or "all" to include fixes that need a look`)
	fs.StringVar(&f.format, "format", "text", `output format: "text" or "json"`)
	fs.StringVar(&f.min, "min", "suggestion", "lowest severity to report: suggestion, warning, or error")
	fs.StringVar(&f.fail, "fail", "error", "lowest severity that makes the exit status 1")
	fs.StringVar(&f.ext, "ext", ".md", "file type to assume for standard input")
	fs.BoolVar(&f.spelling, "spelling", false, "turn on the Red Hat spelling check")
	fs.BoolVar(&f.commit, "commit", false, "the input is a commit message; also check commit-message tells")
	fs.Usage = func() { fmt.Fprint(os.Stderr, usage); fs.PrintDefaults() }
	// Let -fix appear without a value.
	for i, a := range args {
		if a == "-fix" || a == "--fix" {
			args[i] = "--fix=safe"
		}
	}
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if f.fix != "" && f.fix != "safe" && f.fix != "all" {
		fmt.Fprintln(os.Stderr, `--fix takes "safe" or "all"`)
		return 2
	}
	l, err := newLinter(lint.Options{Spelling: f.spelling, Commit: f.commit})
	if err != nil {
		fmt.Fprintln(os.Stderr, "clauditory-lint:", err)
		return 2
	}
	files := fs.Args()
	if len(files) == 0 {
		files = []string{"-"}
	}
	status := 0
	var all []jsonAlert
	for _, path := range files {
		text, ext, err := read(path, f.ext)
		if err != nil {
			fmt.Fprintln(os.Stderr, "clauditory-lint:", err)
			status = 2
			continue
		}
		alerts, fixed, n, err := run(l, text, ext, f.fix)
		if err != nil {
			fmt.Fprintln(os.Stderr, "clauditory-lint:", err)
			status = 2
			continue
		}
		if f.fix != "" {
			if path == "-" {
				fmt.Print(fixed)
			} else if n > 0 {
				if err := os.WriteFile(path, []byte(fixed), 0o644); err != nil {
					fmt.Fprintln(os.Stderr, "clauditory-lint:", err)
					status = 2
					continue
				}
			}
			if n > 0 {
				fmt.Fprintf(os.Stderr, "%s: applied %d %s\n", name(path), n, plural(n, "fix", "fixes"))
			}
		}
		out := os.Stdout
		if f.fix != "" && path == "-" {
			out = os.Stderr
		}
		for _, a := range alerts {
			if !lint.SeverityAtLeast(a.Severity, f.min) {
				continue
			}
			if lint.SeverityAtLeast(a.Severity, f.fail) && status == 0 {
				status = 1
			}
			if f.format == "json" {
				all = append(all, toJSON(name(path), a))
			} else {
				printAlert(out, name(path), a)
			}
		}
	}
	if f.format == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if all == nil {
			all = []jsonAlert{}
		}
		enc.Encode(all)
	}
	return status
}

func newLinter(opt lint.Options) (*lint.Linter, error) {
	rs, err := lint.Load(opt)
	if err != nil {
		return nil, err
	}
	return lint.NewLinter(rs)
}

func read(path, defExt string) (string, string, error) {
	if path == "-" {
		b, err := io.ReadAll(os.Stdin)
		return string(b), defExt, err
	}
	b, err := os.ReadFile(path)
	ext := filepath.Ext(path)
	if ext == "" {
		ext = ".txt"
	}
	return string(b), ext, err
}

func name(path string) string {
	if path == "-" {
		return "<stdin>"
	}
	return path
}

// run lints text and, when asked, fixes it, re-linting after each pass
// until no fix applies.
func run(l *lint.Linter, text, ext, fix string) ([]*lint.Alert, string, int, error) {
	total := 0
	for pass := 0; ; pass++ {
		alerts, err := l.LintText(text, ext)
		if err != nil {
			return nil, text, total, err
		}
		if fix == "" || pass == 5 {
			return alerts, text, total, nil
		}
		fixed, n := lint.Fix(text, alerts, fix)
		if n == 0 {
			return alerts, text, total, nil
		}
		text, total = fixed, total+n
	}
}

func printAlert(w io.Writer, path string, a *lint.Alert) {
	fix := ""
	if a.Fixable("safe") {
		fix = " [fixable]"
	} else if a.Fixable("all") {
		fix = " [fixable with --fix=all]"
	}
	fmt.Fprintf(w, "%s:%d:%d: %s %s: %s%s\n", path, a.Line, a.Span[0], a.Severity,
		a.Check, a.Message, fix)
}

type jsonAlert struct {
	Path     string `json:"path"`
	Line     int    `json:"line"`
	Column   int    `json:"column"`
	End      int    `json:"end"`
	Severity string `json:"severity"`
	Rule     string `json:"rule"`
	Source   string `json:"source"`
	Message  string `json:"message"`
	Match    string `json:"match"`
	Fix      string `json:"fix,omitempty"`
	Link     string `json:"link,omitempty"`
	Entry    string `json:"entry,omitempty"`
}

func toJSON(path string, a *lint.Alert) jsonAlert {
	j := jsonAlert{Path: path, Line: a.Line, Column: a.Span[0], End: a.Span[1], Severity: a.Severity,
		Rule: a.Check, Source: a.Rule.Style,
		Message: a.Message, Match: a.Match, Link: a.Link, Entry: a.Entry}
	if a.Fixable("all") {
		j.Fix, _ = a.Replacement()
		if j.Fix == "" {
			j.Fix = "(delete)"
		}
	}
	return j
}

func cmdRules(args []string) int {
	fs := flag.NewFlagSet("rules", flag.ContinueOnError)
	spelling := fs.Bool("spelling", false, "turn on the Red Hat spelling check")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	rs, err := lint.Load(lint.Options{Spelling: *spelling})
	if err != nil {
		fmt.Fprintln(os.Stderr, "clauditory-lint:", err)
		return 2
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "RULE\tDOMAIN\tRANK\tFIX\tSTATUS")
	for _, r := range rs.Sorted() {
		status := "on"
		if r.Off != "" {
			status = "off: " + r.Off
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", r.Name, r.Domain, r.RankLabel(), r.Fix, status)
	}
	w.Flush()
	return 0
}

// cmdHook reads a Claude Code hook event from standard input. After a Write
// or Edit to a prose file, it lints the file and, if it finds warnings or
// errors, reports them on standard error with exit status 2, which Claude
// Code passes back to the model.
func cmdHook() int {
	if os.Getenv("CLAUDITORY_LINT_HOOK") == "off" {
		return 0
	}
	var ev struct {
		ToolInput struct {
			FilePath string `json:"file_path"`
		} `json:"tool_input"`
	}
	if err := json.NewDecoder(os.Stdin).Decode(&ev); err != nil {
		return 0
	}
	path := ev.ToolInput.FilePath
	switch strings.ToLower(filepath.Ext(path)) {
	case ".md", ".markdown", ".mdx", ".txt", ".rst", ".adoc":
	default:
		return 0
	}
	b, err := os.ReadFile(path)
	if err != nil || strings.Contains(string(b), "clauditory-lint: off") {
		return 0
	}
	l, err := newLinter(lint.Options{})
	if err != nil {
		return 0
	}
	alerts, err := l.LintText(string(b), filepath.Ext(path))
	if err != nil {
		return 0
	}
	var shown []*lint.Alert
	fixable := 0
	for _, a := range alerts {
		if lint.SeverityAtLeast(a.Severity, "warning") {
			shown = append(shown, a)
			if a.Fixable("safe") {
				fixable++
			}
		}
	}
	if len(shown) == 0 {
		return 0
	}
	fmt.Fprintf(os.Stderr, "clauditory-lint found %d style %s in %s.", len(shown), plural(len(shown), "problem", "problems"), path)
	if fixable > 0 {
		cmd := "clauditory-lint"
		if l := os.Getenv("CLAUDITORY_LAUNCHER"); l != "" {
			cmd = l
		}
		fmt.Fprintf(os.Stderr, " Running `%s --fix %s` fixes %d of them.", cmd, path, fixable)
	}
	fmt.Fprintln(os.Stderr, " Fix the rest where the rule applies; skip any that would make the sentence worse.")
	for i, a := range shown {
		if i == 20 {
			fmt.Fprintf(os.Stderr, "...and %d more.\n", len(shown)-20)
			break
		}
		printAlert(os.Stderr, filepath.Base(path), a)
	}
	return 2
}

func plural(n int, one, many string) string {
	if n == 1 {
		return one
	}
	return many
}
