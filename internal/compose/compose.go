package compose

import (
	"errors"
	"orca/infra/applier"
	"orca/ostools"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var _ (composeWriter) = (*applier.ComposeFileWriter)(nil)

type composeInspector interface {
	// docker compose config
	Config(composeDir string) ([]byte, error)
}

// 全てはここから始まる
func GetAllCompose(orcaRoot string,
	c composeInspector) (*ComposeMap, error) {
	result := ComposeMap{}
	dirs, err := ostools.Dirs(orcaRoot)
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		data, err := c.Config(dir)
		if err != nil && errors.Is(err, err) {
			continue
		}

		spec := &ComposeSpec{}
		if err := yaml.Unmarshal(data, spec); err != nil {
			return nil, err
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

type composeWriter interface {
	WriteCompose(string, []byte) (string, error)
}

func (m ComposeMap) DumpAllComposes(cw composeWriter) ([]string, error) {
	result := []string{}
	for name, c := range m {
		b, err := yaml.Marshal(c)
		if err != nil {
			return result, err
		}
		e, err := cw.WriteCompose(name, b)
		if err != nil {
			return result, err
		}
		result = append(result, e)
	}
	return result, nil
}
