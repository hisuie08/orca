package compose

import (
	"orca/internal/ostools"
	"path/filepath"
)

func CollectVolumes(m map[string]*ComposeSpec) []CollectedVolume {
	result := []CollectedVolume{}
	for _, c := range MapToArray(m) {
		for _, v := range c.Spec.Volumes {
			result = append(result, CollectedVolume{
				From: filepath.Base(c.From),
				Spec: v,
			})
		}
	}
	return result
}

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
