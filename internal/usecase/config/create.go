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

func CreateConfig(ctx CreateConfigContext, name string) (string, error) {
	return createConfig(ctx, filesystem.NewExecutor(ctx), name)
}

func createConfig(ctx CreateConfigContext,
	writer filesystem.Executor,
	clusterName string) (string, error) {
	cfg := makeConfig(clusterName)
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}
	return ctx.OrcaYamlFile(), writer.WriteFile(ctx.OrcaYamlFile(), b)
}

func makeConfig(name string) *config.OrcaConfig {
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
