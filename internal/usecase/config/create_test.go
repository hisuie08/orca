package config

import (
	"orca/internal/context"
	"orca/internal/inspector"
	"orca/model/policy"
	"testing"
)

func Test(t *testing.T) {
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
			ctx := context.New().WithRoot(dir).WithPolicy(tC.policy)
			var _ context.WithRoot = &ctx

			_, err := CreateConfig(&ctx, "")
			if err != nil {
				t.Error(err)
			}
			c, err := inspector.NewFilesystem().Files(dir)
			if err != nil {
				t.Error(err)
			}
			if len(c) != tC.expect {
				t.Errorf("expected %d file but created %d files", tC.expect, len(c))
			}
		})
	}
}
func TestConfigCreator(t *testing.T) {

}
