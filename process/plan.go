package process

import (
	"io"
	"orca/internal/compose"
	"orca/internal/context"
	"orca/internal/plan"
)

func BuildPlan(ctx context.OrcaContext) error {
	cmp, err := compose.GetAllCompose(ctx.OrcaRoot, ctx.ComposeInspector)
	if err != nil {
		return err
	}
	vol := cmp.CollectVolumes()
	net := cmp.CollectComposes()
	volumePlan := plan.BuildVolumePlan(vol, &ctx.Config.Volume, ctx.DockerInspector)

	networkPlan := plan.BuildNetworkPlan(net, &ctx.Config.Network)
	if err := ApplyVolumeCompose(*cmp, volumePlan); err != nil {
		return err
	}
	if err := ApplyNetworkCompose(*cmp, networkPlan); err != nil {
		return err
	}
	// o, _ := os.OpenFile("./log.txt", os.O_WRONLY|os.O_CREATE, 0666)
	// printer.W = o
	// printer.C.Enabled = false
	plan.DumpPlan(ctx, volumePlan, networkPlan, ctx.Printer.W)
	return nil
}

// Volume
func ApplyVolumeCompose(m compose.ComposeMap, plans []plan.VolumePlan) error {
	for _, p := range plans {
		for _, u := range p.UsedBy {
			for k, v := range m[u].Volumes {
				// NOTE:
				// docker compose config により volume > Name は一意に正規化されている。
				// map key は保持していないため Name で線形探索する。
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
func ApplyNetworkCompose(m compose.ComposeMap, plans plan.NetworkPlan) error {
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

func Summary(w io.Writer) {

}
