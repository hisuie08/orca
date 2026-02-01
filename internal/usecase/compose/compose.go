package compose

import (
	"orca/internal/context"
	"orca/internal/usecase/compose/internal/getall"
	"orca/internal/usecase/compose/internal/overlay"
	"orca/model/compose"
	"orca/model/plan"
)

type GetAllComposeContext interface {
	context.WithRoot
}

func GetAllCompose(ctx GetAllComposeContext) (compose.ComposeMap, error) {
	return getall.GetAllCompose(ctx)
}

type OverlayContext interface {
	context.WithConfig
}
type overlayer interface {
	OverlayVolume([]plan.VolumePlan)
	OverlayNetwork(plan.NetworkPlan)
}

func ComposeOverlayer(ctx OverlayContext, cm compose.ComposeMap) overlayer {
	return overlay.ComposeOverlayer(ctx, cm)
}
