package compose_test

import (
	"fmt"
	"orca/internal/compose"
	"orca/testdata"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestVolumeSpec_NeedsOrcaOverlay(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		spec compose.VolumeSpec
		want bool
	}{
		{"default", testdata.TestVolSpecDefault, true},
		{"local", testdata.TestVolSpecLocal, true},
		{"external", testdata.TestVolSpecExternal, false},
		{"cache", testdata.TestVolSpecCache, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			got := tt.spec.NeedsOrcaOverlay()
			if tt.want != got {
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
