package internal_test

import (
	"orca/internal/config/internal"
	"orca/model/config"
	"testing"
)

func TestResolve(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		c *config.OrcaConfig
	}{
		// TODO: Add test cases.
		{"test", &config.OrcaConfig{Name: nil, Volume: &config.VolumeConfig{}, Network: &config.NetworkConfig{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := internal.Resolve(tt.c, tt.name)
			// TODO: update the condition below to compare got with tt.want.
			if got.Name != tt.name {
				t.Errorf("Resolve() = %v, want %v", got, tt.name)
			}
		})
	}
}
