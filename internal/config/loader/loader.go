package loader

import (
	"orca/internal/config/internal"
	"orca/internal/context"
	"orca/model/config"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var _ config.ConfigLoader = (*configLoader)(nil)

type configLoader struct {
	context.WithRoot
}

func NewLoader(root string) config.ConfigLoader {
	return &configLoader{WithRoot: context.NewWithRoot(root)}
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
	return internal.Resolve(cfg, filepath.Base(c.Root())), nil
}
