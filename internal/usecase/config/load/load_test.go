package load

import (
	"orca/internal/capability"
	"testing"
)

const fakeYaml string = `name: test
volume:
    volume_root: null
    ensure_path: true
network:
    enabled: true
    internal: false
    name: test_network
`

type fakeInspector struct{ content string }

func (f *fakeInspector) Read(string) ([]byte, error) {
	return []byte(f.content), nil
}
func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name    string
		fi      fsInspector
		wantErr bool
	}{
		{"test", &fakeInspector{content: fakeYaml}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			caps := capability.New().WithRoot(tt.name)
			_, gotErr := loadConfig(&caps, tt.fi)
			if gotErr != nil && !tt.wantErr {
				t.Errorf("LoadConfig() failed: %v", gotErr)
			}
		})
	}
}
