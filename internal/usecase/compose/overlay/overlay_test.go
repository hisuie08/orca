package overlay

import (
	"orca/internal/context"
	"orca/model/compose"
	"orca/model/config"
	"orca/model/plan"
	"testing"
)

var (
	cRef    = "a"
	vRef    = "a_vol"
	nRef    = "default"
	volRoot = "testroot"
	netName = "testnet"
)

func fakeMap() compose.ComposeMap {
	return compose.ComposeMap{cRef: &compose.ComposeSpec{
		Volumes: compose.VolumesSection{
			vRef: &compose.VolumeSpec{Driver: "local", Name: vRef}},
		Networks: compose.NetworksSection{
			nRef: &compose.NetworkSpec{Name: cRef + "_net"}},
	}}
}
func fakeCtx(volume *string, network bool) overlayContext {
	ctx := context.New().WithConfig(&config.OrcaConfig{
		Volume:  config.VolumeConfig{VolumeRoot: volume},
		Network: config.NetworkConfig{Enabled: network, Name: netName}})
	return &ctx
}

func fakePlan(cfg config.OrcaConfig) ([]plan.VolumePlan, plan.NetworkPlan) {
	vp := []plan.VolumePlan{{UsedBy: []plan.VolumeRef{
		{Compose: cRef, Key: vRef}}, Type: plan.VolumeExternal}}
	np := plan.NetworkPlan{SharedName: cfg.Network.Name,
		Actions: []plan.NetworkAction{{Target: plan.NetworkRef{
			Compose: cRef, Key: nRef}, ActionType: plan.NetworkOverrideDefault}}}
	return vp, np
}

func TestOverlay(t *testing.T) {
	testCases := []struct {
		name string
		ctx  overlayContext
		want func(compose.ComposeMap) bool
	}{
		{name: "network_and_volume_enabled", ctx: fakeCtx(&volRoot, true),
			want: func(nm compose.ComposeMap) bool {
				nv := nm[cRef].Volumes[vRef]
				nn := nm[cRef].Networks[nRef]
				return nv.External && nv.Driver == "" &&
					nn.Name == netName && nn.External
			}},
		{name: "network_and_volume_disabled", ctx: fakeCtx(nil, false),
			want: func(nm compose.ComposeMap) bool {
				nv := nm[cRef].Volumes[vRef]
				nn := nm[cRef].Networks[nRef]
				return !nv.External && nv.Driver == "local" &&
					nn.Name == cRef+"_net" && !nn.External
			}},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			cm := fakeMap()
			o := ComposeOverlayer(tt.ctx, cm)
			vps, np := fakePlan(*tt.ctx.Config())
			o.OverlayNetwork(np)
			o.OverlayVolume(vps)
			if !tt.want(cm) {
				t.Error("overlay failed")
			}
		})
	}
}
