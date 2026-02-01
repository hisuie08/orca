package network

import (
	"orca/model/compose"
	"orca/model/plan"
)

func OverlayNetwork(cm compose.ComposeMap,
	np plan.NetworkPlan) {
	for _, a := range np.Actions {
		if np.SharedName == "" {
			panic("invalid NetworkPlan: sharedName is empty")
		}
		switch a.ActionType {
		case plan.NetworkOverrideDefault:
			spec := cm[a.Target.Compose].Networks[a.Target.Key]
			if spec == nil {
				panic("invalid NetworkPlan: target not found")
			}
			o := &compose.NetworkSpec{
				Name: np.SharedName,
				// orcaが作るか、既存のものを使う どちらにしろexternal
				External: true,
				// labelは保持
				Labels: spec.Labels,
			}
			cm[a.Target.Compose].Networks[a.Target.Key] = o

		case plan.NetworkRemoveConflict:
			delete(cm[a.Target.Compose].Networks, a.Target.Key)
		}
	}
}
