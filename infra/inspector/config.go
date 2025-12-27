package inspector

import (
	"orca/consts"
	"os"
	"path/filepath"
)

var _ ConfigReader = (*configFile)(nil)

type ConfigReader interface {
	Root() string
	Read() ([]byte, error)
}
type configFile struct {
	orcaRoot string
}

func ConfigFile(orcaRoot string) *configFile {
	return &configFile{orcaRoot: orcaRoot}
}
func (r configFile) Root() string {
	return r.orcaRoot
}

func (r configFile) Read() ([]byte, error) {
	target := filepath.Join(r.Root(), consts.OrcaYamlFile)
	data, err := os.ReadFile(target)
	if err != nil {
		return nil, err
	}
	return data, nil
}
