package printer

import (
	"fmt"
	"orca/presenter/formatter"
	"strings"
)

// 列を揃えて1行を出力する内部関数
func pRow(cols []string, widths []int) string {
	result := []string{}
	for i, c := range cols {
		vis := formatter.VisibleLen(c)
		pad := max(widths[i]-vis, 0)
		result = append(result, fmt.Sprintf(c),
			fmt.Sprintf(strings.Repeat(" ", pad)), fmt.Sprintf("  "))
	}
	result = append(result, fmt.Sprintf("\n"))
	return strings.Join(result, "")
}

// pSeparator は区切り線を出力する内部関数
func pSeparator(widths []int) string {
	result := []string{}
	for _, wth := range widths {
		result = append(result, fmt.Sprintf(strings.Repeat("-", wth), "  "))
	}
	result = append(result, fmt.Sprintf("\n"))
	return strings.Join(result, "")
}

// PrintTable はタイトルヘッダー付きの表出力
func PTable(title string, headers []string, rows [][]string) string {
	// 各列の最大幅
	widths := make([]int, len(headers))
	// headersを基準に初期値
	for i, h := range headers {
		widths[i] =
			formatter.VisibleLen(h)
	}
	// 内容に応じて拡張
	for _, cols := range rows {
		for i, c := range cols {
			widths[i] = max(widths[i], formatter.VisibleLen(c))
		}
	}
	result := []string{}
	result = append(result, fmt.Sprintf("%s\n\n", title),
		pRow(headers, widths), pSeparator(widths))
	for _, r := range rows {
		result = append(result, pRow(r, widths))
	}
	return strings.Join(result, "")
}
