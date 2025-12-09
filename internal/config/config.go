package config

import (
	"fmt"
	"orca/internal/ostools"
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

var OrcaConfig *OrcaYaml;
type OrcaYaml struct {
	Name     string             `yaml:"name"`
	CacheDir string             `yaml:"cache_dir,omitempty"`
	Volume   *VolumeConfigSpec  `yaml:"volume"`
	Network  *NetworkConfigSpec `yaml:"network"`
}

type VolumeConfigSpec struct {
	ChangeVolumeRoot bool   `yaml:"change_root" default:"false"`
	MakeDirs         bool   `yaml:"make_dirs" default:"false"`
	RootPath         string `yaml:"root_path"`
}

type NetworkConfigSpec struct {
	CreateOrcaNetwork bool   `yaml:"create_network" default:"true"`
	NetworkName       string `yaml:"network_name"`
}

// 動的にデフォルト値を設定する必要があるやつ
func (o *OrcaYaml) SetDefaults() {
	cd, _ := os.Getwd()
	if defaults.CanUpdate(o.Name) {
		o.Name = ostools.DirName(cd)
	}
	if defaults.CanUpdate(o.CacheDir) {
		o.CacheDir = cd + "/.orca"
	}
	if defaults.CanUpdate(o.Volume.RootPath) {
		o.Volume.RootPath = cd + "/volumes"
	}
	if defaults.CanUpdate(o.Network.NetworkName) {
		o.Network.NetworkName = ostools.DirName(cd) + "_network"
	}
}

// 既存のorca.ymlを読み出す。
//
// ostool.LoadFile(ファイルパス)とかを渡す
func Load(data []byte) (*OrcaYaml, error) {
	orcaYaml := &OrcaYaml{}
	yaml.Unmarshal(data, orcaYaml)
	// セクション自体がなければ初期化
	if orcaYaml.Network == nil {
		orcaYaml.Network = &NetworkConfigSpec{}
	}
	if orcaYaml.Volume == nil {
		orcaYaml.Volume = &VolumeConfigSpec{}
	}
	// 値単位でデフォルトをセット
	if err := defaults.Set(orcaYaml); err != nil {
		return nil, fmt.Errorf("unexpected error: %w", err)
	}

	return orcaYaml, nil
}

func LoadF(path string) (*OrcaYaml, error) {
	data, err := ostools.LoadFile(path)
	if err == nil {
		return Load(data)
	}
	return nil, err
}

// ゼロからorca.ymlを生成するやつ
func Create(path string) error {
	orcaYaml := &OrcaYaml{
		Network: &NetworkConfigSpec{},
		Volume:  &VolumeConfigSpec{}}
	if err := defaults.Set(orcaYaml); err != nil {
		panic(err)
	}
	content, _ := yaml.Marshal(orcaYaml)
	return ostools.ToFile(content, path)
}
