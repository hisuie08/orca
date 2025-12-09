package process_test

import(
	"orca/internal/process"
	"testing"
)

func TestCommonProcess(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path string
	}{
		// TODO: Add test cases.
		{"test",""},
		{"test2","/workspace/orca/testdata"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			process.CommonProcess(tt.path)
		})
	}
}
