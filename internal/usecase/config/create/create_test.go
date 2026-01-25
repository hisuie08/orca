package create

import (
	"bytes"
	"orca/internal/context"
	"orca/internal/executor"
	"orca/internal/inspector"
	"orca/internal/logger"
	"orca/model/config"
	"orca/model/policy"
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
				WithLog(logger.LogDebug, new(bytes.Buffer))
			fi := inspector.NewFilesystem()
			fe := executor.NewFilesystem(&ctx)
			_, err := (&creator{ctx: &ctx, fe: fe, fi: fi}).
				Create(config.CfgOption{}, false)
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
