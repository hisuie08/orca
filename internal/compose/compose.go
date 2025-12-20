package compose

import (
	orca "orca/helper"
	"orca/internal/ostools"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)




// 全てはここから始まる
func GetAllCompose(orcaRoot string) (composes *map[string]*ComposeSpec, err error) {
	result := map[string]*ComposeSpec{}
	dirs, err := ostools.Directories(orcaRoot)
	if err != nil {
		return nil, orca.OrcaError("collect compose failed", err)
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

		from := filepath.Base(dir)
		result[from] = cmps
	}
	os.Remove(stopperCompose) //駆け上がり止めcomposeの削除
	return &result, nil
}

// 複数composeをかき集めるやつ
func CollectComposes(m ComposeMap) []CollectedCompose {
	result := []CollectedCompose{}
	for k, v := range m {
		result = append(result, CollectedCompose{From: k, Spec: v})
	}
	return result
}

// Composeを読み出す関数
func ParseCompose(data []byte) (*ComposeSpec, error) {
	cfg := ComposeSpec{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, orca.OrcaError("compose Parse Error", err)
	}
	return &cfg, nil
}
