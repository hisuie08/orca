package compose_test

import (
	"orca/infra/applier"
	"orca/internal/compose"
	"orca/test/fake"

	"testing"
)

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
			got, gotErr := compose.GetAllCompose(tt.orcaRoot, fake.ComposeInspector)
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

func TestComposeMap_DumpAllComposes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		orcaRoot string
		i        compose.ComposeInspector
		// Named input parameters for target function.
		cw      applier.ComposeWriter
		wantErr bool
	}{
		{"test",t.TempDir(),fake.ComposeInspector,fake.ComposeWriter,false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := compose.GetAllCompose(tt.orcaRoot, tt.i)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got, gotErr := m.DumpAllComposes(tt.cw)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("DumpAllComposes() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("DumpAllComposes() succeeded unexpectedly")
			}
			
			// TODO: update the condition below to compare got with tt.want.
			if len(got)!= len(m.CollectComposes()){
				t.Errorf("DumpAllComposes() = got invalid")
			}
			t.Log(got)
		})
	}
}
