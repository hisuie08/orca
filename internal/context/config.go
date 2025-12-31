package context

import "orca/model/config"

type WithConfig interface {
	Config() *config.ResolvedConfig
}

type withConfig struct {
	config *config.ResolvedConfig
}

func NewWithConfig(cfg *config.ResolvedConfig) WithConfig {
	if cfg == nil {
		panic("ResolvedConfig must not be nil")
	}
	return withConfig{config: cfg}
}

func (w withConfig) Config() *config.ResolvedConfig {
	return w.config
}
