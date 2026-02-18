package pinit

import (
	"bytes"
	"orca/internal/capability"
	"orca/internal/inspector"
	"orca/internal/usecase/config"
	. "orca/model/config"
	"path/filepath"

	"orca/model/policy"
	"orca/model/policy/log"
	"testing"
)

func fakeCaps(root string) initProcessCapability {
	caps := capability.New().WithRoot(root).WithPolicy(policy.Real).
		WithLog(log.LogDetail, new(bytes.Buffer))
	return &caps
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
			caps := fakeCaps(root)
			New(caps).Run(InitOption{
				CfgOption:   CfgOption{Name: filepath.Base(root)},
				WriteOption: config.WriteOption{NoCreate: tC.nocreate}})
			exist := inspector.NewFilesystem().FileExists(caps.OrcaYamlFile())
			if exist != tC.expected {
				t.Errorf("files in root expected %t but found %t",
					tC.expected, exist)
			}
		})
	}
}
