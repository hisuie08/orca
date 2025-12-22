package inspector

import (
	"orca/consts"
	"orca/ostools"
	"path/filepath"
)

type ConfigFileReader struct {
}

func (r ConfigFileReader) Read(orcaRoot string) ([]byte, error) {
	return ostools.ReadFile(filepath.Join(orcaRoot, consts.OrcaYamlFile))
}
