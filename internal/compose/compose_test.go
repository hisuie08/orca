package compose_test

import (
	"orca/internal/compose"
	"orca/testdata"
	"testing"

	"github.com/davecgh/go-spew/spew"
)



func TestParseCompose(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", []byte(testdata.TestDataCompose), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := compose.ParseCompose(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ParseCompose() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ParseCompose() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != nil {
				spew.Dump(got)
			}
		})
	}
}
