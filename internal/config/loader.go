package config

import (
	"orca/internal/ostools"
	"path/filepath"
)

type ConfigReader interface {
	Read() ([]byte, error)
}

type ConfigFileReader struct {
	OrcaRoot string
}

func (r ConfigFileReader) Read() ([]byte, error) {
	return ostools.ReadFile(filepath.Join(r.OrcaRoot, OrcaYamlFile))
}

type FakeConfigReader struct {
	Want string
}

func (f FakeConfigReader) Read() ([]byte, error) {
	b := []byte(f.Want)
	return b, nil
}
