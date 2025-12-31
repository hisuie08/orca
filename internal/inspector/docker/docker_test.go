package docker

import (
	"errors"
	"orca/errs"
	"orca/model/compose"
	"testing"
)

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
