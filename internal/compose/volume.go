package compose

import (
	"orca/internal/ostools"
	"path/filepath"
)

// composeのボリュームオプション構造体
type VolumeSpec struct {
	Driver     string            `yaml:"driver"`
	DriverOpts map[string]string `yaml:"driver_opts"`
	External   bool              `yaml:"external"`
	Labels     map[string]string `yaml:"labels"`
	Name       string            `yaml:"name"`
}

// Orcaがボリュームをオーバーレイする必要があるか
//
// local+bind+deviceが存在しないケース
func (v *VolumeSpec) NeedsOrcaOverlay() bool {

	if v.External {
		return false
	}

	// case 1: driver未定義 → defaultの local
	if v.Driver == "" {
		return true
	}

	// case 2: driver=local かつ driver_optsなし
	if v.Driver == "local" && len(v.DriverOpts) == 0 {
		return true
	}

	// case 3: driver=local + bind but deviceのパスが存在しない
	if v.Driver == "local" {
		t := v.DriverOpts["type"]
		o := v.DriverOpts["o"]
		dev := v.DriverOpts["device"]
		if t == "none" && o == "bind" {
			if !ostools.DirExists(dev) {
				return true
			}
		}
	}
	return false
}

// Orcaがボリュームをオプションで上書きする
//
// とりあえずはローカルバインド作成専用
func (v *VolumeSpec) applyOrcaOverlay(spec VolumeSpec) {
	v.Driver = spec.Driver
	v.DriverOpts = spec.DriverOpts
}

// ローカルバインドをオーバーレイ
func (v *VolumeSpec) ApplyLocalBind(volume_root string) *VolumeSpec {
	path := filepath.Join(volume_root, v.Name)
	opts := map[string]string{
		"type": "none", "o": "bind", "device": path,
	}
	spec := VolumeSpec{Driver: "local", DriverOpts: opts}
	v.applyOrcaOverlay(spec)
	return v
}
