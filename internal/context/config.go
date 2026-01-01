package context

import "orca/model/config"

type WithConfig interface {
	Config() *config.ResolvedConfig
}

type withConfig struct {
	config *config.ResolvedConfig
}


func (w withConfig) Config() *config.ResolvedConfig {
	return w.config
}
