package context

import "orca/model/config"

type WithConfig struct {
	config *config.ResolvedConfig
}

func NewWithConfig(cfg *config.ResolvedConfig) WithConfig {

	if cfg == nil {
		panic("ResolvedConfig must not be nil")
	}
	return WithConfig{config: cfg}
}

func (w WithConfig) Config() *config.ResolvedConfig {
	return w.config
}
