package creator

import (
	"orca/infra/executor/fs"
	"orca/internal/context"
	"orca/internal/policy"
	"orca/model/config"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

var _ config.ConfigCreator = (*cfgCreator)(nil)

type cfgCreator struct {
	context.WithRoot
	context.WithPolicy
	writer fs.FileWriter
}

func NewCreator(root string, p policy.ExecPolicy) config.ConfigCreator {
	return &cfgCreator{
		writer:     fs.NewFileWriter(p),
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
	return c.writer.Write(c.OrcaYamlFile(), b)
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
