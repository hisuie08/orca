package volume

import (
	"maps"
	"orca/internal/consts/label"
	"orca/model/compose"
	"orca/model/config"
	"orca/model/plan"
)

func OverlayVolume(cfg config.OrcaConfig, cm compose.ComposeMap, vps []plan.VolumePlan) {
	for _, vp := range vps {
		for _, u := range vp.UsedBy {
			spec := cm[u.Compose].Volumes[u.Key]
			overlay := &compose.VolumeSpec{
				External: spec.External,
				Name:     spec.Name,
			}
			switch vp.Type {
			case plan.VolumeExternal:
				// external
				overlay.External = true
			case plan.VolumeShared:
				// shared external
				overlay.External = true
			case plan.VolumeLocal:
				// local local bind
				overlay.Driver = spec.Driver
				overlay.DriverOpts = maps.Clone(spec.DriverOpts)
				overlay.Labels = maps.Clone(spec.Labels)
				overlay.External = false
				overlay.Driver = "local"
				if overlay.DriverOpts == nil {
					overlay.DriverOpts = map[string]string{}
				}
				overlay.DriverOpts["type"] = "none"
				overlay.DriverOpts["o"] = "bind"
				overlay.DriverOpts["device"] = vp.BindPath
				if overlay.Labels == nil {
					overlay.Labels = map[string]string{}
				}
				overlay.Labels[label.LabelCluster] = cfg.Name
				overlay.Labels[label.LabelVolType] = string(vp.Type)
			default:
			}
			cm[u.Compose].Volumes[u.Key] = overlay
		}
	}
}
