package config

import (
	"orca/internal/context"
	"orca/internal/inspector"
	"testing"
)

var _ inspector.FileSystem = (*fakeFilesystem)(nil)

type fakeFilesystem struct {
	inspector.FileSystem
}

func (f *fakeFilesystem) Read(string) ([]byte, error) {
	return []byte{}, nil
}

func fakeCtx(name string) *context.Context {
	ctx := context.New().WithRoot(name)
	return &ctx
}
func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ctx     LoadConfigContext
		fi      inspector.FileSystem
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", fakeCtx("test"), &fakeFilesystem{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := LoadConfig(tt.ctx, tt.fi)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadConfig() failed: %v", gotErr)
				}
				return
			}

		})
	}
}
