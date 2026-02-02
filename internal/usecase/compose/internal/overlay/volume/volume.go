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
				Driver:     spec.Driver,
				DriverOpts: maps.Clone(spec.DriverOpts),
				External:   spec.External,
				Labels:     maps.Clone(spec.Labels),
				Name:       spec.Name,
			}
			if overlay.DriverOpts == nil {
				overlay.DriverOpts = map[string]string{}
			}
			if overlay.Labels == nil {
				overlay.Labels = map[string]string{}
			}
			overlay.Labels[label.LabelCluster] = cfg.Name
			overlay.Labels[label.LabelVolType] = string(vp.Type)
			switch vp.Type {
			case plan.VolumeExternal:
				// external
				overlay.External = true
				overlay.Driver = ""
				overlay.DriverOpts = map[string]string{}
			case plan.VolumeShared:
				// shared external+label
				overlay.External = true
				overlay.Driver = ""
				overlay.DriverOpts = map[string]string{}
				overlay.Labels[label.LabelVolShareBind] = vp.BindPath
			case plan.VolumeLocal:
				// local local bind
				overlay.External = false
				overlay.Driver = "local"
				overlay.DriverOpts["type"] = "none"
				overlay.DriverOpts["o"] = "bind"
				overlay.DriverOpts["device"] = vp.BindPath
			default:
			}
			cm[u.Compose].Volumes[u.Key] = overlay
		}
	}
}
