package plan

import (
	"orca/internal/context"
	"orca/internal/usecase/plan/network"
	"orca/internal/usecase/plan/volume"
	"orca/model/compose"
	"orca/model/plan"
	"sort"
)

type orcaPlanContext interface {
	context.WithConfig
}

func BuildOrcaPlan(ctx orcaPlanContext, cm compose.ComposeMap) plan.OrcaPlan {
	result := plan.OrcaPlan{
		Name: ctx.Config().Name, ComposeDirs: []string{},
		Volumes: []plan.VolumePlan{}, Networks: plan.NetworkPlan{}}
	for name, _ := range cm {
		result.ComposeDirs = append(result.ComposeDirs, name)
	}
	sort.Strings(result.ComposeDirs)
	if ctx.Config().Volume.Enabled() {
		result.Volumes = volume.BuildVolumePlan(ctx, cm.CollectVolumes())
	}
	if ctx.Config().Network.Enabled {
		result.Networks = network.BuildNetworkPlan(ctx, cm.CollectNetworks())
	}
	return result
}
