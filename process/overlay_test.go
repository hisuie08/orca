package process_test

import (
	"orca/internal/compose"
	"orca/internal/ostools"
	"orca/process"
	"orca/testdata"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
)

func TestApplyLocalBind(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		volume   *compose.VolumeSpec
		bindPath string
	}{
		{"default", &testdata.TestVolSpecDefault, "/workspace/volumeroot"},
		{"local", &testdata.TestVolSpecLocal, "/workspace/volumeroot"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := process.ApplyLocalBind(tt.volume, tt.bindPath)
			if got != nil {
				c, _ := yaml.Marshal(got)
				ostools.CreateFile("./test_volume.yml", c)
			}
		})
	}
}

func TestVolumeSpec_ApplyExternal(t *testing.T) {
	tests := []struct {
		name   string
		volume *compose.VolumeSpec
	}{
		{"default", &testdata.TestVolSpecDefault},
		{"local", &testdata.TestVolSpecLocal},
		{"external", &testdata.TestVolSpecExternal},
		{"cache", &testdata.TestVolSpecCache},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := process.ApplyExternal(tt.volume)
			if got != nil {
				spew.Dump(got)
			}
		})
	}
}
