package config

import (
	"orca/internal/context"
	"orca/internal/usecase/config/create"
	"orca/internal/usecase/config/load"
	"orca/model/config"
)

type LoadConfigContext interface {
	context.WithRoot
}

func Load(ctx LoadConfigContext) (*config.OrcaConfig, error) {
	return load.Load(ctx)
}

type CreateCfgContext interface {
	context.WithRoot
	context.WithPolicy
	context.WithLog
}

type WriteOption struct {
	NoCreate bool
	Force    bool
}

func Create(ctx CreateCfgContext, opt config.CfgOption) *config.OrcaConfig {
	return create.ConfigCreator(ctx).Create(opt)
}

func Write(ctx CreateCfgContext, cfg *config.OrcaConfig, opt WriteOption) error {
	if opt.NoCreate {
		return nil
	}
	return create.ConfigCreator(ctx).Write(cfg, opt.Force)
}
