package load

import (
	"orca/errs"
	"orca/internal/context"
	"orca/internal/inspector"
	"orca/internal/usecase/config/create"
	"orca/model/config"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type LoadConfigContext interface {
	context.WithRoot
}

func Load(ctx LoadConfigContext) (*config.OrcaConfig, error) {
	return loadConfig(ctx, inspector.NewFilesystem())
}

type fsInspector interface {
	Read(string) ([]byte, error)
}

func loadConfig(ctx LoadConfigContext,
	fi fsInspector) (
	*config.OrcaConfig, error) {
	path := ctx.OrcaYamlFile()
	data, err := fi.Read(path)
	if err != nil {
		return nil, &errs.FileError{Path: path, Err: err}
	}
	name := filepath.Base(ctx.Root())
	cfg := create.NewConfig(config.CfgOption{Name: name})
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
