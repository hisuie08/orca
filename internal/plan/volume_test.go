package plan

import (
	"orca/internal/config"
	"orca/internal/ostools"
	"orca/testdata"
	"testing"

	"gopkg.in/yaml.v3"
)

var volumeRoot = "/workspace/orca/testdir"
var testConfig = &config.OrcaConfig{
	Volume: &config.VolumeConfig{
		EnsurePath: true,
		VolumeRoot: &volumeRoot,
	},
}

func TestCollectVolumes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		orcaRoot string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"test", testdata.TestPath, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := collectVolumes(tt.orcaRoot)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CollectVolumes() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CollectVolumes() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				c_, _ := yaml.Marshal(got)
				ostools.CreateFile(testdata.TestPath+"/collect.yml", c_)
				g := groupVolumes(got)
				g_, _ := yaml.Marshal(g)
				ostools.CreateFile(testdata.TestPath+"/group.yml", g_)
				b := buildPlan(g, testConfig.Volume)
				b_, _ := yaml.Marshal(b)
				// 面倒なのでファイルに出力してテスト
				ostools.CreateFile(testdata.TestPath+"/plan.yml", b_)
			}
		})
	}
}
