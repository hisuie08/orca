package create

import (
	"io/fs"
	"orca/errs"
	"orca/internal/context"
	"orca/internal/executor"
	"orca/internal/inspector"
	"orca/model/config"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

var _ Creator = (*creator)(nil)

type createContext interface {
	context.WithRoot
	context.WithPolicy
}
type Creator interface {
	CreateConfig(string, bool) (string, error)
}

type creator struct {
	ctx createContext
	fe  executor.FileSystem
	fi  inspector.FileSystem
}

func ConfigCreator(ctx createContext) Creator {
	return &creator{ctx: ctx}
}

func (c *creator) CreateConfig(clusterName string, force bool) (string, error) {
	if c.fi.FileExists(c.ctx.OrcaYamlFile()) && !force {
		return "", &errs.FileError{Path: c.ctx.OrcaYamlFile(), Err: fs.ErrExist}
	}
	cfg := makeConfig(clusterName)
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}
	return c.ctx.OrcaYamlFile(), c.fe.WriteFile(c.ctx.OrcaYamlFile(), b)
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
