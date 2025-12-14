package plan

import "orca/internal/compose"

// Orcaの変更箇所記録用

// ボリューム
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

type NetworkPlan struct {
	Name       string
	NeedCreate bool
	Internal   bool
	Removed    []string
	Replaced   []string
}

type OrcaPlan struct {
	Volumes  []VolumePlan
	Networks []NetworkPlan
}

func InitPlan(name string, orca_root string) *OrcaPlan {
	return &OrcaPlan{
		Volumes:  []VolumePlan{},
		Networks: []NetworkPlan{}}
}

func (p *OrcaPlan) RegisterVolume(v VolumePlan) {
	p.Volumes = append(p.Volumes, v)
}
func (p *OrcaPlan) RegisterNetwork(n NetworkPlan) {
	p.Networks = append(p.Networks, n)
}

func (p *OrcaPlan) Dump() string {

	return ""
}
