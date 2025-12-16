package compose

import (
	"orca/internal/ostools"
	"path/filepath"
)



// Orcaがボリュームをオーバーレイする必要があるか
//
// local+bind + deviceが存在しないケース
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

	// case 3: driver=local + bind だが deviceのパスが存在しない
	if v.Driver == "local" && len(v.DriverOpts) > 0 {
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

// ローカルバインドをオーバーレイ
func (v *VolumeSpec) ApplyLocalBind(volume_root string) *VolumeSpec {
	var path string
	if device, ok := v.DriverOpts["device"]; ok {
		// もしユーザーが明示したバインド先があればそれを尊重
		path = device
	} else {
		// Nameはdockerが解決してくれる
		path = filepath.Join(volume_root, v.Name)
	}
	v.Driver = "local"
	if v.DriverOpts == nil {
		v.DriverOpts = map[string]string{}
	}
	v.DriverOpts["type"] = "none"
	v.DriverOpts["o"] = "bind"
	v.DriverOpts["device"] = path

	return v
}

func (v *VolumeSpec)ApplyExternal()*VolumeSpec{
	v.Driver=""
	for k := range v.DriverOpts {
		delete(v.DriverOpts, k)
	}
	v.External=true
	return v
}