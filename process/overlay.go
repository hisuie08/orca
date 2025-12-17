package process

import (
	"orca/internal/compose"
	"orca/internal/config"
	"orca/internal/plan"
)

type mapCompose map[string]*compose.ComposeSpec

func PlanProcess(orcaRoot string, cfg *config.OrcaConfig, show bool) error {
	cfg.Resolve(orcaRoot)
	// compose構成ロード
	composeMap, err := compose.ComposeMap(orcaRoot)
	if err != nil {
		return err
	}
	// VolumePlan構築と適用
	volumes := compose.CollectVolumes(*composeMap)
	volPlans := plan.BuildVolumePlan(volumes, cfg.Volume)
	if err := ApplyVolumePlan(*composeMap, volPlans); err != nil {
		return err
	}
	// NetworkPlan構築と適用
	netPlan := plan.BuildNetworkPlan(compose.CollectComposes(*composeMap), cfg.Network)
	if err := ApplyNetworkPlan(*composeMap, netPlan); err != nil {
		return err
	}

	return nil
}

// Volume
func ApplyVolumePlan(m mapCompose, plans []plan.VolumePlan) error {
	for _, p := range plans {
		for _, u := range p.UsedBy {
			for k, v := range m[u].Volumes {
				if v.Name == p.Name {
					switch p.Type {
					case plan.VolumeLocal:
						m[u].Volumes[k] = ApplyLocalBind(v, p.BindPath)
						//
					case plan.VolumeExternal:
						m[u].Volumes[k] = ApplyExternal(v)
						//
					case plan.VolumeShared:
						m[u].Volumes[k] = ApplyExternal(v)
					}
				}
			}

		}

	}
	return nil
}
func ApplyLocalBind(v *compose.VolumeSpec, bindPath string) *compose.VolumeSpec {
	v.Driver = "local"
	if v.DriverOpts == nil {
		v.DriverOpts = map[string]string{}
	}
	v.DriverOpts["type"] = "none"
	v.DriverOpts["o"] = "bind"
	v.DriverOpts["device"] = bindPath
	return v
}

func ApplyExternal(v *compose.VolumeSpec) *compose.VolumeSpec {
	v.Driver = ""
	for k := range v.DriverOpts {
		delete(v.DriverOpts, k)
	}
	v.External = true
	return v
}

// Network
func ApplyNetworkPlan(m mapCompose, plans plan.NetworkPlan) error {
	for c, actions := range plans.Actions {
		for _, action := range actions {
			switch action.Type {
			case plan.NetworkOverrideDefault:
				m[c].Networks[action.Network] = &compose.NetworkSpec{
					Name: plans.SharedName,
					// orcaが作るか、既存のものを使うのでどちらにしろexternal
					External: true,
				}
			case plan.NetworkRemoveConflict:
				delete(m[c].Networks, action.Network)
			}
		}
	}
	return nil
}
