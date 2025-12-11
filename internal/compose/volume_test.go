package compose_test

import (
	"fmt"
	"orca/internal/compose"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var defaultVol = compose.VolumeSpec{Name: "docker_defaultvol"}
var localvol = compose.VolumeSpec{
	Name:   "docker_localvol",
	Driver: "local",
	DriverOpts: map[string]string{
		"type": "none", "o": "bind", "device": "/src/test"}}
var externalvol = compose.VolumeSpec{
	Name: "externalvol", External: true,
}
var cachevol = compose.VolumeSpec{
	Name:   "docker_localvol",
	Driver: "local",
	DriverOpts: map[string]string{
		"type": "tmpfs"}}

func TestVolumeSpec_NeedsOrcaOverlay(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		spec compose.VolumeSpec
		want bool
	}{
		{"default",defaultVol,true},
		{"local",localvol,true},
{"external",externalvol,false},
{"cache",cachevol,false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			got := tt.spec.NeedsOrcaOverlay()
			if tt.want!=got{
				t.Errorf("something wrong,")
			}
				if got {
					fmt.Println("need to create")
					tt.spec.ApplyLocalBind("/workspace/volumeroot")
					spew.Dump(tt.spec)
				} else {
					fmt.Println("dont need create")
				}
		})
	}
}
