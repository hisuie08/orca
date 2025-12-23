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
	config.LoadConfig("def", fake.ConfigReader)
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		orcaRoot string
		cfg      *config.ResolvedVolume

		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", testdata.TestPath, (*config.ResolvedVolume)(testdata.TestConfig.Volume), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cmps, _ := compose.GetAllCompose(testdata.TestPath, fake.ComposeInspector)
			vol := cmps.CollectVolumes()
			got := plan.BuildVolumePlan(vol, tt.cfg, fake.DockerInspector)
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

	comp, _ := compose.GetAllCompose(testdata.TestPath, fake.ComposeInspector)
	vol := comp.CollectVolumes()
	buildPlan := plan.BuildVolumePlan(vol, &cfg.Volume, fake.DockerInspector)
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
