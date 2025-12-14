package plan

import "orca/internal/compose"

// 列挙型系

type PlanStatus string

const (
	StatusOK    PlanStatus = "OK"
	StatusWarn  PlanStatus = "WARN"
	StatusError PlanStatus = "ERROR"
)

// CollectedSpec
// From: 定義されていたcompose
// Spec: 定義
type CollectedVolume struct {
	From string
	Spec *compose.VolumeSpec
}

type CollectedCompose struct {
	From string
	Spec *compose.ComposeSpec
}

type CollectedNetwork struct {
	From string
	Spec *compose.NetworkSpec
}

// Network
type OverlayType string

const (
	Replace OverlayType = "replace"
	Remove  OverlayType = "remove"
)

type NetworkPlan struct {
	Name       string
	NeedCreate bool
	Internal   bool
	Removed    []string
	Replaced   []string
}

// Volume
type VolumeType string

const (
	VolumeLocal    VolumeType = "local"
	VolumeShared   VolumeType = "shared"
	VolumeExternal VolumeType = "external"
)

type VolumePlan struct {
	Name   string
	Type   VolumeType
	UsedBy []string

	// local / shared のみ意味を持つ
	BindPath  string
	NeedMkdir bool

	// 判断理由（ログ・plan表示・デバッグ用）
	Reason string

	// up するとエラーになる可能性がある場合
	Warnings []string
}

type OrcaPlan struct {
	Volumes  []VolumePlan
	Networks []NetworkPlan
}
