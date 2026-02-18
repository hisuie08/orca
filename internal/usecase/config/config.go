package config

import (
	"orca/internal/capability"
	"orca/internal/usecase/config/create"
	"orca/internal/usecase/config/load"
	"orca/model/config"
)

type LoadConfigCapability interface {
	capability.WithRoot
}

func Load(caps LoadConfigCapability) (*config.OrcaConfig, error) {
	return load.Load(caps)
}

type CreateCfgCapability interface {
	capability.WithRoot
	capability.WithPolicy
	capability.WithLog
}

type WriteOption struct {
	NoCreate bool
	Force    bool
}

func Create(caps CreateCfgCapability, opt config.CfgOption) *config.OrcaConfig {
	return create.ConfigCreator(caps).Create(opt)
}

func Write(caps CreateCfgCapability, cfg *config.OrcaConfig, opt WriteOption) error {
	if opt.NoCreate {
		return nil
	}
	return create.ConfigCreator(caps).Write(cfg, opt.Force)
}
