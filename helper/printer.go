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
func PrintTable(w io.Writer, title string, headers []string, rows [][]string) {
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

	fmt.Fprintf(w, "%s\n", title)
	printRow(w, headers, widths)
	printSeparator(w, widths)
	for _, r := range rows {
		printRow(w, r, widths)
	}
}

func printRow(w io.Writer, cols []string, widths []int) {
	for i, c := range cols {
		vis := visibleLen(c)
		pad := max(widths[i] - vis, 0)
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
