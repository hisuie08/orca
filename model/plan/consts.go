package plan

type PlanStatus string

const (
	StatusExist PlanStatus = "OK"
	StatusWarn  PlanStatus = "WARN"
	StatusError PlanStatus = "ERROR"
)

type NetworkActionType = string

const (
	NetworkOverrideDefault NetworkActionType = "override-default"
	NetworkRemoveConflict  NetworkActionType = "remove-conflict"
)

type VolumeType string

const (
	VolumeLocal    VolumeType = "local"
	VolumeShared   VolumeType = "shared"
	VolumeExternal VolumeType = "external"
)
