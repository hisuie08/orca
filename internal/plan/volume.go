package plan

import (
	"fmt"
	orca "orca/helper"
	"orca/internal/config"
	"orca/internal/ostools"
	"path/filepath"
)

// volumeをかき集める
func collectVolumes(orcaRoot string) ([]CollectedVolume, error) {
	result := []CollectedVolume{}

	composes, err := GetComposes(orcaRoot)
	if err != nil {
		return nil, orca.OrcaError("collect volumes failed", err)
	}
	for _, c := range composes {
		for _, v := range c.Spec.Volumes {
			result = append(result, CollectedVolume{
				From: filepath.Base(c.From),
				Spec: v,
			})
		}
	}
	return result, nil
}

// 名前基準にボリュームをグルーピングしなおし。重複やexternalの検出用
func groupVolumes(vols []CollectedVolume) map[string][]CollectedVolume {
	groups := make(map[string][]CollectedVolume)
	for _, v := range vols {
		// orcaがオーバーレイする必要がないボリュームはスキップ
		// 照合のためにexternalは一旦回収
		if !v.Spec.NeedsOrcaOverlay() && !v.Spec.External {
			continue
		}
		name := v.Spec.Name
		groups[name] = append(groups[name], v)
	}
	return groups
}

// ボリュームのPlanを構築する
func buildPlan(
	groups map[string][]CollectedVolume,
	cfg *config.VolumeConfig,
) []VolumePlan {

	plans := []VolumePlan{}

	// 各グループを検証するよ
	for name, vols := range groups {
		usedBy := []string{}
		hasExternal := false
		var customPath string
		// グループ内のボリュームを検証
		for _, v := range vols {
			usedBy = append(usedBy, v.From)
			if v.Spec.External {
				hasExternal = true
			}
			// ユーザー定義のバインド先があるなら回収
			if device, ok := v.Spec.DriverOpts["device"]; ok {
				customPath = device
			}
		}

		exists := ostools.VolumeExists(name)

		plan := VolumePlan{
			Name:     name,
			UsedBy:   usedBy,
			BindPath: customPath,
		}

		switch {
		// 1. すでに存在
		case exists:
			plan.Type = VolumeExternal
			plan.Reason = "docker volume already exists"

		// 2. 重複 + external 指定あり
		case len(vols) > 1 && hasExternal:
			plan.Type = VolumeExternal
			plan.Reason = "external volume defined in compose"

		// 3. 重複 + 未存在
		case len(vols) > 1:
			plan.Type = VolumeShared

			plan.BindPath = filepath.Join(*cfg.VolumeRoot, name)
			plan.NeedMkdir = !ostools.DirExists(plan.BindPath)
			plan.Reason = "duplicated volume across compose"

		// 4. 単一 かつ external
		case len(vols) == 1 && hasExternal:
			plan.Type = VolumeExternal
			plan.NeedMkdir = !ostools.DirExists(plan.BindPath)
			plan.Reason = "external volume defined in compose"
		// 5．単一compose
		default:
			plan.Type = VolumeLocal
			if plan.BindPath == "" {
				plan.BindPath = filepath.Join(*cfg.VolumeRoot, name)
			}
			plan.NeedMkdir = !ostools.DirExists(plan.BindPath)
			plan.Reason = "single compose volume"
		}

		// Warning 判定
		// パスが存在しないのにensure_path=falseで作成が許可されていないとき
		switch {
		case plan.Type != VolumeExternal &&
			plan.NeedMkdir &&
			!cfg.EnsurePath:
			warningMsg := fmt.Sprintf("bind path %s does not exist and ensure_path=false", plan.BindPath)
			plan.Warnings = append(plan.Warnings, warningMsg)
		case plan.Type == VolumeExternal && !ostools.VolumeExists(plan.Name):
			// externalだけどplan時点でボリュームが存在しないとき
			warningMsg := fmt.Sprintf("external volume %s does not exist", plan.Name)
			plan.Warnings = append(plan.Warnings, warningMsg)
		}

		plans = append(plans, plan)
	}
	return plans
}

func BuildVolumePlan(orcaRoot string, cfg *config.VolumeConfig) (
	[]VolumePlan, error) {
	if collect, err := collectVolumes(orcaRoot); err != nil {
		return nil, err
	} else {
		group := groupVolumes(collect)
		return buildPlan(group, cfg), nil
	}
}
