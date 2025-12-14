package plan_test

import (
	"orca/internal/config"
	"orca/internal/ostools"
	"orca/internal/plan"
	"orca/testdata"
	"testing"

	"gopkg.in/yaml.v3"
)

var netname = "orcanet"
var testConfig = &config.OrcaConfig{
	Network: &config.NetworkConfig{
		Enabled:  true,
		Internal: false,
		Name:     &netname,
	},
}

func TestBuildNetworkPlan(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		orcaRoot string
		cfg      *config.NetworkConfig
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"test", testdata.TestPath, testConfig.Network, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := plan.BuildNetworkPlan(tt.orcaRoot, tt.cfg)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("BuildNetworkPlan() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("BuildNetworkPlan() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				c_, _ := yaml.Marshal(got)
				ostools.CreateFile(testdata.TestPath+"/network.yml", c_)
			}
		})
	}
}
