package overlay

import (
	"orca/internal/context"
	"orca/internal/usecase/compose/internal/overlay/network"
	"orca/internal/usecase/compose/internal/overlay/volume"
	"orca/model/compose"
	"orca/model/plan"
)

var _ composeOverlayer = (*overlayer)(nil)

type composeOverlayer interface {
	OverlayVolume([]plan.VolumePlan)
	OverlayNetwork(plan.NetworkPlan)
}

type overlayContext interface {
	context.WithConfig
}

type overlayer struct {
	ctx overlayContext
	cm  compose.ComposeMap
}

func ComposeOverlayer(ctx overlayContext, cm compose.ComposeMap) *overlayer {
	return &overlayer{cm: cm, ctx: ctx}
}

func (o *overlayer) OverlayVolume(vps []plan.VolumePlan) {
	if o.ctx.Config().Volume.VolumeRoot != nil {
		volume.OverlayVolume(o.cm, vps)
	}
}

func (o *overlayer) OverlayNetwork(np plan.NetworkPlan) {
	if o.ctx.Config().Network.Enabled {
		network.OverlayNetwork(o.cm, np)
	}
}
