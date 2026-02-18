package dump

import (
	"bytes"
	"orca/internal/capability"
	"orca/internal/inspector"
	"orca/model/compose"
	"orca/model/config"
	"orca/model/plan"
	"orca/model/policy"
	"orca/model/policy/log"
	"path/filepath"
	"testing"
)

func TestDumpCompose(t *testing.T) {
	tests := []struct {
		name        string
		p           policy.ExecPolicy
		wantFiles   int
		wantWritten int
	}{
		{name: "real policy", p: policy.Real, wantFiles: 1,
			wantWritten: 1},
		{name: "dry policy", p: policy.Dry, wantFiles: 0,
			wantWritten: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := compose.ComposeMap{"a": &compose.ComposeSpec{}}
			tmpdir := t.TempDir()
			caps := capability.New().WithRoot(tmpdir).
				WithConfig(&config.OrcaConfig{}).
				WithPolicy(tt.p).WithLog(log.LogDetail, new(bytes.Buffer))
			dumper := DotOrcaDumper(&caps, false)
			got, err := dumper.DumpComposes(cm)
			if len(got) != tt.wantWritten {
				t.Errorf("expected written 1 but got %d", len(got))
			}
			if tt.p == policy.Real {
				fi := inspector.NewFilesystem()
				files, fe := fi.Files(caps.OrcaDir())
				if fe != nil {
					t.Fatal(fe)
				}
				if len(files) != tt.wantFiles {
					t.Errorf("%d files in dirs", len(files))
				}
				if got[0] != filepath.Join(caps.OrcaDir(), "compose.a.yml") {
					t.Errorf("unexpected path: %s", got[0])
				}
			}
			if err != nil {
				t.Errorf("has %v error", err)
			}
		})
	}
}

func TestDumpPlan(t *testing.T) {
	pl := plan.OrcaPlan{}
	tests := []struct {
		name       string
		policy     policy.ExecPolicy
		wantExists bool
	}{
		{name: "real policy", policy: policy.Real, wantExists: true},
		{name: "dry policy", policy: policy.Dry, wantExists: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			caps := capability.New().WithRoot(t.TempDir()).
				WithConfig(&config.OrcaConfig{}).
				WithPolicy(tt.policy).WithLog(log.LogDetail, new(bytes.Buffer))
			dumper := DotOrcaDumper(&caps, false)
			path, err := dumper.DumpPlan(pl)
			if err != nil {
				t.Error(err)
			}
			if inspector.NewFilesystem().FileExists(path) != tt.wantExists {
				t.Errorf("%s file create failed", path)
			}
		})
	}
}

func TestForceDump(t *testing.T) {
	tests := []struct {
		name    string
		force   bool
		wantErr bool
	}{
		{name: "force=true", force: true, wantErr: false},
		{name: "force=false", force: false, wantErr: true},
	}
	for _, tC := range tests {
		t.Run(tC.name, func(t *testing.T) {
			caps := capability.New().WithRoot(t.TempDir()).
				WithConfig(&config.OrcaConfig{}).
				WithPolicy(policy.Real).WithLog(log.LogDetail, new(bytes.Buffer))
			dumper := DotOrcaDumper(&caps, tC.force)
			cm := compose.ComposeMap{"a": &compose.ComposeSpec{}}
			// first dump
			if _, err := dumper.DumpComposes(cm); err != nil {
				t.Fatal("dump failed on the first try")
			}
			// second dump
			_, err := dumper.DumpComposes(cm)
			if tC.wantErr && err == nil {
				t.Error("expected error but got nil")
			}
			if !tC.wantErr && err != nil {
				t.Error(err)
			}
		})
	}
}
