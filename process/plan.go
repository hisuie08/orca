package process

import (
	"bytes"
	"orca/infra/applier"
	ins "orca/infra/inspector"
	"orca/internal/compose"
	"orca/internal/context"
	"orca/internal/plan"
)

var apl = applier.Applier

func PlanProcess(ctx context.OrcaContext) error {
	cmp, err := compose.GetAllCompose(ins.Compose(ctx.OrcaRoot))
	if err != nil {
		return err
	}
	volumePlan := func() []plan.VolumePlan {
		//DEBUG
		ctx.Printer.Printf("%s\n", *ctx.Config.Volume.VolumeRoot)
		if ctx.Config.Volume.VolumeRoot != nil {
			return plan.BuildVolumePlan(*cmp, &ctx.Config.Volume, ins.Docker)
		} else {
			return []plan.VolumePlan{}
		}
	}()
	networkPlan := func() plan.NetworkPlan {
		if ctx.Config.Network.Enabled {
			return plan.BuildNetworkPlan(*cmp, &ctx.Config.Network)
		}
		return plan.NetworkPlan{}
	}()

	if err := ApplyVolumeCompose(*cmp, volumePlan); err != nil {
		return err
	}
	if err := ApplyNetworkCompose(*cmp, networkPlan); err != nil {
		return err
	}

	b := bytes.Buffer{}
	plan.DumpPlan(ctx, volumePlan, networkPlan, &b)
	b.WriteTo(ctx.Printer.W)
	composes, err := cmp.DumpAllComposes(apl.Compose(ctx.OrcaRoot))
	if err != nil {
		return err
	}
	for _, v := range composes {
		if ctx.RunMode == context.ModeDryRun {
			ctx.Printer.PrintDRY("Created " + v + "\n")
		} else {
			ctx.Printer.Printf("Created %s\n", v)
		}
	}

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
						m[u].Volumes[k] = applyLocalBind(v, p.BindPath)
						//
					case plan.VolumeExternal:
						m[u].Volumes[k] = applyExternal(v)
						//
					case plan.VolumeShared:
						m[u].Volumes[k] = applyExternal(v)
					}
				}
			}

		}

	}
	return nil
}
func applyLocalBind(v *compose.VolumeSpec, bindPath string) *compose.VolumeSpec {
	v.Driver = "local"
	if v.DriverOpts == nil {
		v.DriverOpts = map[string]string{}
	}
	v.DriverOpts["type"] = "none"
	v.DriverOpts["o"] = "bind"
	v.DriverOpts["device"] = bindPath
	return v
}

func applyExternal(v *compose.VolumeSpec) *compose.VolumeSpec {
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
