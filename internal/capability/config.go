package capability

import "orca/model/config"

type WithConfig interface {
	Config() *config.OrcaConfig
}

type withConfig struct {
	config *config.OrcaConfig
}

func (w withConfig) Config() *config.OrcaConfig {
	return w.config
}
