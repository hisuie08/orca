package config

import (
	"orca/internal/context"
	fakeexecutor "orca/internal/executor/fake"
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
			ctx := context.New().WithRoot(dir).WithPolicy(tC.policy)
			writer := &fakeexecutor.FakeFilesystem{
				Issued: []string{}, Done: []string{},
				AllowSideEffect: tC.policy.AllowSideEffect()}
			_, err := createConfig(&ctx, writer, "")
			if err != nil {
				t.Error(err)
			}
			written := len(writer.Done)
			if written != tC.expect {
				t.Errorf("expected %d file but created %d files", tC.expect, written)
			}
		})
	}
}
