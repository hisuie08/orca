package load

import (
	"orca/internal/context"
	"orca/internal/inspector"
	"orca/internal/usecase/config/internal"
	"orca/model/config"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type LoadConfigContext interface {
	context.WithRoot
}

func LoadConfig(ctx LoadConfigContext) (*config.ResolvedConfig, error) {
	return loadConfig(ctx, inspector.NewFilesystem())
}

func loadConfig(ctx LoadConfigContext,
	fi inspector.FileSystem) (
	*config.ResolvedConfig, error) {
	path := ctx.OrcaYamlFile()
	data, err := fi.Read(path)
	if err != nil {
		return nil, err
	}
	cfg := &config.OrcaConfig{
		Volume:  &config.VolumeConfig{},
		Network: &config.NetworkConfig{},
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return internal.Resolve(cfg, filepath.Base(ctx.Root())), nil
}
