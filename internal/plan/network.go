package plan

import (
	"fmt"
	"io"
	orca "orca/helper"
	"orca/internal/compose"
	"orca/internal/config"
	"sort"
)

func BuildNetworkPlan(networks []compose.CollectedNetwork,
	cfg *config.NetworkConfig) *NetworkPlan {
	plan := &NetworkPlan{
		SharedName: *cfg.Name,
		Actions:    map[string][]NetworkAction{},
	}
	for _, n := range networks {
		action := NetworkAction{
			Network: n.Key,
		}
		switch {
		case n.Key == "default" && n.Spec.Name != *cfg.Name:
			// デフォルト上書き
			action.Type = NetworkOverrideDefault
			action.Message = "default network is overridden to use shared network orca_network"
		case n.Key != "default" && n.Spec.Name == *cfg.Name:
			// 競合削除
			action.Type = NetworkRemoveConflict
			action.Message = fmt.Sprintf("network %s conflicts with shared network and will be removed", n.Spec.Name)
		}
		if action.Type != "" { // 変更があるときだけplanに追加
			plan.Actions[n.From] = append(plan.Actions[n.From], action)
		}
	}
	return plan
}

func PrintNetworkPlan(p NetworkPlan, w io.Writer, c *orca.Colorizer) {
	title := "[NETWORK PLAN]"
	fmt.Fprintf(w, "%s\n", title)
	fmt.Fprintf(w, "SHARED NETWORK: %s\n", p.SharedName)

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

		fmt.Fprintf(w, "- %s\n", compose)

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
	}
	fmt.Fprintln(w)
}
