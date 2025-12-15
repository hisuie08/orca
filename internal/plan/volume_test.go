package plan

import (
	"orca/internal/config"
	"orca/internal/ostools"
	"orca/testdata"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestBuildVolumePlan(t *testing.T) {
	testdata.TestConfig.Resolve(testdata.TestPath)
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		orcaRoot string
		cfg      *config.VolumeConfig

		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", testdata.TestPath, testdata.TestConfig.Volume, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := BuildVolumePlan(tt.orcaRoot, tt.cfg)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("BuildVolumePlan() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("BuildVolumePlan() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {

				c_, _ := yaml.Marshal(got)
				ostools.CreateFile(testdata.TestPath+"/plan.yml", c_)
			}
		})
	}
}
