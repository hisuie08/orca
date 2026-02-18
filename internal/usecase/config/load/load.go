package load

import (
	"orca/errs"
	"orca/internal/capability"
	"orca/internal/inspector"
	"orca/model/config"

	"gopkg.in/yaml.v3"
)

type loadCfgCapability interface {
	capability.WithRoot
}

func Load(caps loadCfgCapability) (*config.OrcaConfig, error) {
	return loadConfig(caps, inspector.NewFilesystem())
}

type fsInspector interface {
	Read(string) ([]byte, error)
}

func loadConfig(caps loadCfgCapability,
	fi fsInspector) (
	*config.OrcaConfig, error) {
	path := caps.OrcaYamlFile()
	data, err := fi.Read(path)
	if err != nil {
		return nil, &errs.FileError{Path: path, Err: err}
	}
	cfg := &config.OrcaConfig{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	if cfg.Name == "" {
		return nil, errs.ErrInvalidConfig
	}
	return cfg, nil
}
