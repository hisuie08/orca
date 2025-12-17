package plan

import (
	"io"
	orca "orca/helper"
	"orca/internal/compose"
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

			comp, _ := compose.CollectComposes(testdata.TestPath)
			nets := compose.CollectNetworks(comp)
			got := BuildNetworkPlan(nets, tt.cfg)
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
	comp, _ := compose.CollectComposes(testdata.TestPath)
	nets := compose.CollectNetworks(comp)
	buildPlan := BuildNetworkPlan(nets, testcfg.Network)
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
