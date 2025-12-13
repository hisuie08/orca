package compose

import (
	orca "orca/helper"

	"gopkg.in/yaml.v3"
)

// Orcaが読み出すComposeのルートセクション
type ComposeSpec struct {
	Volumes  VolumesSection  `yaml:"volumes"`
	Networks NetworksSection `yaml:"networks"`
}

// Composeのトップレベルvolumesセクション
type VolumesSection = map[string]*VolumeSpec

// Composeのトップレベルnetworksセクション
type NetworksSection = map[string]*NetworkSpec

// Composeを読み出す関数
func ParseCompose(data []byte) (*ComposeSpec, error) {
	cfg := ComposeSpec{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, orca.OrcaError("compose Parse Error", err)
	}
	// map型のセクションだけnilチェック
	for _, v := range cfg.Volumes {
		if v.DriverOpts == nil {
			v.DriverOpts = make(map[string]string)
		}
		// ラベルをいじる時はここをアンコメント
		// if v.Labels==nil{
		// 	v.Labels=make(map[string]string)
		// }
	}
	return &cfg, nil
}
