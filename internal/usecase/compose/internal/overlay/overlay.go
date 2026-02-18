package overlay

import (
	"orca/internal/capability"
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

type overlayCapability interface {
	capability.WithConfig
}

type overlayer struct {
	caps overlayCapability
	cm   compose.ComposeMap
}

func ComposeOverlayer(caps overlayCapability, cm compose.ComposeMap) *overlayer {
	return &overlayer{cm: cm, caps: caps}
}

func (o *overlayer) OverlayVolume(vps []plan.VolumePlan) {
	if o.caps.Config().Volume.VolumeRoot != nil {
		volume.OverlayVolume(*o.caps.Config(), o.cm, vps)
	}
}

func (o *overlayer) OverlayNetwork(np plan.NetworkPlan) {
	if o.caps.Config().Network.Enabled {
		network.OverlayNetwork(*o.caps.Config(), o.cm, np)
	}
}
