package compose

// CollectedSpec
type CollectedSpec[T any] struct {
	From string // 定義されていたcompose
	Spec T      // 定義
}
type FromRef struct {
	Compose string
	Key     string
}

type CollectedCompose struct {
	From string       // 定義されていたcompose
	Spec *ComposeSpec // 定義
}

type CollectedVolume struct {
	From FromRef     // 定義されていたcompose
	Spec *VolumeSpec // 定義
}

type CollectedNetwork struct {
	From FromRef      // 定義されていたcompose
	Spec *NetworkSpec // 定義
}

func (m ComposeMap) CollectComposes() []CollectedCompose {
	result := []CollectedCompose{}
	for k, v := range m {
		result = append(result, CollectedCompose{From: k, Spec: v})
	}
	return result
}

func (m ComposeMap) CollectNetworks() []CollectedNetwork {
	result := []CollectedNetwork{}
	for name, c := range m {
		for k, v := range c.Networks {
			result = append(result, CollectedNetwork{
				From: FromRef{Compose: name, Key: k},
				Spec: v,
			})
		}
	}
	return result
}

func (m ComposeMap) CollectVolumes() []CollectedVolume {
	result := []CollectedVolume{}
	for name, c := range m {
		for k, v := range c.Volumes {
			result = append(result, CollectedVolume{
				From: FromRef{Compose: name, Key: k},
				Spec: v,
			})
		}
	}
	return result
}
