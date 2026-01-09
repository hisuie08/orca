package plan

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
