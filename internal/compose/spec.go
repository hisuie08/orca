package compose

type SpecMap[T any] map[string]T

func Collect[T any](m SpecMap[T]) []CollectedSpec[T] {
	result := []CollectedSpec[T]{}
	for k, v := range m {
		result = append(result, CollectedSpec[T]{From: k, Spec: v})
	}
	return result
}

type ComposeMap SpecMap[*ComposeSpec]

// ComposeSpec Orcaが読み出すComposeのルートセクション
type ComposeSpec struct {
	Volumes  VolumesSection  `yaml:"volumes"`
	Networks NetworksSection `yaml:"networks"`
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
type CollectedVolume CollectedSpec[*VolumeSpec]

type CollectedCompose CollectedSpec[*ComposeSpec]

type CollectedNetwork CollectedSpec[*NetworkSpec]
