package config

import (
	"orca/internal/context"
	"orca/internal/loader/config/internal"
	"orca/model/config"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var _ config.ConfigLoader = (*configLoader)(nil)

type configLoader struct {
	context.WithRoot
	cfg *config.ResolvedConfig
}

func NewLoader(root string) config.ConfigLoader {
	return &configLoader{cfg: nil, WithRoot: context.NewWithRoot(root)}
}

func (c *configLoader) Load() (*config.ResolvedConfig, error) {
	target := c.OrcaYamlFile()
	data, err := os.ReadFile(target)
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
	c.cfg = internal.Resolve(cfg, filepath.Base(c.Root()))
	return c.cfg, nil
}
