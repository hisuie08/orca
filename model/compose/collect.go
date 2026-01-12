package compose

import (
	"slices"
	"strings"
)

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
	Ref  FromRef     // 定義されていたcompose
	Spec *VolumeSpec // 定義
}

type CollectedNetwork struct {
	Ref  FromRef      // 定義されていたcompose
	Spec *NetworkSpec // 定義
}

func (m ComposeMap) CollectComposes() []CollectedCompose {
	result := []CollectedCompose{}
	for k, v := range m {
		result = append(result, CollectedCompose{From: k, Spec: v})
	}
	slices.SortFunc(result, func(a, b CollectedCompose) int {
		return strings.Compare(a.From, b.From)
	})
	return result
}

func (m ComposeMap) CollectNetworks() []CollectedNetwork {
	result := []CollectedNetwork{}
	for name, c := range m {
		for k, v := range c.Networks {
			result = append(result, CollectedNetwork{
				Ref:  FromRef{Compose: name, Key: k},
				Spec: v,
			})
		}
	}
	slices.SortFunc(result, func(a, b CollectedNetwork) int {
		return strings.Compare(a.Spec.Name, b.Spec.Name)
	})
	return result
}

func (m ComposeMap) CollectVolumes() []CollectedVolume {
	result := []CollectedVolume{}
	for name, c := range m {
		for k, v := range c.Volumes {
			result = append(result, CollectedVolume{
				Ref:  FromRef{Compose: name, Key: k},
				Spec: v,
			})
		}
	}
	slices.SortFunc(result, func(a, b CollectedVolume) int {
		return strings.Compare(a.Spec.Name, b.Spec.Name)
	})
	return result
}
