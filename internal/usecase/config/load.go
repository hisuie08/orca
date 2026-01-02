package config

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
	l := &configLoader{WithRoot: ctx, fi: inspector.NewFilesystem()}
	return l.Load()
}

type configLoader struct {
	context.WithRoot
	fi inspector.FileSystem
}

func (c *configLoader) Load() (*config.ResolvedConfig, error) {
	path := c.OrcaYamlFile()
	data, err := c.fi.Read(path)
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
	return internal.Resolve(cfg, filepath.Base(c.Root())), nil
}
