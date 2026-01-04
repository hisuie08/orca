package volume

import (
	"orca/model/compose"
	"orca/model/plan"
	"testing"

	"github.com/go-test/deep"
)

// テストで変える必要がない定数
const (
	refCompose = "a"
	refKey     = "b"
	volName    = "name"
)

func fakeComposeMap() compose.ComposeMap {
	return compose.ComposeMap{refCompose: &compose.ComposeSpec{
		Volumes: compose.VolumesSection{refKey: &compose.VolumeSpec{
			Name: volName, DriverOpts: map[string]string{},
			// 既存Labelsが保持できていることをテストするため
			Labels: map[string]string{"testlabel": "testvalue"}}}}}
}

func TestVolumeOverlay(t *testing.T) {
	tests := []struct {
		name string
		vp   plan.VolumePlan
		want compose.VolumeSpec
	}{
		{
			name: "external",
			vp: plan.VolumePlan{
				Type:   plan.VolumeExternal,
				UsedBy: []plan.VolumeRef{{Compose: refCompose, Key: refKey}},
				Name:   volName},
			want: compose.VolumeSpec{
				External:   true,
				Driver:     "",
				DriverOpts: map[string]string{},
				Labels:     map[string]string{"testlabel": "testvalue"},
				Name:       volName}},
		{
			name: "shared",
			vp: plan.VolumePlan{
				Type:     plan.VolumeShared,
				UsedBy:   []plan.VolumeRef{{Compose: refCompose, Key: refKey}},
				BindPath: "/path/to/bind",
				Name:     volName},
			want: compose.VolumeSpec{
				External:   true,
				Driver:     "",
				DriverOpts: map[string]string{},
				Labels: map[string]string{"testlabel": "testvalue",
					"orca.volume.shared.bind": "/path/to/bind",
				},
				Name: volName}},
		{
			name: "local",
			vp: plan.VolumePlan{
				Type:     plan.VolumeLocal,
				UsedBy:   []plan.VolumeRef{{Compose: refCompose, Key: refKey}},
				BindPath: "/path/to/bind",
				Name:     volName},
			want: compose.VolumeSpec{
				External: false,
				Driver:   "local",
				DriverOpts: map[string]string{
					"type": "none", "o": "bind",
					"device": "/path/to/bind"},
				Labels: map[string]string{"testlabel": "testvalue"},
				Name:   volName}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := fakeComposeMap()
			vps := []plan.VolumePlan{tt.vp}
			OverlayVolume(cm, vps)
			result := cm[refCompose].Volumes[refKey]
			diff := deep.Equal(result, &tt.want)
			if len(diff) != 0 {
				t.Error(diff)
				return
			}
		})
	}
}

func TestMultipleUsed(t *testing.T) {
	vp :=
		plan.VolumePlan{
			Type: plan.VolumeExternal,
			UsedBy: []plan.VolumeRef{
				{Compose: "a", Key: "b"},
				{Compose: "c", Key: "d"}},
			Name: volName}
	cm := func() compose.ComposeMap {
		return compose.ComposeMap{
			"a": &compose.ComposeSpec{
				Volumes: compose.VolumesSection{"b": &compose.VolumeSpec{
					Name: volName, DriverOpts: map[string]string{},
					Labels: map[string]string{}}}},
			"c": &compose.ComposeSpec{
				Volumes: compose.VolumesSection{"d": &compose.VolumeSpec{
					Name: volName, DriverOpts: map[string]string{},
					Labels: map[string]string{}}}}}
	}()
	OverlayVolume(cm, []plan.VolumePlan{vp})
	if !cm["a"].Volumes["b"].External || !cm["c"].Volumes["d"].External {
		t.Error("multiple overlay failed")
	}
}
