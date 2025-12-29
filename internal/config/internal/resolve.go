package internal

import (
	"orca/model/config"
	"path/filepath"
)

func Resolve(c *config.OrcaConfig, name string) *config.ResolvedConfig {
	result := &config.ResolvedConfig{
		Volume: config.ResolvedVolume{
			VolumeRoot: func() *string {
				if c.Volume.VolumeRoot != nil {
					if path, err := filepath.Abs(*c.Volume.VolumeRoot); err == nil {
						return &path
					}
				}
				return nil
			}(),
			EnsurePath: c.Volume.EnsurePath,
		},
		Network: config.ResolvedNetwork{
			Enabled:  c.Network.Enabled,
			Internal: c.Network.Internal},
	}
	if c.Name == nil {
		result.Name = name
	} else {
		result.Name = *c.Name
	}

	if c.Network != nil && c.Network.Enabled {
		if c.Network.Name == nil {
			name := result.Name + "_network"
			result.Network.Name = name
		} else {
			result.Network.Name = *c.Network.Name
		}
	}
	return result
}
