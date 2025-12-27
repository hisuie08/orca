package config_test

import (
	"orca/infra/applier"
	"orca/infra/inspector"
	"orca/internal/config"
	"os"
	"path/filepath"
	"testing"
)

var _ inspector.ConfigReader = (*fakeCfgReader)(nil)

// FakeConfigReader テスト用
type fakeCfgReader struct {
	root   string
	result string
	ErrOn  bool
}

func (r fakeCfgReader) Root() string {
	return r.root
}
func (f fakeCfgReader) Read() ([]byte, error) {
	if f.ErrOn {
		return nil, os.ErrNotExist
	}
	return []byte(f.result), nil
}

func TestCreate(t *testing.T) {
	root := t.TempDir()
	ok, ng := filepath.Join(root, "ok"), filepath.Join(root, "ng")
	os.Mkdir(ok, 0755)
	os.Mkdir(ng, 0755)
	cfg := applier.DryApplier.ConfigWriter(root)
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path        string
		clusterName string
		wantErr     bool
	}{
		// TODO: Add test cases.
		{"ok", ok, "", false},
		{"ng", ng, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, e := config.Create(tt.clusterName, cfg)

			if e != nil {
				t.Fatal("Failed %w", e)
			}

			if tt.wantErr {
				t.Fatal("Create() succeeded unexpectedly")
			}
			t.Log(p, "\n")

		})
	}
}

func TestLoadConfig(t *testing.T) {
	a, b :=
		fakeCfgReader{result: "", ErrOn: false},
		fakeCfgReader{result: "", ErrOn: true}
	tests := []struct {
		name string // description of this test case
		dir  string
		// Named input parameters for target function.
		r       inspector.ConfigReader
		wantErr bool
	}{
		// TODO: Add test cases.
		{"pass", "a", a, false},
		{"fail", "b", b, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := config.LoadConfig(tt.r)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadConfig() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LoadConfig() succeeded unexpectedly")
			}
		})
	}
}
