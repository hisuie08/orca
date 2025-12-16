package compose

import (
	orca "orca/helper"

	"gopkg.in/yaml.v3"
)


// Composeを読み出す関数
func ParseCompose(data []byte) (*ComposeSpec, error) {
	cfg := ComposeSpec{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, orca.OrcaError("compose Parse Error", err)
	}
	return &cfg, nil
}
