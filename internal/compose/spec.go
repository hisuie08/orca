package compose

type SpecMap[T any] map[string]T

type ComposeMap SpecMap[*ComposeSpec]

// ComposeSpec Orcaが読み出すComposeのルートセクション
type ComposeSpec struct {
	Rest     map[string]any  `yaml:",inline"`
	Volumes  VolumesSection  `yaml:"volumes,omitempty"`
	Networks NetworksSection `yaml:"networks,omitempty"`
}

// Composeのトップレベルvolumesセクション
type VolumesSection SpecMap[*VolumeSpec]

// Composeのトップレベルnetworksセクション
type NetworksSection SpecMap[*NetworkSpec]

// composeのボリュームオプション構造体
type VolumeSpec struct {
	Driver     string            `yaml:"driver"`
	DriverOpts map[string]string `yaml:"driver_opts"`
	External   bool              `yaml:"external"`
	Labels     map[string]string `yaml:"labels"`
	Name       string            `yaml:"name"`
}

// composeのネットワークオプション構造体
type NetworkSpec struct {
	Name     string            `yaml:"name"`
	Driver   string            `yaml:"driver"`
	External bool              `yaml:"external"`
	Labels   map[string]string `yaml:"labels"`
}

// CollectedSpec
type CollectedSpec[T any] struct {
	From string // 定義されていたcompose
	Spec T      // 定義
}
type FromRef struct {
	Compose string
	Key     string
}

type CollectedCompose struct {
	From string       // 定義されていたcompose
	Spec *ComposeSpec // 定義
}

type CollectedVolume struct {
	From FromRef     // 定義されていたcompose
	Spec *VolumeSpec // 定義
}

type CollectedNetwork struct {
	From FromRef      // 定義されていたcompose
	Spec *NetworkSpec // 定義
}
