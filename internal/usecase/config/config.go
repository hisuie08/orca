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

func Create(ctx CreateCfgContext, opt config.CfgOption, force bool) (*config.OrcaConfig, error) {
	return create.ConfigCreator(ctx).Create(opt, force)
}
