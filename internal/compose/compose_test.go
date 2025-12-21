package compose_test

import (
	"orca/infra/inspector"
	"orca/internal/compose"
	"orca/testdata"
	"testing"
)

var fakeInpsector compose.ComposeInspector = inspector.FakeComposeInspector{
	Mock: testdata.TestComposeDir,
}

func TestGetAllCompose(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		orcaRoot string
		i        compose.ComposeInspector
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := compose.GetAllCompose(tt.orcaRoot, fakeInpsector)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAllCompose() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetAllCompose() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != nil {
				if len(got.CollectComposes()) == 0 {
					t.Errorf("compose got 0")
				}
			}
		})
	}
}
