package config

import (
	"orca/internal/context"
	"orca/internal/executor"
	"orca/model/policy"
	"testing"
)

var _ executor.FileSystem = (*fakeFileWriter)(nil)

type fakeFileWriter struct {
	context.WithPolicy
	Files []string
	executor.FileSystem
}

func (f *fakeFileWriter) WriteFile(path string, data []byte) error {
	if f.Policy().AllowSideEffect() {
		f.Files = append(f.Files, path)
	}
	return nil
}
func newFakeWriter(ctx CreateConfigContext) *fakeFileWriter {
	return &fakeFileWriter{WithPolicy: ctx, Files: []string{}}
}

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
			writer := newFakeWriter(&ctx)
			_, err := createConfig(&ctx, writer, "")
			if err != nil {
				t.Error(err)
			}
			if len(writer.Files) != tC.expect {
				t.Errorf("expected %d file but created %d files", tC.expect, len(writer.Files))
			}
		})
	}
}
