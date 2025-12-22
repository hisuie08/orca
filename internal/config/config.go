package config

import (
	orca "orca/helper"
	"path/filepath"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

// テスト用に切り出したパース処理
func parseConfig(cfg *OrcaConfig, data []byte) error {
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return orca.OrcaError("orca.yml unmarshal failed", err)
	}
	return nil
}

// Configの実行時に解決される値を解決する処理 ほぼ全てのコマンドで必要
func (c *OrcaConfig) resolve(name string) *ResolvedConfig {
	result := &ResolvedConfig{
		Volume: ResolvedVolume{
			VolumeRoot: c.Volume.VolumeRoot,
			EnsurePath: c.Volume.EnsurePath,
		},
		Network: ResolvedNetwork{
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

// ファイル読み込みからConfig構築
func LoadConfig(orcaRoot string, r ConfigReader) (*ResolvedConfig, error) {
	data, err := r.Read(orcaRoot)
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

	return cfg.resolve(filepath.Base(orcaRoot)), nil
}

// ゼロからorca.ymlを生成するやつ init コマンドで呼び出される
func Create(clusterName string) *OrcaConfig {
	cfg := &OrcaConfig{
		Volume:  &VolumeConfig{},
		Network: &NetworkConfig{},
	}
	if clusterName != "" {
		cfg.Name = &clusterName
	}
	defaults.Set(cfg)

	return cfg
}
