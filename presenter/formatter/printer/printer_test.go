package printer

import (
	"testing"
)

func TestPrinter_PrintGrid(t *testing.T) {
	testRows := [][]string{{"abcd", "def"}, {"ghi", "jklmn"}}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		// Named input parameters for target function.
		headers []string
		rows    [][]string
	}{
		// TODO: Add test cases.
		{"test", []string{"COL_1", "COL_2"}, testRows},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PTable(tt.name, tt.headers, tt.rows)
		})
	}
}
