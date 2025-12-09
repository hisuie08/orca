package compose_test

import (
	"orca/internal/compose"
	"orca/testdata"
	"testing"
)

func TestLoadMain(t *testing.T) {

	var data string = testdata.TestDataCompose
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data []byte
	}{
		// TODO: Add test cases.
		{"test", []byte(data)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compose.VolumeProcess(tt.data)
		})
	}
}
