package plan

import (
	"fmt"
	orca "orca/helper"
	"orca/internal/config"
)

func BuildNetworkPlan(orcaRoot string,
	cfg *config.NetworkConfig) (*NetworkPlan, error) {
	plan := &NetworkPlan{}
	composes, err := GetComposes(orcaRoot)
	if err != nil {
		return nil, orca.OrcaError("collect networks failed", err)
	}
	for _, c := range composes {
		for k, n := range c.Spec.Networks {
			action := NetworkAction{
				Compose: c.From,
				Network: k,
			}
			switch {
			case k == "default" && n.Name != *cfg.Name:
				// デフォルト上書き
				action.Type = NetworkOverrideDefault
				action.Message = "default network is overridden to use shared network orca_network"
			case k != "default" && n.Name == *cfg.Name:
				// 競合削除
				action.Type = NetworkRemoveConflict
				action.Message = fmt.Sprintf("network %s conflicts with shared network and will be removed", n.Name)
			}
			if action.Type != "" { // 変更があるときだけplanに追加
				plan.Actions = append(plan.Actions, action)
			}
		}
	}
	return plan, nil
}
