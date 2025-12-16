package compose

import (
	orca "orca/helper"
	"orca/internal/ostools"
	"os"
	"path/filepath"
)

// 複数composeをかき集めるやつ

func CollectComposes(orcaRoot string) ([]CollectedCompose, error) {
	result := []CollectedCompose{}
	dirs, err := ostools.Directories(orcaRoot)
	if err != nil {
		return nil, orca.OrcaError("collect volumes failed", err)
	}
	// HACK: 駆け上がり止めcompose
	stopperCompose := filepath.Join(orcaRoot, "compose.yml")
	ostools.CreateFile(stopperCompose, []byte{})
	for _, dir := range dirs {
		composeYml, err := ostools.ComposeConfig(dir)
		if err != nil {
			// compose configが失敗したディレクトリはスキップ
			// （compose.ymlが存在しない等）
			continue
		}
		cmps, err := ParseCompose(composeYml)
		if err != nil {
			return nil, orca.OrcaError("compose parse failed", err)
		}
		result = append(result, CollectedCompose{filepath.Base(dir), cmps})

	}
	os.Remove(stopperCompose) //駆け上がり止めcomposeの削除
	return result, nil
}


func CollectVolumes(composes []CollectedCompose) []CollectedVolume {
	result := []CollectedVolume{}
	for _, c := range composes {
		for _, v := range c.Spec.Volumes {
			result = append(result, CollectedVolume{
				From: filepath.Base(c.From),
				Spec: v,
			})
		}
	}
	return result
}