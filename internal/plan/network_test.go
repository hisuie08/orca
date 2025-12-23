package plan

import (
	orca "orca/helper"
	"orca/internal/compose"
	"orca/internal/config"
	"orca/ostools"
	"orca/test/fake"
	"orca/testdata"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

var fakeComposeInspector = fake.ComposeInspector

func TestBuildNetworkPlan(t *testing.T) {
	netname := "orcanet"
	testcfg := &config.ResolvedConfig{
		Network: config.ResolvedNetwork{
			Enabled:  true,
			Internal: false,
			Name:     netname,
		},
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		orcaRoot string
		cfg      *config.ResolvedNetwork
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"test", testdata.TestPath, &testcfg.Network, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			comp, _ := compose.GetAllCompose(testdata.TestPath, fakeComposeInspector)
			got := BuildNetworkPlan(comp.CollectComposes(), tt.cfg)
			if tt.wantErr {
				t.Fatal("BuildNetworkPlan() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				c_, _ := yaml.Marshal(got)
				ostools.CreateFile(t.TempDir()+"/test_network.yml", c_)
			}
		})
	}
}

func TestPrintNetworkPlan(t *testing.T) {

	netname := "orcanet"
	testcfg := &config.ResolvedConfig{
		Network: config.ResolvedNetwork{
			Enabled:  true,
			Internal: false,
			Name:     netname,
		},
	}

	printer := orca.NewPrinter(os.Stdout, orca.Colorizer{Enabled: true})
	comp, _ := compose.GetAllCompose(testdata.TestPath, fakeComposeInspector)
	buildPlan := BuildNetworkPlan(comp.CollectComposes(), &testcfg.Network)
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		p       NetworkPlan
		printer orca.Printer
	}{
		// TODO: Add test cases.
		{"test", buildPlan, *printer},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintNetworkPlan(tt.p, &tt.printer)
		})
	}
}
