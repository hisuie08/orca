package compose

// Orcaが読み出すComposeのルートセクション
type ComposeSpec struct {
	Volumes  VolumesSection  `yaml:"volumes"`
	Networks NetworksSection `yaml:"networks"`
}

// Composeのトップレベルvolumesセクション
type VolumesSection = map[string]*VolumeSpec

// Composeのトップレベルnetworksセクション
type NetworksSection = map[string]*NetworkSpec

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
	Name     string `yaml:"name"`
	Driver   string `yaml:"driver"`
	External bool   `yaml:"external"`
}


// =================
//
//	CollectedSpec
//
// =================
// From: 定義されていたcompose
// Spec: 定義
type CollectedVolume struct {
	From string
	Spec *VolumeSpec
}

type CollectedCompose struct {
	From string
	Spec *ComposeSpec
}

type CollectedNetwork struct {
	From string
	Spec *NetworkSpec
}