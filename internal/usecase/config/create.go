package config

import (
	"orca/internal/context"
	"orca/internal/executor/filesystem"
	"orca/model/config"
	"orca/model/policy"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

var _ creator = (*cfgCreator)(nil)

type creator interface {
	Create(string) (string, error)
}
type cfgCreator struct {
	context.WithRoot
	context.WithPolicy
	writer filesystem.Executor
}

func NewCreator(root string, p policy.ExecPolicy) creator {
	return &cfgCreator{
		writer:     filesystem.NewExecutor(p),
		WithRoot:   context.NewWithRoot(root),
		WithPolicy: context.NewWithPolicy(p),
	}
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
