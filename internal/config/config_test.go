package config_test

import (
	"orca/internal/config"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var fullYaml = `
name: "myorca"
volume:
  change_root: true
  make_dirs: true
  root_path: "/workspace/orca/volumeroot"
network:
  create_network: true
  network_name: "orcanet"
`

var partYaml = `
volume:
  change_root: true
  make_dirs: true
`

func TestLoad(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    []byte
		want    *config.OrcaYaml
		wantErr bool
	}{
		// TODO: Add test cases.
		{"full", []byte(fullYaml), &config.OrcaYaml{}, false},
		{"一部欠落", []byte(partYaml), &config.OrcaYaml{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := config.Load(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Load() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Load() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != nil {
				spew.Dump(got)
			}
		})
	}
}

func TestLoadF(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		want    *config.OrcaYaml
		wantErr bool
	}{

		{"成功", "/workspace/orca/testdata", &config.OrcaYaml{}, false},
		{"存在しない", "/workspace/orca/testdata2", &config.OrcaYaml{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := config.LoadF(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadF() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LoadF() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != nil {
				spew.Dump(got)
			}
		})
	}
}
