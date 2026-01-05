package overlay

import (
	"orca/internal/context"
	"orca/internal/usecase/compose/overlay/network"
	"orca/internal/usecase/compose/overlay/volume"
	"orca/model/compose"
	"orca/model/plan"
)

var _ composeOverlayer = (*overlayer)(nil)

type composeOverlayer interface {
	OverlayVolume([]plan.VolumePlan)
	OverlayNetwork(plan.NetworkPlan)
}

type OverlayContext interface {
	context.WithConfig
}

type overlayer struct {
	ctx OverlayContext
	cm  compose.ComposeMap
}

func ComposeOverlayer(ctx OverlayContext, cm compose.ComposeMap) *overlayer {
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
