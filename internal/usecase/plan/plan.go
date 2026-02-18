package plan

import (
	"orca/internal/capability"
	"orca/internal/usecase/plan/network"
	"orca/internal/usecase/plan/volume"
	"orca/model/compose"
	"orca/model/plan"
	"sort"
)

type orcaPlanCapability interface {
	capability.WithConfig
}

func BuildOrcaPlan(caps orcaPlanCapability, cm compose.ComposeMap) plan.OrcaPlan {
	result := plan.OrcaPlan{
		Name: caps.Config().Name, ComposeDirs: []string{},
		Volumes: []plan.VolumePlan{}, Networks: plan.NetworkPlan{}}
	for name, _ := range cm {
		result.ComposeDirs = append(result.ComposeDirs, name)
	}
	sort.Strings(result.ComposeDirs)
	if caps.Config().Volume.Enabled() {
		result.Volumes = volume.BuildVolumePlan(caps, cm.CollectVolumes())
	}
	if caps.Config().Network.Enabled {
		result.Networks = network.BuildNetworkPlan(caps, cm.CollectNetworks())
	}
	return result
}
