package plan

import (
	"fmt"
	orca "orca/helper"
	"orca/infra/inspector"
	"orca/internal/compose"
	"orca/internal/config"
	"orca/internal/ostools"
	"path/filepath"
	"strings"
)

// Planに考慮すべきボリュームか判断
// デフォルト/external/local+bind
func shouldConsider(v *compose.VolumeSpec) bool {
	// case 1: external:true か ドライバ設定無し
	if v.External || v.Driver == "" {
		return true
	}

	// case 2: driver=local かつ driver_optsなし
	if v.Driver == "local" && len(v.DriverOpts) == 0 {
		return true
	}

	// case 3: driver=local + bind だが deviceのパスが存在しない
	if v.Driver == "local" && len(v.DriverOpts) > 0 {
		if v.DriverOpts["type"] == "none" && v.DriverOpts["o"] == "bind" {
			if !ostools.DirExists(v.DriverOpts["device"]) {
				return true
			}
		}
	}
	return false
}

// 名前基準にボリュームをグルーピングしなおし。重複やexternalの検出用
func groupVolumes(vols []compose.CollectedVolume) map[string][]compose.CollectedVolume {
	groups := make(map[string][]compose.CollectedVolume)
	for _, v := range vols {
		// orcaが検討する必要がないボリュームはスキップ
		if !shouldConsider(v.Spec) {
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
	cfg *config.ResolvedVolume,
	i inspector.VolumeInspector,
) []VolumePlan {

	plans := []VolumePlan{}

	// 各グループを検証するよ
	for name, vols := range groups {
		usedBy := []string{}
		hasExternal := false
		var customPath string
		// グループ内のボリュームを検証
		for _, v := range vols {
			usedBy = append(usedBy, v.From.Compose)
			if v.Spec.External {
				hasExternal = true
			}
			// ユーザー定義のバインド先があるなら回収
			if device, ok := v.Spec.DriverOpts["device"]; ok {
				customPath = device
			}
		}

		exists := i.VolumeExists(name)

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
		case plan.Type == VolumeExternal && !i.VolumeExists(plan.Name):
			// externalだけどplan時点でボリュームが存在しないとき
			warningMsg := fmt.Sprintf("external volume %s does not exist", plan.Name)
			plan.Warnings = append(plan.Warnings, warningMsg)
		}

		plans = append(plans, plan)
	}
	return plans
}

func BuildVolumePlan(
	collect []compose.CollectedVolume,
	cfg *config.ResolvedVolume,
	i inspector.VolumeInspector) []VolumePlan {
	group := groupVolumes(collect)
	return buildVolPlan(group, cfg, i)
}

func toVolPlanRow(plan VolumePlan, c *orca.Colorizer) []string {
	status := StatusExist
	if plan.NeedMkdir {
		status = StatusCreate
	}
	if len(plan.Warnings) > 0 {
		status = StatusWarn
	}

	bind, err := filepath.Abs(plan.BindPath)
	if err != nil {
		status = StatusError
	}
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
	case StatusExist:
		stat = c.Green(string(StatusExist))
	case StatusCreate:
		stat = c.Green(string(StatusCreate))

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
