package load

import (
	"orca/internal/context"
	"testing"
)

type fakeInspector struct{}

func (f *fakeInspector) Read(string) ([]byte, error) {
	return []byte{}, nil
}
func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fi      fsInspector
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", &fakeInspector{}, false},
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
