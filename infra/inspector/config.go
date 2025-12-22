package inspector

import (
	orca "orca/helper"
	"orca/ostools"
	"path/filepath"
)

type ConfigFileReader struct {
}

func (r ConfigFileReader) Read(orcaRoot string) ([]byte, error) {
	return ostools.ReadFile(filepath.Join(orcaRoot, orca.OrcaYamlFile))
}
