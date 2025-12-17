package compose_test

import (
	"orca/internal/compose"
	"orca/internal/ostools"
	"orca/testdata"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
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
			cmps, gotErr := compose.ComposeMap(tt.orcaRoot)
			got := compose.CollectComposes(*cmps)
			vol := compose.CollectVolumes(*cmps)
			nets := compose.CollectNetworks(*cmps)
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
				c,_:=yaml.Marshal(got)
				v,_:=yaml.Marshal(vol)
				n,_:=yaml.Marshal(nets)
				ostools.CreateFile("./test_compose.yml",c)
				ostools.CreateFile("./test_volume.yml",v)
				ostools.CreateFile("./test_network.yml",n)
			}
		})
	}
}
