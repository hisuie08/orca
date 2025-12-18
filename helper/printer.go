package orca

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"
)

var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func visibleLen(s string) int {
	return utf8.RuneCountInString(stripANSI(s))
}

func stripANSI(s string) string {
	return ansiRegexp.ReplaceAllString(s, "")
}

func printRow(w io.Writer, cols []string, widths []int) {
	for i, c := range cols {
		vis := visibleLen(c)
		pad := max(widths[i]-vis, 0)
		fmt.Fprint(w, c)
		fmt.Fprint(w, strings.Repeat(" ", pad))
		fmt.Fprint(w, "  ")
	}
	fmt.Fprintln(w)
}

func printSeparator(w io.Writer, widths []int) {
	for _, wth := range widths {
		fmt.Fprint(w, strings.Repeat("-", wth), "  ")
	}
	fmt.Fprintln(w)
}

type Printer struct {
	W io.Writer
	C Colorizer
}

func NewPrinter(w io.Writer, c Colorizer) *Printer {
	return &Printer{w, c}
}

func (p *Printer) PrintDRY(s string) {
	label := p.C.Blue("[DRY-RUN]: ")
	fmt.Fprintf(p.W, "%s %s", label, s)
}

func (p *Printer) Printf(format string, a ...any) {
	fmt.Fprintf(p.W, format, a...)
}

// PrintTable はタイトルヘッダー付きの表出力
func (p *Printer) PrintTable(title string, headers []string, rows [][]string) {
	widths := make([]int, len(headers))

	for i, h := range headers {
		widths[i] = visibleLen(h)
	}
	for _, cols := range rows {
		for i, c := range cols {
			if visibleLen(c) > widths[i] {
				widths[i] = visibleLen(c)
			}
		}
	}

	fmt.Fprintf(p.W, "%s\n", title)
	printRow(p.W, headers, widths)
	printSeparator(p.W, widths)
	for _, r := range rows {
		printRow(p.W, r, widths)
	}
}

// PrintGrid はタイトルやヘッダー無しのシンプルな二次元配列を出力
func (p *Printer) PrintGrid(rows [][]string) {
	widths := make([]int, len(rows))
	for _, cols := range rows {
		for i, c := range cols {
			if visibleLen(c) > widths[i] {
				widths[i] = visibleLen(c)
			}
		}
	}
	for _, r := range rows {
		printRow(p.W, r, widths)
	}
}
