package compose

import (
	"orca/internal/capability"
	"orca/internal/usecase/compose/internal/getall"
	"orca/internal/usecase/compose/internal/overlay"
	"orca/model/compose"
	"orca/model/plan"
)

type GetAllComposeCapability interface {
	capability.WithRoot
}

func GetAllCompose(caps GetAllComposeCapability) (compose.ComposeMap, error) {
	return getall.GetAllCompose(caps)
}

type OverlayCapability interface {
	capability.WithConfig
}
type overlayer interface {
	OverlayVolume([]plan.VolumePlan)
	OverlayNetwork(plan.NetworkPlan)
}

func ComposeOverlayer(caps OverlayCapability, cm compose.ComposeMap) overlayer {
	return overlay.ComposeOverlayer(caps, cm)
}
