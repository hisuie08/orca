package config

import (
	"orca/infra/inspector"
	"path/filepath"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

type ConfigWriter interface {
	WriteConfig(b []byte) (string, error)
}

// Configの実行時に解決される値を解決する処理 ほぼ全てのコマンドで必要
func (c *OrcaConfig) resolve(name string) *ResolvedConfig {
	result := &ResolvedConfig{
		Volume: ResolvedVolume{
			VolumeRoot: func() *string {
				if c.Volume.VolumeRoot != nil {
					if path, err := filepath.Abs(*c.Volume.VolumeRoot); err == nil {
						return &path
					}
				}
				return nil
			}(),
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
func LoadConfig(r inspector.ConfigReader) (*ResolvedConfig, error) {
	data, err := r.Read()
	if err != nil {
		return nil, err
	}
	cfg := &OrcaConfig{
		Volume:  &VolumeConfig{},
		Network: &NetworkConfig{},
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg.resolve(filepath.Base(r.Root())), nil
}

// ゼロからorca.ymlを生成するやつ init コマンドで呼び出される
func Create(clusterName string, c ConfigWriter) (string, error) {
	cfg := &OrcaConfig{
		Volume:  &VolumeConfig{},
		Network: &NetworkConfig{},
	}
	if clusterName != "" {
		cfg.Name = &clusterName
	}
	defaults.Set(cfg)
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}
	return c.WriteConfig(b)
}
