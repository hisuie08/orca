package applier

import (
	"orca/consts"
	"orca/internal/config"
	"os"
	"path/filepath"
)

var _ config.ConfigWriter = (*ConfigFile)(nil)

type ConfigFile struct {
	OrcaRoot string
}

func (c ConfigFile) Create(b []byte) (string, error) {
	path, err := filepath.Abs(filepath.Join(c.OrcaRoot, consts.OrcaYamlFile))
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(path, b, 0o644); err != nil {
		return "", err
	}
	return path, nil
}
