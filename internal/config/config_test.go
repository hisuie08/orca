package config

import (
	"os"
	"testing"

	"github.com/creasty/defaults"
	"github.com/davecgh/go-spew/spew"
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
		{"test2", "/path/to/notexist", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := NewDefaultConfig(tt.path)
			// テストで作成したファイルの削除
			defer os.Remove(tt.path + "/" + OrcaYamlFile)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Create() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Create() succeeded unexpectedly")
			}
		})
	}
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		wantErr bool
	}{
		{"test", "/workspace/orca/testdata", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Load(tt.path)
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

func Test_parseConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"デフォルト", []byte(defaultYaml), false},
		{"フル設定", []byte(fullYaml), false},
		{"一部欠落", []byte(partYaml), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotErr error
			got := newConfig()
			gotErr = defaults.Set(got)
			got.Resolve("/test/dir/orca")
			if err := got.parseConfig(tt.data); err != nil {
				if !tt.wantErr {
					t.Errorf("parseConfig() failed: %v", gotErr)
				}
				return
			}
			if err := defaults.Set(got); err != nil {
				if !tt.wantErr {
					t.Errorf("parseConfig() failed: %v", gotErr)
				}
				return
			}

			if tt.wantErr {
				t.Fatal("parseConfig() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != nil {
				spew.Dump(got)
			}
		})
	}
}

func TestOrcaConfig_Resolve(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		baseDir string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", "/path/to/test1", false},
		{"test2", "/path/to/test2", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newConfig()
			gotErr := c.Resolve(tt.baseDir)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Resolve() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Resolve() succeeded unexpectedly")
			}
			if c.Name != nil {
				t.Logf("\n%v\n", *c.Name)
			}
		})
	}
}
