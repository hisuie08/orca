package diff

import (
	"orca/model/compose"

	"gopkg.in/yaml.v3"
)

// CopyMap creates a semantic deep copy of ComposeMap by YAML round-trip.
// This is intended for diffing / snapshot purposes,
// NOT for preserving original YAML formatting, comments, or ordering.
func CopyMap(cm *compose.ComposeMap) (*compose.ComposeMap, error) {
	cp := compose.ComposeMap{}
	for ref, oldSpec := range *cm {
		b, err := yaml.Marshal(oldSpec)
		if err != nil {
			return nil, err
		}
		newSpec := &compose.ComposeSpec{}
		if err := yaml.Unmarshal(b, newSpec); err != nil {
			return nil, err
		}
		cp[ref] = newSpec
	}
	return &cp, nil
}
