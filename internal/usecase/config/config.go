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

func LoadConfig(ctx LoadConfigContext) (*config.ResolvedConfig, error) {
	return load.LoadConfig(ctx)
}

type CreateCfgContext interface {
	context.WithRoot
	context.WithPolicy
}

func CreateConfig(ctx CreateCfgContext, name string, force bool) (string, error) {
	return create.ConfigCreator(ctx).CreateConfig(name, force)
}
