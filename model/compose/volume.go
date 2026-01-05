package compose

import "maps"

func (v *VolumeSpec) IsExternal() bool {
	return v.External
}

func (v *VolumeSpec) IsDefault() bool {
	return (v.Driver == "local" || v.Driver == "") && len(v.DriverOpts) == 0
}

func (v *VolumeSpec) IsLocalBind() bool {
	if v.Driver == "local" && len(v.DriverOpts) > 0 {
		if v.DriverOpts["type"] == "none" && v.DriverOpts["o"] == "bind" {
			return true
		}
	}
	return false
}

func (v *VolumeSpec) HasBindPath() (string, bool) {
	path, ok := v.DriverOpts["device"]
	return path, ok
}

func (v VolumeSpec) Equal(u VolumeSpec) bool {
	l := v.Name == u.Name && v.External == u.External && v.Driver == u.Driver
	m := maps.Equal(v.DriverOpts, u.DriverOpts) && maps.Equal(v.Labels, u.Labels)
	return l && m
}
