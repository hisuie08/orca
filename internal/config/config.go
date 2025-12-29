package config

import (
	"orca/internal/config/creator"
	"orca/internal/config/loader"
	"orca/internal/policy"
	"orca/model/config"
)

func ConfigLoader(root string) config.ConfigLoader {
	return loader.NewLoader(root)
}

func ConfigCreator(root string, p policy.ExecPolicy) config.ConfigCreator {
	return creator.NewCreator(root, p)
}
