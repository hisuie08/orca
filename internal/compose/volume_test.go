package compose

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var defaultVol = VolumeSpec{Name: "docker_defaultvol"}
var localvol = VolumeSpec{
	Name:   "docker_localvol",
	Driver: "local",
	DriverOpts: map[string]string{
		"type": "none", "o": "bind", "device": "/src/test"}}
var externalvol = VolumeSpec{
	Name: "externalvol", External: true,
}
var cachevol = VolumeSpec{
	Name:   "docker_localvol",
	Driver: "local",
	DriverOpts: map[string]string{
		"type": "tmpfs"}}

func TestVolumeSpec_NeedsOrcaOverlay(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		spec VolumeSpec
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
					spew.Dump(tt.spec)
				} else {
					fmt.Println("dont need create")
				}
		})
	}
}

func TestVolumeSpec_ApplyLocalBind(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		volume VolumeSpec
		volume_root string
	}{
		// TODO: Add test cases.
		{"default",defaultVol,"/workspace/volumeroot"},
		{"local",localvol,"/workspace/volumeroot"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.volume.ApplyLocalBind(tt.volume_root)
			// TODO: update the condition below to compare got with tt.want.
			if got!=nil {
				spew.Dump(got)
			}
		})
	}
}
