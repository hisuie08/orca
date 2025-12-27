package applier

import (
	"orca/consts"
	"orca/ostools"
	"path/filepath"
)

type configFileWriter struct {
	OrcaRoot string
}

func (c configFileWriter) WriteConfig(b []byte) (string, error) {
	path := filepath.Join(c.OrcaRoot, consts.OrcaYamlFile)
	return ostools.CreateFile(path, b)
}

type dryConfigWriter struct {
	OrcaRoot string
}

func (d dryConfigWriter) WriteConfig(b []byte) (string, error) {
	path, err := filepath.Abs(filepath.Join(d.OrcaRoot, consts.OrcaYamlFile))
	if err != nil {
		return "", err
	}
	return path, nil
}
