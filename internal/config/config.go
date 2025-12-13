package config

import (
	orca "orca/helper"
	"os"
	"path/filepath"
	"sync"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

var (
	orcaConfig OrcaConfig
	configLock sync.RWMutex
)

const (
	OrcaYamlFile string = "orca.yml"
)

type OrcaConfig struct {
	Name    *string         `yaml:"name"`
	Volume  *VolumeConfig  `yaml:"volume"`
	Network *NetworkConfig `yaml:"network"`
}

type VolumeConfig struct {
	VolumeRoot *string `yaml:"volume_root"`
	EnsurePath bool   `yaml:"ensure_path" default:"true"`
}

type NetworkConfig struct {
	Enabled  bool   `yaml:"enabled" default:"true"`
	Internal bool   `yaml:"internal" default:"false"`
	Name     *string `yaml:"name"`
}

// configを作成
func NewConfig() *OrcaConfig {
	orcaConfig := &OrcaConfig{
		Volume:  &VolumeConfig{},
		Network: &NetworkConfig{},
	}
	return orcaConfig
}

// テスト用に切り出したパース処理
func (cfg *OrcaConfig) parseConfig(data []byte) error {
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return orca.OrcaError("orca.yml unmarshal failed", err)
	}
	return nil
}

// Configの実行時に解決される値を解決する処理
func (c *OrcaConfig) Resolve(baseDir string) error {
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
func Load(path string) (*OrcaConfig, error) {
	data, err := os.ReadFile(filepath.Join(path, OrcaYamlFile))
	if err != nil {
		return nil, orca.OrcaError("file read error", err)
	}
	cfg := NewConfig()
	if err := cfg.parseConfig(data); err != nil {
		return nil, err
	}
	// defaults適用
	if err := defaults.Set(cfg); err != nil {
		return nil, orca.OrcaError("default apply failed", err)
	}
	// pathをもとにして実行時解決部分の補完
	if err := cfg.Resolve(path); err != nil {
		return nil, orca.OrcaError("config resolve failed", err)
	}
	return cfg, nil
}

// ゼロからorca.ymlを生成するやつ
func NewDefaultConfig(clusterName string) *OrcaConfig {
	cfg := NewConfig()
	if clusterName != "" {
    cfg.Name = &clusterName
}
	defaults.Set(cfg)
	return cfg
}
