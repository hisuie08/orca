package plan

import (
	orca "orca/helper"
	"orca/internal/compose"
	"orca/internal/ostools"
	"path/filepath"
)

func GetComposes(orcaRoot string) ([]CollectedCompose, error) {
	result := []CollectedCompose{}
	dirs, err := ostools.Directories(orcaRoot)
	if err != nil {
		return nil, orca.OrcaError("collect volumes failed", err)
	}
	for _, dir := range dirs {
		composeYml, cfgerr := ostools.ComposeConfig(dir)
		if cfgerr != nil {
			// compose configが失敗したディレクトリはスキップ
			// （compose.ymlが存在しない等）
			continue
		}
		cmps, err := compose.ParseCompose(composeYml)
		if err != nil {
			return nil, orca.OrcaError("compose parse failed", err)
		}
		result = append(result, CollectedCompose{filepath.Base(dir), cmps})

	}
	return result, nil
}
