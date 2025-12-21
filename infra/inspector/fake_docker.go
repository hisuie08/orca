package inspector

import (
	"slices"
)

type FakeDockerInspector struct {
	Volumes  []string
	Networks []string
	Dirs     []string
}

// VolumeExists docker volume inspect <name>
func (f FakeDockerInspector) VolumeExists(name string) bool {
	return slices.Contains(f.Volumes, name)
}

// NetworkExists docker network inspect <name>
func (f FakeDockerInspector) NetworkExists(name string) bool {
	return slices.Contains(f.Networks, name)
}

// DirExists ボリュームのマウント先確認用
func (f FakeDockerInspector) DirExists(path string) bool {
	return slices.Contains(f.Dirs, path)
}
