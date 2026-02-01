package create

import (
	"bytes"
	"orca/internal/context"
	"orca/internal/executor"
	"orca/internal/inspector"
	"orca/model/config"
	"orca/model/policy"
	"orca/model/policy/log"
	"testing"
)

func TestCreateConfig(t *testing.T) {
	testCases := []struct {
		desc   string
		policy policy.ExecPolicy
		expect int
	}{
		{desc: "real", policy: policy.Real, expect: 1},
		{desc: "dry", policy: policy.Dry, expect: 0},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			dir := t.TempDir()
			ctx := context.New().WithRoot(dir).WithPolicy(tC.policy).
				WithLog(log.LogDetail, new(bytes.Buffer))
			fi := inspector.NewFilesystem()
			fe := executor.NewFilesystem(&ctx)
			c := &creator{ctx: &ctx, fe: fe, fi: fi}
			cfg := c.Create(config.CfgOption{})
			if err := c.Write(cfg, true); err != nil {
				t.Error(err)
			}
			written, err := fi.Files(dir)
			if err != nil {
				t.Error(err)
			}
			if len(written) != tC.expect {
				t.Errorf("expected %d file but created %d files", tC.expect, len(written))
			}
		})
	}
}
