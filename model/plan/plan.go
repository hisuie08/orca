package plan

import "orca/model/config"

// ===========
//
//	OrcaPlan
//
// ===========
type OrcaPlan struct {
	Config      config.ResolvedConfig
	ComposeDirs []string
	Volumes     []VolumePlan
	Networks    NetworkPlan
}

// ===========
//
//	Network
//
// ===========
type NetworkPlan struct {
	SharedName string // Shared network name used by orca
	Create     bool
	Actions    []NetworkAction // Planned Overlay
}

type NetworkRef struct {
	Compose string
	Key     string // Target network key ("default" or to be deleted)
}

type NetworkAction struct {
	Target     NetworkRef
	ActionType NetworkActionType
}

// =========
//
//	Volume
//
// =========
type VolumeRef struct {
	Compose string
	Key     string
}

type VolumePlan struct {
	Name   string
	Type   VolumeType
	UsedBy []VolumeRef

	// Only meaningful for local/shared
	BindPath  string
	NeedMkdir bool

	// Reason for decision (for logging, displaying plan, and debugging)
	Reason string
}
