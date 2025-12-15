package plan

import (
	"fmt"
	"io"
	orca "orca/helper"
	"orca/internal/config"
	"sort"
)

func BuildNetworkPlan(orcaRoot string,
	cfg *config.NetworkConfig) (*NetworkPlan, error) {
	plan := &NetworkPlan{
		SharedName: *cfg.Name,
		Actions:    map[string][]NetworkAction{},
	}
	composes, err := GetComposes(orcaRoot)
	if err != nil {
		return nil, orca.OrcaError("collect networks failed", err)
	}
	for _, c := range composes {
		for k, n := range c.Spec.Networks {
			action := NetworkAction{
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
				plan.Actions[c.From] = append(plan.Actions[c.From], action)
			}
		}
	}
	return plan, nil
}

func PrintNetworkPlan(p NetworkPlan, w io.Writer, c *orca.Colorizer) {
	title := "NETWORK PLAN"
	fmt.Fprintf(w, "%s\n", title)
	fmt.Fprintf(w, "Shared network: %s\n", p.SharedName)

	// compose名でソート
	composes := make([]string, 0, len(p.Actions))
	for k := range p.Actions {
		composes = append(composes, k)
	}
	sort.Strings(composes)

	for _, compose := range composes {
		actions := p.Actions[compose]
		if len(actions) == 0 {
			continue
		}

		fmt.Fprintf(w, "[%s]\n", compose)

		for _, a := range actions {
			switch a.Type {
			case NetworkOverrideDefault:
				label := c.Blue("override")
				fmt.Fprintf(w, "  %s default → %s\n", label, p.SharedName)
			case NetworkRemoveConflict:
				label := c.Yellow("remove")
				fmt.Fprintf(w, "  %s network  → %s (name conflict)\n", label, a.Network)
			}
		}
		fmt.Fprintln(w)
	}
}
