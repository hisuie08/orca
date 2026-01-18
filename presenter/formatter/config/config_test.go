package config

import (
	"orca/model/config"
	"testing"
)

func TestFmtConfig(t *testing.T) {
	cfg := config.ResolvedConfig{
		Name: "myorca",
		Network: config.ResolvedNetwork{
			Enabled: true, Name: "my-network", Internal: false,
		},
		Volume: config.ResolvedVolume{
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
