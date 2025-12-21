package config_test

import (
	"orca/internal/config"
	"os"
	"testing"
)

var defaultYaml = `
name:
volume:
    volume_root:
    ensure_path: true
network:
    enabled: true 
    internal: false 
    name: 
`
var fullYaml = `
name: MyOrca
volume:
    volume_root: /test/volume
    ensure_path: true
network:
    enabled: true 
    internal: false 
    name: custom_network
`
var partYaml = `
volume:
    volume_root: 
    ensure_path: true
`

func TestCreate(t *testing.T) {
	wd, _ := os.Getwd()
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
			got, gotErr := config.Create("", t.TempDir()+tt.path)
			// テストで作成したファイルの削除
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Create() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Create() succeeded unexpectedly")
			}
			if got.Network.Enabled {

			}
		})
	}
}

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		orca_dir string
		r        config.ConfigReader
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"default", "default", config.FakeConfigReader{Want: defaultYaml}, false},
		{"full", "def", config.FakeConfigReader{Want: fullYaml}, false},
		{"part", "part", config.FakeConfigReader{Want: partYaml}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := config.LoadConfig(tt.orca_dir, tt.r)
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
			if got.Name == nil {
				t.Errorf("name got nil")
			}
		})
	}
}
