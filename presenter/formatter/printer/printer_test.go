package printer

import (
	"testing"
)

func TestPrinter_PrintGrid(t *testing.T) {
	testRows := [][]string{{"abcd", "def"}, {"ghij", "klmn"}}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		// Named input parameters for target function.
		rows [][]string
	}{
		// TODO: Add test cases.
		{"test", testRows},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
		})
	}
}
