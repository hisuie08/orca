package compose_test

import (
	"orca/internal/compose"
	"orca/testdata"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestCollectComposes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		orcaRoot string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"test", testdata.TestPath, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := compose.CollectComposes(tt.orcaRoot)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CollectComposes() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CollectComposes() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				spew.Dump(got)
			}
		})
	}
}
