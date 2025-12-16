package compose

import "orca/internal/config"



func NeedNetworkOverlay(cfg *config.OrcaConfig) bool {
	return cfg.Network.Enabled
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
