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

type CreateConfigContext interface {
	context.WithRoot
	context.WithPolicy
}

func CreateConfig(ctx CreateConfigContext, name string) (string, error) {
	return create.CreateConfig(ctx, name)
}
