package plan_test

import (
	orca "orca/helper"
	"orca/internal/compose"
	"orca/internal/config"
	"orca/internal/plan"
	"orca/test/fake"
	"orca/testdata"
	"os"
	"testing"
)

func TestBuildVolumePlan(t *testing.T) {
	v := t.TempDir()
	cv := []compose.CollectedVolume{}
	cfg := &config.ResolvedVolume{VolumeRoot: &v, EnsurePath: true}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		orcaRoot string
		vol      []compose.CollectedVolume
		cfg      *config.ResolvedVolume

		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", testdata.TestPath, cv, cfg, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plan.BuildVolumePlan(tt.vol, tt.cfg, fake.FakeDocker{})
			if tt.wantErr {
				t.Fatal("BuildVolumePlan() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got == nil {
				t.Error("Error got nil")
			}
		})
	}
}

func TestPrintVolumePlanTable(t *testing.T) {
	// 検証済みモックデータ
	cfg := &config.ResolvedConfig{
		Volume: config.ResolvedVolume{
			EnsurePath: true,
			VolumeRoot: &testdata.TestPath,
		},
	}

	vol := []compose.CollectedVolume{}
	buildPlan := plan.BuildVolumePlan(vol, &cfg.Volume, fake.FakeDocker{})
	printer := orca.NewPrinter(os.Stdout, orca.Colorizer{Enabled: true})
	tests := []struct {
		name  string
		plans []plan.VolumePlan
		p     *orca.Printer
	}{
		// TODO: Add test cases.
		{"test", buildPlan, printer},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plan.PrintVolumePlanTable(tt.plans, printer)
		})
	}
}
