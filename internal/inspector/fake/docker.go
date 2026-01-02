package fakeinspector

import (
	"orca/internal/inspector"
	"orca/model/compose"
	"slices"
)

var _ inspector.Docker = (*FakeDocker)(nil)

type FakeDocker struct {
	Volumes     []string
	Networks    []string
	Binds       []string
	ComposeDirs map[string]*compose.ComposeSpec
}

func (f *FakeDocker) NetworkExists(name string) bool {
	return slices.Contains(f.Networks, name)
}
func (f *FakeDocker) VolumeExists(name string) bool {
	return slices.Contains(f.Volumes, name)
}
func (f *FakeDocker) BindExists(dir string) bool {
	return slices.Contains(f.Binds, dir)
}
func (f *FakeDocker) Compose(dir string) (*compose.ComposeSpec, error) {
	return f.ComposeDirs[dir], nil
}
