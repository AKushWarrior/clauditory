package lint

import (
	"regexp"
	"strings"
)

// commentSyntax describes how one family of languages writes comments.
type commentSyntax struct {
	line      []string    // markers that start a comment running to the end of the line.
	block     [][2]string // start and end markers of block comments.
	quotes    string      // characters that delimit string literals.
	docstring bool        // a triple-quoted string that opens a statement is a comment.
}

var (
	cFamily    = &commentSyntax{line: []string{"//"}, block: [][2]string{{"/*", "*/"}}, quotes: "\"'`"}
	hashFamily = &commentSyntax{line: []string{"#"}, quotes: "\"'"}
	python     = &commentSyntax{line: []string{"#"}, quotes: "\"'", docstring: true}
	dashFamily = &commentSyntax{line: []string{"--"}, block: [][2]string{{"--[[", "]]"}, {"{-", "-}"}, {"/*", "*/"}}, quotes: "\"'"}
)

// codeComments maps a source file's extension to its comment syntax. Files
// with these extensions are linted on their comments alone.
var codeComments = map[string]*commentSyntax{
	".go": cFamily, ".c": cFamily, ".h": cFamily, ".cc": cFamily, ".cpp": cFamily, ".cxx": cFamily,
	".hpp": cFamily, ".cs": cFamily, ".java": cFamily, ".js": cFamily, ".jsx": cFamily,
	".mjs": cFamily, ".cjs": cFamily, ".ts": cFamily, ".tsx": cFamily, ".mts": cFamily,
	".cts": cFamily, ".kt": cFamily, ".kts": cFamily, ".rs": cFamily, ".swift": cFamily,
	".scala": cFamily, ".dart": cFamily, ".php": cFamily, ".proto": cFamily, ".css": cFamily,
	".scss": cFamily, ".less": cFamily, ".m": cFamily, ".zig": cFamily,
	".py": python, ".pyi": python,
	".sh": hashFamily, ".bash": hashFamily, ".zsh": hashFamily, ".fish": hashFamily,
	".rb": hashFamily, ".pl": hashFamily, ".pm": hashFamily, ".r": hashFamily,
	".ps1": hashFamily, ".ex": hashFamily, ".exs": hashFamily, ".jl": hashFamily,
	".nix": hashFamily, ".tf": hashFamily, ".mk": hashFamily,
	".lua": dashFamily, ".sql": dashFamily, ".hs": dashFamily,
}

// IsCode reports whether files with this extension are linted on their
// comments alone.
func IsCode(ext string) bool { return codeComments[strings.ToLower(ext)] != nil }

// commentView returns text with everything but comment prose blanked to
// spaces: code, string literals, comment markers, and a leading shebang.
// Line breaks stay, so positions map to the same lines and columns, and code
// between comments separates them into paragraphs.
func commentView(text string, cs *commentSyntax) string {
	src := []byte(text)
	out := make([]byte, len(src))
	for i, c := range src {
		if c == '\n' {
			out[i] = '\n'
		} else {
			out[i] = ' '
		}
	}
	keep := func(lo, hi int) {
		for k := lo; k < hi; k++ {
			if src[k] != '\n' && src[k] != '\r' {
				out[k] = src[k]
			}
		}
	}
	at := func(i int, s string) bool { return strings.HasPrefix(text[i:], s) }
	lineStart := func(i int) bool {
		for k := i - 1; k >= 0 && src[k] != '\n'; k-- {
			if src[k] != ' ' && src[k] != '\t' {
				return false
			}
		}
		return true
	}
	// keepComment keeps a comment's text, without decoration such as the
	// leading stars of a block comment or a run of comment characters.
	keepComment := func(lo, hi int) {
		for lo < hi {
			end := strings.IndexByte(text[lo:hi], '\n')
			if end < 0 {
				end = hi
			} else {
				end += lo
			}
			s := lo
			for s < end && strings.ContainsRune(" \t*/#-!;=", rune(src[s])) {
				s++
			}
			keep(s, end)
			lo = end + 1
		}
	}
	i := 0
	if at(0, "#!") {
		i = strings.IndexByte(text, '\n')
		if i < 0 {
			return string(out)
		}
	}
outer:
	for i < len(src) {
		if cs.docstring && (at(i, `"""`) || at(i, "'''")) {
			q := text[i : i+3]
			end := strings.Index(text[i+3:], q)
			stop := len(src)
			if end >= 0 {
				stop = i + 3 + end
			}
			if lineStart(i) {
				keepComment(i+3, stop)
			}
			i = stop + 3
			continue
		}
		for _, b := range cs.block {
			if at(i, b[0]) {
				end := strings.Index(text[i+len(b[0]):], b[1])
				stop := len(src)
				if end >= 0 {
					stop = i + len(b[0]) + end
				}
				keepComment(i+len(b[0]), stop)
				i = stop + len(b[1])
				continue outer
			}
		}
		for _, m := range cs.line {
			if at(i, m) {
				end := strings.IndexByte(text[i:], '\n')
				stop := len(src)
				if end >= 0 {
					stop = i + end
				}
				keepComment(i+len(m), stop)
				i = stop
				continue outer
			}
		}
		if c := src[i]; strings.IndexByte(cs.quotes, c) >= 0 {
			for i++; i < len(src) && src[i] != c; i++ {
				if src[i] == '\\' {
					i++
				} else if src[i] == '\n' && c != '`' {
					break
				}
			}
		}
		i++
	}
	// Trim each line's trailing blanks, so a line of code becomes empty and
	// separates the comments around it into paragraphs.
	// Blank code spans and URLs, which aren't prose.
	for _, re := range []*regexp.Regexp{reInline, reURL} {
		for _, m := range re.FindAllIndex(out, -1) {
			for k := m[0]; k < m[1]; k++ {
				if out[k] != '\n' {
					out[k] = ' '
				}
			}
		}
	}
	lines := strings.Split(string(out), "\n")
	for k, ln := range lines {
		lines[k] = strings.TrimRight(ln, " ")
	}
	return strings.Join(lines, "\n")
}
