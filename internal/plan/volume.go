package plan

import (
	"fmt"
	orca "orca/helper"
	"orca/internal/compose"
	"orca/internal/config"
	"orca/internal/ostools"
	"path/filepath"
	"strings"
)

// volumeをかき集める

// 名前基準にボリュームをグルーピングしなおし。重複やexternalの検出用
func groupVolumes(vols []compose.CollectedVolume) map[string][]compose.CollectedVolume {
	groups := make(map[string][]compose.CollectedVolume)
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
func buildVolPlan(
	groups map[string][]compose.CollectedVolume,
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

func BuildVolumePlan(collect []compose.CollectedVolume, cfg *config.VolumeConfig) []VolumePlan {
	group := groupVolumes(collect)
	return buildVolPlan(group, cfg)
}

func toVolPlanRow(plan VolumePlan, c *orca.Colorizer) []string {
	status := StatusOK
	if len(plan.Warnings) > 0 {
		status = StatusWarn
	}

	bind := plan.BindPath
	if bind == "" {
		bind = "-"
	}
	typ := string(plan.Type)
	switch plan.Type {
	case VolumeShared:
		typ = c.Blue("shared")
	case VolumeLocal:
		typ = c.Green("local")
	case VolumeExternal:
		typ = c.Gray("external")
	}
	stat := string(status)
	// TODO: NeedMkDirを反映したstatusの細かい出し分け
	switch status {
	case StatusOK:
		stat = c.Green(string(StatusOK))
	case StatusWarn:
		stat = c.Yellow(string(StatusWarn))
	}
	return []string{
		plan.Name,
		typ,
		strings.Join(plan.UsedBy, ","),
		bind,
		stat,
	}
}

func PrintVolumePlanTable(plans []VolumePlan, printer *orca.Printer) {
	title := "[VOLUME PLAN]"
	headers := []string{"NAME", "TYPE", "USED BY", "BIND PATH", "STATUS"}

	rows := make([][]string, 0, len(plans))
	for _, p := range plans {
		rows = append(rows, toVolPlanRow(p, &printer.C))
	}
	printer.PrintTable(title, headers, rows)

}
