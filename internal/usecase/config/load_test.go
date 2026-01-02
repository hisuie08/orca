package config

import (
	"orca/internal/context"
	"orca/internal/inspector"
	"testing"
)

var _ inspector.FileSystem = (*fakeFileReader)(nil)

type fakeFileReader struct {
	inspector.FileSystem
}

func (f *fakeFileReader) Read(string) ([]byte, error) {
	return []byte{}, nil
}

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fi      inspector.FileSystem
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", &fakeFileReader{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.New().WithRoot(tt.name)
			_, gotErr := loadConfig(&ctx, tt.fi)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadConfig() failed: %v", gotErr)
				}
				return
			}

		})
	}
}
