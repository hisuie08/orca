package compose

import (
	orca "orca/helper"
	"orca/infra/applier"
	"orca/internal/context"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// 全てはここから始まる
func GetAllCompose(orcaRoot string,
	i ComposeInspector) (composes *ComposeMap, err error) {
	result := ComposeMap{}
	dirs, err := i.Directories()
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		data, err := i.Config(dir)
		if err != nil {
			continue
		}

		spec := &ComposeSpec{}
		if err := yaml.Unmarshal(data, spec); err != nil {
			return nil, orca.OrcaError("compose Parse Error", err)
		}

		result[filepath.Base(dir)] = spec
	}
	return &result, nil
}

// 複数composeをかき集めるやつ
func (m ComposeMap) CollectComposes() []CollectedCompose {
	result := []CollectedCompose{}
	for k, v := range m {
		result = append(result, CollectedCompose{From: k, Spec: v})
	}
	return result
}

func (m ComposeMap) CollectNetworks() []CollectedNetwork {
	result := []CollectedNetwork{}
	for name, c := range m {
		for k, v := range c.Networks {
			result = append(result, CollectedNetwork{
				From: FromRef{Compose: name, Key: k},
				Spec: v,
			})
		}
	}
	return result
}

func (m ComposeMap) CollectVolumes() []CollectedVolume {
	result := []CollectedVolume{}
	for name, c := range m {
		for k, v := range c.Volumes {
			result = append(result, CollectedVolume{
				From: FromRef{Compose: name, Key: k},
				Spec: v,
			})
		}
	}
	return result
}

func (m ComposeMap) DumpAllComposes(ctx context.OrcaContext) ([]string, error) {
	result:=[]string{}
	for name, c := range m {
		b, err := yaml.Marshal(c)
		if err != nil {
			return []string{}, err
		}
		d := func() applier.ComposeWriter {
			switch ctx.RunMode {
			case context.ModeDryRun:
				return applier.FakeDotOrcaDumper{FakeRoot: ctx.OrcaRoot, FakeDir: map[string][]byte{}}
			default:
				return applier.NewDotOrcaDumper(ctx.OrcaRoot)
			}
		}()
		if e,err:=d.DumpCompose(name,b); err != nil {
			return nil, err
		}else{result= append(result,e)}
	}
	return result,nil
}
