package network

import (
	"orca/internal/capability"
	"orca/internal/inspector"
	"orca/model/compose"
	"orca/model/config"
	"orca/model/plan"
)

type NetworkPlanCapability interface {
	capability.WithConfig
}

func BuildNetworkPlan(caps NetworkPlanCapability,
	cn []compose.CollectedNetwork) plan.NetworkPlan {
	return buildNetworkPlan(caps, cn, inspector.NewDocker())
}

type dockerInspector interface {
	NetworkExists(name string) bool
}

func buildNetworkPlan(
	caps NetworkPlanCapability,
	cn []compose.CollectedNetwork,
	di dockerInspector) plan.NetworkPlan {
	cfg := caps.Config().Network
	np := plan.NetworkPlan{Actions: []plan.NetworkAction{}}
	if !cfg.Enabled {
		return np
	}
	np.SharedName = cfg.Name
	exists := di.NetworkExists(cfg.Name)
	np.Create = !exists
	for _, net := range cn {
		if o, ok := buildAction(net, cfg); ok {
			np.Actions = append(np.Actions, o)
		}
	}
	return np
}

func buildAction(net compose.CollectedNetwork, cfg config.NetworkConfig,
) (plan.NetworkAction, bool) {
	action := plan.NetworkAction{}
	action.Target = plan.NetworkRef{
		Compose: net.Ref.Compose,
		Key:     net.Ref.Key,
	}
	switch {
	case net.Ref.Key == "default" && net.Spec.Name != cfg.Name:
		// デフォルトは上書き
		action.ActionType = plan.NetworkOverrideDefault

	case net.Ref.Key != "default" && net.Spec.Name == cfg.Name:
		// 名前が競合するnetworkはcomposeから削除
		action.ActionType = plan.NetworkRemoveConflict
	}
	return action, action.ActionType != ""
}
