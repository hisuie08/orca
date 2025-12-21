package config

import (
	orca "orca/helper"
	"orca/internal/ostools"
	"path/filepath"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

const (
	OrcaYamlFile string = "orca.yml"
)

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
	VolumeConfig
}
type ResolvedNetwork struct {
	Enabled  bool
	Internal bool
	Name     string
}

// configを作成
func newConfig() *OrcaConfig {
	orcaConfig := &OrcaConfig{
		Volume:  &VolumeConfig{},
		Network: &NetworkConfig{},
	}
	return orcaConfig
}

// テスト用に切り出したパース処理
func parseConfig(cfg *OrcaConfig, data []byte) error {
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return orca.OrcaError("orca.yml unmarshal failed", err)
	}
	return nil
}

// Configの実行時に解決される値を解決する処理 ほぼ全てのコマンドで必要
func (c *OrcaConfig) resolve(baseDir string) error {
	if c.Name == nil {
		name := filepath.Base(baseDir)
		c.Name = &name
	}

	if c.Network != nil && c.Network.Enabled {
		if c.Network.Name == nil {
			name := *c.Name + "_network"
			c.Network.Name = &name
		}
	}
	return nil
}

// ファイル読み込みからConfig構築
func LoadConfig(orca_dir string, r ConfigReader) (*OrcaConfig, error) {
	data, err := r.Read()
	if err != nil {
		return nil, err
	}
	cfg := &OrcaConfig{
		Volume:  &VolumeConfig{},
		Network: &NetworkConfig{},
	}
	if err := parseConfig(cfg, data); err != nil {
		return nil, err
	}
	// defaults適用
	if err := defaults.Set(cfg); err != nil {
		return nil, orca.OrcaError("default apply failed", err)
	}
	// pathをもとにして実行時解決部分の補完
	if err := cfg.resolve(orca_dir); err != nil {
		return nil, orca.OrcaError("config resolve failed", err)
	}
	return cfg, nil
}

// ゼロからorca.ymlを生成するやつ init コマンドで呼び出される
func Create(clusterName string, path string) (*OrcaConfig, error) {
	cfg := &OrcaConfig{
		Volume:  &VolumeConfig{},
		Network: &NetworkConfig{},
	}
	if clusterName != "" {
		cfg.Name = &clusterName
	}
	defaults.Set(cfg)
	if data, err := yaml.Marshal(cfg); err != nil {
		return nil, err
	} else {
		if err := ostools.CreateFile(path, data); err != nil {
			return nil, err
		}
	}
	return cfg, nil
}
