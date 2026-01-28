package config

import (
	"orca/model/config"
	"testing"
)

func TestFmtConfig(t *testing.T) {
	cfg := config.OrcaConfig{
		Name: "myorca",
		Network: config.NetworkConfig{
			Enabled: true, Name: "my-network", Internal: false,
		},
		Volume: config.VolumeConfig{
			VolumeRoot: nil,
			EnsurePath: false,
		},
	}
	t.Run("format config", func(t *testing.T) {
		got := FmtConfig(cfg)
		// TODO: update the condition below to compare got with tt.want.
		t.Log(got)
		return
	})
}
