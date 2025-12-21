package config_test

import (
	"orca/infra/inspector"
	"orca/internal/config"
	"orca/testdata"
	"testing"
)

var fakeReader = inspector.FakeConfigReader{
	Mock: testdata.TestConfigYaml,
}

func TestCreate(t *testing.T) {
	wd := t.TempDir()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", wd, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.Create("")

			if tt.wantErr {
				t.Fatal("Create() succeeded unexpectedly")
			}
			if got.Network.Enabled {

			}
		})
	}
}

func TestLoadConfig(t *testing.T) {
	d := "/workspace/orca/testdata"
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		orca_dir string
		r        config.ConfigReader
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"default", "def", fakeReader, false},
		{"full", "full", fakeReader, false},
		{"part", "part", fakeReader, false},
		{"notexist", "notexist", fakeReader, true},
		{"fin", d, inspector.ConfigFileReader{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := config.LoadConfig(tt.orca_dir, tt.r)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadConfig() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LoadConfig() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
		})
	}
}
