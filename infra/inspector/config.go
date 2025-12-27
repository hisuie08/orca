package inspector

import (
	"orca/consts"
	"orca/internal/config"
	"os"
	"path/filepath"
)

var _ config.ConfigReader = (*ConfigFile)(nil)

type ConfigFile struct {
	OrcaRoot string
}

func (r ConfigFile) Read() ([]byte, error) {
	target := filepath.Join(r.OrcaRoot, consts.OrcaYamlFile)
	data, err := os.ReadFile(target)
	if err != nil {
		return nil, err
	}
	return data, nil
}
