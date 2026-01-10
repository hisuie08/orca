package volume

import (
	"orca/internal/context"
	"orca/internal/inspector"
	"orca/model/compose"
	"orca/model/plan"
	"path/filepath"
)

type VolumePlanContext interface {
	context.WithConfig
}

func BuildVolumePlan(ctx VolumePlanContext,
	cv []compose.CollectedVolume) []plan.VolumePlan {
	return buildVolumePlan(ctx, cv, inspector.NewDocker())
}

type dockerInspector interface {
	VolumeExists(name string) bool
	BindExists(dir string) bool
}

/*
VolumeRoot が nil または空文字列の場合、orca は volume 管理を無効とみなす
無効状態で BuildVolumePlan が呼ばれた場合は panic する
volume 管理を有効にするには、明示的に 非空の VolumeRoot を指定する必要がある
*/
func buildVolumePlan(
	ctx VolumePlanContext,
	cv []compose.CollectedVolume,
	di dockerInspector) []plan.VolumePlan {
	cfg := ctx.Config().Volume
	plans := []plan.VolumePlan{}
	if !cfg.Enabled() {
		panic("BuildVolumePlan called while volume management is disabled (VolumeRoot is nil)")
	}
	volumeRoot, err := filepath.Abs(*cfg.VolumeRoot)
	if err != nil {
		panic("volumeRoot could not be resolved")
	}
	// ボリュームは名前基準に処理
	for name, vols := range groupVolsByName(cv) {
		vp := plan.VolumePlan{
			Name:     name,
			UsedBy:   []plan.VolumeRef{},
			BindPath: "",
		}
		hasExternal := false
		// <name> を持つ各ボリューム定義を検証
		for _, v := range vols {
			vp.UsedBy = append(vp.UsedBy,
				plan.VolumeRef{
					Compose: v.Ref.Compose, Key: v.Ref.Key})
			if v.Spec.External {
				hasExternal = true
			}
			// ユーザー定義のバインド先があるなら回収
			// NOTE: 複数のバインドパスが定義されている場合、最後のパスが優先される
			if path, ok := v.Spec.HasBindPath(); ok {
				vp.BindPath = path
				vp.Reason = "user-defined bind path detected"
			}
		}
		switch {
		// 1. docker Volumeがすでに存在 -> externalとして 既存ボリュームを使用
		case di.VolumeExists(name):
			vp.Type = plan.VolumeExternal
			vp.BindPath = ""
			vp.NeedMkdir = false
			vp.Reason = "docker volume already exists"
		// 2. 少なくとも1つ external 指定がある -> 全て external に沿う
		case hasExternal:
			vp.Type = plan.VolumeExternal
			vp.BindPath = ""
			vp.NeedMkdir = false
			vp.Reason = "external volume defined in compose"
		// 3. 複数個所に定義 かつ 未存在 -> shared volumeとして orca管轄に昇格
		case len(vols) > 1 && !hasExternal:
			vp.Type = plan.VolumeShared
			vp.BindPath = filepath.Join(volumeRoot, name)
			vp.NeedMkdir = !di.BindExists(vp.BindPath)
			vp.Reason = "duplicated volume across compose"

		// 4. それ以外はConfigに従いlocal bind化
		default:
			vp.Type = plan.VolumeLocal
			if vp.BindPath == "" {
				abs, err := filepath.Abs(filepath.Join(volumeRoot, name))
				if err != nil {
					panic(err)
				}
				vp.BindPath = abs
			}
			vp.NeedMkdir = !di.BindExists(vp.BindPath)
			vp.Reason = "single compose volume"
		}

		plans = append(plans, vp)
	}
	return plans
}

// 名前基準にボリュームをグルーピングしなおす。重複やexternalの検出用
func groupVolsByName(vols []compose.CollectedVolume) map[string][]compose.CollectedVolume {
	groups := make(map[string][]compose.CollectedVolume)
	for _, v := range vols {
		// orcaが検討する必要がないボリュームはスキップ (tmpfs等)
		if !v.Spec.IsExternal() && !v.Spec.IsDefault() && !v.Spec.IsLocalBind() {
			continue
		}
		name := v.Spec.Name
		groups[name] = append(groups[name], v)
	}
	return groups
}
