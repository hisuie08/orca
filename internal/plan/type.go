package plan

// ===========
//
//	Network
//
// ===========

type NetworkPlan struct {
	SharedName string // orcaが使う共有ネットワーク名
	Actions    map[string][]NetworkAction
}

type NetworkActionType = string

const (
	NetworkOverrideDefault NetworkActionType = "override-default"
	NetworkRemoveConflict  NetworkActionType = "remove-conflict"
)

type NetworkAction struct {
	Type    NetworkActionType
	Network string // 対象ネットワーク名（default or 削除対象）
	Message string // 人間向け補足
}

// =========
//
//	Volume
//
// =========
type VolumeType string

const (
	VolumeLocal    VolumeType = "local"
	VolumeShared   VolumeType = "shared"
	VolumeExternal VolumeType = "external"
)

type PlanStatus string

const (
	StatusExist  PlanStatus = "OK"
	StatusCreate PlanStatus = "CREATE"
	StatusWarn   PlanStatus = "WARN"
	StatusError  PlanStatus = "ERROR"
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

// ===========
//
//	OrcaPlan
//
// ===========
type OrcaPlan struct {
	Volumes  []VolumePlan
	Networks NetworkPlan
}
