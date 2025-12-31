package docker

import (
	"orca/errs"
	"orca/model/compose"
	"slices"
)

var _ Inspector = (*fakeInspector)(nil)

type fakeInspector struct {
	Volumes     []string
	Networks    []string
	Binds       []string
	ComposeDirs map[string]*compose.ComposeSpec
}

func newFakeInspector() Inspector {
	return &fakeInspector{
		Volumes:     []string{},
		Networks:    []string{},
		Binds:       []string{},
		ComposeDirs: map[string]*compose.ComposeSpec{},
	}
}

func (f *fakeInspector) NetworkExists(name string) bool {
	return slices.Contains(f.Networks, name)
}
func (f *fakeInspector) VolumeExists(name string) bool {
	return slices.Contains(f.Volumes, name)
}
func (f *fakeInspector) BindExists(dir string) bool {
	return slices.Contains(f.Binds, dir)
}
func (f *fakeInspector) Compose(dir string) (*compose.ComposeSpec, error) {
	c, ok := f.ComposeDirs[dir]
	if !ok || c == nil {
		// In fact, `compose config` returns the same error whether which is invalid the dir or compose file.
		// This ErrComposeNotFound follows the specifications
		return nil, errs.ErrComposeNotFound
	}
	return c, nil
}
