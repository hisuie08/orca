package docker

import (
	"errors"
	"orca/errs"
	"orca/model/compose"
	"slices"
	"testing"
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

func TestExistsBool(t *testing.T) {

	ins := newFakeInspector().(*fakeInspector)
	ins.Volumes = []string{"a"}
	ins.Networks = []string{"b"}
	ins.Binds = []string{"c"}

	type testcase map[string]bool
	caseVol := testcase{"a": true, "b": false, "c": false}
	caseNet := testcase{"a": false, "b": true, "c": false}
	caseBnd := testcase{"a": false, "b": false, "c": true}

	for k, v := range caseVol {
		t.Run("volume/"+k, func(t *testing.T) {
			if got := ins.VolumeExists(k); got != v {
				t.Errorf("VolumeExists(%s) expected %t but got %t", k, v, got)
			}
		})
	}
	for k, v := range caseNet {
		t.Run("network/"+k, func(t *testing.T) {
			if got := ins.NetworkExists(k); got != v {
				t.Errorf("NetworkExists(%s) expected %t but got %t", k, v, got)
			}
		})
	}
	for k, v := range caseBnd {
		t.Run("bind/"+k, func(t *testing.T) {
			if got := ins.BindExists(k); got != v {
				t.Errorf("BindExists(%s) expected %t but got %t", k, v, got)
			}
		})
	}
}

func TestInspectorCompose(t *testing.T) {
	ins := newFakeInspector().(*fakeInspector)
	ins.ComposeDirs = map[string]*compose.ComposeSpec{
		"a": &compose.ComposeSpec{},
		"b": &compose.ComposeSpec{},
		"c": nil,
	}
	t.Run("compose", func(t *testing.T) {
		expected := 2
		result := []*compose.ComposeSpec{}
		for _, dir := range []string{"a", "b", "c"} {
			if got, err := ins.Compose(dir); err == nil {
				result = append(result, got)

			}
		}
		if len(result) != expected {
			t.Errorf("expected detects was %d but result contains %d", expected, len(result))
		}
	})
	t.Run("compose_not_found", func(t *testing.T) {
		if _, err := ins.Compose("not-exist"); !errors.Is(err, errs.ErrComposeNotFound) {
			t.Fatalf("expected ErrComposeNotFound but got %v", err)
		}
	})
}
