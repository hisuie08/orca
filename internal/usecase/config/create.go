package config

import (
	"orca/internal/context"
	"orca/internal/executor/filesystem"
	"orca/model/config"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

type CreateConfigContext interface {
	context.WithRoot
	context.WithPolicy
}

type cfgCreator struct {
	context.WithRoot
	context.WithPolicy
	writer filesystem.Executor
}

func CreateConfig(ctx CreateConfigContext, name string) (string, error) {
	c := &cfgCreator{
		writer:     filesystem.NewExecutor(ctx),
		WithRoot:   ctx,
		WithPolicy: ctx,
	}
	return c.Create(name)
}

func (c *cfgCreator) Create(clusterName string) (string, error) {
	cfg := c.makeConfig(clusterName)
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}
	return c.OrcaYamlFile(), c.writer.WriteFile(c.OrcaYamlFile(), b)
}

func (c *cfgCreator) makeConfig(name string) *config.OrcaConfig {
	cfg := &config.OrcaConfig{
		Volume:  &config.VolumeConfig{},
		Network: &config.NetworkConfig{},
	}
	if name != "" {
		cfg.Name = &name
	}
	defaults.Set(cfg)
	return cfg
}
