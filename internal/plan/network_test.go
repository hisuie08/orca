package plan

import (
	orca "orca/helper"
	"orca/internal/compose"
	"orca/internal/config"
	"os"
	"testing"
)

func TestBuildNetworkPlan(t *testing.T) {
	cc := compose.ComposeMap{}
	cfg := config.ResolvedNetwork{
		Enabled: true, Name: "test", Internal: false}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cc      compose.ComposeMap
		cfg     config.ResolvedNetwork
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", cc, cfg, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := BuildNetworkPlan(tt.cc, &tt.cfg)
			if tt.wantErr {
				t.Fatal("BuildNetworkPlan() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got.SharedName != tt.cfg.Name {
				t.Errorf("expected name %s but got %s", tt.cfg.Name, got.SharedName)
			}
		})
	}
}

func TestPrintNetworkPlan(t *testing.T) {

	printer := orca.NewPrinter(os.Stdout, orca.Colorizer{Enabled: true})
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		p       NetworkPlan
		printer orca.Printer
	}{
		// TODO: Add test cases.
		{"test", NetworkPlan{}, *printer},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintNetworkPlan(tt.p, &tt.printer)
		})
	}
}
