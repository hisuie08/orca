package plan_test

import (
	orca "orca/helper"
	"orca/internal/compose"
	"orca/internal/config"
	"orca/internal/ostools"
	"orca/internal/plan"
	"orca/testdata"
	"os"
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

			cmps, _ := compose.GetAllCompose(testdata.TestPath)
			vol := compose.CollectVolumes(*cmps)
			got := plan.BuildVolumePlan(vol, tt.cfg)
			if tt.wantErr {
				t.Fatal("BuildVolumePlan() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {

				c_, _ := yaml.Marshal(got)
				ostools.CreateFile("./test_volume.yml", c_)
			}
		})
	}
}

func TestPrintVolumePlanTable(t *testing.T) {
	// 検証済みモックデータ
	cfg := &config.OrcaConfig{
		Volume: &config.VolumeConfig{
			EnsurePath: true,
			VolumeRoot: &testdata.TestPath,
		},
	}

	comp, _ := compose.GetAllCompose(testdata.TestPath)
	vol := compose.CollectVolumes(*comp)
	buildPlan := plan.BuildVolumePlan(vol, cfg.Volume)
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
