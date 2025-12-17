package compose

import "path/filepath"

func CollectNetworks(m map[string]*ComposeSpec) []CollectedNetwork {
	result := []CollectedNetwork{}
	for _, c := range MapToArray(m) {
		for _, v := range c.Spec.Networks {
			result = append(result, CollectedNetwork{
				From: filepath.Base(c.From),
				Spec: v,
			})
		}
	}
	return result
}

func ApplyNetworkOverlay(cmps *ComposeSpec, name string) {
	for k, n := range cmps.Networks {
		if n.Name == name {
			// 万が一競合する同名ネットワークがcompose内にあったら削除
			delete(cmps.Networks, k)
		}
	}
	cmps.Networks["default"] = &NetworkSpec{
		Name: name,
		// orcaが作るか、既存のものを使うのでどちらにしろexternal
		External: true,
	}
}
