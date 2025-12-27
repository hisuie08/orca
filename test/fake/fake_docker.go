package fake

import (
	"orca/infra/inspector"
	"orca/ostools"
	"slices"
)

var _ inspector.NetworkInspector = (*FakeDocker)(nil)
var _ inspector.Volumenspector = (*FakeDocker)(nil)
var _ inspector.BindInspector = (*FakeDocker)(nil)

type FakeDocker struct {
	Volumes  []string
	Networks []string
}

// VolumeExists docker volume inspect <name>
func (f FakeDocker) VolumeExists(name string) bool {
	return slices.Contains(f.Volumes, name)
}

// NetworkExists docker network inspect <name>
func (f FakeDocker) NetworkExists(name string) bool {
	return slices.Contains(f.Networks, name)
}

func (f FakeDocker) BindExists(dir string) bool {
	return ostools.DirExists(dir)
}
