package create

import (
	"fmt"
	"orca/errs"
	"orca/internal/capability"
	"orca/internal/executor"
	"orca/internal/inspector"
	"orca/model/config"
	"path/filepath"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

var _ Creator = (*creator)(nil)

type createCapability interface {
	capability.WithRoot
	capability.WithPolicy
	capability.WithLog
}
type Creator interface {
	Create(config.CfgOption) *config.OrcaConfig
	Write(*config.OrcaConfig, bool) error
}

type creator struct {
	caps createCapability
	fe   executor.FileSystem
	fi   inspector.FileSystem
}

func ConfigCreator(caps createCapability) Creator {
	return &creator{caps: caps, fi: inspector.NewFilesystem(),
		fe: executor.NewFilesystem(caps)}
}

func (c *creator) Create(opt config.CfgOption) *config.OrcaConfig {
	return NewConfig(opt)
}

func (c *creator) Write(cfg *config.OrcaConfig, force bool) error {
	if c.fi.FileExists(c.caps.OrcaYamlFile()) {
		if !force {
			return errs.ErrAlreadyInitialized
		}
	}
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return c.fe.WriteFile(c.caps.OrcaYamlFile(), b)
}

func NewConfig(c config.CfgOption) *config.OrcaConfig {
	cfg := &config.OrcaConfig{
		Name: c.Name,
		Volume: config.VolumeConfig{
			VolumeRoot: func() *string {
				if c.Volume.Path != "" {
					if path, err := filepath.Abs(c.Volume.Path); err == nil {
						return &path
					}
				}
				return nil
			}(),
			EnsurePath: c.Volume.EnsurePath},
		Network: config.NetworkConfig{
			Name: func() string {
				if c.Network.Name != "" {
					return c.Network.Name
				}
				return fmt.Sprintf("%s_network", c.Name)
			}(),
			Enabled:  c.Network.Enabled,
			Internal: c.Network.Internal},
	}
	defaults.Set(cfg)
	return cfg
}
