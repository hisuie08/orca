package load

import (
	"orca/errs"
	"orca/internal/context"
	"orca/internal/inspector"
	"orca/model/config"

	"gopkg.in/yaml.v3"
)

type loadCfgContext interface {
	context.WithRoot
}

func Load(ctx loadCfgContext) (*config.OrcaConfig, error) {
	return loadConfig(ctx, inspector.NewFilesystem())
}

type fsInspector interface {
	Read(string) ([]byte, error)
}

func loadConfig(ctx loadCfgContext,
	fi fsInspector) (
	*config.OrcaConfig, error) {
	path := ctx.OrcaYamlFile()
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
