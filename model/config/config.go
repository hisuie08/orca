package config

type OrcaConfig struct {
	Name    *string        `yaml:"name"`
	Volume  *VolumeConfig  `yaml:"volume"`
	Network *NetworkConfig `yaml:"network"`
}

type VolumeConfig struct {
	VolumeRoot *string `yaml:"volume_root"`
	EnsurePath bool    `yaml:"ensure_path" default:"true"`
}

type NetworkConfig struct {
	Enabled  bool    `yaml:"enabled" default:"true"`
	Internal bool    `yaml:"internal" default:"false"`
	Name     *string `yaml:"name"`
}

type ResolvedConfig struct {
	Name    string
	Volume  ResolvedVolume
	Network ResolvedNetwork
}
type ResolvedVolume struct {
	VolumeRoot *string
	EnsurePath bool
}
type ResolvedNetwork struct {
	Enabled  bool
	Internal bool
	Name     string
}

type ConfigCreator interface {
	Create(string) (string, error)
}

type ConfigLoader interface {
	Load() (*ResolvedConfig, error)
}
