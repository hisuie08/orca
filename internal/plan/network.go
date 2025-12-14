package plan

import (
	orca "orca/helper"
	"orca/internal/config"
	"orca/internal/ostools"
)

type typeOverlay string

const (
	Replace typeOverlay = "replace"
	Remove  typeOverlay = "remove"
)

func collectNeedOverlay(orcaRoot string, name string) (
	map[typeOverlay][]string, error) {
	result := make(map[typeOverlay][]string)
	composes, err := GetComposes(orcaRoot)
	if err != nil {
		return nil, orca.OrcaError("collect networks failed", err)
	}
	// ネットワークセクションのエントリを回収
	for _, c := range composes {
		for k, n := range c.Spec.Networks {
			switch {
			case k != "default" && n.Name == name:
				// defaultでない競合する同名ネットワークがcompose内にあったら置換
				result[Remove] = append(result[Remove], c.From)
			case k == "default" && n.Name != name:
				// defaultが共通ネットワークと違うなら置き換え
				result[Replace] = append(result[Replace], c.From)
			default:
			}

		}
	}
	return result, nil

}
func BuildNetworkPlan(orcaRoot string, cfg *config.NetworkConfig) (
	*NetworkPlan, error) {
	plan := &NetworkPlan{}
	if !cfg.Enabled {
		return plan, nil
	}
	plan.Name = *cfg.Name
	exists := ostools.NetworkExists(*cfg.Name)
	plan.NeedCreate = !exists
	collect, err := collectNeedOverlay(orcaRoot, plan.Name)
	if err != nil {
		return nil, err
	}
	plan.Removed = collect[Remove]
	plan.Replaced = collect[Replace]
	return plan, err
}
