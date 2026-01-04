package compose

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
