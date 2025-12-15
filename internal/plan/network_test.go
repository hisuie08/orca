package plan

import (
	"io"
	orca "orca/helper"
	"orca/internal/config"
	"orca/internal/ostools"
	"orca/testdata"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestBuildNetworkPlan(t *testing.T) {
	netname := "orcanet"
	testcfg := &config.OrcaConfig{
		Network: &config.NetworkConfig{
			Enabled:  true,
			Internal: false,
			Name:     &netname,
		},
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		orcaRoot string
		cfg      *config.NetworkConfig
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"test", testdata.TestPath, testcfg.Network, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := BuildNetworkPlan(tt.orcaRoot, tt.cfg)
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

func TestPrintNetworkPlan(t *testing.T) {

	netname := "orcanet"
	testcfg := &config.OrcaConfig{
		Network: &config.NetworkConfig{
			Enabled:  true,
			Internal: false,
			Name:     &netname,
		},
	}
	buildPlan, _ := BuildNetworkPlan(testdata.TestPath, testcfg.Network)
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		p NetworkPlan
		w io.Writer
	}{
		// TODO: Add test cases.
		{"test", *buildPlan, os.Stdout},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintNetworkPlan(tt.p, tt.w, &orca.Colorizer{Enabled: true})
		})
	}
}
