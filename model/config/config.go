package config

type CfgOption struct {
	Name   string
	Volume struct {
		Path       string
		EnsurePath bool
	}
	Network struct {
		Enabled  bool
		Name     string
		Internal bool
	}
}

type OrcaConfig struct {
	Name    string        `yaml:"name"`
	Volume  VolumeConfig  `yaml:"volume"`
	Network NetworkConfig `yaml:"network"`
}
type VolumeConfig struct {
	VolumeRoot *string `yaml:"volume_root"`
	EnsurePath bool    `yaml:"ensure_path" default:"true"`
}
type NetworkConfig struct {
	Enabled  bool   `yaml:"enabled" default:"true"`
	Internal bool   `yaml:"internal" default:"false"`
	Name     string `yaml:"name"`
}

func (v *VolumeConfig) Enabled() bool {
	return v.VolumeRoot != nil && *v.VolumeRoot != ""
}
