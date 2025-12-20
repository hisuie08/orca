package compose

import "path/filepath"

func CollectNetworks(m map[string]*ComposeSpec) []CollectedNetwork {
	result := []CollectedNetwork{}
	for _, c := range Collect(m) {
		for _, v := range c.Spec.Networks {
			result = append(result, CollectedNetwork{
				From: filepath.Base(c.From),
				Spec: v,
			})
		}
	}
	return result
}
