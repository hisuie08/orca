package plan

import (
	"orca/internal/capability"
	"orca/model/compose"
	"orca/model/config"
	"slices"
	"testing"
)

// 子関数を束ねるだけなので最小限のIOチェックのみ
func TestBuildOrcaPlan(t *testing.T) {
	root := t.TempDir()
	caps := capability.New().WithRoot(root).
		WithConfig(&config.OrcaConfig{Name: "test",
			Volume:  config.VolumeConfig{VolumeRoot: &root},
			Network: config.NetworkConfig{Name: "test_net", Enabled: true},
		})
	cm := compose.ComposeMap{"b": &compose.ComposeSpec{},
		"a": &compose.ComposeSpec{}}
	got := BuildOrcaPlan(&caps, cm)
	if len(got.ComposeDirs) != 2 || got.ComposeDirs[0] != "a" ||
		got.ComposeDirs[1] != "b" {
		t.Errorf("unexpected ComposeDirs: %#v", got.ComposeDirs)
	}
	if got.Name != "test" || got.Networks.SharedName != "test_net" {
		t.Errorf("unexpected value, %#v", got)
	}
	if !slices.Equal(got.ComposeDirs, []string{"a", "b"}) {
		t.Errorf("composeDirs is not sorted")
	}
}
