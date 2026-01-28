package pinit

import (
	"bytes"
	"orca/internal/context"
	"orca/internal/inspector"
	"orca/internal/usecase/config"
	. "orca/model/config"
	"path/filepath"

	"orca/model/policy"
	"orca/model/policy/log"
	"testing"
)

func fakeCtx(root string) initProcessContext {
	ctx := context.New().WithRoot(root).WithPolicy(policy.Real).
		WithLog(log.LogDebug, new(bytes.Buffer))
	return &ctx
}

func TestCreateOrNocreate(t *testing.T) {
	testCases := []struct {
		desc     string
		nocreate bool
		expected bool
	}{
		{desc: "Nocreate", nocreate: true, expected: false},
		{desc: "Create", nocreate: false, expected: true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			root := t.TempDir()
			ctx := fakeCtx(root)
			New(ctx).Run(InitOption{
				CfgOption:   CfgOption{Name: filepath.Base(root)},
				WriteOption: config.WriteOption{NoCreate: tC.nocreate}})
			exist := inspector.NewFilesystem().FileExists(ctx.OrcaYamlFile())
			if exist != tC.expected {
				t.Errorf("files in root expected %t but found %t",
					tC.expected, exist)
			}
		})
	}
}
